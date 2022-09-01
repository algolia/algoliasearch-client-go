// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

export type NextPageTokenParam = {
  /**
   * The token is used to navigate forward in the user list. To navigate from the current user list to the next page, the API generates the next page token and it sends the token in the response, beside the current user list. NOTE: This body param cannot be used with `previousPageToken` at the same time.
   */
  nextPageToken?: string;
};
