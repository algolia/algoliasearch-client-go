// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { ClientOptions } from '@algolia/client-common';
import {
  DEFAULT_CONNECT_TIMEOUT_NODE,
  DEFAULT_READ_TIMEOUT_NODE,
  DEFAULT_WRITE_TIMEOUT_NODE,
  createMemoryCache,
  createNullCache,
} from '@algolia/client-common';
import { createHttpRequester } from '@algolia/requester-node-http';

import type {
  QuerySuggestionsClient,
  Region,
} from '../src/querySuggestionsClient';
import {
  createQuerySuggestionsClient,
  REGIONS,
} from '../src/querySuggestionsClient';

export {
  apiClientVersion,
  QuerySuggestionsClient,
  Region,
} from '../src/querySuggestionsClient';
export * from '../model';

export function querySuggestionsClient(
  appId: string,
  apiKey: string,
  region: Region,
  options?: ClientOptions
): QuerySuggestionsClient {
  if (!appId || typeof appId !== 'string') {
    throw new Error('`appId` is missing.');
  }

  if (!apiKey || typeof apiKey !== 'string') {
    throw new Error('`apiKey` is missing.');
  }

  if (
    !region ||
    (region && (typeof region !== 'string' || !REGIONS.includes(region)))
  ) {
    throw new Error(
      `\`region\` is required and must be one of the following: ${REGIONS.join(
        ', '
      )}`
    );
  }

  return createQuerySuggestionsClient({
    appId,
    apiKey,
    region,
    timeouts: {
      connect: DEFAULT_CONNECT_TIMEOUT_NODE,
      read: DEFAULT_READ_TIMEOUT_NODE,
      write: DEFAULT_WRITE_TIMEOUT_NODE,
    },
    requester: createHttpRequester(),
    algoliaAgents: [{ segment: 'Node.js', version: process.versions.node }],
    responsesCache: createNullCache(),
    requestsCache: createNullCache(),
    hostsCache: createMemoryCache(),
    ...options,
  });
}
