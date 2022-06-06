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
  if (commitMessage.startsWith(text.commitReleaseMessage)) {
    return LANGUAGES;
  }

  const result = commitMessage.match(/(.+)\((.+)\):/);
  if (!result) {
    // no scope
    return LANGUAGES;
  }

  const scope = result[2] as Language;
  return LANGUAGES.includes(scope) ? [scope] : LANGUAGES;
}

export function cleanUpCommitMessage(
  commitMessage: string,
  version: string
): string {
  if (commitMessage.startsWith(text.commitReleaseMessage)) {
    return `chore: release ${version}`;
  }

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

  console.log('Starting spread generation script...');
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
    text.commitReleaseMessage
  );
  const langs = decideWhereToSpread(lastCommitMessage);
  console.log(
    'Spreading code to the following repositories:',
    langs.join(' | ')
  );

  // At this point, we know the release will happen on at least one client
  // So we want to set the released tag at the monorepo level too.
  if (IS_RELEASE_COMMIT) {
    console.log(
      'Processing release commit, removing old `released` tag on the monorepo'
    );
    await run(
      `git fetch origin refs/tags/${RELEASED_TAG}:refs/tags/${RELEASED_TAG}`
    );
    await run(`git tag -d ${RELEASED_TAG}`);
    await run(`git push --delete origin ${RELEASED_TAG}`);

    console.log('Creating new `released` tag for latest commit');
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

    // We want to ensure we have an up to date `yarn.lock` in the JS client repository
    if (lang === 'javascript') {
      await run('YARN_ENABLE_IMMUTABLE_INSTALLS=false yarn install', {
        cwd: tempGitDir,
      });
    }

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
    } else {
      console.log(`✅ Spreading code to the ${lang} repository.`);
    }

    const version = getPackageVersionDefault(lang);
    const commitMessage = cleanUpCommitMessage(lastCommitMessage, version);

    await configureGitHubAuthor(tempGitDir);
    await run('git add .', { cwd: tempGitDir });
    await gitCommit({
      message: commitMessage,
      coAuthors: [author, ...coAuthors],
      cwd: tempGitDir,
    });
    await execa('git', ['tag', version], {
      cwd: tempGitDir,
    });
    await run(IS_RELEASE_COMMIT ? 'git push --follow-tags' : 'git push', {
      cwd: tempGitDir,
    });
    console.log(
      `✅ Code generation successfully pushed to ${lang} repository.`
    );
  }
}

if (require.main === module) {
  spreadGeneration();
}
