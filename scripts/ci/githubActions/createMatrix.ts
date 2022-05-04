/* eslint-disable no-console */
import { CLIENTS, GENERATORS, LANGUAGES } from '../../common';
import { getLanguageApiFolder, getLanguageModelFolder } from '../../config';
import { camelize, createClientName } from '../../cts/utils';
import type { Language } from '../../types';
import { getNbGitDiff } from '../utils';

import { DEPENDENCIES } from './setRunVariables';
import type { CreateMatrix, ClientMatrix, Matrix, SpecMatrix } from './types';
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

async function getClientMatrix(baseBranch: string): Promise<void> {
  const matrix: Record<Language, Matrix<ClientMatrix>> = {
    java: { client: [] },
    php: { client: [] },
    javascript: { client: [] },
  };

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

    const clientName = createClientName(client, language);
    const camelizedClientName = camelize(client);
    const pathToApi = `${output}/${getLanguageApiFolder(language)}`;
    const pathToModel = `${output}/${getLanguageModelFolder(language)}`;

    const clientMatrix: ClientMatrix = {
      language,
      name: client,
      path: output,

      configName: `${clientName}Config`,
      apiName: `${clientName}Client`,

      apiPath: pathToApi,
      modelPath: pathToModel,

      cacheKey: await computeCacheKey(`client-${client}`, [
        `specs/bundled/${bundledSpec}.yml`,
        `templates/${language}`,
        `generators/src`,
      ]),
    };

    // While JavaScript have it's own package per client, other language have
    // API and models in folders common to all clients, so we scope it.
    switch (language) {
      case 'java':
        clientMatrix.apiPath = `${pathToApi}/${clientMatrix.apiName}.java`;
        clientMatrix.modelPath = `${pathToModel}/${camelizedClientName}/`;
        break;
      case 'php':
        clientMatrix.apiPath = `${pathToApi}/${clientMatrix.apiName}.php`;
        clientMatrix.modelPath = `${pathToModel}/${clientName}/`;
        break;
      default:
        break;
    }

    matrix[language].client.push(clientMatrix);
  }

  // Set output variables for the CI
  for (const language of LANGUAGES) {
    const lang = language.toLocaleUpperCase();
    const shouldRun = matrix[language].client.length > 0;

    console.log(`::set-output name=RUN_${lang}::${shouldRun}`);
    console.log(
      `::set-output name=${lang}_MATRIX::${JSON.stringify(
        shouldRun ? matrix[language] : EMPTY_MATRIX
      )}`
    );
  }
}

async function getSpecMatrix(baseBranch: string): Promise<void> {
  const matrix: Matrix<SpecMatrix> = { client: [] };

  for (const client of CLIENTS) {
    // The `algoliasearch-lite` spec is created by the `search` spec
    const bundledSpecName = client === 'algoliasearch-lite' ? 'search' : client;
    const specChanges = await getNbGitDiff({
      branch: baseBranch,
      path: `specs/${bundledSpecName}`,
    });
    const baseChanged = await isBaseChanged(
      baseBranch,
      MATRIX_DEPENDENCIES.common
    );

    // No changes found, we don't put this job in the matrix
    if (specChanges === 0 && !baseChanged) {
      continue;
    }

    const path = `specs/${bundledSpecName}`;

    matrix.client.push({
      name: client,
      path,
      bundledPath: `specs/bundled/${client}.yml`,
      cacheKey: await computeCacheKey(`spec-${client}`, ['specs/common', path]),
    });
  }

  const shouldRun = matrix.client.length > 0;

  console.log(`::set-output name=RUN_SPECS::${shouldRun}`);
  console.log(
    `::set-output name=MATRIX::${JSON.stringify(
      shouldRun ? matrix : EMPTY_MATRIX
    )}`
  );
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
