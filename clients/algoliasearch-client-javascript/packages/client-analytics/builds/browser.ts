import type { InitClientOptions } from '@experimental-api-clients-automation/client-common';
import {
  createMemoryCache,
  createFallbackableCache,
  createBrowserLocalStorageCache,
} from '@experimental-api-clients-automation/client-common';
import { createXhrRequester } from '@experimental-api-clients-automation/requester-browser-xhr';

import {
  createAnalyticsClient,
  apiClientVersion,
} from '../src/analyticsClient';
import type { AnalyticsClient, Region } from '../src/analyticsClient';

export * from '../src/analyticsClient';

export function analyticsClient(
  appId: string,
  apiKey: string,
  region?: Region,
  options?: InitClientOptions
): AnalyticsClient {
  if (!appId) {
    throw new Error('`appId` is missing.');
  }

  if (!apiKey) {
    throw new Error('`apiKey` is missing.');
  }

  return createAnalyticsClient({
    appId,
    apiKey,
    region,
    timeouts: {
      connect: 1,
      read: 2,
      write: 30,
    },
    requester: options?.requester ?? createXhrRequester(),
    userAgents: [{ segment: 'Browser' }],
    authMode: 'WithinQueryParameters',
    responsesCache: options?.responsesCache ?? createMemoryCache(),
    requestsCache:
      options?.requestsCache ?? createMemoryCache({ serializable: false }),
    hostsCache:
      options?.hostsCache ??
      createFallbackableCache({
        caches: [
          createBrowserLocalStorageCache({
            key: `${apiClientVersion}-${appId}`,
          }),
          createMemoryCache(),
        ],
      }),
    ...options,
  });
}
