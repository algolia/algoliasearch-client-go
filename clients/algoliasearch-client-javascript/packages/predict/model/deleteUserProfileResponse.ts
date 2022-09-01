// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

export type DeleteUserProfileResponse = {
  /**
   * The ID of the user that was deleted.
   */
  user: string;

  /**
   * The time the same user ID will be imported again when the data is ingested.
   */
  deletedUntil: string;
};
