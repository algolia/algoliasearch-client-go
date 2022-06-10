/* eslint-disable @typescript-eslint/no-unused-vars, require-await */
// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { SearchClient } from '@experimental-api-clients-automation/client-search';
import { searchClient } from '@experimental-api-clients-automation/client-search';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';
import type { EchoResponse } from '@experimental-api-clients-automation/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): SearchClient {
  return searchClient(appId, apiKey, { requester: echoRequester() });
}

describe('api', () => {
  test('calls api with correct read host', async () => {
    const $client = searchClient('test-app-id', 'test-api-key', {
      requester: echoRequester(),
    });

    const result = (await $client.get({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(result.host).toEqual('test-app-id-dsn.algolia.net');
  });

  test('calls api with correct write host', async () => {
    const $client = searchClient('test-app-id', 'test-api-key', {
      requester: echoRequester(),
    });

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(result.host).toEqual('test-app-id.algolia.net');
  });
});

describe('commonApi', () => {
  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(decodeURI(result.algoliaAgent)).toMatch(
      /^Algolia for JavaScript \(\d+\.\d+\.\d+(-.*)?\)(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*(; Search (\(\d+\.\d+\.\d+(-.*)?\)))(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*$/
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
  test('client throws with invalid parameters', async () => {
    try {
      const $client = searchClient('', '', { requester: echoRequester() });

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch('`appId` is missing.');
    }
    try {
      const $client = searchClient('', 'my-api-key', {
        requester: echoRequester(),
      });

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch('`appId` is missing.');
    }
    try {
      const $client = searchClient('my-app-id', '', {
        requester: echoRequester(),
      });

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch('`apiKey` is missing.');
    }
  });

  test('`addApiKey` throws with invalid parameters', async () => {
    const $client = createClient();

    try {
      const result = (await $client.addApiKey(null)) as unknown as EchoResponse;

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch(
        'Parameter `apiKey` is required when calling `addApiKey`.'
      );
    }
  });

  test('`addOrUpdateObject` throws with invalid parameters', async () => {
    const $client = createClient();

    try {
      const result = (await $client.addOrUpdateObject({
        objectID: 'my-object-id',
        body: {},
      })) as unknown as EchoResponse;

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch(
        'Parameter `indexName` is required when calling `addOrUpdateObject`.'
      );
    }
    try {
      const result = (await $client.addOrUpdateObject({
        indexName: 'my-index-name',
        body: {},
      })) as unknown as EchoResponse;

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch(
        'Parameter `objectID` is required when calling `addOrUpdateObject`.'
      );
    }
    try {
      const result = (await $client.addOrUpdateObject({
        indexName: 'my-index-name',
        objectID: 'my-object-id',
      })) as unknown as EchoResponse;

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect((e as Error).message).toMatch(
        'Parameter `body` is required when calling `addOrUpdateObject`.'
      );
    }
  });
});
