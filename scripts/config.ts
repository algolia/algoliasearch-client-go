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
