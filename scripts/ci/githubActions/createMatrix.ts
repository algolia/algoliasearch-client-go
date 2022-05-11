/* eslint-disable no-console */
import { CLIENTS, GENERATORS, LANGUAGES } from '../../common';
import { getLanguageFolder, getTestOutputFolder } from '../../config';
import type { Language } from '../../types';
import { getNbGitDiff } from '../utils';

import { DEPENDENCIES, COMMON_DEPENDENCIES } from './setRunVariables';
import type { ClientMatrix, CreateMatrix, Matrix, SpecMatrix } from './types';
import { computeCacheKey, isBaseChanged } from './utils';

// This empty matrix is required by the CI, otherwise it throws
const EMPTY_MATRIX = { client: ['no-run'] };

type ToRunMatrix = {
  path: string;
  toRun: string[];
  cacheToCompute: string[];
};

async function getClientMatrix(baseBranch: string): Promise<void> {
  const matrix = LANGUAGES.reduce(
    (curr, lang) => ({
      ...curr,
      [lang]: {
        path: getLanguageFolder(lang),
        toRun: [],
        cacheToCompute: [],
      },
    }),
    {} as Record<Language, ToRunMatrix>
  );

  for (const { language, client, output } of Object.values(GENERATORS)) {
    // `algoliasearch` is an aggregation of generated clients.
    if (client === 'algoliasearch') {
      continue;
    }

    const languageDependencies = Object.entries(DEPENDENCIES).reduce(
      (finalDeps, [key, deps]) => {
        if (key.startsWith(`${language.toUpperCase()}_`)) {
          return {
            ...finalDeps,
            [key]: deps,
          };
        }

        return finalDeps;
      },
      {} as Record<string, string[]>
    );

    const bundledSpec = client === 'algoliasearch-lite' ? 'search' : client;
    const specChanges = await getNbGitDiff({
      branch: baseBranch,
      path: `specs/${bundledSpec}`,
    });
    const clientChanges = await getNbGitDiff({
      branch: baseBranch,
      path: output,
    });
    const baseChanged = await isBaseChanged(baseBranch, {
      ...COMMON_DEPENDENCIES,
      ...languageDependencies,
    });

    // No changes found, we don't put this job in the matrix
    if (clientChanges === 0 && specChanges === 0 && !baseChanged) {
      continue;
    }

    matrix[language].toRun.push(client);
    matrix[language].cacheToCompute.push(`specs/${bundledSpec}`);
  }

  const clientMatrix: Matrix<ClientMatrix> = {
    client: [],
  };

  // Set output variables for the CI
  for (const language of LANGUAGES) {
    if (matrix[language].toRun.length === 0) {
      continue;
    }

    const testOutputBase = `./tests/output/${language}/${getTestOutputFolder(
      language
    )}`;

    clientMatrix.client.push({
      language,
      path: matrix[language].path,
      toRun: matrix[language].toRun.join(' '),
      cacheKey: await computeCacheKey(`clients-${language}`, [
        ...matrix[language].cacheToCompute,
        'specs/common',
        `templates/${language}`,
        `generators/src`,
      ]),
      testsOutputPath: `${testOutputBase}/client ${testOutputBase}/methods`,
    });
    console.log(`::set-output name=RUN_GEN_${language.toUpperCase()}::true`);
  }

  const shouldRun = clientMatrix.client.length > 0;

  console.log(`::set-output name=RUN_GEN::${shouldRun}`);
  console.log(
    `::set-output name=GEN_MATRIX::${JSON.stringify(
      shouldRun ? clientMatrix : EMPTY_MATRIX
    )}`
  );
}

async function getSpecMatrix(): Promise<void> {
  const matrix: ToRunMatrix = {
    path: 'specs/bundled',
    toRun: [],
    cacheToCompute: [],
  };

  for (const client of CLIENTS) {
    // The `algoliasearch-lite` spec is created by the `search` spec
    const bundledSpecName = client === 'algoliasearch-lite' ? 'search' : client;

    matrix.toRun.push(client);
    matrix.cacheToCompute.push(`specs/${bundledSpecName}`);
  }

  const ciMatrix: Matrix<SpecMatrix> = {
    client: [
      {
        path: matrix.path,
        bundledPath: 'specs/bundled',
        toRun: matrix.toRun.join(' '),
        cacheKey: await computeCacheKey('specs', [
          ...matrix.cacheToCompute,
          'specs/common',
        ]),
      },
    ],
  };

  console.log(`::set-output name=MATRIX::${JSON.stringify(ciMatrix)}`);
}

/**
 * Creates a matrix for the CI jobs based on the files that changed.
 */
async function createMatrix(opts: CreateMatrix): Promise<void> {
  if (opts.forClients) {
    return await getClientMatrix(opts.baseBranch);
  }

  return await getSpecMatrix();
}

if (require.main === module) {
  const args = process.argv.slice(2);

  createMatrix({
    baseBranch: args[0],
    forClients: args[1] === 'clients',
  });
}
