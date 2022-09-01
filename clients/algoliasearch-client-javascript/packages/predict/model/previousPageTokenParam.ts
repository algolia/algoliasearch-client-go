// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

export type PreviousPageTokenParam = {
  /**
   * The token is used to navigate backward in the user list. To navigate from the current user list to the previous page, the API generates the previous page token and it sends the token in the response, beside the current user list. NOTE: This body param cannot be used with `nextPageToken` at the same time.
   */
  previousPageToken?: string;
};
