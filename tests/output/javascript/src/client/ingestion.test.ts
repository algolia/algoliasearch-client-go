/* eslint-disable @typescript-eslint/no-unused-vars, require-await */
// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { IngestionClient } from '@algolia/ingestion';
import { ingestionClient } from '@algolia/ingestion';
import { echoRequester } from '@algolia/requester-node-http';
import type { EchoResponse } from '@algolia/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): IngestionClient {
  return ingestionClient(appId, apiKey, 'us', { requester: echoRequester() });
}

describe('commonApi', () => {
  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(decodeURIComponent(result.algoliaAgent)).toMatch(
      /^Algolia for JavaScript \(\d+\.\d+\.\d+(-.*)?\)(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*(; Ingestion (\(\d+\.\d+\.\d+(-.*)?\)))(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*$/
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
  test('uses the correct region', async () => {
    const $client = ingestionClient('my-app-id', 'my-api-key', 'us', {
      requester: echoRequester(),
    });

    const result = (await $client.getSource({
      sourceID: '6c02aeb1-775e-418e-870b-1faccd4b2c0f',
    })) as unknown as EchoResponse;

    expect(result.host).toEqual('data.us.algolia.com');
  });

  test('throws when incorrect region is given', async () => {
    try {
      const $client = ingestionClient(
        'my-app-id',
        'my-api-key',
        'not_a_region',
        { requester: echoRequester() }
      );

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch(
        '`region` is required and must be one of the following: eu, us'
      );
    }
  });
});
