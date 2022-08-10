/* eslint-disable @typescript-eslint/no-unused-vars, require-await */
// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { SourcesClient } from '@algolia/client-sources';
import { sourcesClient } from '@algolia/client-sources';
import { echoRequester } from '@algolia/requester-node-http';
import type { EchoResponse } from '@algolia/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): SourcesClient {
  return sourcesClient(appId, apiKey, 'us', { requester: echoRequester() });
}

describe('api', () => {
  test('calls api with correct host', async () => {
    const $client = createClient();

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(result.host).toEqual('data.us.algolia.com');
  });
});

describe('commonApi', () => {
  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(decodeURI(result.algoliaAgent)).toMatch(
      /^Algolia for JavaScript \(\d+\.\d+\.\d+(-.*)?\)(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*(; Sources (\(\d+\.\d+\.\d+(-.*)?\)))(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*$/
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
  test('throws when region is not given', async () => {
    try {
      const $client = sourcesClient('my-app-id', 'my-api-key', '', {
        requester: echoRequester(),
      });

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch(
        '`region` is required and must be one of the following: de, us'
      );
    }
  });

  test('throws when incorrect region is given', async () => {
    try {
      const $client = sourcesClient('my-app-id', 'my-api-key', 'not_a_region', {
        requester: echoRequester(),
      });

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch(
        '`region` is required and must be one of the following: de, us'
      );
    }
  });

  test('does not throw when region is given', async () => {
    const $client = sourcesClient('my-app-id', 'my-api-key', 'us', {
      requester: echoRequester(),
    });
  });
});
