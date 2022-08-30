/* eslint-disable @typescript-eslint/no-unused-vars, require-await */
// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { AnalyticsClient } from '@algolia/client-analytics';
import { analyticsClient } from '@algolia/client-analytics';
import { echoRequester } from '@algolia/requester-node-http';
import type { EchoResponse } from '@algolia/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): AnalyticsClient {
  return analyticsClient(appId, apiKey, 'us', { requester: echoRequester() });
}

describe('commonApi', () => {
  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(decodeURIComponent(result.algoliaAgent)).toMatch(
      /^Algolia for JavaScript \(\d+\.\d+\.\d+(-.*)?\)(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*(; Analytics (\(\d+\.\d+\.\d+(-.*)?\)))(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*$/
    );
  });

  test('calls api with default read timeouts', async () => {
    const $client = createClient();

    const result = (await $client.get({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(result).toEqual(
      expect.objectContaining({ connectTimeout: 2000, responseTimeout: 5000 })
    );
  });

  test('calls api with default write timeouts', async () => {
    const $client = createClient();

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(result).toEqual(
      expect.objectContaining({ connectTimeout: 2000, responseTimeout: 30000 })
    );
  });
});

describe('parameters', () => {
  test('fallbacks to the alias when region is not given', async () => {
    const $client = analyticsClient('my-app-id', 'my-api-key', '', {
      requester: echoRequester(),
    });

    const result = (await $client.getAverageClickPosition({
      index: 'my-index',
    })) as unknown as EchoResponse;

    expect(result.host).toEqual('analytics.algolia.com');
  });

  test('uses the correct region', async () => {
    const $client = analyticsClient('my-app-id', 'my-api-key', 'de', {
      requester: echoRequester(),
    });

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(result.host).toEqual('analytics.de.algolia.com');
  });

  test('throws when incorrect region is given', async () => {
    try {
      const $client = analyticsClient(
        'my-app-id',
        'my-api-key',
        'not_a_region',
        { requester: echoRequester() }
      );

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch(
        '`region` must be one of the following: de, us'
      );
    }
  });

  test('getAverageClickPosition throws without index', async () => {
    const $client = createClient();

    try {
      const result = (await $client.getClickPositions(
        {}
      )) as unknown as EchoResponse;

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch(
        'Parameter `index` is required when calling `getClickPositions`.'
      );
    }
  });
});
