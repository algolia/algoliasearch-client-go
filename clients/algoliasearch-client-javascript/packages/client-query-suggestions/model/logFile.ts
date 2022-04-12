import type { LogLevel } from './logLevel';

export type LogFile = {
  /**
   * Date and time of creation of the record.
   */
  timestamp: string;
  level: LogLevel;
  /**
   * Detailed description of what happened.
   */
  message: string;
  /**
   * Indicates the hierarchy of the records. For example, a record with contextLevel=1 belongs to a preceding record with contextLevel=0.
   */
  contextLevel: number;
};
