import { Argument, program } from 'commander';

import { buildClients } from '../buildClients';
import { buildSpecs } from '../buildSpecs';
import { CI, DOCKER, LANGUAGES } from '../common';
import { ctsGenerateMany } from '../cts/generate';
import { runCts } from '../cts/runCts';
import { formatter } from '../formatter';
import { generate } from '../generate';
import { playground } from '../playground';

import type { Job, LangArg } from './utils';
import {
  ALL,
  getClientChoices,
  generatorList,
  prompt,
  PROMPT_CLIENTS,
  PROMPT_LANGUAGES,
} from './utils';

if (!CI && !DOCKER) {
  // eslint-disable-next-line no-console
  console.log('You should run scripts via the docker container, see README.md');
  // eslint-disable-next-line no-process-exit
  process.exit(1);
}

const args = {
  language: new Argument('[language]', 'The language').choices(
    PROMPT_LANGUAGES
  ),
  clients: (job: Job): Argument =>
    new Argument('[client...]', 'The client').choices(getClientChoices(job)),
  client: new Argument('[client]', 'The client').choices(PROMPT_CLIENTS),
};

const flags = {
  verbose: {
    flag: '-v, --verbose',
    description: 'make the generation verbose',
  },
  interactive: {
    flag: '-i, --interactive',
    description: 'open prompt to query parameters',
  },
  skipCache: {
    flag: '-s, --skip-cache',
    description: 'skip cache checking to force building specs',
  },
  skipUtils: {
    flag: '-su, --skip-utils',
    description: 'skip utils build when building a JavaScript client',
  },
  outputType: {
    flag: '-json, --output-json',
    description: 'outputs the spec in JSON instead of yml',
  },
};

program.name('cli');

program
  .command('generate')
  .description('Generate a specified client')
  .addArgument(args.language)
  .addArgument(args.clients('generate'))
  .option(flags.verbose.flag, flags.verbose.description)
  .option(flags.interactive.flag, flags.interactive.description)
  .action(
    async (langArg: LangArg, clientArg: string[], { verbose, interactive }) => {
      const { language, client, clientList } = await prompt({
        langArg,
        clientArg,
        job: 'generate',
        interactive,
      });

      await generate(
        generatorList({ language, client, clientList }),
        Boolean(verbose)
      );
    }
  );

const buildCommand = program.command('build');

buildCommand
  .command('clients')
  .description('Build a specified client')
  .addArgument(args.language)
  .addArgument(args.clients('build'))
  .option(flags.verbose.flag, flags.verbose.description)
  .option(flags.interactive.flag, flags.interactive.description)
  .option(flags.skipUtils.flag, flags.skipUtils.description)
  .action(
    async (
      langArg: LangArg,
      clientArg: string[],
      { verbose, interactive, skipUtils }
    ) => {
      const { language, client, clientList } = await prompt({
        langArg,
        clientArg,
        job: 'build',
        interactive,
      });

      await buildClients(generatorList({ language, client, clientList }), {
        verbose: Boolean(verbose),
        skipUtils: Boolean(skipUtils),
      });
    }
  );

buildCommand
  .command('specs')
  .description('Build a specified spec')
  .addArgument(args.clients('specs'))
  .option(flags.verbose.flag, flags.verbose.description)
  .option(flags.interactive.flag, flags.interactive.description)
  .option(flags.skipCache.flag, flags.skipCache.description)
  .option(flags.outputType.flag, flags.outputType.description)
  .action(
    async (
      clientArg: string[],
      { verbose, interactive, skipCache, outputJson }
    ) => {
      const { client, clientList } = await prompt({
        langArg: ALL,
        clientArg,
        job: 'specs',
        interactive,
      });

      const outputFormat = outputJson ? 'json' : 'yml';

      // ignore cache when building from cli
      await buildSpecs(
        client[0] === ALL ? clientList : client,
        outputFormat,
        Boolean(verbose),
        !skipCache
      );
    }
  );

const ctsCommand = program.command('cts');

ctsCommand
  .command('generate')
  .description('Generate the CTS tests')
  .addArgument(args.language)
  .addArgument(args.clients('generate'))
  .option(flags.verbose.flag, flags.verbose.description)
  .option(flags.interactive.flag, flags.interactive.description)
  .action(
    async (langArg: LangArg, clientArg: string[], { verbose, interactive }) => {
      const { language, client, clientList } = await prompt({
        langArg,
        clientArg,
        job: 'generate',
        interactive,
      });

      await ctsGenerateMany(
        generatorList({ language, client, clientList }),
        Boolean(verbose)
      );
    }
  );

ctsCommand
  .command('run')
  .description('Run the tests for the CTS')
  .addArgument(args.language)
  .option(flags.verbose.flag, flags.verbose.description)
  .option(flags.interactive.flag, flags.interactive.description)
  .action(async (langArg: LangArg, { verbose, interactive }) => {
    const { language } = await prompt({
      langArg,
      clientArg: [ALL],
      job: 'generate',
      interactive,
    });

    await runCts(language === ALL ? LANGUAGES : [language], Boolean(verbose));
  });

program
  .command('playground')
  .description('Run the playground')
  .addArgument(args.language)
  .addArgument(args.client)
  .option(flags.interactive.flag, flags.interactive.description)
  .action(async (langArg: LangArg, cliClient: string, { interactive }) => {
    const { language, client } = await prompt({
      langArg,
      clientArg: [cliClient],
      job: 'build',
      interactive,
    });

    await playground({
      language,
      client: client[0],
    });
  });

program
  .command('format')
  .description('Format the specified folder for a specific language')
  .addArgument(args.language)
  .argument('folder', 'The folder to format')
  .option(flags.verbose.flag, flags.verbose.description)
  .action(async (language: string, folder: string, { verbose }) => {
    await formatter(language, folder, verbose);
  });

program.parse();
