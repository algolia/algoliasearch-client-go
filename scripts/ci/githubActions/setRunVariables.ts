/* eslint-disable no-console */
import { getLanguageFolder } from '../../config';

import { isBaseChanged } from './utils';

const JS_CLIENT_FOLDER = getLanguageFolder('javascript');
const JAVA_CLIENT_FOLDER = getLanguageFolder('java');
const PHP_CLIENT_FOLDER = getLanguageFolder('php');

// Files that are common to every clients
const CLIENTS_COMMON_FILES = [
  'config/openapitools.json',
  'config/clients.config.json',
  'generators/src/main/java/com/algolia/codegen/Utils.java',
];

/**
 * Dependencies that are common to every specs, clients or CTS jobs.
 */
export const COMMON_DEPENDENCIES = {
  GITHUB_ACTIONS_CHANGED: [
    '.github/actions',
    '.github/workflows',
    '.github/.cache_version',
  ],
  SCRIPTS_CHANGED: ['scripts', 'eslint', 'yarn.lock'],
  COMMON_SPECS_CHANGED: ['specs/common'],
};

/**
 * Exhaustive list of output variables to use in the CI.
 *
 * Those variables are used to determine if jobs should run, based on the changes
 * made in their respective dependencies.
 *
 * Negative paths should start with `:!`.
 *
 * The variable will be accessible in the CI via `steps.diff.outputs.<name>`.
 *
 * Variables starting by `LANGUAGENAME_` will be used in the `createMatrix` to determine
 * if a job should be added.
 */
export const DEPENDENCIES = {
  ...COMMON_DEPENDENCIES,
  JS_ALGOLIASEARCH_CHANGED: [
    `${JS_CLIENT_FOLDER}/packages/algoliasearch`,
    `${JS_CLIENT_FOLDER}/packages/client-search`,
    `${JS_CLIENT_FOLDER}/packages/client-analytics`,
    `${JS_CLIENT_FOLDER}/packages/client-personalization`,
  ],
  JS_COMMON_TESTS_CHANGED: [
    `${JS_CLIENT_FOLDER}/packages/client-common/src/__tests__`,
  ],
  JAVASCRIPT_UTILS_CHANGED: [
    `${JS_CLIENT_FOLDER}/packages/client-common`,
    `${JS_CLIENT_FOLDER}/packages/requester-browser-xhr`,
    `${JS_CLIENT_FOLDER}/packages/requester-node-http`,
  ],
  JAVASCRIPT_CLIENT_CHANGED: [
    ...CLIENTS_COMMON_FILES,
    JS_CLIENT_FOLDER,
    'templates/javascript',
    'generators/src/main/java/com/algolia/codegen/AlgoliaJavaScriptGenerator.java',
    `:!${JS_CLIENT_FOLDER}/.github`,
    `:!${JS_CLIENT_FOLDER}/README.md`,
    'tests/CTS/methods/requests/templates/javascript',
  ],
  JAVA_CLIENT_CHANGED: [
    ...CLIENTS_COMMON_FILES,
    JAVA_CLIENT_FOLDER,
    'templates/java',
    'generators/src/main/java/com/algolia/codegen/AlgoliaJavaGenerator.java',
    'tests/CTS/methods/requests/templates/java',
  ],
  PHP_CLIENT_CHANGED: [
    ...CLIENTS_COMMON_FILES,
    PHP_CLIENT_FOLDER,
    'templates/php',
    'generators/src/main/java/com/algolia/codegen/AlgoliaPhpGenerator.java',
    'tests/CTS/methods/requests/templates/php',
  ],
};

/**
 * Outputs variables used in the CI to determine if a job should run.
 */
async function setRunVariables({
  originBranch,
}: {
  originBranch: string;
}): Promise<void> {
  console.log(`Checking diff between ${originBranch} and HEAD`);

  console.log(`::set-output name=ORIGIN_BRANCH::${originBranch}`);

  await isBaseChanged(originBranch, DEPENDENCIES, true);
}

if (require.main === module) {
  const [originBranch] = process.argv.slice(2);

  if (!originBranch) {
    throw new Error(
      `Unable to retrieve the origin branch: ${JSON.stringify(originBranch)}`
    );
  }

  setRunVariables({ originBranch });
}
