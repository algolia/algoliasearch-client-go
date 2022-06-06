/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable require-await */
// @ts-nocheck Failing tests will have type errors, but we cannot suppress them even with @ts-expect-error because it doesn't work for a block of lines.
import type { PredictClient } from '@experimental-api-clients-automation/client-predict';
import { predictClient } from '@experimental-api-clients-automation/client-predict';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';

const appId = 'test-app-id';
const apiKey = 'test-api-key';

function createClient(): PredictClient {
  return predictClient(appId, apiKey, 'ew', { requester: echoRequester() });
}

describe('api', () => {
  test('calls api with correct user agent', async () => {
    const $client = createClient();

    const result = await $client.fetchUserProfile({
      userID: 'user1',
      params: { modelsToRetrieve: ['funnel_stage'] },
    });

    expect(result.algoliaAgent).toMatch(
      /Algolia%20for%20(.+)%20\(\d+\.\d+\.\d+\)/
    );
  });

  test('calls api with correct timeouts', async () => {
    const $client = createClient();

    const result = await $client.fetchUserProfile({
      userID: 'user1',
      params: { modelsToRetrieve: ['funnel_stage'] },
    });

    expect(result).toEqual(
      expect.objectContaining({ connectTimeout: 2, responseTimeout: 30 })
    );
  });
});

describe('parameters', () => {
  test('throws when region is not given', async () => {
    try {
      const $client = predictClient('my-app-id', 'my-api-key', '', {
        requester: echoRequester(),
      });

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect(e.message).toMatch('`region` is missing.');
    }
  });

  test('throws when incorrect region is given', async () => {
    try {
      const $client = predictClient('my-app-id', 'my-api-key', 'not_a_region', {
        requester: echoRequester(),
      });

      throw new Error('test is expected to throw error');
    } catch (e) {
      expect(e.message).toMatch(
        '`region` must be one of the following: ue, ew'
      );
    }
  });

  test('does not throw when region is given', async () => {
    const $client = predictClient('my-app-id', 'my-api-key', 'ew', {
      requester: echoRequester(),
    });
  });
});
