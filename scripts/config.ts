import clientsConfig from '../config/clients.config.json';
import openapiConfig from '../config/openapitools.json';

import type { Language, LanguageConfig } from './types';

export function getClientsConfigField(
  language: Language,
  pathToField: string[] | string
): any {
  const config: LanguageConfig = clientsConfig[language];
  const path = Array.isArray(pathToField) ? pathToField : [pathToField];

  return path.reduce((current, key) => {
    if (!current || !current[key]) {
      throw new Error(`Unable to find '${pathToField}' for '${language}'`);
    }

    return current[key];
  }, config);
}

export function getLanguageFolder(language: Language): string {
  return getClientsConfigField(language, 'folder');
}

export function getLanguageApiFolder(language: Language): string {
  return getClientsConfigField(language, 'apiFolder');
}

export function getLanguageModelFolder(language: Language): string {
  return getClientsConfigField(language, 'modelFolder');
}

export function getTestExtension(language: Language): string {
  return getClientsConfigField(language, ['tests', 'extension']);
}

export function getTestOutputFolder(language: Language): string {
  return getClientsConfigField(language, ['tests', 'outputFolder']);
}

export function getCustomGenerator(language: Language): string {
  return getClientsConfigField(language, 'customGenerator');
}

/**
 * Returns the version of the package from clients.config.json, except for JavaScript where it returns the version of javascript-search.
 */
export function getPackageVersionDefault(language: Language): string {
  if (language === 'javascript') {
    return openapiConfig['generator-cli'].generators['javascript-search']
      .additionalProperties.packageVersion;
  }

  return getClientsConfigField(language, 'packageVersion');
}

export function getGitHubUrl(
  language: Language,
  options?: { token: string }
): string {
  const { gitRepoId } = clientsConfig[language];

  // GitHub Action provides a default token for authentication
  // https://docs.github.com/en/actions/security-guides/automatic-token-authentication
  // But it has access to only the self repository.
  // If we want to do something like pushing commits to other repositories,
  // we need to specify a token with more access.
  return options?.token
    ? `https://${options.token}:${options.token}@github.com/algolia/${gitRepoId}`
    : `https://github.com/algolia/${gitRepoId}`;
}
