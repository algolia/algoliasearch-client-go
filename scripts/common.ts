import fsp from 'fs/promises';
import path from 'path';

import execa from 'execa'; // https://github.com/sindresorhus/execa/tree/v5.1.1
import { hashElement } from 'folder-hash';
import { remove } from 'fs-extra';

import openapiConfig from '../config/openapitools.json';
import releaseConfig from '../config/release.config.json';

import { getPackageVersionDefault } from './config';
import { createSpinner } from './oraLog';
import type {
  CheckForCache,
  CheckForCacheOptions,
  Generator,
  Language,
  RunOptions,
} from './types';

export const MAIN_BRANCH = releaseConfig.mainBranch;
export const OWNER = releaseConfig.owner;
export const REPO = releaseConfig.repo;
export const REPO_URL = `https://github.com/${OWNER}/${REPO}`;

export const CI = Boolean(process.env.CI);
export const DOCKER = Boolean(process.env.DOCKER);
export const BUNDLE_WITH_DOC = process.env.BUNDLE_WITH_DOC === 'true';

// This script is run by `yarn workspace ...`, which means the current working directory is `./script`
export const ROOT_DIR = path.resolve(process.cwd(), '..');

export const ROOT_ENV_PATH = path.resolve(ROOT_DIR, '.env');

export const GENERATORS: Record<string, Generator> = {
  // Default `algoliasearch` package as it's built similarly to generated clients
  'javascript-algoliasearch': {
    language: 'javascript',
    client: 'algoliasearch',
    key: 'javascript-algoliasearch',
    additionalProperties: {
      buildFile: 'algoliasearch',
      packageName: '@experimental-api-clients-automation/algoliasearch',
      packageVersion: getPackageVersionDefault('javascript'),
    },
  },
};

// Build `GENERATORS` from the openapitools file
Object.entries(openapiConfig['generator-cli'].generators).forEach(
  ([key, gen]) => {
    GENERATORS[key] = {
      ...gen,
      output: gen.output.replace('#{cwd}/', ''),
      ...splitGeneratorKey(key),
    };
  }
);

export const LANGUAGES = [
  ...new Set(Object.values(GENERATORS).map((gen) => gen.language)),
];

export const CLIENTS_JS_UTILS = [
  'client-common',
  'requester-browser-xhr',
  'requester-node-http',
];

export const CLIENTS_JS = [
  ...new Set(Object.values(GENERATORS).map((gen) => gen.client)),
];

export const CLIENTS = CLIENTS_JS.filter(
  (client) => client !== 'algoliasearch'
);

/**
 * Takes a generator key in the form 'language-client' and returns the Generator object.
 */
export function splitGeneratorKey(
  generatorKey: string
): Pick<Generator, 'client' | 'key' | 'language'> {
  const language = generatorKey.slice(0, generatorKey.indexOf('-')) as Language;
  const client = generatorKey.slice(generatorKey.indexOf('-') + 1);
  return { language, client, key: generatorKey };
}

export function createGeneratorKey({
  language,
  client,
}: {
  language: Language | 'all';
  client: string;
}): string {
  return `${language}-${client}`;
}

export async function run(
  command: string,
  { errorMessage, verbose, cwd = ROOT_DIR }: RunOptions = {}
): Promise<string> {
  try {
    if (verbose) {
      return (
        (
          await execa.command(command, {
            stdout: 'inherit',
            stderr: 'inherit',
            stdin: 'inherit',
            all: true,
            shell: 'bash',
            cwd,
          })
        ).all ?? ''
      );
    }
    return (
      (await execa.command(command, { shell: 'bash', all: true, cwd })).all ??
      ''
    );
  } catch (err) {
    if (errorMessage) {
      throw new Error(`[ERROR] ${errorMessage}`);
    } else {
      throw err;
    }
  }
}

export async function exists(ppath: string): Promise<boolean> {
  try {
    await fsp.stat(ppath);
    return true;
  } catch {
    return false;
  }
}

export function toAbsolutePath(ppath: string): string {
  return path.resolve(ROOT_DIR, ppath);
}

export async function runIfExists(
  scriptFile: string,
  args: string,
  opts: RunOptions = {}
): Promise<string> {
  if (await exists(toAbsolutePath(scriptFile))) {
    return await run(`${scriptFile} ${args}`, opts);
  }
  return '';
}

export async function gitCommit({
  message,
  coAuthors,
  cwd = ROOT_DIR,
}: {
  message: string;
  coAuthors?: string[];
  cwd?: string;
}): Promise<void> {
  const messageWithCoAuthors = coAuthors
    ? `${message}\n\n\n${coAuthors.join('\n')}`
    : message;

  await execa('git', ['commit', '-m', messageWithCoAuthors], {
    cwd,
  });
}

export async function checkForCache({
  folder,
  generatedFiles,
  filesToCache,
  cacheFile,
}: CheckForCacheOptions): Promise<CheckForCache> {
  const cache: CheckForCache = {
    cacheExists: false,
    hash: '',
  };
  const generatedFilesExists = (
    await Promise.all(
      generatedFiles.map((generatedFile) =>
        exists(`${folder}/${generatedFile}`)
      )
    )
  ).every((exist) => exist);

  for (const fileToCache of filesToCache) {
    const fileHash = (await hashElement(`${folder}/${fileToCache}`)).hash;

    cache.hash = `${cache.hash}-${fileHash}`;
  }

  // We only skip if both the cache and the generated file exists
  if (generatedFilesExists && (await exists(cacheFile))) {
    const storedHash = (await fsp.readFile(cacheFile)).toString();
    if (storedHash === cache.hash) {
      return {
        cacheExists: true,
        hash: cache.hash,
      };
    }
  }

  return cache;
}

export async function buildCustomGenerators(verbose: boolean): Promise<void> {
  const spinner = createSpinner('building custom generators', verbose).start();

  const cacheFile = toAbsolutePath('generators/.cache');
  const { cacheExists, hash } = await checkForCache({
    folder: toAbsolutePath('generators/'),
    generatedFiles: ['build'],
    filesToCache: ['src', 'build.gradle', 'settings.gradle'],
    cacheFile,
  });

  if (cacheExists) {
    spinner.succeed('job skipped, cache found for custom generators');
    return;
  }

  await run('./gradle/gradlew --no-daemon -p generators assemble', {
    verbose,
  });

  if (hash) {
    spinner.text = 'storing custom generators cache';
    await fsp.writeFile(cacheFile, hash);
  }

  spinner.succeed();
}

export async function gitBranchExists(branchName: string): Promise<boolean> {
  return Boolean(await run(`git ls-remote --heads origin ${branchName}`));
}

export async function emptyDirExceptForDotGit(dir: string): Promise<void> {
  for (const file of await fsp.readdir(dir)) {
    if (file !== '.git') {
      await remove(path.resolve(dir, file));
    }
  }
}

export async function runComposerUpdate(verbose: boolean): Promise<void> {
  if (!CI) {
    await run(
      'composer update --working-dir=clients/algoliasearch-client-php && composer dump-autoload --working-dir=clients/algoliasearch-client-php',
      {
        verbose,
      }
    );
  }
}
