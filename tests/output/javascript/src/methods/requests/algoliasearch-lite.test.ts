import { algoliasearchLiteClient } from '@experimental-api-clients-automation/algoliasearch-lite';
import type {
  EchoResponse,
  RequestOptions,
} from '@experimental-api-clients-automation/client-common';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';

const appId = process.env.ALGOLIA_APPLICATION_ID || 'test_app_id';
const apiKey = process.env.ALGOLIA_SEARCH_KEY || 'test_api_key';

const client = algoliasearchLiteClient(appId, apiKey, {
  requester: echoRequester(),
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

describe('search', () => {
  test('search for a single hits request with minimal parameters', async () => {
    const req = (await client.search({
      requests: [{ indexName: 'theIndexName' }],
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/indexes/*/queries');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({ requests: [{ indexName: 'theIndexName' }] });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('search for a single facet request with minimal parameters', async () => {
    const req = (await client.search({
      requests: [
        { indexName: 'theIndexName', type: 'facet', facet: 'theFacet' },
      ],
      strategy: 'stopIfEnoughMatches',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/indexes/*/queries');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      requests: [
        { indexName: 'theIndexName', type: 'facet', facet: 'theFacet' },
      ],
      strategy: 'stopIfEnoughMatches',
    });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('search for a single hits request with all parameters', async () => {
    const req = (await client.search({
      requests: [
        {
          indexName: 'theIndexName',
          query: 'myQuery',
          hitsPerPage: 50,
          type: 'default',
        },
      ],
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/indexes/*/queries');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      requests: [
        {
          indexName: 'theIndexName',
          query: 'myQuery',
          hitsPerPage: 50,
          type: 'default',
        },
      ],
    });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('search for a single facet request with all parameters', async () => {
    const req = (await client.search({
      requests: [
        {
          indexName: 'theIndexName',
          type: 'facet',
          facet: 'theFacet',
          facetQuery: 'theFacetQuery',
          query: 'theQuery',
          maxFacetHits: 50,
        },
      ],
      strategy: 'stopIfEnoughMatches',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/indexes/*/queries');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      requests: [
        {
          indexName: 'theIndexName',
          type: 'facet',
          facet: 'theFacet',
          facetQuery: 'theFacetQuery',
          query: 'theQuery',
          maxFacetHits: 50,
        },
      ],
      strategy: 'stopIfEnoughMatches',
    });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('search for multiple mixed requests in multiple indices with minimal parameters', async () => {
    const req = (await client.search({
      requests: [
        { indexName: 'theIndexName' },
        { indexName: 'theIndexName2', type: 'facet', facet: 'theFacet' },
        { indexName: 'theIndexName', type: 'default' },
      ],
      strategy: 'stopIfEnoughMatches',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/indexes/*/queries');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      requests: [
        { indexName: 'theIndexName' },
        { indexName: 'theIndexName2', type: 'facet', facet: 'theFacet' },
        { indexName: 'theIndexName', type: 'default' },
      ],
      strategy: 'stopIfEnoughMatches',
    });
    expect(req.searchParams).toStrictEqual(undefined);
  });

  test('search for multiple mixed requests in multiple indices with all parameters', async () => {
    const req = (await client.search({
      requests: [
        {
          indexName: 'theIndexName',
          type: 'facet',
          facet: 'theFacet',
          facetQuery: 'theFacetQuery',
          query: 'theQuery',
          maxFacetHits: 50,
        },
        {
          indexName: 'theIndexName',
          query: 'myQuery',
          hitsPerPage: 50,
          type: 'default',
        },
      ],
      strategy: 'stopIfEnoughMatches',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/1/indexes/*/queries');
    expect(req.method).toEqual('POST');
    expect(req.data).toEqual({
      requests: [
        {
          indexName: 'theIndexName',
          type: 'facet',
          facet: 'theFacet',
          facetQuery: 'theFacetQuery',
          query: 'theQuery',
          maxFacetHits: 50,
        },
        {
          indexName: 'theIndexName',
          query: 'myQuery',
          hitsPerPage: 50,
          type: 'default',
        },
      ],
      strategy: 'stopIfEnoughMatches',
    });
    expect(req.searchParams).toStrictEqual(undefined);
  });
});
