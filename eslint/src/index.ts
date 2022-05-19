import { endWithDot } from './rules/endWithDot';
import { createOutOfLineRule } from './rules/outOfLineRule';
import { singleQuoteRef } from './rules/singleQuoteRef';

const rules = {
  'end-with-dot': endWithDot,
  'out-of-line-enum': createOutOfLineRule({ property: 'enum' }),
  'out-of-line-one-of': createOutOfLineRule({ property: 'oneOf' }),
  'out-of-line-all-of': createOutOfLineRule({ property: 'allOf' }),
  'out-of-line-any-of': createOutOfLineRule({ property: 'anyOf' }),
  'single-quote-ref': singleQuoteRef,
};

export { rules };
