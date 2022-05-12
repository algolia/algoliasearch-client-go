/* eslint-disable no-console */
import fsp from 'fs/promises';

import dotenv from 'dotenv';
import execa from 'execa';
import { copy } from 'fs-extra';
import semver from 'semver';

import clientsConfig from '../../config/clients.config.json';
import openapiConfig from '../../config/openapitools.json';
import {
  ROOT_ENV_PATH,
  toAbsolutePath,
  run,
  exists,
  getGitHubUrl,
  gitCommit,
  OWNER,
  REPO,
  emptyDirExceptForDotGit,
  GENERATORS,
  LANGUAGES,
} from '../common';
import { getLanguageFolder, getPackageVersionDefault } from '../config';
import type { Language } from '../types';

import {
  RELEASED_TAG,
  TEAM_SLUG,
  getMarkdownSection,
  configureGitHubAuthor,
  cloneRepository,
  getOctokit,
} from './common';
import TEXT from './text';
import type {
  VersionsToRelease,
  BeforeClientGenerationCommand,
  BeforeClientCommitCommand,
} from './types';

dotenv.config({ path: ROOT_ENV_PATH });

const BEFORE_CLIENT_GENERATION: {
  [lang in Language]?: BeforeClientGenerationCommand;
} = {
  javascript: async ({ releaseType, dir }) => {
    await run(`yarn release:bump ${releaseType}`, { cwd: dir });
  },
};

const BEFORE_CLIENT_COMMIT: { [lang: string]: BeforeClientCommitCommand } = {
  javascript: async ({ dir }) => {
    // https://github.com/yarnpkg/berry/issues/2948
    await run(`YARN_ENABLE_IMMUTABLE_INSTALLS=false yarn`, { cwd: dir }); // generate `yarn.lock` file
  },
};

async function getIssueBody(): Promise<string> {
  const octokit = getOctokit(process.env.GITHUB_TOKEN!);
  const {
    data: { body },
  } = await octokit.rest.issues.get({
    owner: OWNER,
    repo: REPO,
    issue_number: Number(process.env.EVENT_NUMBER),
  });

  if (!body) {
    throw new Error(
      `Unexpected \`body\` of the release issue: ${JSON.stringify(body)}`
    );
  }
  return body;
}

function getDateStamp(): string {
  return new Date().toISOString().split('T')[0];
}

export function getVersionsToRelease(issueBody: string): VersionsToRelease {
  const versionsToRelease: VersionsToRelease = {};

  getMarkdownSection(issueBody, TEXT.versionChangeHeader)
    .split('\n')
    .forEach((line) => {
      const result = line.match(/- \[x\] (.+): (.+) -> `(.+)`/);
      if (!result) {
        return;
      }
      const [, lang, current, releaseType] = result;
      if (!['major', 'minor', 'patch', 'prerelease'].includes(releaseType)) {
        throw new Error(
          `\`${releaseType}\` is unknown release type. Allowed: major, minor, patch, prerelease`
        );
      }
      versionsToRelease[lang] = {
        current,
        releaseType,
      };
    });

  return versionsToRelease;
}

// Bump each client version of the JavaScript client in openapitools.json
async function updateVersionForJavascript(
  versionsToRelease: VersionsToRelease
): Promise<void> {
  if (!versionsToRelease.javascript) {
    return;
  }
  const jsVersion = versionsToRelease.javascript;
  const nextUtilsPackageVersion =
    semver.inc(
      openapiConfig['generator-cli'].generators['javascript-search']
        .additionalProperties.utilsPackageVersion,
      jsVersion.releaseType
    ) || '';
  Object.values(GENERATORS)
    .filter((gen) => gen.language === 'javascript')
    .forEach((gen) => {
      const additionalProperties =
        openapiConfig['generator-cli'].generators[gen.key].additionalProperties;

      const newVersion = semver.inc(
        additionalProperties.packageVersion,
        jsVersion.releaseType
      );
      if (!newVersion) {
        throw new Error(
          `Failed to bump version ${additionalProperties.packageVersion} by ${jsVersion.releaseType}.`
        );
      }
      additionalProperties.packageVersion = newVersion;
      additionalProperties.utilsPackageVersion = nextUtilsPackageVersion;
    });
  await fsp.writeFile(
    toAbsolutePath('config/openapitools.json'),
    JSON.stringify(openapiConfig, null, 2)
  );
}

async function updateConfigFiles(
  versionsToRelease: VersionsToRelease
): Promise<void> {
  await updateVersionForJavascript(versionsToRelease);

  // update the other versions in clients.config.json
  LANGUAGES.forEach((lang) => {
    if (lang === 'javascript' || !versionsToRelease[lang]) {
      return;
    }
    const releaseType = versionsToRelease[lang]!.releaseType;

    const newVersion = semver.inc(getPackageVersionDefault(lang), releaseType);
    if (!newVersion) {
      throw new Error(
        `Failed to bump version ${getPackageVersionDefault(
          lang
        )} by ${releaseType}.`
      );
    }
    clientsConfig[lang].packageVersion = newVersion;
  });
  await fsp.writeFile(
    toAbsolutePath('config/clients.config.json'),
    JSON.stringify(clientsConfig, null, 2)
  );
}

async function updateChangelog({
  lang,
  issueBody,
  current,
  next,
}: {
  lang: Language;
  issueBody: string;
  current: string;
  next: string;
}): Promise<void> {
  const changelogPath = toAbsolutePath(
    `${getLanguageFolder(lang)}/CHANGELOG.md`
  );
  const existingContent = (await exists(changelogPath))
    ? (await fsp.readFile(changelogPath)).toString()
    : '';
  const changelogHeader = `## [${next}](${getGitHubUrl(
    lang
  )}/compare/${current}...${next})`;
  const newChangelog = getMarkdownSection(
    getMarkdownSection(issueBody, TEXT.changelogHeader),
    `### ${lang}`
  );
  await fsp.writeFile(
    changelogPath,
    [changelogHeader, newChangelog, existingContent].join('\n\n')
  );
}

function formatGitTag({
  lang,
  version,
}: {
  lang: string;
  version: string;
}): string {
  return lang === 'go' ? `v${version}` : version;
}

async function isAuthorizedRelease(): Promise<boolean> {
  const octokit = getOctokit(process.env.GITHUB_TOKEN!);
  const { data: members } = await octokit.rest.teams.listMembersInOrg({
    org: OWNER,
    team_slug: TEAM_SLUG,
  });

  const { data: comments } = await octokit.rest.issues.listComments({
    owner: OWNER,
    repo: REPO,
    issue_number: Number(process.env.EVENT_NUMBER),
  });

  return comments.some(
    (comment) =>
      comment.body?.toLowerCase().trim() === 'approved' &&
      members.find((member) => member.login === comment.user?.login)
  );
}

async function processRelease(): Promise<void> {
  if (!process.env.GITHUB_TOKEN) {
    throw new Error('Environment variable `GITHUB_TOKEN` does not exist.');
  }

  if (!process.env.EVENT_NUMBER) {
    throw new Error('Environment variable `EVENT_NUMBER` does not exist.');
  }

  if (!(await isAuthorizedRelease())) {
    throw new Error(
      'The issue was not approved.\nA team member must leave a comment "approved" in the release issue.'
    );
  }

  const issueBody = await getIssueBody();
  const versionsToRelease = getVersionsToRelease(issueBody);

  await updateConfigFiles(versionsToRelease);

  for (const [lang, { current, releaseType }] of Object.entries(
    versionsToRelease
  )) {
    /*
    About bumping versions of JS clients:

    There are generated clients in JS repo, and non-generated clients like `algoliasearch`, `client-common`, etc.
    Now that the versions of generated clients are updated in `openapitools.json`,
    the generation output will have correct new versions.
    
    However, we need to manually update versions of the non-generated (a.k.a. manually written) clients.
    In order to do that, we run `yarn release:bump <releaseType>` in this monorepo first.
    It will update the versions of the non-generated clients which exists in this monorepo.
    After that, we generate clients with new versions. And then, we copy all of them over to JS repository.
    */
    await BEFORE_CLIENT_GENERATION[lang]?.({
      releaseType,
      dir: toAbsolutePath(getLanguageFolder(lang as Language)),
    });

    console.log(`Generating ${lang} client(s)...`);
    console.log(await run(`yarn cli generate ${lang}`));

    const next = semver.inc(current, releaseType);
    await updateChangelog({
      lang: lang as Language,
      issueBody,
      current,
      next: next!,
    });
  }

  // We push commits to each repository AFTER all the generations are done.
  // Otherwise, we will end up having broken release.
  for (const [lang, { current, releaseType }] of Object.entries(
    versionsToRelease
  )) {
    const { tempGitDir } = await cloneRepository({
      lang,
      githubToken: process.env.GITHUB_TOKEN,
      tempDir: process.env.RUNNER_TEMP!,
    });

    const clientPath = toAbsolutePath(getLanguageFolder(lang as Language));
    await emptyDirExceptForDotGit(tempGitDir);
    await copy(clientPath, tempGitDir, { preserveTimestamps: true });

    await configureGitHubAuthor(tempGitDir);
    await BEFORE_CLIENT_COMMIT[lang]?.({
      dir: tempGitDir,
    });
    await run(`git add .`, { cwd: tempGitDir });

    const next = semver.inc(current, releaseType);
    const tag = formatGitTag({ lang, version: next! });
    await gitCommit({
      message: `chore: release ${tag}`,
      cwd: tempGitDir,
    });
    await execa('git', ['tag', tag], {
      cwd: tempGitDir,
    });
    await run(`git push --follow-tags`, { cwd: tempGitDir });
  }

  // Commit and push from the monorepo level.
  await configureGitHubAuthor();
  await run(`git add .`);
  const dateStamp = getDateStamp();
  await gitCommit({
    message: `chore: release ${dateStamp}`,
  });
  await run(`git push`);

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

// JS version of `if __name__ == '__main__'`
if (require.main === module) {
  processRelease();
}
