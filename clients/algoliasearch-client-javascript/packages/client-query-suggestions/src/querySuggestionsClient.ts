// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import {
  createAuth,
  createTransporter,
  getAlgoliaAgent,
} from '@algolia/client-common';
import type {
  CreateClientOptions,
  Headers,
  Host,
  Request,
  RequestOptions,
  QueryParameters,
} from '@algolia/client-common';

import type {
  DelProps,
  DeleteConfigProps,
  GetProps,
  GetConfigProps,
  GetConfigStatusProps,
  GetLogFileProps,
  PostProps,
  PutProps,
  UpdateConfigProps,
} from '../model/clientMethodProps';
import type { LogFile } from '../model/logFile';
import type { QuerySuggestionsIndex } from '../model/querySuggestionsIndex';
import type { QuerySuggestionsIndexWithIndexParam } from '../model/querySuggestionsIndexWithIndexParam';
import type { Status } from '../model/status';
import type { SuccessResponse } from '../model/successResponse';

export const apiClientVersion = '5.0.0-alpha.42';

export const REGIONS = ['eu', 'us'] as const;
export type Region = typeof REGIONS[number];

function getDefaultHosts(region: Region): Host[] {
  const url = 'query-suggestions.{region}.algolia.com'.replace(
    '{region}',
    region
  );

  return [{ url, accept: 'readWrite', protocol: 'https' }];
}

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export function createQuerySuggestionsClient({
  appId: appIdOption,
  apiKey: apiKeyOption,
  authMode,
  algoliaAgents,
  region: regionOption,
  ...options
}: CreateClientOptions & { region: Region }) {
  const auth = createAuth(appIdOption, apiKeyOption, authMode);
  const transporter = createTransporter({
    hosts: getDefaultHosts(regionOption),
    ...options,
    algoliaAgent: getAlgoliaAgent({
      algoliaAgents,
      client: 'QuerySuggestions',
      version: apiClientVersion,
    }),
    baseHeaders: {
      'content-type': 'text/plain',
      ...auth.headers(),
      ...options.baseHeaders,
    },
    baseQueryParameters: {
      ...auth.queryParameters(),
      ...options.baseQueryParameters,
    },
  });

  return {
    transporter,

    /**
     * The `appId` currently in use.
     */
    appId: appIdOption,

    /**
     * Clears the cache of the transporter for the `requestsCache` and `responsesCache` properties.
     */
    clearCache(): Promise<void> {
      return Promise.all([
        transporter.requestsCache.clear(),
        transporter.responsesCache.clear(),
      ]).then(() => undefined);
    },

    /**
     * Get the value of the `algoliaAgent`, used by our libraries internally and telemetry system.
     */
    get _ua(): string {
      return transporter.algoliaAgent.value;
    },

    /**
     * Adds a `segment` to the `x-algolia-agent` sent with every requests.
     *
     * @param segment - The algolia agent (user-agent) segment to add.
     * @param version - The version of the agent.
     */
    addAlgoliaAgent(segment: string, version?: string): void {
      transporter.algoliaAgent.add({ segment, version });
    },

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
    ): Promise<SuccessResponse> {
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
        queryParameters,
        headers,
        data: querySuggestionsIndexWithIndexParam,
      };

      return transporter.request(request, requestOptions);
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
      const queryParameters: QueryParameters = parameters ? parameters : {};

      const request: Request = {
        method: 'DELETE',
        path: requestPath,
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Delete a configuration of a Query Suggestion\'s index. By deleting a configuration, you stop all updates to the underlying query suggestion index. Note that when doing this, the underlying index does not change - existing suggestions remain untouched.
     *
     * @summary Delete a configuration.
     * @param deleteConfig - The deleteConfig object.
     * @param deleteConfig.indexName - The index in which to perform the request.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    deleteConfig(
      { indexName }: DeleteConfigProps,
      requestOptions?: RequestOptions
    ): Promise<SuccessResponse> {
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
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
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
      const queryParameters: QueryParameters = parameters ? parameters : {};

      const request: Request = {
        method: 'GET',
        path: requestPath,
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
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
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
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
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
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
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
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
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
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
      const queryParameters: QueryParameters = parameters ? parameters : {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: body ? body : {},
      };

      return transporter.request(request, requestOptions);
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
      const queryParameters: QueryParameters = parameters ? parameters : {};

      const request: Request = {
        method: 'PUT',
        path: requestPath,
        queryParameters,
        headers,
        data: body ? body : {},
      };

      return transporter.request(request, requestOptions);
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
    ): Promise<SuccessResponse> {
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
        queryParameters,
        headers,
        data: querySuggestionsIndexParam,
      };

      return transporter.request(request, requestOptions);
    },
  };
}

/**
 * The client type.
 */
export type QuerySuggestionsClient = ReturnType<
  typeof createQuerySuggestionsClient
>;
