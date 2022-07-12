// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

/**
 * Single objectID to promote as hits.
 */
export type PromoteObjectID = {
  /**
   * Unique identifier of the object to promote.
   */
  objectID: string;
  /**
   * The position to promote the objects to (zero-based). If you pass objectIDs, the objects are placed at this position as a group. For example, if you pass four objectIDs to position 0, the objects take the first four positions.
   */
  position: number;
};
