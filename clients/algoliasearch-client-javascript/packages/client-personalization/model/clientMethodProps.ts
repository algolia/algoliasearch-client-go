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
 * Properties for the `deleteUserProfile` method.
 */
export type DeleteUserProfileProps = {
  /**
   * UserToken representing the user for which to fetch the Personalization profile.
   */
  userToken: string;
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
 * Properties for the `getUserTokenProfile` method.
 */
export type GetUserTokenProfileProps = {
  /**
   * UserToken representing the user for which to fetch the Personalization profile.
   */
  userToken: string;
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
