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
  DeleteUserProfileProps,
  GetProps,
  GetUserTokenProfileProps,
  PostProps,
  PutProps,
} from '../model/clientMethodProps';
import type { DeleteUserProfileResponse } from '../model/deleteUserProfileResponse';
import type { GetUserTokenResponse } from '../model/getUserTokenResponse';
import type { PersonalizationStrategyParams } from '../model/personalizationStrategyParams';
import type { SetPersonalizationStrategyResponse } from '../model/setPersonalizationStrategyResponse';

export const apiClientVersion = '5.0.0-alpha.10';

export const REGIONS = ['eu', 'us'] as const;
export type Region = typeof REGIONS[number];

function getDefaultHosts(region: Region): Host[] {
  const url = 'personalization.{region}.algolia.com'.replace(
    '{region}',
    region
  );

  return [{ url, accept: 'readWrite', protocol: 'https' }];
}

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export function createPersonalizationClient({
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
      client: 'Personalization',
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
     * Delete the user profile and all its associated data.  Returns, as part of the response, a date until which the data can safely be considered as deleted for the given user. This means if you send events for the given user before this date, they will be ignored. Any data received after the deletedUntil date will start building a new user profile.  It might take a couple hours for the deletion request to be fully processed.
     *
     * @summary Delete a user profile.
     * @param deleteUserProfile - The deleteUserProfile object.
     * @param deleteUserProfile.userToken - UserToken representing the user for which to fetch the Personalization profile.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    deleteUserProfile(
      { userToken }: DeleteUserProfileProps,
      requestOptions?: RequestOptions
    ): Promise<DeleteUserProfileResponse> {
      if (!userToken) {
        throw new Error(
          'Parameter `userToken` is required when calling `deleteUserProfile`.'
        );
      }

      const requestPath = '/1/profiles/{userToken}'.replace(
        '{userToken}',
        encodeURIComponent(userToken)
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
     * The strategy contains information on the events and facets that impact user profiles and personalized search results.
     *
     * @summary Get the current strategy.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getPersonalizationStrategy(
      requestOptions?: RequestOptions
    ): Promise<PersonalizationStrategyParams> {
      const requestPath = '/1/strategies/personalization';
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
     * Get the user profile built from Personalization strategy.  The profile is structured by facet name used in the strategy. Each facet value is mapped to its score. Each score represents the user affinity for a specific facet value given the userToken past events and the Personalization strategy defined. Scores are bounded to 20. The last processed event timestamp is provided using the ISO 8601 format for debugging purposes.
     *
     * @summary Get a user profile.
     * @param getUserTokenProfile - The getUserTokenProfile object.
     * @param getUserTokenProfile.userToken - UserToken representing the user for which to fetch the Personalization profile.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getUserTokenProfile(
      { userToken }: GetUserTokenProfileProps,
      requestOptions?: RequestOptions
    ): Promise<GetUserTokenResponse> {
      if (!userToken) {
        throw new Error(
          'Parameter `userToken` is required when calling `getUserTokenProfile`.'
        );
      }

      const requestPath = '/1/profiles/personalization/{userToken}'.replace(
        '{userToken}',
        encodeURIComponent(userToken)
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
     * A strategy defines the events and facets that impact user profiles and personalized search results.
     *
     * @summary Set a new strategy.
     * @param personalizationStrategyParams - The personalizationStrategyParams object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    setPersonalizationStrategy(
      personalizationStrategyParams: PersonalizationStrategyParams,
      requestOptions?: RequestOptions
    ): Promise<SetPersonalizationStrategyResponse> {
      if (!personalizationStrategyParams) {
        throw new Error(
          'Parameter `personalizationStrategyParams` is required when calling `setPersonalizationStrategy`.'
        );
      }

      if (!personalizationStrategyParams.eventScoring) {
        throw new Error(
          'Parameter `personalizationStrategyParams.eventScoring` is required when calling `setPersonalizationStrategy`.'
        );
      }
      if (!personalizationStrategyParams.facetScoring) {
        throw new Error(
          'Parameter `personalizationStrategyParams.facetScoring` is required when calling `setPersonalizationStrategy`.'
        );
      }
      if (!personalizationStrategyParams.personalizationImpact) {
        throw new Error(
          'Parameter `personalizationStrategyParams.personalizationImpact` is required when calling `setPersonalizationStrategy`.'
        );
      }

      const requestPath = '/1/strategies/personalization';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: personalizationStrategyParams,
      };

      return transporter.request(request, requestOptions);
    },
  };
}

/**
 * The client type.
 */
export type PersonalizationClient = ReturnType<
  typeof createPersonalizationClient
>;
