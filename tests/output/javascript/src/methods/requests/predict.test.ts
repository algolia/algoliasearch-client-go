import type { EchoResponse, RequestOptions } from '@algolia/client-common';
import { predictClient } from '@algolia/predict';
import { echoRequester } from '@algolia/requester-node-http';

const appId = process.env.ALGOLIA_APPLICATION_ID || 'test_app_id';
const apiKey = process.env.ALGOLIA_SEARCH_KEY || 'test_api_key';

const client = predictClient(appId, apiKey, 'eu', {
  requester: echoRequester(),
});

describe('activateModelInstance', () => {
  test('activate a model instance', async () => {
    const req = (await client.activateModelInstance({
      type: 'funnel_stage',
      name: 'Shopping stage for EU users',
      sourceID: '0200030-129930',
      index: 'Products Production',
      affinities: [],
      contentAttributes: ['title', 'description'],
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/predict/models');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      type: 'funnel_stage',
      name: 'Shopping stage for EU users',
      sourceID: '0200030-129930',
      index: 'Products Production',
      affinities: [],
      contentAttributes: ['title', 'description'],
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

describe('deleteModelInstance', () => {
  test('delete a model instance', async () => {
    const req = (await client.deleteModelInstance({
      modelID: 'model1',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/predict/models/model1');
    expect(req.method).toEqual('DELETE');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });
});

describe('deleteUserProfile', () => {
  test('deleteUserProfile', async () => {
    const req = (await client.deleteUserProfile({
      userID: 'user1',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/users/user1');
    expect(req.method).toEqual('DELETE');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });
});

describe('fetchAllUserProfiles', () => {
  test('fetchAllUserProfiles with minimal parameters for modelsToRetrieve', async () => {
    const req = (await client.fetchAllUserProfiles({
      modelsToRetrieve: ['funnel_stage', 'order_value', 'affinities'],
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/users');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      modelsToRetrieve: ['funnel_stage', 'order_value', 'affinities'],
    });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('fetchAllUserProfiles with minimal parameters for typesToRetrieve', async () => {
    const req = (await client.fetchAllUserProfiles({
      typesToRetrieve: ['properties', 'segments'],
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/users');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ typesToRetrieve: ['properties', 'segments'] });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('fetchAllUserProfiles with a limit', async () => {
    const req = (await client.fetchAllUserProfiles({
      limit: 10,
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/users');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ limit: 10 });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('fetchAllUserProfiles with a nextPageToken', async () => {
    const req = (await client.fetchAllUserProfiles({
      nextPageToken: 'nextPageTokenExample123',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/users');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ nextPageToken: 'nextPageTokenExample123' });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('fetchAllUserProfiles with a previousPageToken', async () => {
    const req = (await client.fetchAllUserProfiles({
      previousPageToken: 'previousPageTokenExample123',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/users');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      previousPageToken: 'previousPageTokenExample123',
    });
    expect(req.searchParams).toStrictEqual(undefined);
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

describe('getAvailableModelTypes', () => {
  test('get available model types', async () => {
    const req =
      (await client.getAvailableModelTypes()) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/predict/modeltypes');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });
});

describe('getModelInstanceConfig', () => {
  test('get configurations for a model instance', async () => {
    const req = (await client.getModelInstanceConfig({
      modelID: 'model1',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/predict/models/model1');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });
});

describe('getModelInstances', () => {
  test('get a list of model instances', async () => {
    const req = (await client.getModelInstances()) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/predict/models');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });
});

describe('getModelMetrics', () => {
  test('get metrics for a model instance', async () => {
    const req = (await client.getModelMetrics({
      modelID: 'model1',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/predict/models/model1/metrics');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual(undefined);
  });
});

describe('post', () => {
  test('allow post method for a custom path with minimal parameters', async () => {
    const req = (await client.post({
      path: '/test/minimal',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/minimal');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({});
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

  test('requestOptions can override default query parameters', async () => {
    const requestOptions: RequestOptions = {
      queryParameters: { query: 'myQueryParameter' },
    };

    const req = (await client.post(
      {
        path: '/test/requestOptions',
        parameters: { query: 'parameters' },
        body: { facet: 'filters' },
      },
      requestOptions
    )) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/requestOptions');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ facet: 'filters' });
    expect(req.searchParams).toStrictEqual({ query: 'myQueryParameter' });
  });

  test('requestOptions merges query parameters with default ones', async () => {
    const requestOptions: RequestOptions = {
      queryParameters: { query2: 'myQueryParameter' },
    };

    const req = (await client.post(
      {
        path: '/test/requestOptions',
        parameters: { query: 'parameters' },
        body: { facet: 'filters' },
      },
      requestOptions
    )) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/requestOptions');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ facet: 'filters' });
    expect(req.searchParams).toStrictEqual({
      query: 'parameters',
      query2: 'myQueryParameter',
    });
  });

  test('requestOptions can override default headers', async () => {
    const requestOptions: RequestOptions = {
      headers: { 'x-algolia-api-key': 'myApiKey' },
    };

    const req = (await client.post(
      {
        path: '/test/requestOptions',
        parameters: { query: 'parameters' },
        body: { facet: 'filters' },
      },
      requestOptions
    )) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/requestOptions');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ facet: 'filters' });
    expect(req.searchParams).toStrictEqual({ query: 'parameters' });
    expect(req.headers).toEqual(
      expect.objectContaining({ 'x-algolia-api-key': 'myApiKey' })
    );
  });

  test('requestOptions merges headers with default ones', async () => {
    const requestOptions: RequestOptions = {
      headers: { 'x-algolia-api-key': 'myApiKey' },
    };

    const req = (await client.post(
      {
        path: '/test/requestOptions',
        parameters: { query: 'parameters' },
        body: { facet: 'filters' },
      },
      requestOptions
    )) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/requestOptions');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ facet: 'filters' });
    expect(req.searchParams).toStrictEqual({ query: 'parameters' });
    expect(req.headers).toEqual(
      expect.objectContaining({ 'x-algolia-api-key': 'myApiKey' })
    );
  });

  test('requestOptions queryParameters accepts booleans', async () => {
    const requestOptions: RequestOptions = {
      queryParameters: { isItWorking: true },
    };

    const req = (await client.post(
      {
        path: '/test/requestOptions',
        parameters: { query: 'parameters' },
        body: { facet: 'filters' },
      },
      requestOptions
    )) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/requestOptions');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ facet: 'filters' });
    expect(req.searchParams).toStrictEqual({
      query: 'parameters',
      isItWorking: 'true',
    });
  });

  test('requestOptions queryParameters accepts integers', async () => {
    const requestOptions: RequestOptions = {
      queryParameters: { myParam: 2 },
    };

    const req = (await client.post(
      {
        path: '/test/requestOptions',
        parameters: { query: 'parameters' },
        body: { facet: 'filters' },
      },
      requestOptions
    )) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/requestOptions');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ facet: 'filters' });
    expect(req.searchParams).toStrictEqual({
      query: 'parameters',
      myParam: '2',
    });
  });

  test('requestOptions queryParameters accepts list of string', async () => {
    const requestOptions: RequestOptions = {
      queryParameters: { myParam: ['c', 'd'] },
    };

    const req = (await client.post(
      {
        path: '/test/requestOptions',
        parameters: { query: 'parameters' },
        body: { facet: 'filters' },
      },
      requestOptions
    )) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/requestOptions');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ facet: 'filters' });
    expect(req.searchParams).toStrictEqual({
      query: 'parameters',
      myParam: 'c,d',
    });
  });

  test('requestOptions queryParameters accepts list of booleans', async () => {
    const requestOptions: RequestOptions = {
      queryParameters: { myParam: [true, true, false] },
    };

    const req = (await client.post(
      {
        path: '/test/requestOptions',
        parameters: { query: 'parameters' },
        body: { facet: 'filters' },
      },
      requestOptions
    )) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/requestOptions');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ facet: 'filters' });
    expect(req.searchParams).toStrictEqual({
      query: 'parameters',
      myParam: 'true,true,false',
    });
  });

  test('requestOptions queryParameters accepts list of integers', async () => {
    const requestOptions: RequestOptions = {
      queryParameters: { myParam: [1, 2] },
    };

    const req = (await client.post(
      {
        path: '/test/requestOptions',
        parameters: { query: 'parameters' },
        body: { facet: 'filters' },
      },
      requestOptions
    )) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/requestOptions');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ facet: 'filters' });
    expect(req.searchParams).toStrictEqual({
      query: 'parameters',
      myParam: '1,2',
    });
  });
});

describe('put', () => {
  test('allow put method for a custom path with minimal parameters', async () => {
    const req = (await client.put({
      path: '/test/minimal',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/test/minimal');
    expect(req.method).toEqual('PUT');
    expect(req.data).toEqual({});
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

describe('updateModelInstance', () => {
  test('update a model instance', async () => {
    const req = (await client.updateModelInstance({
      modelID: 'model1',
      updateModelParams: {
        name: 'Shopping stage for EU users',
        affinities: ['brand', 'color', 'category_level0', 'category_level1'],
        contentAttributes: ['title', 'description'],
        status: 'inactive',
      },
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/predict/models/model1');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      name: 'Shopping stage for EU users',
      affinities: ['brand', 'color', 'category_level0', 'category_level1'],
      contentAttributes: ['title', 'description'],
      status: 'inactive',
    });
    expect(req.searchParams).toStrictEqual(undefined);
  });
});
