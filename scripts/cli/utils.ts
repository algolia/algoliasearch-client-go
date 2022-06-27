import inquirer from 'inquirer';

import { CLIENTS, GENERATORS, LANGUAGES } from '../common';
import type { Generator, Language } from '../types';

export const ALL = 'all';
export const PROMPT_LANGUAGES = [ALL, ...LANGUAGES];
export const PROMPT_CLIENTS = [ALL, ...CLIENTS];

export type AllLanguage = Language | typeof ALL;
export type LangArg = AllLanguage | undefined;

export type PromptDecision = {
  language: AllLanguage;
  client: string[];
  clientList: string[];
};

export type Job = 'build' | 'generate' | 'specs';

type Prompt = {
  langArg: LangArg;
  clientArg: string[];
  job: Job;
  interactive: boolean;
};

export function getClientChoices(
  job: Job,
  language?: LangArg,
  clientList = PROMPT_CLIENTS
): string[] {
  const withoutAlgoliaSearch = clientList.filter(
    (client) => client !== 'algoliasearch'
  );

  if (!language) {
    return job === 'specs' ? withoutAlgoliaSearch : clientList;
  }

  const isJavaScript = language === ALL || language === 'javascript';

  switch (job) {
    // We don't need to build `lite` client as it's a subset of the `algoliasearch` one
    case 'build':
      // Only `JavaScript` provide a lite client, others can build anything but it.
      if (isJavaScript) {
        return clientList.filter((client) => client !== 'lite');
      }

      return withoutAlgoliaSearch.filter((client) => client !== 'lite');
    // `algoliasearch` is not built from specs, it's an aggregation of clients
    case 'specs':
      return withoutAlgoliaSearch;
    case 'generate':
      // Only `JavaScript` provide a lite client, others can build anything but it.
      if (isJavaScript) {
        return withoutAlgoliaSearch;
      }

      return withoutAlgoliaSearch.filter((client) => client !== 'lite');
    default:
      return clientList;
  }
}

export function generatorList({
  language,
  client,
  clientList,
}: {
  language: AllLanguage;
  client: string[];
  clientList: string[];
}): Generator[] {
  const langsTodo = language === ALL ? LANGUAGES : [language];
  const clientsTodo = client[0] === ALL ? clientList : client;

  return langsTodo
    .flatMap((lang) => clientsTodo.map((cli) => GENERATORS[`${lang}-${cli}`]))
    .filter(Boolean);
}

export async function prompt({
  langArg,
  clientArg,
  job,
  interactive,
}: Prompt): Promise<PromptDecision> {
  const decision: PromptDecision = {
    client: [ALL],
    language: ALL,
    clientList: [],
  };

  if (!langArg) {
    if (interactive) {
      const { language } = await inquirer.prompt<PromptDecision>([
        {
          type: 'list',
          name: 'language',
          message: 'Select a language',
          default: ALL,
          choices: LANGUAGES,
        },
      ]);

      decision.language = language;
    }
  } else {
    decision.language = langArg;
  }

  decision.clientList = getClientChoices(job, decision.language, CLIENTS);

  if (!clientArg || !clientArg.length) {
    if (interactive) {
      const { client } = await inquirer.prompt<{ client: string }>([
        {
          type: 'list',
          name: 'client',
          message: 'Select a client',
          default: ALL,
          choices: getClientChoices(job, decision.language),
        },
      ]);

      decision.client = [client];
    }
  } else {
    clientArg.forEach((client) => {
      if (!PROMPT_CLIENTS.includes(client)) {
        throw new Error(
          `The '${clientArg}' client can't run with the given job: '${job}'.\n\nAllowed choices are: ${decision.clientList.join(
            ', '
          )}`
        );
      }
    });

    decision.client = clientArg;
  }

  return decision;
}
