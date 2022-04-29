import { endWithDot } from './rules/endWithDot';
import { outOfLineEnum } from './rules/outOfLineEnum';
import { singleQuoteRef } from './rules/singleQuoteRef';

const rules = {
  'end-with-dot': endWithDot,
  'out-of-line-enum': outOfLineEnum,
  'single-quote-ref': singleQuoteRef,
};

export { rules };
