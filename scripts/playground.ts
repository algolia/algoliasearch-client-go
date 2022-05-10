import { run } from './common';
import { capitalize } from './cts/utils';
import type { Language } from './types';

export async function playground({
  language,
  client,
}: {
  language: Language | 'all';
  client: string;
}): Promise<void> {
  const verbose = true;
  switch (language) {
    case 'javascript':
      await run(`yarn workspace javascript-playground start:${client}`, {
        verbose,
      });
      break;
    case 'java':
      await run(
        `./gradle/gradlew -p playground/java -PmainClass=com.algolia.playground.${capitalize(
          client
        )} run`,
        {
          verbose,
        }
      );
      break;
    case 'php':
      await run(
        `cd playground/php && \
       composer update && \
       composer dump-autoload && \
       cd src && \
       php8 ${client}.php`,
        { verbose }
      );
      break;
    default:
  }
}
