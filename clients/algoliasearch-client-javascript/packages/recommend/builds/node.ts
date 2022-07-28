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

import { createRecommendClient } from '../src/recommendClient';
import type { RecommendClient } from '../src/recommendClient';

export { apiClientVersion, RecommendClient } from '../src/recommendClient';
export * from '../model';

export function recommendClient(
  appId: string,
  apiKey: string,
  options?: ClientOptions
): RecommendClient {
  if (!appId || typeof appId !== 'string') {
    throw new Error('`appId` is missing.');
  }

  if (!apiKey || typeof apiKey !== 'string') {
    throw new Error('`apiKey` is missing.');
  }

  return createRecommendClient({
    appId,
    apiKey,
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
