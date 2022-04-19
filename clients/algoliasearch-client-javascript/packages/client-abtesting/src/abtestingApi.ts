import {
  createAuth,
  createTransporter,
  getUserAgent,
} from '@experimental-api-clients-automation/client-common';
import type {
  CreateClientOptions,
  Headers,
  Host,
  Request,
  RequestOptions,
  QueryParameters,
} from '@experimental-api-clients-automation/client-common';

import type { ABTest } from '../model/aBTest';
import type { ABTestResponse } from '../model/aBTestResponse';
import type { AddABTestsRequest } from '../model/addABTestsRequest';
import type { ListABTestsResponse } from '../model/listABTestsResponse';

export * from '../model/models';
export const apiClientVersion = '0.0.5';

export type Region = 'de' | 'us';

function getDefaultHosts(region?: Region): Host[] {
  const regionHost = region ? `.${region}.` : '.';

  return [
    {
      url: `analytics${regionHost}algolia.com`,
      accept: 'readWrite',
      protocol: 'https',
    },
  ];
}

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export function createAbtestingApi(
  options: CreateClientOptions & { region?: Region }
) {
  const auth = createAuth(options.appId, options.apiKey, options.authMode);
  const transporter = createTransporter({
    hosts: options?.hosts ?? getDefaultHosts(options.region),
    hostsCache: options.hostsCache,
    requestsCache: options.requestsCache,
    responsesCache: options.responsesCache,
    baseHeaders: {
      'content-type': 'application/x-www-form-urlencoded',
      ...auth.headers(),
    },
    baseQueryParameters: auth.queryParameters(),
    userAgent: getUserAgent({
      userAgents: options.userAgents,
      client: 'Abtesting',
      version: apiClientVersion,
    }),
    timeouts: options.timeouts,
    requester: options.requester,
  });

  function addUserAgent(segment: string, version?: string): void {
    transporter.userAgent.add({ segment, version });
  }

  return {
    addUserAgent,
    /**
     * Creates a new A/B test with provided configuration. You can set an A/B test on two different indices with different settings, or on the same index with different search parameters by providing a customSearchParameters setting on one of the variants.
     *
     * @summary Creates a new A/B test with provided configuration.
     * @param addABTestsRequest - The addABTestsRequest object.
     */
    addABTests(
      addABTestsRequest: AddABTestsRequest,
      requestOptions?: RequestOptions
    ): Promise<ABTestResponse> {
      if (!addABTestsRequest) {
        throw new Error(
          'Parameter `addABTestsRequest` is required when calling `addABTests`.'
        );
      }

      if (!addABTestsRequest.name) {
        throw new Error(
          'Parameter `addABTestsRequest.name` is required when calling `addABTests`.'
        );
      }
      if (!addABTestsRequest.variant) {
        throw new Error(
          'Parameter `addABTestsRequest.variant` is required when calling `addABTests`.'
        );
      }
      if (!addABTestsRequest.endAt) {
        throw new Error(
          'Parameter `addABTestsRequest.endAt` is required when calling `addABTests`.'
        );
      }

      const requestPath = '/2/abtests';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        data: addABTestsRequest,
      };

      return transporter.request(
        request,
        {
          queryParameters,
          headers,
        },
        requestOptions
      );
    },

    /**
     * This method allow you to send requests to the Algolia REST API.
     *
     * @summary Send requests to the Algolia REST API.
     * @param del - The del object.
     * @param del.path - The path of the API endpoint to target, anything after the /1 needs to be specified.
     * @param del.parameters - Query parameters to be applied to the current query.
     */
    del(
      { path, parameters }: DelProps,
      requestOptions?: RequestOptions
    ): Promise<Record<string, any>> {
      if (!path) {
        throw new Error('Parameter `path` is required when calling `del`.');
      }

      const requestPath = '/1{path}'.replace('{path}', path);
      const headers: Headers = {};
      const queryParameters: QueryParameters = parameters || {};

      const request: Request = {
        method: 'DELETE',
        path: requestPath,
      };

      return transporter.request(
        request,
        {
          queryParameters,
          headers,
        },
        requestOptions
      );
    },

    /**
     * Deletes the A/B Test and removes all associated metadata & metrics.
     *
     * @summary Deletes the A/B Test.
     * @param deleteABTest - The deleteABTest object.
     * @param deleteABTest.id - The A/B test ID.
     */
    deleteABTest(
      { id }: DeleteABTestProps,
      requestOptions?: RequestOptions
    ): Promise<ABTestResponse> {
      if (!id) {
        throw new Error(
          'Parameter `id` is required when calling `deleteABTest`.'
        );
      }

      const requestPath = '/2/abtests/{id}'.replace(
        '{id}',
        encodeURIComponent(id)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'DELETE',
        path: requestPath,
      };

      return transporter.request(
        request,
        {
          queryParameters,
          headers,
        },
        requestOptions
      );
    },

    /**
     * This method allow you to send requests to the Algolia REST API.
     *
     * @summary Send requests to the Algolia REST API.
     * @param get - The get object.
     * @param get.path - The path of the API endpoint to target, anything after the /1 needs to be specified.
     * @param get.parameters - Query parameters to be applied to the current query.
     */
    get(
      { path, parameters }: GetProps,
      requestOptions?: RequestOptions
    ): Promise<Record<string, any>> {
      if (!path) {
        throw new Error('Parameter `path` is required when calling `get`.');
      }

      const requestPath = '/1{path}'.replace('{path}', path);
      const headers: Headers = {};
      const queryParameters: QueryParameters = parameters || {};

      const request: Request = {
        method: 'GET',
        path: requestPath,
      };

      return transporter.request(
        request,
        {
          queryParameters,
          headers,
        },
        requestOptions
      );
    },

    /**
     * Returns metadata and metrics for A/B test id. Behaves in the same way as GET /2/abtests however the endpoint will return 403.
     *
     * @summary Returns metadata and metrics for A/B test id.
     * @param getABTest - The getABTest object.
     * @param getABTest.id - The A/B test ID.
     */
    getABTest(
      { id }: GetABTestProps,
      requestOptions?: RequestOptions
    ): Promise<ABTest> {
      if (!id) {
        throw new Error('Parameter `id` is required when calling `getABTest`.');
      }

      const requestPath = '/2/abtests/{id}'.replace(
        '{id}',
        encodeURIComponent(id)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'GET',
        path: requestPath,
      };

      return transporter.request(
        request,
        {
          queryParameters,
          headers,
        },
        requestOptions
      );
    },

    /**
     * Fetch all existing A/B tests for App that are available for the current API Key. Returns an array of metadata and metrics. When no data has been processed, the metrics will be returned as null.
     *
     * @summary Fetch all existing A/B tests for App that are available for the current API Key.
     * @param listABTests - The listABTests object.
     * @param listABTests.offset - Position of the starting record. Used for paging. 0 is the first record.
     * @param listABTests.limit - Number of records to return. Limit is the size of the page.
     */
    listABTests(
      { offset, limit }: ListABTestsProps,
      requestOptions?: RequestOptions
    ): Promise<ListABTestsResponse> {
      const requestPath = '/2/abtests';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (offset !== undefined) {
        queryParameters.offset = offset.toString();
      }

      if (limit !== undefined) {
        queryParameters.limit = limit.toString();
      }

      const request: Request = {
        method: 'GET',
        path: requestPath,
      };

      return transporter.request(
        request,
        {
          queryParameters,
          headers,
        },
        requestOptions
      );
    },

    /**
     * This method allow you to send requests to the Algolia REST API.
     *
     * @summary Send requests to the Algolia REST API.
     * @param post - The post object.
     * @param post.path - The path of the API endpoint to target, anything after the /1 needs to be specified.
     * @param post.parameters - Query parameters to be applied to the current query.
     * @param post.body - The parameters to send with the custom request.
     */
    post(
      { path, parameters, body }: PostProps,
      requestOptions?: RequestOptions
    ): Promise<Record<string, any>> {
      if (!path) {
        throw new Error('Parameter `path` is required when calling `post`.');
      }

      const requestPath = '/1{path}'.replace('{path}', path);
      const headers: Headers = {};
      const queryParameters: QueryParameters = parameters || {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        data: body,
      };

      return transporter.request(
        request,
        {
          queryParameters,
          headers,
        },
        requestOptions
      );
    },

    /**
     * This method allow you to send requests to the Algolia REST API.
     *
     * @summary Send requests to the Algolia REST API.
     * @param put - The put object.
     * @param put.path - The path of the API endpoint to target, anything after the /1 needs to be specified.
     * @param put.parameters - Query parameters to be applied to the current query.
     * @param put.body - The parameters to send with the custom request.
     */
    put(
      { path, parameters, body }: PutProps,
      requestOptions?: RequestOptions
    ): Promise<Record<string, any>> {
      if (!path) {
        throw new Error('Parameter `path` is required when calling `put`.');
      }

      const requestPath = '/1{path}'.replace('{path}', path);
      const headers: Headers = {};
      const queryParameters: QueryParameters = parameters || {};

      const request: Request = {
        method: 'PUT',
        path: requestPath,
        data: body,
      };

      return transporter.request(
        request,
        {
          queryParameters,
          headers,
        },
        requestOptions
      );
    },

    /**
     * Marks the A/B test as stopped. At this point, the test is over and cannot be restarted. As a result, your application is back to normal: index A will perform as usual, receiving 100% of all search requests. Associated metadata and metrics are still stored.
     *
     * @summary Marks the A/B test as stopped.
     * @param stopABTest - The stopABTest object.
     * @param stopABTest.id - The A/B test ID.
     */
    stopABTest(
      { id }: StopABTestProps,
      requestOptions?: RequestOptions
    ): Promise<ABTestResponse> {
      if (!id) {
        throw new Error(
          'Parameter `id` is required when calling `stopABTest`.'
        );
      }

      const requestPath = '/2/abtests/{id}/stop'.replace(
        '{id}',
        encodeURIComponent(id)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
      };

      return transporter.request(
        request,
        {
          queryParameters,
          headers,
        },
        requestOptions
      );
    },
  };
}

export type AbtestingApi = ReturnType<typeof createAbtestingApi>;

export type DelProps = {
  /**
   * The path of the API endpoint to target, anything after the /1 needs to be specified.
   */
  path: string;
  /**
   * Query parameters to be applied to the current query.
   */
  parameters?: Record<string, any>;
};

export type DeleteABTestProps = {
  /**
   * The A/B test ID.
   */
  id: number;
};

export type GetProps = {
  /**
   * The path of the API endpoint to target, anything after the /1 needs to be specified.
   */
  path: string;
  /**
   * Query parameters to be applied to the current query.
   */
  parameters?: Record<string, any>;
};

export type GetABTestProps = {
  /**
   * The A/B test ID.
   */
  id: number;
};

export type ListABTestsProps = {
  /**
   * Position of the starting record. Used for paging. 0 is the first record.
   */
  offset?: number;
  /**
   * Number of records to return. Limit is the size of the page.
   */
  limit?: number;
};

export type PostProps = {
  /**
   * The path of the API endpoint to target, anything after the /1 needs to be specified.
   */
  path: string;
  /**
   * Query parameters to be applied to the current query.
   */
  parameters?: Record<string, any>;
  /**
   * The parameters to send with the custom request.
   */
  body?: Record<string, any>;
};

export type PutProps = {
  /**
   * The path of the API endpoint to target, anything after the /1 needs to be specified.
   */
  path: string;
  /**
   * Query parameters to be applied to the current query.
   */
  parameters?: Record<string, any>;
  /**
   * The parameters to send with the custom request.
   */
  body?: Record<string, any>;
};

export type StopABTestProps = {
  /**
   * The A/B test ID.
   */
  id: number;
};
