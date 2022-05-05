import clientsConfig from '../config/clients.config.json';

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
