import { buildCustomGenerators, run, toAbsolutePath } from '../common';
import { getTestOutputFolder } from '../config';
import { formatter } from '../formatter';
import { createSpinner } from '../oraLog';
import { generateOpenapitools } from '../pre-gen';
import type { Generator } from '../types';

async function ctsGenerate(gen: Generator, verbose: boolean): Promise<void> {
  const spinner = createSpinner(
    `generating CTS for ${gen.key}`,
    verbose
  ).start();
  await run(
    `yarn openapi-generator-cli --custom-generator=generators/build/libs/algolia-java-openapi-generator-1.0.0.jar generate \
     -g algolia-cts -i specs/bundled/${gen.client}.yml --additional-properties="language=${gen.language},client=${gen.client}"`,
    { verbose }
  );
  spinner.succeed();
}

export async function ctsGenerateMany(
  generators: Generator[],
  verbose: boolean
): Promise<void> {
  await buildCustomGenerators(verbose);
  await generateOpenapitools(generators);

  for (const gen of generators) {
    if (!getTestOutputFolder(gen.language)) {
      continue;
    }
    await ctsGenerate(gen, verbose);
  }

  const langs = [...new Set(generators.map((gen) => gen.language))];
  for (const lang of langs) {
    if (!getTestOutputFolder(lang)) {
      continue;
    }
    await formatter(lang, toAbsolutePath(`tests/output/${lang}`), verbose);
  }
}
