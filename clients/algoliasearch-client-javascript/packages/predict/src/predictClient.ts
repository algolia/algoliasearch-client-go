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

import type { ActivateModelInstanceResponse } from '../model/activateModelInstanceResponse';
import type { ActivateModelParams } from '../model/activateModelParams';
import type {
  DelProps,
  DeleteModelInstanceProps,
  DeleteSegmentProps,
  DeleteUserProfileProps,
  FetchAllSegmentsProps,
  FetchSegmentProps,
  FetchUserProfileProps,
  GetProps,
  GetModelInstanceConfigProps,
  GetModelMetricsProps,
  GetSegmentUsersProps,
  PostProps,
  PutProps,
  UpdateModelInstanceProps,
  UpdateSegmentProps,
} from '../model/clientMethodProps';
import type { CreateSegmentParams } from '../model/createSegmentParams';
import type { CreateSegmentResponse } from '../model/createSegmentResponse';
import type { DeleteModelInstanceResponse } from '../model/deleteModelInstanceResponse';
import type { DeleteSegmentResponse } from '../model/deleteSegmentResponse';
import type { DeleteUserProfileResponse } from '../model/deleteUserProfileResponse';
import type { FetchAllUserProfilesParams } from '../model/fetchAllUserProfilesParams';
import type { FetchAllUserProfilesResponse } from '../model/fetchAllUserProfilesResponse';
import type { GetAvailableModelTypesResponseInner } from '../model/getAvailableModelTypesResponseInner';
import type { GetModelMetricsResponse } from '../model/getModelMetricsResponse';
import type { GetSegmentUsersResponse } from '../model/getSegmentUsersResponse';
import type { ModelInstance } from '../model/modelInstance';
import type { Segment } from '../model/segment';
import type { UpdateModelInstanceResponse } from '../model/updateModelInstanceResponse';
import type { UpdateSegmentResponse } from '../model/updateSegmentResponse';
import type { UserProfile } from '../model/userProfile';

export const apiClientVersion = '1.0.0-alpha.51';

export const REGIONS = ['eu', 'us'] as const;
export type Region = (typeof REGIONS)[number];

function getDefaultHosts(region: Region): Host[] {
  const url = 'predict.{region}.algolia.com'.replace('{region}', region);

  return [{ url, accept: 'readWrite', protocol: 'https' }];
}

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export function createPredictClient({
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
      client: 'Predict',
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
     * Activate an existing model template. This action triggers the training and inference pipelines for the selected model.  The model is added with `modelStatus=pending`. If a model with the exact same source & index already exists, the API endpoint returns an error.
     *
     * @summary Activate a model instance.
     * @param activateModelParams - The activateModelParams object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    activateModelInstance(
      activateModelParams: ActivateModelParams,
      requestOptions?: RequestOptions
    ): Promise<ActivateModelInstanceResponse> {
      if (!activateModelParams) {
        throw new Error(
          'Parameter `activateModelParams` is required when calling `activateModelInstance`.'
        );
      }

      if (!activateModelParams.type) {
        throw new Error(
          'Parameter `activateModelParams.type` is required when calling `activateModelInstance`.'
        );
      }
      if (!activateModelParams.name) {
        throw new Error(
          'Parameter `activateModelParams.name` is required when calling `activateModelInstance`.'
        );
      }
      if (!activateModelParams.sourceID) {
        throw new Error(
          'Parameter `activateModelParams.sourceID` is required when calling `activateModelInstance`.'
        );
      }
      if (!activateModelParams.index) {
        throw new Error(
          'Parameter `activateModelParams.index` is required when calling `activateModelInstance`.'
        );
      }

      const requestPath = '/1/predict/models';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: activateModelParams,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Create a new segment. All segments added by this endpoint will have a computed type. The endpoint receives a filters parameter, with a syntax similar to filters for Rules.
     *
     * @summary Create a segment.
     * @param createSegmentParams - The createSegmentParams object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    createSegment(
      createSegmentParams: CreateSegmentParams,
      requestOptions?: RequestOptions
    ): Promise<CreateSegmentResponse> {
      if (!createSegmentParams) {
        throw new Error(
          'Parameter `createSegmentParams` is required when calling `createSegment`.'
        );
      }

      if (!createSegmentParams.name) {
        throw new Error(
          'Parameter `createSegmentParams.name` is required when calling `createSegment`.'
        );
      }
      if (!createSegmentParams.conditions) {
        throw new Error(
          'Parameter `createSegmentParams.conditions` is required when calling `createSegment`.'
        );
      }

      const requestPath = '/1/segments';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: createSegmentParams,
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
     * Delete the model’s configuration, pipelines and generated predictions.
     *
     * @summary Delete a model instance.
     * @param deleteModelInstance - The deleteModelInstance object.
     * @param deleteModelInstance.modelID - The ID of the model to retrieve.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    deleteModelInstance(
      { modelID }: DeleteModelInstanceProps,
      requestOptions?: RequestOptions
    ): Promise<DeleteModelInstanceResponse> {
      if (!modelID) {
        throw new Error(
          'Parameter `modelID` is required when calling `deleteModelInstance`.'
        );
      }

      const requestPath = '/1/predict/models/{modelID}'.replace(
        '{modelID}',
        encodeURIComponent(modelID)
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
     * Delete the segment’s configuration. User intents (predictions) from the segment are not deleted. All segment types (computed or custom) can be deleted.  When the query is successful, the HTTP response is 200 OK and returns the date until which you can safely consider the data as being deleted.
     *
     * @summary Delete a segment\'s configuration.
     * @param deleteSegment - The deleteSegment object.
     * @param deleteSegment.segmentID - The ID of the Segment to fetch.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    deleteSegment(
      { segmentID }: DeleteSegmentProps,
      requestOptions?: RequestOptions
    ): Promise<DeleteSegmentResponse> {
      if (!segmentID) {
        throw new Error(
          'Parameter `segmentID` is required when calling `deleteSegment`.'
        );
      }

      const requestPath = '/1/segments/{segmentID}'.replace(
        '{segmentID}',
        encodeURIComponent(segmentID)
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
     * Delete all data and predictions associated with an authenticated user (userID) or an anonymous user (cookieID, sessionID).
     *
     * @summary Delete user profile.
     * @param deleteUserProfile - The deleteUserProfile object.
     * @param deleteUserProfile.userID - User ID for authenticated users or cookie ID for non-authenticated repeated users (visitors).
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    deleteUserProfile(
      { userID }: DeleteUserProfileProps,
      requestOptions?: RequestOptions
    ): Promise<DeleteUserProfileResponse> {
      if (!userID) {
        throw new Error(
          'Parameter `userID` is required when calling `deleteUserProfile`.'
        );
      }

      const requestPath = '/1/users/{userID}'.replace(
        '{userID}',
        encodeURIComponent(userID)
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
     * Get the list of segments with their configuration.
     *
     * @summary Get all segments.
     * @param fetchAllSegments - The fetchAllSegments object.
     * @param fetchAllSegments.type - The type of segments to fetch.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    fetchAllSegments(
      { type }: FetchAllSegmentsProps = {},
      requestOptions: RequestOptions | undefined = undefined
    ): Promise<Segment[]> {
      const requestPath = '/1/segments';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (type !== undefined) {
        queryParameters.type = type.toString();
      }

      const request: Request = {
        method: 'GET',
        path: requestPath,
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Get all users with predictions in the provided application.
     *
     * @summary Get all user profiles.
     * @param fetchAllUserProfilesParams - The fetchAllUserProfilesParams object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    fetchAllUserProfiles(
      fetchAllUserProfilesParams: FetchAllUserProfilesParams,
      requestOptions?: RequestOptions
    ): Promise<FetchAllUserProfilesResponse> {
      if (!fetchAllUserProfilesParams) {
        throw new Error(
          'Parameter `fetchAllUserProfilesParams` is required when calling `fetchAllUserProfiles`.'
        );
      }

      const requestPath = '/1/users';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: fetchAllUserProfilesParams,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Get the segment configuration.
     *
     * @summary Get the segment configuration.
     * @param fetchSegment - The fetchSegment object.
     * @param fetchSegment.segmentID - The ID of the Segment to fetch.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    fetchSegment(
      { segmentID }: FetchSegmentProps,
      requestOptions?: RequestOptions
    ): Promise<Segment> {
      if (!segmentID) {
        throw new Error(
          'Parameter `segmentID` is required when calling `fetchSegment`.'
        );
      }

      const requestPath = '/1/segments/{segmentID}'.replace(
        '{segmentID}',
        encodeURIComponent(segmentID)
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
     * Get predictions, properties (raw, computed or custom) and segments (computed or custom) for a user profile.
     *
     * @summary Get user profile.
     * @param fetchUserProfile - The fetchUserProfile object.
     * @param fetchUserProfile.userID - User ID for authenticated users or cookie ID for non-authenticated repeated users (visitors).
     * @param fetchUserProfile.params - The params object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    fetchUserProfile(
      { userID, params }: FetchUserProfileProps,
      requestOptions?: RequestOptions
    ): Promise<UserProfile> {
      if (!userID) {
        throw new Error(
          'Parameter `userID` is required when calling `fetchUserProfile`.'
        );
      }

      if (!params) {
        throw new Error(
          'Parameter `params` is required when calling `fetchUserProfile`.'
        );
      }

      const requestPath = '/1/users/{userID}/fetch'.replace(
        '{userID}',
        encodeURIComponent(userID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: params,
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
     * Get a list of all available model types. Each model type can be activated more than once, by selecting a different data source.
     *
     * @summary Get a list of available model types.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getAvailableModelTypes(
      requestOptions?: RequestOptions
    ): Promise<GetAvailableModelTypesResponseInner[]> {
      const requestPath = '/1/predict/modeltypes';
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
     * Get the configuration for a model that was activated.
     *
     * @summary Get a model’s instance configuration.
     * @param getModelInstanceConfig - The getModelInstanceConfig object.
     * @param getModelInstanceConfig.modelID - The ID of the model to retrieve.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getModelInstanceConfig(
      { modelID }: GetModelInstanceConfigProps,
      requestOptions?: RequestOptions
    ): Promise<ModelInstance> {
      if (!modelID) {
        throw new Error(
          'Parameter `modelID` is required when calling `getModelInstanceConfig`.'
        );
      }

      const requestPath = '/1/predict/models/{modelID}'.replace(
        '{modelID}',
        encodeURIComponent(modelID)
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
     * Get a list of all model instances.
     *
     * @summary Get model instances.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getModelInstances(
      requestOptions?: RequestOptions
    ): Promise<ModelInstance[]> {
      const requestPath = '/1/predict/models';
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
     * Get the model instance’ training metrics.
     *
     * @summary Get a model’s instance metrics.
     * @param getModelMetrics - The getModelMetrics object.
     * @param getModelMetrics.modelID - The ID of the model to retrieve.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getModelMetrics(
      { modelID }: GetModelMetricsProps,
      requestOptions?: RequestOptions
    ): Promise<GetModelMetricsResponse> {
      if (!modelID) {
        throw new Error(
          'Parameter `modelID` is required when calling `getModelMetrics`.'
        );
      }

      const requestPath = '/1/predict/models/{modelID}/metrics'.replace(
        '{modelID}',
        encodeURIComponent(modelID)
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
     * Get the profiles of users that belong to a segment.
     *
     * @summary Get segment users.
     * @param getSegmentUsers - The getSegmentUsers object.
     * @param getSegmentUsers.segmentID - The ID of the Segment to fetch.
     * @param getSegmentUsers.fetchAllUserProfilesParams - The fetchAllUserProfilesParams object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getSegmentUsers(
      { segmentID, fetchAllUserProfilesParams }: GetSegmentUsersProps,
      requestOptions?: RequestOptions
    ): Promise<GetSegmentUsersResponse> {
      if (!segmentID) {
        throw new Error(
          'Parameter `segmentID` is required when calling `getSegmentUsers`.'
        );
      }

      if (!fetchAllUserProfilesParams) {
        throw new Error(
          'Parameter `fetchAllUserProfilesParams` is required when calling `getSegmentUsers`.'
        );
      }

      const requestPath = '/1/segments/{segmentID}/users'.replace(
        '{segmentID}',
        encodeURIComponent(segmentID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: fetchAllUserProfilesParams,
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
     * Update a model’s configuration.
     *
     * @summary Update a model instance.
     * @param updateModelInstance - The updateModelInstance object.
     * @param updateModelInstance.modelID - The ID of the model to retrieve.
     * @param updateModelInstance.updateModelParams - The updateModelParams object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    updateModelInstance(
      { modelID, updateModelParams }: UpdateModelInstanceProps,
      requestOptions?: RequestOptions
    ): Promise<UpdateModelInstanceResponse> {
      if (!modelID) {
        throw new Error(
          'Parameter `modelID` is required when calling `updateModelInstance`.'
        );
      }

      if (!updateModelParams) {
        throw new Error(
          'Parameter `updateModelParams` is required when calling `updateModelInstance`.'
        );
      }

      const requestPath = '/1/predict/models/{modelID}'.replace(
        '{modelID}',
        encodeURIComponent(modelID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: updateModelParams,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Update a segment’s configuration.
     *
     * @summary Update segment.
     * @param updateSegment - The updateSegment object.
     * @param updateSegment.segmentID - The ID of the Segment to fetch.
     * @param updateSegment.updateSegmentParams - The updateSegmentParams object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    updateSegment(
      { segmentID, updateSegmentParams }: UpdateSegmentProps,
      requestOptions?: RequestOptions
    ): Promise<UpdateSegmentResponse> {
      if (!segmentID) {
        throw new Error(
          'Parameter `segmentID` is required when calling `updateSegment`.'
        );
      }

      if (!updateSegmentParams) {
        throw new Error(
          'Parameter `updateSegmentParams` is required when calling `updateSegment`.'
        );
      }

      const requestPath = '/1/segments/{segmentID}'.replace(
        '{segmentID}',
        encodeURIComponent(segmentID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: updateSegmentParams,
      };

      return transporter.request(request, requestOptions);
    },
  };
}

/**
 * The client type.
 */
export type PredictClient = ReturnType<typeof createPredictClient>;
