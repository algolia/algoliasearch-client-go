import path from 'path';

import { createClientName, run, toAbsolutePath } from '../common';
import { getLanguageApiFolder, getLanguageModelFolder } from '../config';
import type { Generator } from '../types';

/**
 * Remove `model` folder for the current language and client.
 */
export async function removeExistingCodegen(
  { language, client, output }: Generator,
  verbose?: boolean
): Promise<void> {
  const baseModelFolder = getLanguageModelFolder(language);
  const baseApiFolder = getLanguageApiFolder(language);
  const clientName = createClientName(client, language);

  let clientModel = '';
  let clientApi = '';

  switch (language) {
    case 'java':
      clientModel = client;
      clientApi = `${clientName}*.java`;
      break;
    case 'php':
      clientModel = clientName;
      clientApi = `${clientName}*.php`;
      break;
    case 'javascript':
      // We want to also delete the nested `lite` client or folders that only exists in JS
      if (clientName === 'algoliasearch') {
        await run(
          `rm -rf ${toAbsolutePath(path.resolve('..', output, 'lite'))}`,
          {
            verbose,
          }
        );
      }

      // Delete `builds` folder
      await run(
        `rm -rf ${toAbsolutePath(path.resolve('..', output, 'builds'))}`,
        {
          verbose,
        }
      );

      break;
    default:
      break;
  }

  // Delete client model folder/file
  await run(
    `rm -rf ${toAbsolutePath(
      path.resolve('..', output, baseModelFolder, clientModel)
    )}`,
    {
      verbose,
    }
  );

  // Delete client api folder/file
  await run(
    `rm -rf ${toAbsolutePath(
      path.resolve('..', output, baseApiFolder, clientApi)
    )}`,
    {
      verbose,
    }
  );
}
