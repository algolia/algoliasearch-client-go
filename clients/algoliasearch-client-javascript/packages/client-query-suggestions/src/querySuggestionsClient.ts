import {
  createAuth,
  createTransporter,
  getAlgoliaAgent,
} from '@experimental-api-clients-automation/client-common';
import type {
  CreateClientOptions,
  Headers,
  Host,
  Request,
  RequestOptions,
  QueryParameters,
} from '@experimental-api-clients-automation/client-common';

import type { LogFile } from '../model/logFile';
import type { QuerySuggestionsIndex } from '../model/querySuggestionsIndex';
import type { QuerySuggestionsIndexParam } from '../model/querySuggestionsIndexParam';
import type { QuerySuggestionsIndexWithIndexParam } from '../model/querySuggestionsIndexWithIndexParam';
import type { Status } from '../model/status';
import type { SucessResponse } from '../model/sucessResponse';

export * from '../model';
export const apiClientVersion = '0.2.0';

export type Region = 'eu' | 'us';

function getDefaultHosts(region: Region): Host[] {
  const url = 'query-suggestions.{region}.algolia.com'.replace(
    '{region}',
    region
  );

  return [{ url, accept: 'readWrite', protocol: 'https' }];
}

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export function createQuerySuggestionsClient(
  options: CreateClientOptions & { region: Region }
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
    algoliaAgent: getAlgoliaAgent({
      algoliaAgents: options.algoliaAgents,
      client: 'QuerySuggestions',
      version: apiClientVersion,
    }),
    timeouts: options.timeouts,
    requester: options.requester,
  });

  function addAlgoliaAgent(segment: string, version?: string): void {
    transporter.algoliaAgent.add({ segment, version });
  }

  return {
    addAlgoliaAgent,
    /**
     * Create a configuration of a Query Suggestions index. There\'s a limit of 100 configurations per application.
     *
     * @summary Create a configuration.
     * @param querySuggestionsIndexWithIndexParam - The querySuggestionsIndexWithIndexParam object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    createConfig(
      querySuggestionsIndexWithIndexParam: QuerySuggestionsIndexWithIndexParam,
      requestOptions?: RequestOptions
    ): Promise<SucessResponse> {
      if (!querySuggestionsIndexWithIndexParam) {
        throw new Error(
          'Parameter `querySuggestionsIndexWithIndexParam` is required when calling `createConfig`.'
        );
      }

      const requestPath = '/1/configs';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        data: querySuggestionsIndexWithIndexParam,
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
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
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
     * Delete a configuration of a Query Suggestion\'s index. By deleting a configuraton, you stop all updates to the underlying query suggestion index. Note that when doing this, the underlying index does not change - existing suggestions remain untouched.
     *
     * @summary Delete a configuration.
     * @param deleteConfig - The deleteConfig object.
     * @param deleteConfig.indexName - The index in which to perform the request.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    deleteConfig(
      { indexName }: DeleteConfigProps,
      requestOptions?: RequestOptions
    ): Promise<SucessResponse> {
      if (!indexName) {
        throw new Error(
          'Parameter `indexName` is required when calling `deleteConfig`.'
        );
      }

      const requestPath = '/1/configs/{indexName}'.replace(
        '{indexName}',
        encodeURIComponent(indexName)
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
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
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
     * Get all the configurations of Query Suggestions. For each index, you get a block of JSON with a list of its configuration settings.
     *
     * @summary List configurations.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getAllConfigs(
      requestOptions?: RequestOptions
    ): Promise<QuerySuggestionsIndex[]> {
      const requestPath = '/1/configs';
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
     * Get the configuration of a single Query Suggestions index.
     *
     * @summary Get a single configuration.
     * @param getConfig - The getConfig object.
     * @param getConfig.indexName - The index in which to perform the request.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getConfig(
      { indexName }: GetConfigProps,
      requestOptions?: RequestOptions
    ): Promise<QuerySuggestionsIndex> {
      if (!indexName) {
        throw new Error(
          'Parameter `indexName` is required when calling `getConfig`.'
        );
      }

      const requestPath = '/1/configs/{indexName}'.replace(
        '{indexName}',
        encodeURIComponent(indexName)
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
     * Get the status of a Query Suggestion\'s index. The status includes whether the Query Suggestions index is currently in the process of being built, and the last build time.
     *
     * @summary Get configuration status.
     * @param getConfigStatus - The getConfigStatus object.
     * @param getConfigStatus.indexName - The index in which to perform the request.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getConfigStatus(
      { indexName }: GetConfigStatusProps,
      requestOptions?: RequestOptions
    ): Promise<Status> {
      if (!indexName) {
        throw new Error(
          'Parameter `indexName` is required when calling `getConfigStatus`.'
        );
      }

      const requestPath = '/1/configs/{indexName}/status'.replace(
        '{indexName}',
        encodeURIComponent(indexName)
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
     * Get the log file of the last build of a single Query Suggestion index.
     *
     * @summary Get a log file.
     * @param getLogFile - The getLogFile object.
     * @param getLogFile.indexName - The index in which to perform the request.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getLogFile(
      { indexName }: GetLogFileProps,
      requestOptions?: RequestOptions
    ): Promise<LogFile[]> {
      if (!indexName) {
        throw new Error(
          'Parameter `indexName` is required when calling `getLogFile`.'
        );
      }

      const requestPath = '/1/logs/{indexName}'.replace(
        '{indexName}',
        encodeURIComponent(indexName)
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
     * This method allow you to send requests to the Algolia REST API.
     *
     * @summary Send requests to the Algolia REST API.
     * @param post - The post object.
     * @param post.path - The path of the API endpoint to target, anything after the /1 needs to be specified.
     * @param post.parameters - Query parameters to be applied to the current query.
     * @param post.body - The parameters to send with the custom request.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
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
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
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
     * Update the configuration of a Query Suggestions index.
     *
     * @summary Update a configuration.
     * @param updateConfig - The updateConfig object.
     * @param updateConfig.indexName - The index in which to perform the request.
     * @param updateConfig.querySuggestionsIndexParam - The querySuggestionsIndexParam object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    updateConfig(
      { indexName, querySuggestionsIndexParam }: UpdateConfigProps,
      requestOptions?: RequestOptions
    ): Promise<SucessResponse> {
      if (!indexName) {
        throw new Error(
          'Parameter `indexName` is required when calling `updateConfig`.'
        );
      }

      if (!querySuggestionsIndexParam) {
        throw new Error(
          'Parameter `querySuggestionsIndexParam` is required when calling `updateConfig`.'
        );
      }

      if (!querySuggestionsIndexParam.sourceIndices) {
        throw new Error(
          'Parameter `querySuggestionsIndexParam.sourceIndices` is required when calling `updateConfig`.'
        );
      }

      const requestPath = '/1/configs/{indexName}'.replace(
        '{indexName}',
        encodeURIComponent(indexName)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'PUT',
        path: requestPath,
        data: querySuggestionsIndexParam,
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

export type QuerySuggestionsClient = ReturnType<
  typeof createQuerySuggestionsClient
>;

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

export type DeleteConfigProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
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

export type GetConfigProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
};

export type GetConfigStatusProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
};

export type GetLogFileProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
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

export type UpdateConfigProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  querySuggestionsIndexParam: QuerySuggestionsIndexParam;
};
