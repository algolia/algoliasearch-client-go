// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { RecommendClient } from '@experimental-api-clients-automation/recommend';
import { recommendClient } from '@experimental-api-clients-automation/recommend';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): RecommendClient {
  return recommendClient(appId, apiKey, { requester: echoRequester() });
}

describe('api', () => {
  test('calls api with correct host', async () => {
    const $client = createClient();

    const result = await $client.getRecommendations({ requests: [] });

    expect(result).toEqual(
      expect.objectContaining({ host: 'test-app-id-dsn.algolia.net' })
    );
  });

  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result = await $client.getRecommendations({ requests: [] });

    expect(result.algoliaAgent).toMatch(
      /Algolia%20for%20(.+)%20\(\d+\.\d+\.\d+\)/
    );
  });

  test('calls api with correct timeouts', async () => {
    const $client = createClient();

    const result = await $client.getRecommendations({ requests: [] });

    expect(result).toEqual(
      expect.objectContaining({ connectTimeout: 2000, responseTimeout: 5000 })
    );
  });
});
