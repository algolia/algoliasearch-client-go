/* eslint-disable no-console */
import { CLIENTS, GENERATORS, LANGUAGES } from '../../common';
import { getLanguageFolder, getTestOutputFolder } from '../../config';
import type { Language } from '../../types';
import { getNbGitDiff } from '../utils';

import { DEPENDENCIES } from './setRunVariables';
import type { ClientMatrix, CreateMatrix, Matrix, SpecMatrix } from './types';
import { computeCacheKey, isBaseChanged } from './utils';

// This empty matrix is required by the CI, otherwise it throws
const EMPTY_MATRIX = { client: ['no-run'] };

/**
 * List of dependencies based on the language, inherited from `./setRunVariables.ts` in a more dynamic form.
 */
const MATRIX_DEPENDENCIES = {
  common: {
    GITHUB_ACTIONS_CHANGED: DEPENDENCIES.GITHUB_ACTIONS_CHANGED,
    SCRIPTS_CHANGED: DEPENDENCIES.SCRIPTS_CHANGED,
    COMMON_SPECS_CHANGED: DEPENDENCIES.COMMON_SPECS_CHANGED,
  },
  clients: {
    common: {
      GENERATORS_CHANGED: DEPENDENCIES.GENERATORS_CHANGED,
    },
    javascript: {
      JS_UTILS_CHANGED: DEPENDENCIES.JS_UTILS_CHANGED,
      JS_TEMPLATE_CHANGED: DEPENDENCIES.JS_TEMPLATE_CHANGED,
    },
    php: {
      PHP_TEMPLATE_CHANGED: DEPENDENCIES.PHP_TEMPLATE_CHANGED,
    },
    java: {
      JAVA_TEMPLATE_CHANGED: DEPENDENCIES.JAVA_TEMPLATE_CHANGED,
    },
  },
};

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
      ...MATRIX_DEPENDENCIES.common,
      ...MATRIX_DEPENDENCIES.clients.common,
      ...MATRIX_DEPENDENCIES.clients[language],
    });

    // No changes found, we don't put this job in the matrix
    if (clientChanges === 0 && specChanges === 0 && !baseChanged) {
      continue;
    }

    console.log(`::set-output name=RUN_GEN_${language.toUpperCase()}::true`);
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
      testsOutputPath: `./tests/output/${language}/${getTestOutputFolder(
        language
      )}`,
    });
  }

  const shouldRun = clientMatrix.client.length > 0;

  console.log(`::set-output name=RUN_GEN::${shouldRun}`);
  console.log(
    `::set-output name=GEN_MATRIX::${JSON.stringify(
      shouldRun ? clientMatrix : EMPTY_MATRIX
    )}`
  );
}

async function getSpecMatrix(baseBranch: string): Promise<void> {
  const matrix: ToRunMatrix = {
    path: 'specs/bundled',
    toRun: [],
    cacheToCompute: [],
  };

  for (const client of CLIENTS) {
    // The `algoliasearch-lite` spec is created by the `search` spec
    const bundledSpecName = client === 'algoliasearch-lite' ? 'search' : client;
    const path = `specs/${bundledSpecName}`;
    const specChanges = await getNbGitDiff({
      branch: baseBranch,
      path,
    });
    const baseChanged = await isBaseChanged(
      baseBranch,
      MATRIX_DEPENDENCIES.common
    );

    // No changes found, we don't put this job in the matrix
    if (specChanges === 0 && !baseChanged) {
      continue;
    }

    matrix.toRun.push(client);
    matrix.cacheToCompute.push(path);
  }

  // We have nothing to run
  if (matrix.toRun.length === 0) {
    console.log('::set-output name=RUN_SPECS::false');
    console.log(`::set-output name=MATRIX::${JSON.stringify(EMPTY_MATRIX)}`);

    return;
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

  console.log('::set-output name=RUN_SPECS::true');
  console.log(`::set-output name=MATRIX::${JSON.stringify(ciMatrix)}`);
}

/**
 * Creates a matrix for the CI jobs based on the files that changed.
 */
async function createMatrix(opts: CreateMatrix): Promise<void> {
  if (opts.forClients) {
    return await getClientMatrix(opts.baseBranch);
  }

  return await getSpecMatrix(opts.baseBranch);
}

if (require.main === module) {
  const args = process.argv.slice(2);

  createMatrix({
    baseBranch: args[0],
    forClients: args[1] === 'clients',
  });
}
