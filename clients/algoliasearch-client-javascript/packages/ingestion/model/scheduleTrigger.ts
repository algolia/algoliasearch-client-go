// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { ScheduleTriggerType } from './scheduleTriggerType';

/**
 * The trigger information for a task of type \'schedule\'.
 */
export type ScheduleTrigger = {
  type: ScheduleTriggerType;

  /**
   * A cron expression that represent at which regularity the task should run.
   */
  cron: string;

  /**
   * The last time the scheduled task ran (RFC3339 format).
   */
  lastRun?: string;

  /**
   * The next scheduled run of the task (RFC3339 format).
   */
  nextRun: string;
};
