/* eslint-disable no-console */
import execa from 'execa';
import { copy } from 'fs-extra';

import {
  emptyDirExceptForDotGit,
  gitCommit,
  LANGUAGES,
  run,
  toAbsolutePath,
  REPO_URL,
  ensureGitHubToken,
} from '../../common';
import { getLanguageFolder, getPackageVersionDefault } from '../../config';
import {
  cloneRepository,
  configureGitHubAuthor,
  RELEASED_TAG,
} from '../../release/common';
import type { Language } from '../../types';
import { getNbGitDiff } from '../utils';

import text from './text';

export function decideWhereToSpread(commitMessage: string): Language[] {
  if (commitMessage.startsWith('chore: release')) {
    return [];
  }

  const result = commitMessage.match(/(.+)\((.+)\):/);
  if (!result) {
    // no scope
    return LANGUAGES;
  }

  const scope = result[2] as Language;
  return LANGUAGES.includes(scope) ? [scope] : LANGUAGES;
}

export function cleanUpCommitMessage(commitMessage: string): string {
  const isCodeGenCommit = commitMessage.startsWith(text.commitStartMessage);

  if (isCodeGenCommit) {
    const hash = commitMessage
      .split(text.commitStartMessage)[1]
      .replace('. [skip ci]', '')
      .trim();

    if (!hash) {
      return commitMessage;
    }

    return [
      `${text.commitStartMessage} ${hash.substring(0, 8)}. [skip ci]`,
      `${REPO_URL}/commit/${hash}`,
    ].join('\n\n');
  }

  const prCommit = commitMessage.match(/(.+)\s\(#(\d+)\)$/);

  if (!prCommit) {
    return commitMessage;
  }

  return [prCommit[1], `${REPO_URL}/pull/${prCommit[2]}`].join('\n\n');
}

async function spreadGeneration(): Promise<void> {
  const githubToken = ensureGitHubToken();

  const lastCommitMessage = await run('git log -1 --format="%s"');
  const author = (
    await run('git log -1 --format="Co-authored-by: %an <%ae>"')
  ).trim();
  const coAuthors = (
    await run('git log -1 --format="%(trailers:key=Co-authored-by)"')
  )
    .split('\n')
    .map((coAuthor) => coAuthor.trim())
    .filter(Boolean);

  const IS_RELEASE_COMMIT = lastCommitMessage.startsWith(
    text.commitPrepareReleaseMessage
  );
  const commitMessage = cleanUpCommitMessage(lastCommitMessage);
  const langs = decideWhereToSpread(lastCommitMessage);

  // At this point, we know the release will happen on at least one client
  // So we want to set the released tag at the monorepo level too.
  if (IS_RELEASE_COMMIT) {
    console.log('Processing release commit');

    // remove old `released` tag
    await run(
      `git fetch origin refs/tags/${RELEASED_TAG}:refs/tags/${RELEASED_TAG}`
    );
    await run(`git tag -d ${RELEASED_TAG}`);
    await run(`git push --delete origin ${RELEASED_TAG}`);

    // create new `released` tag
    await run(`git tag released`);
    await run(`git push --tags`);
  }

  for (const lang of langs) {
    const { tempGitDir } = await cloneRepository({
      lang,
      githubToken,
      tempDir: process.env.RUNNER_TEMP!,
    });

    const clientPath = toAbsolutePath(getLanguageFolder(lang));
    await emptyDirExceptForDotGit(tempGitDir);
    await copy(clientPath, tempGitDir, { preserveTimestamps: true });

    if (
      (await getNbGitDiff({
        head: null,
        cwd: tempGitDir,
      })) === 0
    ) {
      console.log(
        `❎ Skipping ${lang} repository, because there is no change.`
      );
      continue;
    }

    const version = getPackageVersionDefault(lang);
    const message = IS_RELEASE_COMMIT
      ? `chore: release ${version}`
      : commitMessage;

    await configureGitHubAuthor(tempGitDir);
    await run(`git add .`, { cwd: tempGitDir });
    await gitCommit({
      message,
      coAuthors: [author, ...coAuthors],
      cwd: tempGitDir,
    });
    await execa('git', ['tag', version], {
      cwd: tempGitDir,
    });
    await run(IS_RELEASE_COMMIT ? 'git push --follow-tags' : 'git push', {
      cwd: tempGitDir,
    });
    console.log(`✅ Spread the generation to ${lang} repository.`);
  }
}

if (require.main === module) {
  spreadGeneration();
}
