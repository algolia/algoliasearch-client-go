import { abtestingClient } from '@experimental-api-clients-automation/client-abtesting';
import type { EchoResponse } from '@experimental-api-clients-automation/client-common';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';

const appId = process.env.ALGOLIA_APPLICATION_ID || 'test_app_id';
const apiKey = process.env.ALGOLIA_SEARCH_KEY || 'test_api_key';

const client = abtestingClient(appId, apiKey, 'us', {
  requester: echoRequester(),
});

describe('addABTests', () => {
  test('addABTests with minimal parameters', async () => {
    const req = (await client.addABTests({
      endAt: '2022-12-31T00:00:00.000Z',
      name: 'myABTest',
      variant: [
        { index: 'AB_TEST_1', trafficPercentage: 30 },
        { index: 'AB_TEST_2', trafficPercentage: 50 },
      ],
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/abtests');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      endAt: '2022-12-31T00:00:00.000Z',
      name: 'myABTest',
      variant: [
        { index: 'AB_TEST_1', trafficPercentage: 30 },
        { index: 'AB_TEST_2', trafficPercentage: 50 },
      ],
    });
    expect(req.searchParams).toStrictEqual(undefined);
  });
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

describe('deleteABTest', () => {
  test('deleteABTest', async () => {
    const req = (await client.deleteABTest({
      id: 42,
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/abtests/42');
    expect(req.method).toEqual('DELETE');
    expect(req.data).toEqual(undefined);
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

describe('getABTest', () => {
  test('getABTest', async () => {
    const req = (await client.getABTest({ id: 42 })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/abtests/42');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });
});

describe('listABTests', () => {
  test('listABTests with minimal parameters', async () => {
    const req = (await client.listABTests({
      offset: 42,
      limit: 21,
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/abtests');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ offset: '42', limit: '21' });
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

describe('stopABTest', () => {
  test('stopABTest', async () => {
    const req = (await client.stopABTest({
      id: 42,
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/abtests/42/stop');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });
});
