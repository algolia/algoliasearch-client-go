// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

export type ClickPosition = {
  /**
   * Range of positions with the following pattern: - Positions from 1 to 10 included are displayed in separated groups. - Positions from 11 to 20 included are grouped together. - Positions from 21 and up are grouped together.
   */
  position: number[];
  /**
   * The number of click event.
   */
  clickCount: number;
};
