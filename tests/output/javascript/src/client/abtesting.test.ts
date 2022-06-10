/* eslint-disable @typescript-eslint/no-unused-vars, require-await */
// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { AbtestingClient } from '@experimental-api-clients-automation/client-abtesting';
import { abtestingClient } from '@experimental-api-clients-automation/client-abtesting';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';
import type { EchoResponse } from '@experimental-api-clients-automation/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): AbtestingClient {
  return abtestingClient(appId, apiKey, 'us', { requester: echoRequester() });
}

describe('commonApi', () => {
  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(decodeURI(result.algoliaAgent)).toMatch(
      /^Algolia for JavaScript \(\d+\.\d+\.\d+(-.*)?\)(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*(; Abtesting (\(\d+\.\d+\.\d+(-.*)?\)))(; [a-zA-Z. ]+ (\(\d+((\.\d+)?\.\d+)?(-.*)?\))?)*$/
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
    const $client = abtestingClient('my-app-id', 'my-api-key', '', {
      requester: echoRequester(),
    });

    const result = (await $client.getABTest({
      id: 123,
    })) as unknown as EchoResponse;

    expect(result.host).toEqual('analytics.algolia.com');
  });

  test('uses the correct region', async () => {
    const $client = abtestingClient('my-app-id', 'my-api-key', 'us', {
      requester: echoRequester(),
    });

    const result = (await $client.getABTest({
      id: 123,
    })) as unknown as EchoResponse;

    expect(result.host).toEqual('analytics.us.algolia.com');
  });

  test('throws when incorrect region is given', async () => {
    try {
      const $client = abtestingClient(
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
});
