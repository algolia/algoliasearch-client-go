import type {
  AnalyticsClient,
  Region as AnalyticsRegion,
} from '@experimental-api-clients-automation/client-analytics/src/analyticsClient';
import { createAnalyticsClient } from '@experimental-api-clients-automation/client-analytics/src/analyticsClient';
import type {
  CreateClientOptions,
  Host,
  Requester,
} from '@experimental-api-clients-automation/client-common';
import {
  createMemoryCache,
  createFallbackableCache,
  createBrowserLocalStorageCache,
} from '@experimental-api-clients-automation/client-common';
import type {
  PersonalizationClient,
  Region as PersonalizationRegion,
} from '@experimental-api-clients-automation/client-personalization/src/personalizationClient';
import { createPersonalizationClient } from '@experimental-api-clients-automation/client-personalization/src/personalizationClient';
import {
  createSearchClient,
  apiClientVersion,
} from '@experimental-api-clients-automation/client-search/src/searchClient';
import { createXhrRequester } from '@experimental-api-clients-automation/requester-browser-xhr';

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export function algoliasearch(
  appId: string,
  apiKey: string,
  options?: { requester?: Requester; hosts?: Host[] }
) {
  if (!appId) {
    throw new Error('`appId` is missing.');
  }

  if (!apiKey) {
    throw new Error('`apiKey` is missing.');
  }

  const commonOptions: Omit<CreateClientOptions, 'apiKey' | 'appId'> = {
    timeouts: {
      connect: 1,
      read: 2,
      write: 30,
    },
    requester: options?.requester ?? createXhrRequester(),
    userAgents: [{ segment: 'Browser' }],
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
  };

  function initAnalytics(
    analyticsAppId: string,
    analyticsApiKey: string,
    region?: AnalyticsRegion,
    analyticsOptions?: { requester?: Requester; hosts?: Host[] }
  ): AnalyticsClient {
    return createAnalyticsClient({
      appId: analyticsAppId,
      apiKey: analyticsApiKey,
      region,
      ...analyticsOptions,
      ...commonOptions,
    });
  }

  function initPersonalization(
    personalizationAppId: string,
    personalizationApiKey: string,
    region: PersonalizationRegion,
    personalizationOptions?: { requester?: Requester; hosts?: Host[] }
  ): PersonalizationClient {
    if (!region) {
      throw new Error('`region` is missing.');
    }

    return createPersonalizationClient({
      appId: personalizationAppId,
      apiKey: personalizationApiKey,
      region,
      ...personalizationOptions,
      ...commonOptions,
    });
  }

  return {
    ...createSearchClient({ appId, apiKey, ...commonOptions }),
    initAnalytics,
    initPersonalization,
  };
}
