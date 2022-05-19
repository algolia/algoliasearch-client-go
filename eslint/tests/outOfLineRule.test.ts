import { RuleTester } from 'eslint';

import { createOutOfLineRule } from '../src/rules/outOfLineRule';

const ruleTester = new RuleTester({
  parser: require.resolve('yaml-eslint-parser'),
});

// this test is enough for oneOf, allOf, anyOf, as they use the same rule.
ruleTester.run('out-of-line-enum', createOutOfLineRule({ property: 'enum' }), {
  valid: [
    `
simple:
  type: string
  enum: [bla, blabla]
`,
    `
simple:
  type: string
  enum:
    - bla
    - blabla
`,
    `
simple:
  type: string
  enum: [bla, blabla]

useIt:
  $ref: '#/simple'
`,
    `
servers:
  - url: http://test-server.com
    variables:
      region:
        default: us
        enum:
          - us
          - de
`,
  ],
  invalid: [
    {
      code: `
root:
  inside:
    type: string
    enum: [bla, blabla]
  `,
      errors: [{ messageId: 'enumNotOutOfLine' }],
    },
    {
      code: `
root:
  inside:
    deeper:
      type: string
      enum: [bla, blabla]

useIt:
  $ref: '#/root/inside/deeper'
  `,
      errors: [{ messageId: 'enumNotOutOfLine' }],
    },
  ],
});
