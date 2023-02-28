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

import type { Authentication } from '../model/authentication';
import type { AuthenticationCreate } from '../model/authenticationCreate';
import type { AuthenticationCreateResponse } from '../model/authenticationCreateResponse';
import type { AuthenticationSearch } from '../model/authenticationSearch';
import type { AuthenticationUpdateResponse } from '../model/authenticationUpdateResponse';
import type { AuthenticationWithInput } from '../model/authenticationWithInput';
import type {
  DeleteAuthenticationProps,
  DeleteDestinationProps,
  DeleteSourceProps,
  DeleteTaskProps,
  DisableTaskProps,
  EnableTaskProps,
  GetAuthenticationProps,
  GetAuthenticationsProps,
  GetDestinationProps,
  GetDestinationsProps,
  GetEventProps,
  GetEventsProps,
  GetRunProps,
  GetRunsProps,
  GetSourceProps,
  GetSourcesProps,
  GetTaskProps,
  GetTasksProps,
  RunTaskProps,
  UpdateAuthenticationProps,
  UpdateDestinationProps,
  UpdateSourceProps,
  UpdateTaskProps,
} from '../model/clientMethodProps';
import type { DeleteResponse } from '../model/deleteResponse';
import type { Destination } from '../model/destination';
import type { DestinationCreate } from '../model/destinationCreate';
import type { DestinationCreateResponse } from '../model/destinationCreateResponse';
import type { DestinationSearch } from '../model/destinationSearch';
import type { DestinationUpdateResponse } from '../model/destinationUpdateResponse';
import type { Event } from '../model/event';
import type { ListAuthenticationsResponse } from '../model/listAuthenticationsResponse';
import type { ListDestinationsResponse } from '../model/listDestinationsResponse';
import type { ListEventsResponse } from '../model/listEventsResponse';
import type { ListSourcesResponse } from '../model/listSourcesResponse';
import type { ListTasksResponse } from '../model/listTasksResponse';
import type { OnDemandTrigger } from '../model/onDemandTrigger';
import type { Run } from '../model/run';
import type { RunListResponse } from '../model/runListResponse';
import type { RunResponse } from '../model/runResponse';
import type { ScheduleTrigger } from '../model/scheduleTrigger';
import type { Source } from '../model/source';
import type { SourceCreate } from '../model/sourceCreate';
import type { SourceCreateResponse } from '../model/sourceCreateResponse';
import type { SourceSearch } from '../model/sourceSearch';
import type { SourceUpdateResponse } from '../model/sourceUpdateResponse';
import type { SubscriptionTrigger } from '../model/subscriptionTrigger';
import type { Task } from '../model/task';
import type { TaskCreate } from '../model/taskCreate';
import type { TaskCreateResponse } from '../model/taskCreateResponse';
import type { TaskSearch } from '../model/taskSearch';
import type { TaskUpdateResponse } from '../model/taskUpdateResponse';
import type { Trigger } from '../model/trigger';

export const apiClientVersion = '1.0.0-alpha.24';

export const REGIONS = ['us'] as const;
export type Region = typeof REGIONS[number];

function getDefaultHosts(region: Region): Host[] {
  const url = 'data.{region}.algolia.com'.replace('{region}', region);

  return [{ url, accept: 'readWrite', protocol: 'https' }];
}

/**
 * Guard: Return strongly typed specific OnDemandTrigger for a given Trigger.
 *
 * @summary Guard method that returns a strongly typed specific OnDemandTrigger for a given Trigger.
 * @param trigger - The given Task Trigger.
 */
export function isOnDemandTrigger(
  trigger: Trigger
): trigger is OnDemandTrigger {
  return trigger.type === 'onDemand';
}

/**
 * Guard: Return strongly typed specific ScheduleTrigger for a given Trigger.
 *
 * @summary Guard method that returns a strongly typed specific ScheduleTrigger for a given Trigger.
 * @param trigger - The given Task Trigger.
 */
export function isScheduleTrigger(
  trigger: Trigger
): trigger is ScheduleTrigger {
  return trigger.type === 'schedule';
}

/**
 * Guard: Return strongly typed specific SubscriptionTrigger for a given Trigger.
 *
 * @summary Guard method that returns a strongly typed specific SubscriptionTrigger for a given Trigger.
 * @param trigger - The given Task Trigger.
 */
export function isSubscriptionTrigger(
  trigger: Trigger
): trigger is SubscriptionTrigger {
  return trigger.type === 'subscription';
}

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export function createIngestionClient({
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
    algoliaAgent: {
      ...getAlgoliaAgent({
        algoliaAgents,
        client: 'Ingestion',
        version: apiClientVersion,
      }),
      value: '',
    },
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
     * Create a authentication.
     *
     * @summary Create a authentication.
     * @param authenticationCreate -.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    createAuthentication(
      authenticationCreate: AuthenticationCreate,
      requestOptions?: RequestOptions
    ): Promise<AuthenticationCreateResponse> {
      if (!authenticationCreate) {
        throw new Error(
          'Parameter `authenticationCreate` is required when calling `createAuthentication`.'
        );
      }

      const requestPath = '/1/authentications';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: authenticationCreate,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Create a destination.
     *
     * @summary Create a destination.
     * @param destinationCreate -.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    createDestination(
      destinationCreate: DestinationCreate,
      requestOptions?: RequestOptions
    ): Promise<DestinationCreateResponse> {
      if (!destinationCreate) {
        throw new Error(
          'Parameter `destinationCreate` is required when calling `createDestination`.'
        );
      }

      const requestPath = '/1/destinations';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: destinationCreate,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Create a source.
     *
     * @summary Create a source.
     * @param sourceCreate -.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    createSource(
      sourceCreate: SourceCreate,
      requestOptions?: RequestOptions
    ): Promise<SourceCreateResponse> {
      if (!sourceCreate) {
        throw new Error(
          'Parameter `sourceCreate` is required when calling `createSource`.'
        );
      }

      if (!sourceCreate.type) {
        throw new Error(
          'Parameter `sourceCreate.type` is required when calling `createSource`.'
        );
      }
      if (!sourceCreate.input) {
        throw new Error(
          'Parameter `sourceCreate.input` is required when calling `createSource`.'
        );
      }

      const requestPath = '/1/sources';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: sourceCreate,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Create a task.
     *
     * @summary Create a task.
     * @param taskCreate -.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    createTask(
      taskCreate: TaskCreate,
      requestOptions?: RequestOptions
    ): Promise<TaskCreateResponse> {
      if (!taskCreate) {
        throw new Error(
          'Parameter `taskCreate` is required when calling `createTask`.'
        );
      }

      if (!taskCreate.sourceID) {
        throw new Error(
          'Parameter `taskCreate.sourceID` is required when calling `createTask`.'
        );
      }
      if (!taskCreate.trigger) {
        throw new Error(
          'Parameter `taskCreate.trigger` is required when calling `createTask`.'
        );
      }
      if (!taskCreate.action) {
        throw new Error(
          'Parameter `taskCreate.action` is required when calling `createTask`.'
        );
      }

      const requestPath = '/1/tasks';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: taskCreate,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Soft delete the authentication of the given authenticationID.
     *
     * @summary Delete a authentication.
     * @param deleteAuthentication - The deleteAuthentication object.
     * @param deleteAuthentication.authenticationID - The authentication UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    deleteAuthentication(
      { authenticationID }: DeleteAuthenticationProps,
      requestOptions?: RequestOptions
    ): Promise<DeleteResponse> {
      if (!authenticationID) {
        throw new Error(
          'Parameter `authenticationID` is required when calling `deleteAuthentication`.'
        );
      }

      const requestPath = '/1/authentications/{authenticationID}'.replace(
        '{authenticationID}',
        encodeURIComponent(authenticationID)
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
     * Soft delete the destination of the given destinationID.
     *
     * @summary Delete a destination.
     * @param deleteDestination - The deleteDestination object.
     * @param deleteDestination.destinationID - The destination UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    deleteDestination(
      { destinationID }: DeleteDestinationProps,
      requestOptions?: RequestOptions
    ): Promise<DeleteResponse> {
      if (!destinationID) {
        throw new Error(
          'Parameter `destinationID` is required when calling `deleteDestination`.'
        );
      }

      const requestPath = '/1/destinations/{destinationID}'.replace(
        '{destinationID}',
        encodeURIComponent(destinationID)
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
     * Soft delete the source of the given sourceID.
     *
     * @summary Delete a source.
     * @param deleteSource - The deleteSource object.
     * @param deleteSource.sourceID - The source UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    deleteSource(
      { sourceID }: DeleteSourceProps,
      requestOptions?: RequestOptions
    ): Promise<DeleteResponse> {
      if (!sourceID) {
        throw new Error(
          'Parameter `sourceID` is required when calling `deleteSource`.'
        );
      }

      const requestPath = '/1/sources/{sourceID}'.replace(
        '{sourceID}',
        encodeURIComponent(sourceID)
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
     * Soft delete the task of the given taskID.
     *
     * @summary Delete a task.
     * @param deleteTask - The deleteTask object.
     * @param deleteTask.taskID - The task UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    deleteTask(
      { taskID }: DeleteTaskProps,
      requestOptions?: RequestOptions
    ): Promise<DeleteResponse> {
      if (!taskID) {
        throw new Error(
          'Parameter `taskID` is required when calling `deleteTask`.'
        );
      }

      const requestPath = '/1/tasks/{taskID}'.replace(
        '{taskID}',
        encodeURIComponent(taskID)
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
     * Disable the task of the given taskID.
     *
     * @summary Disable a task.
     * @param disableTask - The disableTask object.
     * @param disableTask.taskID - The task UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    disableTask(
      { taskID }: DisableTaskProps,
      requestOptions?: RequestOptions
    ): Promise<TaskUpdateResponse> {
      if (!taskID) {
        throw new Error(
          'Parameter `taskID` is required when calling `disableTask`.'
        );
      }

      const requestPath = '/1/tasks/{taskID}/disable'.replace(
        '{taskID}',
        encodeURIComponent(taskID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'PUT',
        path: requestPath,
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Enable the task of the given taskID.
     *
     * @summary Enable a task.
     * @param enableTask - The enableTask object.
     * @param enableTask.taskID - The task UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    enableTask(
      { taskID }: EnableTaskProps,
      requestOptions?: RequestOptions
    ): Promise<TaskUpdateResponse> {
      if (!taskID) {
        throw new Error(
          'Parameter `taskID` is required when calling `enableTask`.'
        );
      }

      const requestPath = '/1/tasks/{taskID}/enable'.replace(
        '{taskID}',
        encodeURIComponent(taskID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'PUT',
        path: requestPath,
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Get the authentication of the given authenticationID.
     *
     * @summary Get a authentication.
     * @param getAuthentication - The getAuthentication object.
     * @param getAuthentication.authenticationID - The authentication UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getAuthentication(
      { authenticationID }: GetAuthenticationProps,
      requestOptions?: RequestOptions
    ): Promise<AuthenticationWithInput> {
      if (!authenticationID) {
        throw new Error(
          'Parameter `authenticationID` is required when calling `getAuthentication`.'
        );
      }

      const requestPath = '/1/authentications/{authenticationID}'.replace(
        '{authenticationID}',
        encodeURIComponent(authenticationID)
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
     * Get a list of authentications for the given query parameters, with pagination details.
     *
     * @summary Get a list of authentications.
     * @param getAuthentications - The getAuthentications object.
     * @param getAuthentications.itemsPerPage - The number of items per page to return.
     * @param getAuthentications.page - The page number to fetch, starting at 1.
     * @param getAuthentications.type - The type of the authentications to retrieve.
     * @param getAuthentications.platform - The platform of the authentications to retrieve.
     * @param getAuthentications.sort - The key by which the list should be sorted.
     * @param getAuthentications.order - The order of the returned list.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getAuthentications(
      {
        itemsPerPage,
        page,
        type,
        platform,
        sort,
        order,
      }: GetAuthenticationsProps = {},
      requestOptions: RequestOptions | undefined = undefined
    ): Promise<ListAuthenticationsResponse> {
      const requestPath = '/1/authentications';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (itemsPerPage !== undefined) {
        queryParameters.itemsPerPage = itemsPerPage.toString();
      }

      if (page !== undefined) {
        queryParameters.page = page.toString();
      }

      if (type !== undefined) {
        queryParameters.type = type.toString();
      }

      if (platform !== undefined) {
        queryParameters.platform = platform.toString();
      }

      if (sort !== undefined) {
        queryParameters.sort = sort.toString();
      }

      if (order !== undefined) {
        queryParameters.order = order.toString();
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
     * Get the destination of the given destinationID.
     *
     * @summary Get a destination.
     * @param getDestination - The getDestination object.
     * @param getDestination.destinationID - The destination UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getDestination(
      { destinationID }: GetDestinationProps,
      requestOptions?: RequestOptions
    ): Promise<Destination> {
      if (!destinationID) {
        throw new Error(
          'Parameter `destinationID` is required when calling `getDestination`.'
        );
      }

      const requestPath = '/1/destinations/{destinationID}'.replace(
        '{destinationID}',
        encodeURIComponent(destinationID)
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
     * Get a list of destinations for the given query parameters, with pagination details.
     *
     * @summary Get a list of destinations.
     * @param getDestinations - The getDestinations object.
     * @param getDestinations.itemsPerPage - The number of items per page to return.
     * @param getDestinations.page - The page number to fetch, starting at 1.
     * @param getDestinations.type - The type of the destinations to retrive.
     * @param getDestinations.authenticationID - The authenticationIDs of the destinations to retrive.
     * @param getDestinations.sort - The key by which the list should be sorted.
     * @param getDestinations.order - The order of the returned list.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getDestinations(
      {
        itemsPerPage,
        page,
        type,
        authenticationID,
        sort,
        order,
      }: GetDestinationsProps = {},
      requestOptions: RequestOptions | undefined = undefined
    ): Promise<ListDestinationsResponse> {
      const requestPath = '/1/destinations';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (itemsPerPage !== undefined) {
        queryParameters.itemsPerPage = itemsPerPage.toString();
      }

      if (page !== undefined) {
        queryParameters.page = page.toString();
      }

      if (type !== undefined) {
        queryParameters.type = type.toString();
      }

      if (authenticationID !== undefined) {
        queryParameters.authenticationID = authenticationID.toString();
      }

      if (sort !== undefined) {
        queryParameters.sort = sort.toString();
      }

      if (order !== undefined) {
        queryParameters.order = order.toString();
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
     * Get a single event for a specific runID.
     *
     * @summary Get an event.
     * @param getEvent - The getEvent object.
     * @param getEvent.runID - The run UUID.
     * @param getEvent.eventID - The event UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getEvent(
      { runID, eventID }: GetEventProps,
      requestOptions?: RequestOptions
    ): Promise<Event> {
      if (!runID) {
        throw new Error(
          'Parameter `runID` is required when calling `getEvent`.'
        );
      }

      if (!eventID) {
        throw new Error(
          'Parameter `eventID` is required when calling `getEvent`.'
        );
      }

      const requestPath = '/1/runs/{runID}/events/{eventID}'
        .replace('{runID}', encodeURIComponent(runID))
        .replace('{eventID}', encodeURIComponent(eventID));
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
     * Get a list of events associated to the given runID, for the given query parameters.
     *
     * @summary Get a list of events.
     * @param getEvents - The getEvents object.
     * @param getEvents.runID - The run UUID.
     * @param getEvents.itemsPerPage - The number of items per page to return.
     * @param getEvents.page - The page number to fetch, starting at 1.
     * @param getEvents.status - Filter the status of the events.
     * @param getEvents.type - Filter the type of the events.
     * @param getEvents.sort - The key by which the list should be sorted.
     * @param getEvents.order - The order of the returned list.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getEvents(
      { runID, itemsPerPage, page, status, type, sort, order }: GetEventsProps,
      requestOptions?: RequestOptions
    ): Promise<ListEventsResponse> {
      if (!runID) {
        throw new Error(
          'Parameter `runID` is required when calling `getEvents`.'
        );
      }

      const requestPath = '/1/runs/{runID}/events'.replace(
        '{runID}',
        encodeURIComponent(runID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (itemsPerPage !== undefined) {
        queryParameters.itemsPerPage = itemsPerPage.toString();
      }

      if (page !== undefined) {
        queryParameters.page = page.toString();
      }

      if (status !== undefined) {
        queryParameters.status = status.toString();
      }

      if (type !== undefined) {
        queryParameters.type = type.toString();
      }

      if (sort !== undefined) {
        queryParameters.sort = sort.toString();
      }

      if (order !== undefined) {
        queryParameters.order = order.toString();
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
     * Get a single run for the given ID.
     *
     * @summary Get a run.
     * @param getRun - The getRun object.
     * @param getRun.runID - The run UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getRun(
      { runID }: GetRunProps,
      requestOptions?: RequestOptions
    ): Promise<Run> {
      if (!runID) {
        throw new Error('Parameter `runID` is required when calling `getRun`.');
      }

      const requestPath = '/1/runs/{runID}'.replace(
        '{runID}',
        encodeURIComponent(runID)
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
     * Get a list of runs for the given query parameters, with pagination details.
     *
     * @summary Get a list of runs.
     * @param getRuns - The getRuns object.
     * @param getRuns.itemsPerPage - The number of items per page to return.
     * @param getRuns.page - The page number to fetch, starting at 1.
     * @param getRuns.status - Filter the status of the runs.
     * @param getRuns.taskID - Filter by taskID.
     * @param getRuns.sort - The key by which the list should be sorted.
     * @param getRuns.order - The order of the returned list.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getRuns(
      { itemsPerPage, page, status, taskID, sort, order }: GetRunsProps = {},
      requestOptions: RequestOptions | undefined = undefined
    ): Promise<RunListResponse> {
      const requestPath = '/1/runs';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (itemsPerPage !== undefined) {
        queryParameters.itemsPerPage = itemsPerPage.toString();
      }

      if (page !== undefined) {
        queryParameters.page = page.toString();
      }

      if (status !== undefined) {
        queryParameters.status = status.toString();
      }

      if (taskID !== undefined) {
        queryParameters.taskID = taskID.toString();
      }

      if (sort !== undefined) {
        queryParameters.sort = sort.toString();
      }

      if (order !== undefined) {
        queryParameters.order = order.toString();
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
     * Get the source of the given sourceID.
     *
     * @summary Get a source.
     * @param getSource - The getSource object.
     * @param getSource.sourceID - The source UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getSource(
      { sourceID }: GetSourceProps,
      requestOptions?: RequestOptions
    ): Promise<Source> {
      if (!sourceID) {
        throw new Error(
          'Parameter `sourceID` is required when calling `getSource`.'
        );
      }

      const requestPath = '/1/sources/{sourceID}'.replace(
        '{sourceID}',
        encodeURIComponent(sourceID)
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
     * Get a list of sources for the given query parameters, with pagination details.
     *
     * @summary Get a list of sources.
     * @param getSources - The getSources object.
     * @param getSources.itemsPerPage - The number of items per page to return.
     * @param getSources.page - The page number to fetch, starting at 1.
     * @param getSources.type - The type of the sources to retrieve.
     * @param getSources.authenticationID - The authenticationIDs of the sources to retrieve. \'none\' returns sources that doesn\'t have an authentication.
     * @param getSources.sort - The key by which the list should be sorted.
     * @param getSources.order - The order of the returned list.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getSources(
      {
        itemsPerPage,
        page,
        type,
        authenticationID,
        sort,
        order,
      }: GetSourcesProps = {},
      requestOptions: RequestOptions | undefined = undefined
    ): Promise<ListSourcesResponse> {
      const requestPath = '/1/sources';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (itemsPerPage !== undefined) {
        queryParameters.itemsPerPage = itemsPerPage.toString();
      }

      if (page !== undefined) {
        queryParameters.page = page.toString();
      }

      if (type !== undefined) {
        queryParameters.type = type.toString();
      }

      if (authenticationID !== undefined) {
        queryParameters.authenticationID = authenticationID.toString();
      }

      if (sort !== undefined) {
        queryParameters.sort = sort.toString();
      }

      if (order !== undefined) {
        queryParameters.order = order.toString();
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
     * Get the task of the given taskID.
     *
     * @summary Get a task.
     * @param getTask - The getTask object.
     * @param getTask.taskID - The task UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getTask(
      { taskID }: GetTaskProps,
      requestOptions?: RequestOptions
    ): Promise<Task> {
      if (!taskID) {
        throw new Error(
          'Parameter `taskID` is required when calling `getTask`.'
        );
      }

      const requestPath = '/1/tasks/{taskID}'.replace(
        '{taskID}',
        encodeURIComponent(taskID)
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
     * Get a list of tasks for the given query parameters, with pagination details.
     *
     * @summary Get a list of tasks.
     * @param getTasks - The getTasks object.
     * @param getTasks.itemsPerPage - The number of items per page to return.
     * @param getTasks.page - The page number to fetch, starting at 1.
     * @param getTasks.action - The action of the tasks to retrieve.
     * @param getTasks.enabled - Whether the task is enabled or not.
     * @param getTasks.sourceID - The sourceIDs of the tasks to retrive.
     * @param getTasks.destinationID - The destinationIDs of the tasks to retrive.
     * @param getTasks.triggerType - The trigger type of the task.
     * @param getTasks.sort - The key by which the list should be sorted.
     * @param getTasks.order - The order of the returned list.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getTasks(
      {
        itemsPerPage,
        page,
        action,
        enabled,
        sourceID,
        destinationID,
        triggerType,
        sort,
        order,
      }: GetTasksProps = {},
      requestOptions: RequestOptions | undefined = undefined
    ): Promise<ListTasksResponse> {
      const requestPath = '/1/tasks';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (itemsPerPage !== undefined) {
        queryParameters.itemsPerPage = itemsPerPage.toString();
      }

      if (page !== undefined) {
        queryParameters.page = page.toString();
      }

      if (action !== undefined) {
        queryParameters.action = action.toString();
      }

      if (enabled !== undefined) {
        queryParameters.enabled = enabled.toString();
      }

      if (sourceID !== undefined) {
        queryParameters.sourceID = sourceID.toString();
      }

      if (destinationID !== undefined) {
        queryParameters.destinationID = destinationID.toString();
      }

      if (triggerType !== undefined) {
        queryParameters.triggerType = triggerType.toString();
      }

      if (sort !== undefined) {
        queryParameters.sort = sort.toString();
      }

      if (order !== undefined) {
        queryParameters.order = order.toString();
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
     * Run the task of the given taskID.
     *
     * @summary Run a task.
     * @param runTask - The runTask object.
     * @param runTask.taskID - The task UUID.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    runTask(
      { taskID }: RunTaskProps,
      requestOptions?: RequestOptions
    ): Promise<RunResponse> {
      if (!taskID) {
        throw new Error(
          'Parameter `taskID` is required when calling `runTask`.'
        );
      }

      const requestPath = '/1/tasks/{taskID}/run'.replace(
        '{taskID}',
        encodeURIComponent(taskID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Search among authentications with a defined set of parameters.
     *
     * @summary Search among authentications.
     * @param authenticationSearch - The authenticationSearch object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    searchAuthentications(
      authenticationSearch: AuthenticationSearch,
      requestOptions?: RequestOptions
    ): Promise<Authentication[]> {
      if (!authenticationSearch) {
        throw new Error(
          'Parameter `authenticationSearch` is required when calling `searchAuthentications`.'
        );
      }

      if (!authenticationSearch.authenticationIDs) {
        throw new Error(
          'Parameter `authenticationSearch.authenticationIDs` is required when calling `searchAuthentications`.'
        );
      }

      const requestPath = '/1/authentications/search';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: authenticationSearch,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Search among destinations with a defined set of parameters.
     *
     * @summary Search among destinations.
     * @param destinationSearch - The destinationSearch object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    searchDestinations(
      destinationSearch: DestinationSearch,
      requestOptions?: RequestOptions
    ): Promise<Destination[]> {
      if (!destinationSearch) {
        throw new Error(
          'Parameter `destinationSearch` is required when calling `searchDestinations`.'
        );
      }

      if (!destinationSearch.destinationIDs) {
        throw new Error(
          'Parameter `destinationSearch.destinationIDs` is required when calling `searchDestinations`.'
        );
      }

      const requestPath = '/1/destinations/search';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: destinationSearch,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Search among sources with a defined set of parameters.
     *
     * @summary Search among sources.
     * @param sourceSearch - The sourceSearch object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    searchSources(
      sourceSearch: SourceSearch,
      requestOptions?: RequestOptions
    ): Promise<Source[]> {
      if (!sourceSearch) {
        throw new Error(
          'Parameter `sourceSearch` is required when calling `searchSources`.'
        );
      }

      if (!sourceSearch.sourceIDs) {
        throw new Error(
          'Parameter `sourceSearch.sourceIDs` is required when calling `searchSources`.'
        );
      }

      const requestPath = '/1/sources/search';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: sourceSearch,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Search among tasks with a defined set of parameters.
     *
     * @summary Search among tasks.
     * @param taskSearch - The taskSearch object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    searchTasks(
      taskSearch: TaskSearch,
      requestOptions?: RequestOptions
    ): Promise<Task[]> {
      if (!taskSearch) {
        throw new Error(
          'Parameter `taskSearch` is required when calling `searchTasks`.'
        );
      }

      if (!taskSearch.taskIDs) {
        throw new Error(
          'Parameter `taskSearch.taskIDs` is required when calling `searchTasks`.'
        );
      }

      const requestPath = '/1/tasks/search';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'POST',
        path: requestPath,
        queryParameters,
        headers,
        data: taskSearch,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Update the authentication of the given authenticationID.
     *
     * @summary Update a authentication.
     * @param updateAuthentication - The updateAuthentication object.
     * @param updateAuthentication.authenticationID - The authentication UUID.
     * @param updateAuthentication.authenticationUpdate - The authenticationUpdate object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    updateAuthentication(
      { authenticationID, authenticationUpdate }: UpdateAuthenticationProps,
      requestOptions?: RequestOptions
    ): Promise<AuthenticationUpdateResponse> {
      if (!authenticationID) {
        throw new Error(
          'Parameter `authenticationID` is required when calling `updateAuthentication`.'
        );
      }

      if (!authenticationUpdate) {
        throw new Error(
          'Parameter `authenticationUpdate` is required when calling `updateAuthentication`.'
        );
      }

      const requestPath = '/1/authentications/{authenticationID}'.replace(
        '{authenticationID}',
        encodeURIComponent(authenticationID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'PUT',
        path: requestPath,
        queryParameters,
        headers,
        data: authenticationUpdate,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Update the destination of the given destinationID.
     *
     * @summary Update a destination.
     * @param updateDestination - The updateDestination object.
     * @param updateDestination.destinationID - The destination UUID.
     * @param updateDestination.destinationUpdate - The destinationUpdate object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    updateDestination(
      { destinationID, destinationUpdate }: UpdateDestinationProps,
      requestOptions?: RequestOptions
    ): Promise<DestinationUpdateResponse> {
      if (!destinationID) {
        throw new Error(
          'Parameter `destinationID` is required when calling `updateDestination`.'
        );
      }

      if (!destinationUpdate) {
        throw new Error(
          'Parameter `destinationUpdate` is required when calling `updateDestination`.'
        );
      }

      const requestPath = '/1/destinations/{destinationID}'.replace(
        '{destinationID}',
        encodeURIComponent(destinationID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'PUT',
        path: requestPath,
        queryParameters,
        headers,
        data: destinationUpdate,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Update the source of the given sourceID.
     *
     * @summary Update a source.
     * @param updateSource - The updateSource object.
     * @param updateSource.sourceID - The source UUID.
     * @param updateSource.sourceUpdate - The sourceUpdate object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    updateSource(
      { sourceID, sourceUpdate }: UpdateSourceProps,
      requestOptions?: RequestOptions
    ): Promise<SourceUpdateResponse> {
      if (!sourceID) {
        throw new Error(
          'Parameter `sourceID` is required when calling `updateSource`.'
        );
      }

      if (!sourceUpdate) {
        throw new Error(
          'Parameter `sourceUpdate` is required when calling `updateSource`.'
        );
      }

      const requestPath = '/1/sources/{sourceID}'.replace(
        '{sourceID}',
        encodeURIComponent(sourceID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'PUT',
        path: requestPath,
        queryParameters,
        headers,
        data: sourceUpdate,
      };

      return transporter.request(request, requestOptions);
    },

    /**
     * Update the task of the given taskID.
     *
     * @summary Update a task.
     * @param updateTask - The updateTask object.
     * @param updateTask.taskID - The task UUID.
     * @param updateTask.taskUpdate - The taskUpdate object.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    updateTask(
      { taskID, taskUpdate }: UpdateTaskProps,
      requestOptions?: RequestOptions
    ): Promise<TaskUpdateResponse> {
      if (!taskID) {
        throw new Error(
          'Parameter `taskID` is required when calling `updateTask`.'
        );
      }

      if (!taskUpdate) {
        throw new Error(
          'Parameter `taskUpdate` is required when calling `updateTask`.'
        );
      }

      const requestPath = '/1/tasks/{taskID}'.replace(
        '{taskID}',
        encodeURIComponent(taskID)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      const request: Request = {
        method: 'PUT',
        path: requestPath,
        queryParameters,
        headers,
        data: taskUpdate,
      };

      return transporter.request(request, requestOptions);
    },
  };
}

/**
 * The client type.
 */
export type IngestionClient = ReturnType<typeof createIngestionClient>;
