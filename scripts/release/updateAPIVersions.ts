/* eslint-disable no-console */
import fsp from 'fs/promises';

import dotenv from 'dotenv';
import semver from 'semver';

import clientsConfig from '../../config/clients.config.json';
import openapiConfig from '../../config/openapitools.json';
import generationCommitText from '../ci/codegen/text';
import {
  ROOT_ENV_PATH,
  toAbsolutePath,
  run,
  exists,
  GENERATORS,
  LANGUAGES,
  MAIN_BRANCH,
  gitBranchExists,
  CLIENTS_JS_UTILS,
} from '../common';
import {
  getClientsConfigField,
  getGitHubUrl,
  getLanguageFolder,
} from '../config';
import type { Language } from '../types';

import { readJsonFile, writeJsonFile } from './common';
import type { Changelog, Versions, VersionsToRelease } from './types';

dotenv.config({ path: ROOT_ENV_PATH });

/**
 * Bump each client version of the JavaScript client in workspace places and config files.
 *
 * We don't use the pre-computed `next` version for JavaScript, because the packages have independent versioning.
 */
async function updateVersionForJavascript(
  jsVersion: NonNullable<VersionsToRelease['javascript']>
): Promise<void> {
  // Sets the new version of the utils package
  const utilsPackageVersion = getClientsConfigField(
    'javascript',
    'utilsPackageVersion'
  );
  const nextUtilsPackageVersion = semver.inc(
    utilsPackageVersion,
    jsVersion.releaseType
  );

  if (!nextUtilsPackageVersion) {
    throw new Error(
      `Failed to bump version ${utilsPackageVersion} by ${jsVersion.releaseType}.`
    );
  }

  clientsConfig.javascript.utilsPackageVersion = nextUtilsPackageVersion;

  // update local playground deps
  const nodePgPackageFile = await readJsonFile(
    toAbsolutePath('playground/javascript/node/package.json')
  );
  const browserPgPackageFile = await readJsonFile(
    toAbsolutePath('playground/javascript/browser/package.json')
  );

  if (!nodePgPackageFile || !browserPgPackageFile) {
    throw new Error('Failed to read playground package files');
  }

  // Sets the new version of the JavaScript client
  Object.values(GENERATORS)
    .filter((gen) => gen.language === 'javascript')
    .forEach((gen) => {
      const additionalProperties = gen.additionalProperties!;
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

      if (!additionalProperties.packageName) {
        throw new Error(
          `Package name is missing for JavaScript - ${gen.client}.`
        );
      }

      if (nodePgPackageFile.dependencies[additionalProperties.packageName]) {
        nodePgPackageFile.dependencies[additionalProperties.packageName] =
          newVersion;
      }

      if (browserPgPackageFile.dependencies[additionalProperties.packageName]) {
        browserPgPackageFile.dependencies[additionalProperties.packageName] =
          newVersion;
      }
    });

  CLIENTS_JS_UTILS.forEach((util) => {
    const utilPackageName = `${clientsConfig.javascript.npmNamespace}/${util}`;

    if (nodePgPackageFile.dependencies[utilPackageName]) {
      nodePgPackageFile.dependencies[utilPackageName] = nextUtilsPackageVersion;
    }

    if (browserPgPackageFile.dependencies[utilPackageName]) {
      browserPgPackageFile.dependencies[utilPackageName] =
        nextUtilsPackageVersion;
    }
  });

  // update `openapitools.json` config file
  await writeJsonFile(
    toAbsolutePath('config/openapitools.json'),
    openapiConfig
  );

  // update `package.json` node playground file
  await writeJsonFile(
    toAbsolutePath('playground/javascript/node/package.json'),
    nodePgPackageFile
  );

  // update `package.json` browser playground file
  await writeJsonFile(
    toAbsolutePath('playground/javascript/browser/package.json'),
    browserPgPackageFile
  );

  // update `clients.config.json` file for the utils version
  await writeJsonFile(
    toAbsolutePath('config/clients.config.json'),
    clientsConfig
  );
}

async function updateConfigFiles(
  versionsToRelease: VersionsToRelease
): Promise<void> {
  if (versionsToRelease.javascript) {
    await updateVersionForJavascript(versionsToRelease.javascript);
  }

  // update the other versions in clients.config.json
  LANGUAGES.forEach((lang) => {
    if (lang === 'javascript' || !versionsToRelease[lang]) {
      return;
    }

    clientsConfig[lang].packageVersion = versionsToRelease[lang]!.next;
  });

  await writeJsonFile(
    toAbsolutePath('config/clients.config.json'),
    clientsConfig
  );
}

async function updateChangelog({
  lang,
  changelog,
  current,
  next,
}: {
  lang: Language;
  changelog: Changelog;
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
  await fsp.writeFile(
    changelogPath,
    [changelogHeader, changelog[lang], existingContent].join('\n\n')
  );
}

export function getVersionsToRelease(versions: Versions): VersionsToRelease {
  const versionsToRelease: VersionsToRelease = {};

  Object.entries(versions).forEach(
    ([lang, { noCommit, current, skipRelease, releaseType, next }]) => {
      if (noCommit || skipRelease || !current || !next) {
        return;
      }

      if (
        !releaseType ||
        !['major', 'minor', 'patch', 'prerelease'].includes(releaseType)
      ) {
        throw new Error(
          `\`${releaseType}\` is unknown release type. Allowed: major, minor, patch, prerelease`
        );
      }

      versionsToRelease[lang] = {
        current,
        releaseType,
        next,
      };
    }
  );

  return versionsToRelease;
}

/**
 * Updates the changelogs and the config files containing versions of the API clients, then pushes the changes to the `headBranch`.
 *
 * @param versions - A summary of the version changes, with their new version and release type.
 * @param changelog - The changelog of all the languages, which is generated by `createReleasePR`.
 * @param headBranch - The branch to push the changes to.
 */
export async function updateAPIVersions(
  versions: Versions,
  changelog: Changelog,
  headBranch: string
): Promise<void> {
  if (await gitBranchExists(headBranch)) {
    await run(`git fetch origin ${headBranch}`);
    await run(`git push -d origin ${headBranch}`);
  }

  await run(`git checkout -b ${headBranch}`);

  const versionsToRelease = getVersionsToRelease(versions);

  await updateConfigFiles(versionsToRelease);

  for (const [lang, { current, releaseType, next }] of Object.entries(
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
    if (lang === 'javascript') {
      await run(
        `yarn workspace algoliasearch-client-javascript release:bump ${releaseType}`,
        {
          verbose: true,
        }
      );
    }

    await updateChangelog({
      lang: lang as Language,
      changelog,
      current,
      next,
    });
  }

  console.log(`Pushing updated changes to ${headBranch}`);
  const commitMessage = generationCommitText.commitPrepareReleaseMessage;
  await run(`git add clients config`, { verbose: true });
  if (process.env.LOCAL_TEST_DEV) {
    await run(`CI=true git commit -m "${commitMessage} [skip ci]"`, {
      verbose: true,
    });
  } else {
    await run(`CI=true git commit -m "${commitMessage}"`, { verbose: true });
  }

  await run(`git push origin ${headBranch}`, { verbose: true });
  await run(`git checkout ${MAIN_BRANCH}`, { verbose: true });
}
