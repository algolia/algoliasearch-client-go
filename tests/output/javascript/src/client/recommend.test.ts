// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { RecommendClient } from '@experimental-api-clients-automation/recommend';
import { recommendClient } from '@experimental-api-clients-automation/recommend';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';
import type { EchoResponse } from '@experimental-api-clients-automation/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): RecommendClient {
  return recommendClient(appId, apiKey, { requester: echoRequester() });
}

describe('api', () => {
  test('calls api with correct read host', async () => {
    const $client = recommendClient('test-app-id', 'test-api-key', {
      requester: echoRequester(),
    });

    const result = (await $client.get({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(result.host).toEqual('test-app-id-dsn.algolia.net');
  });
});

describe('commonApi', () => {
  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(decodeURI(result.algoliaAgent)).toMatch(
      /^Algolia for JavaScript \(\d+\.\d+\.\d+(-.*)?\)(; [a-zA-Z. ]+ (\(\d+\.\d+\.\d+(-.*)?\))?)*(; Recommend (\(\d+\.\d+\.\d+(-.*)?\)))(; [a-zA-Z. ]+ (\(\d+\.\d+\.\d+(-.*)?\))?)*$/
    );
  });

  test('calls api with correct timeouts', async () => {
    const $client = createClient();

    const result = (await $client.post({
      path: '/test',
    })) as unknown as EchoResponse;

    expect(result).toEqual(
      expect.objectContaining({ connectTimeout: 2000, responseTimeout: 30000 })
    );
  });
});
