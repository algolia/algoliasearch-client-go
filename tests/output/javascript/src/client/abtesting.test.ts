// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { AbtestingClient } from '@experimental-api-clients-automation/client-abtesting';
import { abtestingClient } from '@experimental-api-clients-automation/client-abtesting';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): AbtestingClient {
  return abtestingClient(appId, apiKey, 'us', { requester: echoRequester() });
}

describe('api', () => {
  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result = await $client.addABTests({
      name: 'test',
      variant: [{ index: 'my-test-index', trafficPercentage: 90 }],
      endAt: '2022-02-01T13:37:01Z',
    });

    expect(result.algoliaAgent).toMatch(
      /Algolia%20for%20(.+)%20\(\d+\.\d+\.\d+\)/
    );
  });

  test('calls api with correct timeouts', async () => {
    const $client = createClient();

    const result = await $client.addABTests({
      name: 'test',
      variant: [{ index: 'my-test-index', trafficPercentage: 90 }],
      endAt: '2022-02-01T13:37:01Z',
    });

    expect(result).toEqual(
      expect.objectContaining({ connectTimeout: 2000, responseTimeout: 30000 })
    );
  });
});

describe('parameters', () => {
  test('fallbacks to the alias when region is not given', async () => {
    const $client = abtestingClient('my-app-id', 'my-api-key', '', {
      requester: echoRequester(),
    });

    const result = await $client.getABTest({ id: 'test' });

    expect(result).toEqual(
      expect.objectContaining({ host: 'analytics.algolia.com' })
    );
  });
});
