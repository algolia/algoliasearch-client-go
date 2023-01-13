// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { ClientOptions } from '@algolia/client-common';
import {
  DEFAULT_CONNECT_TIMEOUT_BROWSER,
  DEFAULT_READ_TIMEOUT_BROWSER,
  DEFAULT_WRITE_TIMEOUT_BROWSER,
  createMemoryCache,
  createFallbackableCache,
  createBrowserLocalStorageCache,
} from '@algolia/client-common';
import { createXhrRequester } from '@algolia/requester-browser-xhr';

import type { IngestionClient, Region } from '../src/ingestionClient';
import {
  createIngestionClient,
  apiClientVersion,
  REGIONS,
} from '../src/ingestionClient';

export {
  apiClientVersion,
  IngestionClient,
  Region,
} from '../src/ingestionClient';
export * from '../model';

export function ingestionClient(
  appId: string,
  apiKey: string,
  region: Region,
  options?: ClientOptions
): IngestionClient {
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

  return createIngestionClient({
    appId,
    apiKey,
    region,
    timeouts: {
      connect: DEFAULT_CONNECT_TIMEOUT_BROWSER,
      read: DEFAULT_READ_TIMEOUT_BROWSER,
      write: DEFAULT_WRITE_TIMEOUT_BROWSER,
    },
    requester: createXhrRequester(),
    algoliaAgents: [{ segment: 'Browser' }],
    authMode: 'WithinQueryParameters',
    responsesCache: createMemoryCache(),
    requestsCache: createMemoryCache({ serializable: false }),
    hostsCache: createFallbackableCache({
      caches: [
        createBrowserLocalStorageCache({ key: `${apiClientVersion}-${appId}` }),
        createMemoryCache(),
      ],
    }),
    ...options,
  });
}