// eslint-disable-next-line import/no-commonjs
module.exports = {
  patterns: [
    // Ignore the roots and go down the tree by negating hand written files
    'specs/bundled/*.yml',

    'clients/**',
    '!clients/README.md',
    '!clients/**/.openapi-generator-ignore',

    // Java
    '!clients/algoliasearch-client-java-2/**',
    'clients/algoliasearch-client-java-2/gradle.properties',
    'clients/algoliasearch-client-java-2/algoliasearch-core/src/main/java/com/algolia/api/**',
    'clients/algoliasearch-client-java-2/algoliasearch-core/src/main/java/com/algolia/model/**',

    'tests/output/java/build.gradle',

    // JavaScript
    '!clients/algoliasearch-client-javascript/*',
    '!clients/algoliasearch-client-javascript/.github/**',
    '!clients/algoliasearch-client-javascript/.yarn/**',
    '!clients/algoliasearch-client-javascript/scripts/**',
    '!clients/algoliasearch-client-javascript/packages/requester-*/**',
    '!clients/algoliasearch-client-javascript/packages/client-common/**',
    '!clients/algoliasearch-client-javascript/packages/algoliasearch/__tests__/**',
    '!clients/algoliasearch-client-javascript/packages/algoliasearch/jest.config.ts',

    // PHP
    '!clients/algoliasearch-client-php/*',
    'clients/algoliasearch-client-php/lib/Api/*',
    'clients/algoliasearch-client-php/lib/Model/**',
    'clients/algoliasearch-client-php/lib/Configuration/*',
    'clients/algoliasearch-client-php/lib/ApiException.php',
    'clients/algoliasearch-client-php/lib/ObjectSerializer.php',
  ],
};
