/* eslint-disable no-console */
import { getLanguageFolder } from '../../config';

import { isBaseChanged } from './utils';

const JS_CLIENT_FOLDER = getLanguageFolder('javascript');
const JAVA_CLIENT_FOLDER = getLanguageFolder('java');
const PHP_CLIENT_FOLDER = getLanguageFolder('php');

// Files that are common to every clients
const CLIENTS_COMMON_FILES = [
  'openapitools.json',
  'config/clients.config.json',
];

/**
 * Exhaustive list of output variables to use in the CI.
 *
 * Those variables are used to determine if jobs should run, based on the changes
 * made in their respective `path`s.
 *
 * Negative paths should start with `:!`.
 *
 * The variable will be accessible in the CI via `steps.diff.outputs.<name>`.
 */
export const DEPENDENCIES = {
  GITHUB_ACTIONS_CHANGED: [
    '.github/actions',
    '.github/workflows',
    '.github/.cache_version',
  ],
  SPECS_CHANGED: ['specs', ':!specs/bundled'],
  COMMON_SPECS_CHANGED: ['specs/common'],
  TESTS_CHANGED: ['tests'],
  SCRIPTS_CHANGED: ['scripts'],
  GENERATORS_CHANGED: ['generators'],
  JS_CLIENT_CHANGED: [
    ...CLIENTS_COMMON_FILES,
    JS_CLIENT_FOLDER,
    `:!${JS_CLIENT_FOLDER}/.github`,
    `:!${JS_CLIENT_FOLDER}/README.md`,
  ],
  JS_ALGOLIASEARCH_CHANGED: [
    `${JS_CLIENT_FOLDER}/packages/algoliasearch`,
    `${JS_CLIENT_FOLDER}/packages/client-search`,
    `${JS_CLIENT_FOLDER}/packages/client-analytics`,
    `${JS_CLIENT_FOLDER}/packages/client-personalization`,
  ],
  JS_UTILS_CHANGED: [
    `${JS_CLIENT_FOLDER}/packages/client-common`,
    `${JS_CLIENT_FOLDER}/packages/requester-browser-xhr`,
    `${JS_CLIENT_FOLDER}/packages/requester-node-http`,
  ],
  JS_COMMON_TESTS_CHANGED: [
    `${JS_CLIENT_FOLDER}/packages/client-common/src/__tests__`,
  ],
  JS_TEMPLATE_CHANGED: ['templates/javascript'],
  JAVA_CLIENT_CHANGED: [...CLIENTS_COMMON_FILES, JAVA_CLIENT_FOLDER],
  JAVA_TEMPLATE_CHANGED: ['templates/java'],
  PHP_CLIENT_CHANGED: [...CLIENTS_COMMON_FILES, PHP_CLIENT_FOLDER],
  PHP_TEMPLATE_CHANGED: ['templates/php'],
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
