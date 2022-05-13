import clientsConfig from '../config/clients.config.json';
import openapiConfig from '../config/openapitools.json';

import type { Language } from './types';

export function getLanguageFolder(language: Language): string {
  return clientsConfig[language].folder;
}

export function getLanguageApiFolder(language: Language): string {
  return clientsConfig[language].apiFolder;
}

export function getLanguageModelFolder(language: Language): string {
  return clientsConfig[language].modelFolder;
}

export function getTestExtension(language: Language): string {
  return clientsConfig[language].tests.extension;
}

export function getTestOutputFolder(language: Language): string {
  return clientsConfig[language].tests.outputFolder;
}

export function getCustomGenerator(language: Language): string {
  return clientsConfig[language].customGenerator;
}

// Returns the version of the package from clients.config.json, except for JavaScript where it returns the version of javascript-search
export function getPackageVersionDefault(language: Language): string {
  if (language === 'javascript') {
    return openapiConfig['generator-cli'].generators['javascript-search']
      .additionalProperties.packageVersion;
  }
  return clientsConfig[language].packageVersion;
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
