/* eslint-disable no-console */
import chalk from 'chalk';
import dotenv from 'dotenv';
import semver from 'semver';

import generationCommitText from '../ci/codegen/text';
import { getNbGitDiff } from '../ci/utils';
import {
  LANGUAGES,
  ROOT_ENV_PATH,
  run,
  MAIN_BRANCH,
  OWNER,
  REPO,
  getOctokit,
  ensureGitHubToken,
  TODAY,
} from '../common';
import { getPackageVersionDefault } from '../config';

import { RELEASED_TAG } from './common';
import TEXT from './text';
import type {
  Versions,
  VersionsWithoutReleaseType,
  PassedCommit,
  Commit,
  Scope,
  Changelog,
} from './types';
import { updateAPIVersions } from './updateAPIVersions';

dotenv.config({ path: ROOT_ENV_PATH });

const COMMON_SCOPES = ['specs'];

export function readVersions(): VersionsWithoutReleaseType {
  return Object.fromEntries(
    LANGUAGES.map((lang) => [lang, { current: getPackageVersionDefault(lang) }])
  );
}

export function getVersionChangesText(versions: Versions): string {
  return LANGUAGES.map((lang) => {
    const { current, releaseType, noCommit, skipRelease } = versions[lang];

    if (noCommit) {
      return `- ~${lang}: ${current} (${TEXT.noCommit})~`;
    }

    if (!current) {
      return `- ~${lang}: (${TEXT.currentVersionNotFound})~`;
    }

    const next = semver.inc(current, releaseType!);

    if (skipRelease) {
      return [
        `- ~${lang}: ${current} -> **\`${releaseType}\` _(e.g. ${next})_**~`,
        TEXT.descriptionForSkippedLang,
      ].join('\n');
    }

    return `- ${lang}: ${current} -> **\`${releaseType}\` _(e.g. ${next})_**`;
  }).join('\n');
}

export function getSkippedCommitsText({
  commitsWithoutLanguageScope,
  commitsWithUnknownLanguageScope,
}: {
  commitsWithoutLanguageScope: string[];
  commitsWithUnknownLanguageScope: string[];
}): string {
  if (
    commitsWithoutLanguageScope.length === 0 &&
    commitsWithUnknownLanguageScope.length === 0
  ) {
    return '_(None)_';
  }

  // GitHub API restrict the size of a PR body, if we send too many commits
  // we might end up with 502 errors when trying to send the pull request
  // So we limit the size of the missed commits
  return `
<p>${TEXT.skippedCommitsDesc}</p>

<details>
  <summary>
    <i>Commits without language scope:</i>
  </summary>

  ${commitsWithoutLanguageScope
    .slice(0, 15)
    .map((commit) => `- ${commit}`)
    .join('\n')}
</details>

<details>
  <summary>
    <i>Commits with unknown language scope:</i>
  </summary>

  ${commitsWithUnknownLanguageScope
    .slice(0, 15)
    .map((commit) => `- ${commit}`)
    .join('\n')}
</details>`;
}

export function parseCommit(commit: string): Commit {
  const LENGTH_SHA1 = 8;
  const hash = commit.slice(0, LENGTH_SHA1);
  let message = commit.slice(LENGTH_SHA1 + 1);
  let type = message.slice(0, message.indexOf(':'));
  const matchResult = type.match(/(.+)\((.+)\)/);

  if (
    message
      .toLocaleLowerCase()
      .startsWith(generationCommitText.commitStartMessage)
  ) {
    return {
      error: 'generation-commit',
    };
  }

  if (!matchResult) {
    return {
      error: 'missing-language-scope',
    };
  }
  message = message.slice(message.indexOf(':') + 1).trim();
  type = matchResult[1];
  const scope = matchResult[2] as Scope;
  // A spec commit should be added to every clients, as it mostly imply a client change.
  const allowedScopes = [...LANGUAGES, ...COMMON_SCOPES];

  if (!allowedScopes.includes(scope)) {
    return { error: 'unknown-language-scope' };
  }

  return {
    hash,
    type, // `fix` | `feat` | `chore` | ...
    scope, // `specs` | `javascript` | `php` | `java` | ...
    message,
    raw: commit,
  };
}

/* eslint-disable no-param-reassign */
export function decideReleaseStrategy({
  versions,
  commits,
}: {
  versions: VersionsWithoutReleaseType;
  commits: PassedCommit[];
}): Versions {
  return Object.entries(versions).reduce(
    (versionsWithReleaseType: Versions, [lang, version]) => {
      const commitsPerLang = commits.filter(
        (commit) =>
          commit.scope === lang || COMMON_SCOPES.includes(commit.scope)
      );
      const currentVersion = versions[lang].current;

      if (commitsPerLang.length === 0) {
        versionsWithReleaseType[lang] = {
          ...version,
          noCommit: true,
          releaseType: null,
        };
        return versionsWithReleaseType;
      }

      if (semver.prerelease(currentVersion)) {
        // if version is like 0.1.2-beta.1, it increases to 0.1.2-beta.2, even if there's a breaking change.
        versionsWithReleaseType[lang] = {
          ...version,
          releaseType: 'prerelease',
        };
        return versionsWithReleaseType;
      }

      if (
        commitsPerLang.some((commit) =>
          commit.message.includes('BREAKING CHANGE')
        )
      ) {
        versionsWithReleaseType[lang] = {
          ...version,
          releaseType: 'major',
        };
        return versionsWithReleaseType;
      }

      const commitTypes = new Set(commitsPerLang.map(({ type }) => type));
      if (commitTypes.has('feat')) {
        versionsWithReleaseType[lang] = {
          ...version,
          releaseType: 'minor',
        };
        return versionsWithReleaseType;
      }

      versionsWithReleaseType[lang] = {
        ...version,
        releaseType: 'patch',
        ...(commitTypes.has('fix') ? undefined : { skipRelease: true }),
      };
      return versionsWithReleaseType;
    },
    {}
  );
}
/* eslint-enable no-param-reassign */

/**
 * Returns commits separated in categories used to compute the next release version.
 *
 * Gracefully exits if there is none.
 */
async function getCommits(): Promise<{
  validCommits: PassedCommit[];
  skippedCommits: string;
}> {
  // Reading commits since last release
  const latestCommits = (
    await run(`git log --oneline --abbrev=8 ${RELEASED_TAG}..${MAIN_BRANCH}`)
  )
    .split('\n')
    .filter(Boolean);

  const commitsWithoutLanguageScope: string[] = [];
  const commitsWithUnknownLanguageScope: string[] = [];

  const validCommits = latestCommits
    .map((commitMessage) => {
      const commit = parseCommit(commitMessage);

      if ('error' in commit) {
        // We don't do anything in that case, as we don't really care about
        // those commits
        if (commit.error === 'generation-commit') {
          return undefined;
        }

        if (commit.error === 'missing-language-scope') {
          commitsWithoutLanguageScope.push(commitMessage);
          return undefined;
        }

        if (commit.error === 'unknown-language-scope') {
          commitsWithUnknownLanguageScope.push(commitMessage);
          return undefined;
        }
      }

      return commit;
    })
    .filter(Boolean) as PassedCommit[];

  if (validCommits.length === 0) {
    console.log(
      chalk.black.bgYellow('[INFO]'),
      `Skipping release because no valid commit has been added since \`released\` tag.`
    );
    // eslint-disable-next-line no-process-exit
    process.exit(0);
  }

  return {
    validCommits,
    skippedCommits: getSkippedCommitsText({
      commitsWithoutLanguageScope,
      commitsWithUnknownLanguageScope,
    }),
  };
}

async function createReleasePR(): Promise<void> {
  ensureGitHubToken();

  if (!process.env.LOCAL_TEST_DEV) {
    if ((await run('git rev-parse --abbrev-ref HEAD')) !== MAIN_BRANCH) {
      throw new Error(
        `You can run this script only from \`${MAIN_BRANCH}\` branch.`
      );
    }

    if (
      (await getNbGitDiff({
        head: null,
      })) !== 0
    ) {
      throw new Error(
        'Working directory is not clean. Commit all the changes first.'
      );
    }
  }

  await run(`git rev-parse --verify refs/tags/${RELEASED_TAG}`, {
    errorMessage: '`released` tag is missing in this repository.',
  });

  console.log('Pulling from origin...');
  await run('git fetch origin');
  await run('git pull');

  // Remove the local tag, and fetch it from the remote.
  // We move the `released` tag as we release, so we need to make it up-to-date.
  await run(`git tag -d ${RELEASED_TAG}`);
  await run(
    `git fetch origin refs/tags/${RELEASED_TAG}:refs/tags/${RELEASED_TAG}`
  );

  console.log('Search for commits since last release...');
  const { validCommits, skippedCommits } = await getCommits();

  const versions = decideReleaseStrategy({
    versions: readVersions(),
    commits: validCommits,
  });
  const versionChanges = getVersionChangesText(versions);

  console.log('Creating changelogs for all languages...');
  const changelog: Changelog = LANGUAGES.reduce((newChangelog, lang) => {
    if (versions[lang].noCommit) {
      return newChangelog;
    }

    return {
      ...newChangelog,
      [lang]: validCommits
        .filter(
          (commit) =>
            commit.scope === lang || COMMON_SCOPES.includes(commit.scope)
        )
        .map((commit) => `- ${commit.raw}`)
        .join('\n'),
    };
  }, {} as Changelog);

  const headBranch = `chore/prepare-release-${TODAY}`;

  console.log('Updating config files...');
  await updateAPIVersions(versionChanges, changelog, headBranch);

  console.log('Creating pull request...');
  const octokit = getOctokit();

  const {
    data: { number, html_url: url },
  } = await octokit.pulls.create({
    owner: OWNER,
    repo: REPO,
    title: generationCommitText.commitPrepareReleaseMessage,
    body: [
      TEXT.header,
      TEXT.summary,
      TEXT.versionChangeHeader,
      versionChanges,
      TEXT.skippedCommitsHeader,
      skippedCommits,
    ].join('\n\n'),
    base: 'main',
    head: headBranch,
  });

  console.log('Assigning team members to the PR...');
  await octokit.pulls.requestReviewers({
    owner: OWNER,
    repo: REPO,
    pull_number: number,
    team_reviewers: ['api-clients-automation'],
  });

  console.log(`Release PR #${number} is ready for review.`);
  console.log(`  > ${url}`);
}

// JS version of `if __name__ == '__main__'`
if (require.main === module) {
  createReleasePR();
}
