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
  GetProps,
  GetAverageClickPositionProps,
  GetClickPositionsProps,
  GetClickThroughRateProps,
  GetConversationRateProps,
  GetNoClickRateProps,
  GetNoResultsRateProps,
  GetSearchesCountProps,
  GetSearchesNoClicksProps,
  GetSearchesNoResultsProps,
  GetStatusProps,
  GetTopCountriesProps,
  GetTopFilterAttributesProps,
  GetTopFilterForAttributeProps,
  GetTopFiltersNoResultsProps,
  GetTopHitsProps,
  GetTopSearchesProps,
  GetUsersCountProps,
  PostProps,
  PutProps,
} from '../model/clientMethodProps';
import type { GetAverageClickPositionResponse } from '../model/getAverageClickPositionResponse';
import type { GetClickPositionsResponse } from '../model/getClickPositionsResponse';
import type { GetClickThroughRateResponse } from '../model/getClickThroughRateResponse';
import type { GetConversationRateResponse } from '../model/getConversationRateResponse';
import type { GetNoClickRateResponse } from '../model/getNoClickRateResponse';
import type { GetNoResultsRateResponse } from '../model/getNoResultsRateResponse';
import type { GetSearchesCountResponse } from '../model/getSearchesCountResponse';
import type { GetSearchesNoClicksResponse } from '../model/getSearchesNoClicksResponse';
import type { GetSearchesNoResultsResponse } from '../model/getSearchesNoResultsResponse';
import type { GetStatusResponse } from '../model/getStatusResponse';
import type { GetTopCountriesResponse } from '../model/getTopCountriesResponse';
import type { GetTopFilterAttributesResponse } from '../model/getTopFilterAttributesResponse';
import type { GetTopFilterForAttributeResponse } from '../model/getTopFilterForAttributeResponse';
import type { GetTopFiltersNoResultsResponse } from '../model/getTopFiltersNoResultsResponse';
import type { GetTopHitsResponse } from '../model/getTopHitsResponse';
import type { GetTopSearchesResponse } from '../model/getTopSearchesResponse';
import type { GetUsersCountResponse } from '../model/getUsersCountResponse';

export const apiClientVersion = '5.0.0-alpha.0';

export const REGIONS = ['de', 'us'] as const;
export type Region = typeof REGIONS[number];

function getDefaultHosts(region?: Region): Host[] {
  const url = !region
    ? 'analytics.algolia.com'
    : 'analytics.{region}.algolia.com'.replace('{region}', region);

  return [{ url, accept: 'readWrite', protocol: 'https' }];
}

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export function createAnalyticsClient({
  appId: appIdOption,
  apiKey: apiKeyOption,
  authMode,
  algoliaAgents,
  region: regionOption,
  ...options
}: CreateClientOptions & { region?: Region }) {
  const auth = createAuth(appIdOption, apiKeyOption, authMode);
  const transporter = createTransporter({
    hosts: getDefaultHosts(regionOption),
    ...options,
    algoliaAgent: getAlgoliaAgent({
      algoliaAgents,
      client: 'Analytics',
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
     * Returns the average click position. The endpoint returns a value for the complete given time range, as well as a value per day.
     *
     * @summary Get average click position.
     * @param getAverageClickPosition - The getAverageClickPosition object.
     * @param getAverageClickPosition.index - The index name to target.
     * @param getAverageClickPosition.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getAverageClickPosition.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getAverageClickPosition.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getAverageClickPosition(
      { index, startDate, endDate, tags }: GetAverageClickPositionProps,
      requestOptions?: RequestOptions
    ): Promise<GetAverageClickPositionResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getAverageClickPosition`.'
        );
      }

      const requestPath = '/2/clicks/averageClickPosition';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns the distribution of clicks per range of positions.  If the groups all have a count of 0, it means Algolia didn’t receive any click events for the queries with the clickAnalytics search parameter set to true. The count is 0 until Algolia receives at least one click event.
     *
     * @summary Get clicks per positions.
     * @param getClickPositions - The getClickPositions object.
     * @param getClickPositions.index - The index name to target.
     * @param getClickPositions.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getClickPositions.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getClickPositions.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getClickPositions(
      { index, startDate, endDate, tags }: GetClickPositionsProps,
      requestOptions?: RequestOptions
    ): Promise<GetClickPositionsResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getClickPositions`.'
        );
      }

      const requestPath = '/2/clicks/positions';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns a click-through rate (CTR). The endpoint returns a value for the complete given time range, as well as a value per day. It also returns the count of clicks and searches used to compute the rates.
     *
     * @summary Get click-through rate (CTR).
     * @param getClickThroughRate - The getClickThroughRate object.
     * @param getClickThroughRate.index - The index name to target.
     * @param getClickThroughRate.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getClickThroughRate.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getClickThroughRate.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getClickThroughRate(
      { index, startDate, endDate, tags }: GetClickThroughRateProps,
      requestOptions?: RequestOptions
    ): Promise<GetClickThroughRateResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getClickThroughRate`.'
        );
      }

      const requestPath = '/2/clicks/clickThroughRate';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns a conversion rate (CR). The endpoint returns a value for the complete given time range, as well as a value per day. It also returns the count of conversion and searches used to compute the rates.
     *
     * @summary Get conversion rate (CR).
     * @param getConversationRate - The getConversationRate object.
     * @param getConversationRate.index - The index name to target.
     * @param getConversationRate.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getConversationRate.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getConversationRate.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getConversationRate(
      { index, startDate, endDate, tags }: GetConversationRateProps,
      requestOptions?: RequestOptions
    ): Promise<GetConversationRateResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getConversationRate`.'
        );
      }

      const requestPath = '/2/conversions/conversionRate';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns the rate at which searches didn\'t lead to any clicks. The endpoint returns a value for the complete given time range, as well as a value per day. It also returns the count of searches and searches without clicks.
     *
     * @summary Get no click rate.
     * @param getNoClickRate - The getNoClickRate object.
     * @param getNoClickRate.index - The index name to target.
     * @param getNoClickRate.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getNoClickRate.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getNoClickRate.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getNoClickRate(
      { index, startDate, endDate, tags }: GetNoClickRateProps,
      requestOptions?: RequestOptions
    ): Promise<GetNoClickRateResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getNoClickRate`.'
        );
      }

      const requestPath = '/2/searches/noClickRate';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns the rate at which searches didn\'t return any results. The endpoint returns a value for the complete given time range, as well as a value per day. It also returns the count of searches and searches without results used to compute the rates.
     *
     * @summary Get no results rate.
     * @param getNoResultsRate - The getNoResultsRate object.
     * @param getNoResultsRate.index - The index name to target.
     * @param getNoResultsRate.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getNoResultsRate.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getNoResultsRate.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getNoResultsRate(
      { index, startDate, endDate, tags }: GetNoResultsRateProps,
      requestOptions?: RequestOptions
    ): Promise<GetNoResultsRateResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getNoResultsRate`.'
        );
      }

      const requestPath = '/2/searches/noResultRate';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns the number of searches across the given time range. The endpoint returns a value for the complete given time range, as well as a value per day.
     *
     * @summary Get searches count.
     * @param getSearchesCount - The getSearchesCount object.
     * @param getSearchesCount.index - The index name to target.
     * @param getSearchesCount.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getSearchesCount.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getSearchesCount.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getSearchesCount(
      { index, startDate, endDate, tags }: GetSearchesCountProps,
      requestOptions?: RequestOptions
    ): Promise<GetSearchesCountResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getSearchesCount`.'
        );
      }

      const requestPath = '/2/searches/count';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns top searches that didn\'t lead to any clicks. Limited to the 1000 most frequent ones. For each search, also returns the average number of found hits.
     *
     * @summary Get top searches with no clicks.
     * @param getSearchesNoClicks - The getSearchesNoClicks object.
     * @param getSearchesNoClicks.index - The index name to target.
     * @param getSearchesNoClicks.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getSearchesNoClicks.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getSearchesNoClicks.limit - Number of records to return. Limit is the size of the page.
     * @param getSearchesNoClicks.offset - Position of the starting record. Used for paging. 0 is the first record.
     * @param getSearchesNoClicks.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getSearchesNoClicks(
      {
        index,
        startDate,
        endDate,
        limit,
        offset,
        tags,
      }: GetSearchesNoClicksProps,
      requestOptions?: RequestOptions
    ): Promise<GetSearchesNoClicksResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getSearchesNoClicks`.'
        );
      }

      const requestPath = '/2/searches/noClicks';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (limit !== undefined) {
        queryParameters.limit = limit.toString();
      }

      if (offset !== undefined) {
        queryParameters.offset = offset.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns top searches that didn\'t return any results. Limited to the 1000 most frequent ones.
     *
     * @summary Get top searches with no results.
     * @param getSearchesNoResults - The getSearchesNoResults object.
     * @param getSearchesNoResults.index - The index name to target.
     * @param getSearchesNoResults.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getSearchesNoResults.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getSearchesNoResults.limit - Number of records to return. Limit is the size of the page.
     * @param getSearchesNoResults.offset - Position of the starting record. Used for paging. 0 is the first record.
     * @param getSearchesNoResults.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getSearchesNoResults(
      {
        index,
        startDate,
        endDate,
        limit,
        offset,
        tags,
      }: GetSearchesNoResultsProps,
      requestOptions?: RequestOptions
    ): Promise<GetSearchesNoResultsResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getSearchesNoResults`.'
        );
      }

      const requestPath = '/2/searches/noResults';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (limit !== undefined) {
        queryParameters.limit = limit.toString();
      }

      if (offset !== undefined) {
        queryParameters.offset = offset.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns the latest update time of the analytics API for a given index. If the index has been recently created and/or no search has been performed yet the updated time will be null.
     *
     * @summary Get Analytics API status.
     * @param getStatus - The getStatus object.
     * @param getStatus.index - The index name to target.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getStatus(
      { index }: GetStatusProps,
      requestOptions?: RequestOptions
    ): Promise<GetStatusResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getStatus`.'
        );
      }

      const requestPath = '/2/status';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
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
     * Returns top countries. Limited to the 1000 most frequent ones.
     *
     * @summary Get top countries.
     * @param getTopCountries - The getTopCountries object.
     * @param getTopCountries.index - The index name to target.
     * @param getTopCountries.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopCountries.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopCountries.limit - Number of records to return. Limit is the size of the page.
     * @param getTopCountries.offset - Position of the starting record. Used for paging. 0 is the first record.
     * @param getTopCountries.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getTopCountries(
      { index, startDate, endDate, limit, offset, tags }: GetTopCountriesProps,
      requestOptions?: RequestOptions
    ): Promise<GetTopCountriesResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getTopCountries`.'
        );
      }

      const requestPath = '/2/countries';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (limit !== undefined) {
        queryParameters.limit = limit.toString();
      }

      if (offset !== undefined) {
        queryParameters.offset = offset.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns top filter attributes. Limited to the 1000 most used filters.
     *
     * @summary Get top filter attributes.
     * @param getTopFilterAttributes - The getTopFilterAttributes object.
     * @param getTopFilterAttributes.index - The index name to target.
     * @param getTopFilterAttributes.search - The query term to search for. Must match the exact user input.
     * @param getTopFilterAttributes.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopFilterAttributes.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopFilterAttributes.limit - Number of records to return. Limit is the size of the page.
     * @param getTopFilterAttributes.offset - Position of the starting record. Used for paging. 0 is the first record.
     * @param getTopFilterAttributes.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getTopFilterAttributes(
      {
        index,
        search,
        startDate,
        endDate,
        limit,
        offset,
        tags,
      }: GetTopFilterAttributesProps,
      requestOptions?: RequestOptions
    ): Promise<GetTopFilterAttributesResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getTopFilterAttributes`.'
        );
      }

      const requestPath = '/2/filters';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (search !== undefined) {
        queryParameters.search = search.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (limit !== undefined) {
        queryParameters.limit = limit.toString();
      }

      if (offset !== undefined) {
        queryParameters.offset = offset.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns top filters for the given attribute. Limited to the 1000 most used filters.
     *
     * @summary Get top filters for the an attribute.
     * @param getTopFilterForAttribute - The getTopFilterForAttribute object.
     * @param getTopFilterForAttribute.attribute - The exact name of the attribute.
     * @param getTopFilterForAttribute.index - The index name to target.
     * @param getTopFilterForAttribute.search - The query term to search for. Must match the exact user input.
     * @param getTopFilterForAttribute.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopFilterForAttribute.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopFilterForAttribute.limit - Number of records to return. Limit is the size of the page.
     * @param getTopFilterForAttribute.offset - Position of the starting record. Used for paging. 0 is the first record.
     * @param getTopFilterForAttribute.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getTopFilterForAttribute(
      {
        attribute,
        index,
        search,
        startDate,
        endDate,
        limit,
        offset,
        tags,
      }: GetTopFilterForAttributeProps,
      requestOptions?: RequestOptions
    ): Promise<GetTopFilterForAttributeResponse> {
      if (!attribute) {
        throw new Error(
          'Parameter `attribute` is required when calling `getTopFilterForAttribute`.'
        );
      }

      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getTopFilterForAttribute`.'
        );
      }

      const requestPath = '/2/filters/{attribute}'.replace(
        '{attribute}',
        encodeURIComponent(attribute)
      );
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (search !== undefined) {
        queryParameters.search = search.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (limit !== undefined) {
        queryParameters.limit = limit.toString();
      }

      if (offset !== undefined) {
        queryParameters.offset = offset.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns top filters with no results. Limited to the 1000 most used filters.
     *
     * @summary Get top filters for a no result search.
     * @param getTopFiltersNoResults - The getTopFiltersNoResults object.
     * @param getTopFiltersNoResults.index - The index name to target.
     * @param getTopFiltersNoResults.search - The query term to search for. Must match the exact user input.
     * @param getTopFiltersNoResults.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopFiltersNoResults.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopFiltersNoResults.limit - Number of records to return. Limit is the size of the page.
     * @param getTopFiltersNoResults.offset - Position of the starting record. Used for paging. 0 is the first record.
     * @param getTopFiltersNoResults.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getTopFiltersNoResults(
      {
        index,
        search,
        startDate,
        endDate,
        limit,
        offset,
        tags,
      }: GetTopFiltersNoResultsProps,
      requestOptions?: RequestOptions
    ): Promise<GetTopFiltersNoResultsResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getTopFiltersNoResults`.'
        );
      }

      const requestPath = '/2/filters/noResults';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (search !== undefined) {
        queryParameters.search = search.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (limit !== undefined) {
        queryParameters.limit = limit.toString();
      }

      if (offset !== undefined) {
        queryParameters.offset = offset.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns top hits. Limited to the 1000 most frequent ones.
     *
     * @summary Get top hits.
     * @param getTopHits - The getTopHits object.
     * @param getTopHits.index - The index name to target.
     * @param getTopHits.search - The query term to search for. Must match the exact user input.
     * @param getTopHits.clickAnalytics - Whether to include the click-through and conversion rates for a search.
     * @param getTopHits.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopHits.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopHits.limit - Number of records to return. Limit is the size of the page.
     * @param getTopHits.offset - Position of the starting record. Used for paging. 0 is the first record.
     * @param getTopHits.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getTopHits(
      {
        index,
        search,
        clickAnalytics,
        startDate,
        endDate,
        limit,
        offset,
        tags,
      }: GetTopHitsProps,
      requestOptions?: RequestOptions
    ): Promise<GetTopHitsResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getTopHits`.'
        );
      }

      const requestPath = '/2/hits';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (search !== undefined) {
        queryParameters.search = search.toString();
      }

      if (clickAnalytics !== undefined) {
        queryParameters.clickAnalytics = clickAnalytics.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (limit !== undefined) {
        queryParameters.limit = limit.toString();
      }

      if (offset !== undefined) {
        queryParameters.offset = offset.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns top searches. Limited to the 1000 most frequent ones. For each search, also returns the average number of hits returned.
     *
     * @summary Get top searches.
     * @param getTopSearches - The getTopSearches object.
     * @param getTopSearches.index - The index name to target.
     * @param getTopSearches.clickAnalytics - Whether to include the click-through and conversion rates for a search.
     * @param getTopSearches.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopSearches.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getTopSearches.orderBy - Reorder the results.
     * @param getTopSearches.direction - The sorting of the result.
     * @param getTopSearches.limit - Number of records to return. Limit is the size of the page.
     * @param getTopSearches.offset - Position of the starting record. Used for paging. 0 is the first record.
     * @param getTopSearches.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getTopSearches(
      {
        index,
        clickAnalytics,
        startDate,
        endDate,
        orderBy,
        direction,
        limit,
        offset,
        tags,
      }: GetTopSearchesProps,
      requestOptions?: RequestOptions
    ): Promise<GetTopSearchesResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getTopSearches`.'
        );
      }

      const requestPath = '/2/searches';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (clickAnalytics !== undefined) {
        queryParameters.clickAnalytics = clickAnalytics.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (orderBy !== undefined) {
        queryParameters.orderBy = orderBy.toString();
      }

      if (direction !== undefined) {
        queryParameters.direction = direction.toString();
      }

      if (limit !== undefined) {
        queryParameters.limit = limit.toString();
      }

      if (offset !== undefined) {
        queryParameters.offset = offset.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
     * Returns the distinct count of users across the given time range. The endpoint returns a value for the complete given time range, as well as a value per day.
     *
     * @summary Get users count.
     * @param getUsersCount - The getUsersCount object.
     * @param getUsersCount.index - The index name to target.
     * @param getUsersCount.startDate - The lower bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getUsersCount.endDate - The upper bound timestamp (a date, a string like \"2006-01-02\") of the period to analyze.
     * @param getUsersCount.tags - Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded.
     * @param requestOptions - The requestOptions to send along with the query, they will be merged with the transporter requestOptions.
     */
    getUsersCount(
      { index, startDate, endDate, tags }: GetUsersCountProps,
      requestOptions?: RequestOptions
    ): Promise<GetUsersCountResponse> {
      if (!index) {
        throw new Error(
          'Parameter `index` is required when calling `getUsersCount`.'
        );
      }

      const requestPath = '/2/users/count';
      const headers: Headers = {};
      const queryParameters: QueryParameters = {};

      if (index !== undefined) {
        queryParameters.index = index.toString();
      }

      if (startDate !== undefined) {
        queryParameters.startDate = startDate.toString();
      }

      if (endDate !== undefined) {
        queryParameters.endDate = endDate.toString();
      }

      if (tags !== undefined) {
        queryParameters.tags = tags.toString();
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
  };
}

export type AnalyticsClient = ReturnType<typeof createAnalyticsClient>;
