// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { InsightsClient } from '@experimental-api-clients-automation/client-insights';
import { insightsClient } from '@experimental-api-clients-automation/client-insights';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): InsightsClient {
  return insightsClient(appId, apiKey, 'us', { requester: echoRequester() });
}

describe('api', () => {
  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result = await $client.pushEvents({ events: [] });

    expect(result.algoliaAgent).toMatch(
      /Algolia%20for%20(.+)%20\(\d+\.\d+\.\d+\)/
    );
  });

  test('calls api with correct timeouts', async () => {
    const $client = createClient();

    const result = await $client.pushEvents({ events: [] });

    expect(result).toEqual(
      expect.objectContaining({ connectTimeout: 2000, responseTimeout: 30000 })
    );
  });
});

describe('parameters', () => {
  test('fallbacks to the alias when region is not given', async () => {
    const $client = insightsClient('my-app-id', 'my-api-key', '', {
      requester: echoRequester(),
    });

    const result = await $client.pushEvents({ events: [] });

    expect(result).toEqual(
      expect.objectContaining({ host: 'insights.algolia.io' })
    );
  });
});
