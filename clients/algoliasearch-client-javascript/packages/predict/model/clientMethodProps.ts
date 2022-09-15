// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { FetchAllUserProfilesParams } from './fetchAllUserProfilesParams';
import type { Params } from './params';
import type { SegmentType } from './segmentType';
import type { UpdateModelParams } from './updateModelParams';
import type { UpdateSegmentParams } from './updateSegmentParams';

/**
 * Properties for the `del` method.
 */
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

/**
 * Properties for the `deleteModelInstance` method.
 */
export type DeleteModelInstanceProps = {
  /**
   * The ID of the model to retrieve.
   */
  modelID: string;
};

/**
 * Properties for the `deleteSegment` method.
 */
export type DeleteSegmentProps = {
  /**
   * The ID of the Segment to fetch.
   */
  segmentID: string;
};

/**
 * Properties for the `deleteUserProfile` method.
 */
export type DeleteUserProfileProps = {
  /**
   * User ID for authenticated users or cookie ID for non-authenticated repeated users (visitors).
   */
  userID: string;
};

/**
 * Properties for the `fetchAllSegments` method.
 */
export type FetchAllSegmentsProps = {
  /**
   * The type of segments to fetch.
   */
  type?: SegmentType;
};

/**
 * Properties for the `fetchSegment` method.
 */
export type FetchSegmentProps = {
  /**
   * The ID of the Segment to fetch.
   */
  segmentID: string;
};

/**
 * Properties for the `fetchUserProfile` method.
 */
export type FetchUserProfileProps = {
  /**
   * User ID for authenticated users or cookie ID for non-authenticated repeated users (visitors).
   */
  userID: string;
  params: Params;
};

/**
 * Properties for the `get` method.
 */
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

/**
 * Properties for the `getModelInstanceConfig` method.
 */
export type GetModelInstanceConfigProps = {
  /**
   * The ID of the model to retrieve.
   */
  modelID: string;
};

/**
 * Properties for the `getModelMetrics` method.
 */
export type GetModelMetricsProps = {
  /**
   * The ID of the model to retrieve.
   */
  modelID: string;
};

/**
 * Properties for the `getSegmentUsers` method.
 */
export type GetSegmentUsersProps = {
  /**
   * The ID of the Segment to fetch.
   */
  segmentID: string;
  fetchAllUserProfilesParams: FetchAllUserProfilesParams;
};

/**
 * Properties for the `post` method.
 */
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

/**
 * Properties for the `put` method.
 */
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

/**
 * Properties for the `updateModelInstance` method.
 */
export type UpdateModelInstanceProps = {
  /**
   * The ID of the model to retrieve.
   */
  modelID: string;
  updateModelParams: UpdateModelParams;
};

/**
 * Properties for the `updateSegment` method.
 */
export type UpdateSegmentProps = {
  /**
   * The ID of the Segment to fetch.
   */
  segmentID: string;
  updateSegmentParams: UpdateSegmentParams;
};
