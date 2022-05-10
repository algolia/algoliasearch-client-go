import type { EchoResponse } from '@experimental-api-clients-automation/client-common';
import { predictClient } from '@experimental-api-clients-automation/client-predict';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';

const appId = process.env.ALGOLIA_APPLICATION_ID || 'test_app_id';
const apiKey = process.env.ALGOLIA_SEARCH_KEY || 'test_api_key';

const client = predictClient(appId, apiKey, 'ew', {
  requester: echoRequester(),
});

describe('del', () => {
  test('allow del method for a custom path with minimal parameters', async () => {
    const req = (await client.del({
      path: '/test/minimal',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/minimal');
    expect(req.method).toEqual('DELETE');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('allow del method for a custom path with all parameters', async () => {
    const req = (await client.del({
      path: '/test/all',
      parameters: { query: 'parameters' },
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/all');
    expect(req.method).toEqual('DELETE');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ query: 'parameters' });
  });
});

describe('fetchUserProfile', () => {
  test('fetchUserProfile with minimal parameters for modelsToRetrieve', async () => {
    const req = (await client.fetchUserProfile({
      userID: 'user1',
      params: {
        modelsToRetrieve: ['funnel_stage', 'order_value', 'affinities'],
      },
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/users/user1/fetch');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      modelsToRetrieve: ['funnel_stage', 'order_value', 'affinities'],
    });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('fetchUserProfile with minimal parameters for typesToRetrieve', async () => {
    const req = (await client.fetchUserProfile({
      userID: 'user1',
      params: { typesToRetrieve: ['properties', 'segments'] },
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/users/user1/fetch');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ typesToRetrieve: ['properties', 'segments'] });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('fetchUserProfile with all parameters', async () => {
    const req = (await client.fetchUserProfile({
      userID: 'user1',
      params: {
        modelsToRetrieve: ['funnel_stage', 'order_value', 'affinities'],
        typesToRetrieve: ['properties', 'segments'],
      },
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/users/user1/fetch');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      modelsToRetrieve: ['funnel_stage', 'order_value', 'affinities'],
      typesToRetrieve: ['properties', 'segments'],
    });
    expect(req.searchParams).toStrictEqual(undefined);
  });
});

describe('get', () => {
  test('allow get method for a custom path with minimal parameters', async () => {
    const req = (await client.get({
      path: '/test/minimal',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/minimal');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('allow get method for a custom path with all parameters', async () => {
    const req = (await client.get({
      path: '/test/all',
      parameters: { query: 'parameters' },
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/all');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ query: 'parameters' });
  });
});

describe('post', () => {
  test('allow post method for a custom path with minimal parameters', async () => {
    const req = (await client.post({
      path: '/test/minimal',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/minimal');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('allow post method for a custom path with all parameters', async () => {
    const req = (await client.post({
      path: '/test/all',
      parameters: { query: 'parameters' },
      body: { body: 'parameters' },
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/all');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ body: 'parameters' });
    expect(req.searchParams).toStrictEqual({ query: 'parameters' });
  });
});

describe('put', () => {
  test('allow put method for a custom path with minimal parameters', async () => {
    const req = (await client.put({
      path: '/test/minimal',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/minimal');
    expect(req.method).toEqual('PUT');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('allow put method for a custom path with all parameters', async () => {
    const req = (await client.put({
      path: '/test/all',
      parameters: { query: 'parameters' },
      body: { body: 'parameters' },
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/all');
    expect(req.method).toEqual('PUT');
    expect(req.data).toEqual({ body: 'parameters' });
    expect(req.searchParams).toStrictEqual({ query: 'parameters' });
  });
});
