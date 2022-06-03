/* eslint-disable @typescript-eslint/no-unused-vars */

// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { AnalyticsClient } from '@experimental-api-clients-automation/client-analytics';
import { analyticsClient } from '@experimental-api-clients-automation/client-analytics';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): AnalyticsClient {
  return analyticsClient(appId, apiKey, 'us', { requester: echoRequester() });
}

describe('api', () => {
  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result0 = await $client.getAverageClickPosition({
      index: 'my-index',
    });

    expect(result0.algoliaAgent).toMatch(
      /Algolia%20for%20(.+)%20\(\d+\.\d+\.\d+\)/
    );
  });

  test('calls api with correct timeouts', async () => {
    const $client = createClient();

    const result0 = await $client.getAverageClickPosition({
      index: 'my-index',
    });

    expect(result0).toEqual(
      expect.objectContaining({ connectTimeout: 2, responseTimeout: 5 })
    );
  });
});

describe('parameters', () => {
  test('fallbacks to the alias when region is not given', async () => {
    const $client = analyticsClient('my-app-id', 'my-api-key', '', {
      requester: echoRequester(),
    });

    const result1 = await $client.getAverageClickPosition({
      index: 'my-index',
    });

    expect(result1).toEqual(
      expect.objectContaining({ host: 'analytics.algolia.com' })
    );
  });

  test('getAverageClickPosition throws without index', async () => {
    const $client = createClient();

    try {
      const result0 = await $client.getClickPositions({});

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect(e.message).toMatch(
        'Parameter `index` is required when calling `getClickPositions`.'
      );
    }
  });
});
