// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { UserProfile } from './userProfile';

export type GetSegmentUsersResponse = {
  /**
   * The ID of the segment.
   */
  segmentID: string;

  users: UserProfile[];

  /**
   * The token is used to navigate backward in the user list. To navigate from the current user list to the previous page, the API generates the previous page token and it sends the token in the response, beside the current user list. NOTE: This body param cannot be used with `nextPageToken` at the same time.
   */
  previousPageToken?: string;

  /**
   * The token is used to navigate forward in the user list. To navigate from the current user list to the next page, the API generates the next page token and it sends the token in the response, beside the current user list. NOTE: This body param cannot be used with `previousPageToken` at the same time.
   */
  nextPageToken?: string;
};
