import { analyticsClient } from '@experimental-api-clients-automation/client-analytics';
import type {
  EchoResponse,
  RequestOptions,
} from '@experimental-api-clients-automation/client-common';
import { echoRequester } from '@experimental-api-clients-automation/requester-node-http';

const appId = process.env.ALGOLIA_APPLICATION_ID || 'test_app_id';
const apiKey = process.env.ALGOLIA_SEARCH_KEY || 'test_api_key';

const client = analyticsClient(appId, apiKey, 'us', {
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

describe('getAverageClickPosition', () => {
  test('get getAverageClickPosition with minimal parameters', async () => {
    const req = (await client.getAverageClickPosition({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/clicks/averageClickPosition');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getAverageClickPosition with all parameters', async () => {
    const req = (await client.getAverageClickPosition({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/clicks/averageClickPosition');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    });
  });
});

describe('getClickPositions', () => {
  test('get getClickPositions with minimal parameters', async () => {
    const req = (await client.getClickPositions({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/clicks/positions');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getClickPositions with all parameters', async () => {
    const req = (await client.getClickPositions({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/clicks/positions');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    });
  });
});

describe('getClickThroughRate', () => {
  test('get getClickThroughRate with minimal parameters', async () => {
    const req = (await client.getClickThroughRate({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/clicks/clickThroughRate');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getClickThroughRate with all parameters', async () => {
    const req = (await client.getClickThroughRate({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/clicks/clickThroughRate');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    });
  });
});

describe('getConversationRate', () => {
  test('get getConversationRate with minimal parameters', async () => {
    const req = (await client.getConversationRate({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/conversions/conversionRate');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getConversationRate with all parameters', async () => {
    const req = (await client.getConversationRate({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/conversions/conversionRate');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    });
  });
});

describe('getNoClickRate', () => {
  test('get getNoClickRate with minimal parameters', async () => {
    const req = (await client.getNoClickRate({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches/noClickRate');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getNoClickRate with all parameters', async () => {
    const req = (await client.getNoClickRate({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches/noClickRate');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    });
  });
});

describe('getNoResultsRate', () => {
  test('get getNoResultsRate with minimal parameters', async () => {
    const req = (await client.getNoResultsRate({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches/noResultRate');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getNoResultsRate with all parameters', async () => {
    const req = (await client.getNoResultsRate({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches/noResultRate');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    });
  });
});

describe('getSearchesCount', () => {
  test('get getSearchesCount with minimal parameters', async () => {
    const req = (await client.getSearchesCount({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches/count');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getSearchesCount with all parameters', async () => {
    const req = (await client.getSearchesCount({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches/count');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    });
  });
});

describe('getSearchesNoClicks', () => {
  test('get getSearchesNoClicks with minimal parameters', async () => {
    const req = (await client.getSearchesNoClicks({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches/noClicks');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getSearchesNoClicks with all parameters', async () => {
    const req = (await client.getSearchesNoClicks({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: 21,
      offset: 42,
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches/noClicks');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: '21',
      offset: '42',
      tags: 'tag',
    });
  });
});

describe('getSearchesNoResults', () => {
  test('get getSearchesNoResults with minimal parameters', async () => {
    const req = (await client.getSearchesNoResults({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches/noResults');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getSearchesNoResults with all parameters', async () => {
    const req = (await client.getSearchesNoResults({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: 21,
      offset: 42,
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches/noResults');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: '21',
      offset: '42',
      tags: 'tag',
    });
  });
});

describe('getStatus', () => {
  test('get getStatus with minimal parameters', async () => {
    const req = (await client.getStatus({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/status');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });
});

describe('getTopCountries', () => {
  test('get getTopCountries with minimal parameters', async () => {
    const req = (await client.getTopCountries({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/countries');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getTopCountries with all parameters', async () => {
    const req = (await client.getTopCountries({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: 21,
      offset: 42,
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/countries');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: '21',
      offset: '42',
      tags: 'tag',
    });
  });
});

describe('getTopFilterAttributes', () => {
  test('get getTopFilterAttributes with minimal parameters', async () => {
    const req = (await client.getTopFilterAttributes({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/filters');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getTopFilterAttributes with all parameters', async () => {
    const req = (await client.getTopFilterAttributes({
      index: 'index',
      search: 'mySearch',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: 21,
      offset: 42,
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/filters');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      search: 'mySearch',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: '21',
      offset: '42',
      tags: 'tag',
    });
  });
});

describe('getTopFilterForAttribute', () => {
  test('get getTopFilterForAttribute with minimal parameters', async () => {
    const req = (await client.getTopFilterForAttribute({
      attribute: 'myAttribute',
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/filters/myAttribute');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getTopFilterForAttribute with minimal parameters and multiple attributes', async () => {
    const req = (await client.getTopFilterForAttribute({
      attribute: 'myAttribute1,myAttribute2',
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/filters/myAttribute1%2CmyAttribute2');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getTopFilterForAttribute with all parameters', async () => {
    const req = (await client.getTopFilterForAttribute({
      attribute: 'myAttribute',
      index: 'index',
      search: 'mySearch',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: 21,
      offset: 42,
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/filters/myAttribute');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      search: 'mySearch',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: '21',
      offset: '42',
      tags: 'tag',
    });
  });

  test('get getTopFilterForAttribute with all parameters and multiple attributes', async () => {
    const req = (await client.getTopFilterForAttribute({
      attribute: 'myAttribute1,myAttribute2',
      index: 'index',
      search: 'mySearch',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: 21,
      offset: 42,
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/filters/myAttribute1%2CmyAttribute2');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      search: 'mySearch',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: '21',
      offset: '42',
      tags: 'tag',
    });
  });
});

describe('getTopFiltersNoResults', () => {
  test('get getTopFiltersNoResults with minimal parameters', async () => {
    const req = (await client.getTopFiltersNoResults({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/filters/noResults');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getTopFiltersNoResults with all parameters', async () => {
    const req = (await client.getTopFiltersNoResults({
      index: 'index',
      search: 'mySearch',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: 21,
      offset: 42,
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/filters/noResults');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      search: 'mySearch',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: '21',
      offset: '42',
      tags: 'tag',
    });
  });
});

describe('getTopHits', () => {
  test('get getTopHits with minimal parameters', async () => {
    const req = (await client.getTopHits({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/hits');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getTopHits with all parameters', async () => {
    const req = (await client.getTopHits({
      index: 'index',
      search: 'mySearch',
      clickAnalytics: true,
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: 21,
      offset: 42,
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/hits');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      search: 'mySearch',
      clickAnalytics: 'true',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      limit: '21',
      offset: '42',
      tags: 'tag',
    });
  });
});

describe('getTopSearches', () => {
  test('get getTopSearches with minimal parameters', async () => {
    const req = (await client.getTopSearches({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getTopSearches with all parameters', async () => {
    const req = (await client.getTopSearches({
      index: 'index',
      clickAnalytics: true,
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      orderBy: 'searchCount',
      direction: 'asc',
      limit: 21,
      offset: 42,
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/searches');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      clickAnalytics: 'true',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      orderBy: 'searchCount',
      direction: 'asc',
      limit: '21',
      offset: '42',
      tags: 'tag',
    });
  });
});

describe('getUsersCount', () => {
  test('get getUsersCount with minimal parameters', async () => {
    const req = (await client.getUsersCount({
      index: 'index',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/users/count');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({ index: 'index' });
  });

  test('get getUsersCount with all parameters', async () => {
    const req = (await client.getUsersCount({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    })) as unknown as EchoResponse;

    expect(req.path).toEqual('/2/users/count');
    expect(req.method).toEqual('GET');
    expect(req.data).toEqual(undefined);
    expect(req.searchParams).toStrictEqual({
      index: 'index',
      startDate: '1999-09-19',
      endDate: '2001-01-01',
      tags: 'tag',
    });
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
    const requestOptions: RequestOptions = { queryParameters: { myParam: 2 } };

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
