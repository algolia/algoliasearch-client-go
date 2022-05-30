import fsp from 'fs/promises';
import path from 'path';

import { exists, toAbsolutePath } from '../common';
import { getTestExtension, getTestOutputFolder } from '../config';
import type { Language } from '../types';

export async function* walk(
  dir: string
): AsyncGenerator<{ path: string; name: string }> {
  for await (const d of await fsp.opendir(dir)) {
    const entry = path.join(dir, d.name);
    if (d.isDirectory()) yield* walk(entry);
    else if (d.isFile()) yield { path: entry, name: d.name };
  }
}

/**
 * Sets the first letter of the given string in capital.
 *
 * `searchClient` -> `SearchClient`.
 */
export function capitalize(str: string): string {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

/**
 * Splits a string for a given `delimiter` (defaults to `-`) and capitalize each
 * parts except the first letter.
 *
 * `search-client` -> `searchClient`.
 */
export function camelize(str: string, delimiter: string = '-'): string {
  return str
    .split(delimiter)
    .map((part, i) => {
      if (i === 0) {
        return part;
      }

      return capitalize(part);
    })
    .join('');
}

/**
 * Returns the client name with the correct casing for its language.
 *
 * `search-client`, `java` -> `SearchClient`.
 *
 * `search-client`, `javascript` -> `searchClient`.
 */
export function createClientName(client: string, language: string): string {
  if (language === 'javascript') {
    return camelize(client);
  }

  return capitalize(camelize(client));
}

export async function createOutputDir({
  language,
  testPath,
}: {
  language: Language;
  testPath: string;
}): Promise<void> {
  await fsp.mkdir(
    toAbsolutePath(
      `tests/output/${language}/${getTestOutputFolder(language)}/${testPath}`
    ),
    {
      recursive: true,
    }
  );
}

export function getOutputPath({
  language,
  client,
  testPath,
}: {
  language: Language;
  client: string;
  testPath: string;
}): string {
  return toAbsolutePath(
    `tests/output/${language}/${getTestOutputFolder(
      language
    )}/${testPath}/${client}${getTestExtension(language)}`
  );
}

export async function loadTemplates({
  language,
}: {
  language: string;
}): Promise<Record<string, string>> {
  const templates: Record<string, string> = {};
  const templatePath = toAbsolutePath(`templates/${language}/tests/client`);

  if (!(await exists(templatePath))) {
    return {};
  }

  for await (const file of walk(templatePath)) {
    if (!file.name.endsWith('.mustache')) {
      continue;
    }
    const name = file.name.replace('.mustache', '');
    const fileContent = (await fsp.readFile(file.path)).toString();
    templates[name] = fileContent;
  }
  return templates;
}
