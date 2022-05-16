// eslint-disable-next-line import/no-commonjs
module.exports = {
  patterns: [
    // Ignore the roots and go down the tree by negating hand written files
    'specs/bundled/*.yml',

    'clients/**',
    'clients/**/.*', // hidden files are not ignored by default
    '!clients/README.md',
    '!clients/**/.openapi-generator-ignore',

    // Java
    '!clients/algoliasearch-client-java-2/*.gradle',
    '!clients/algoliasearch-client-java-2/gradlew',
    '!clients/algoliasearch-client-java-2/.gitignore',
    '!clients/algoliasearch-client-java-2/gradle/wrapper/**',
    '!clients/algoliasearch-client-java-2/algoliasearch-core/build.gradle',
    '!clients/algoliasearch-client-java-2/algoliasearch-core/gradle.properties',
    '!clients/algoliasearch-client-java-2/algoliasearch-core/src/main/java/com/algolia/exceptions/**',
    '!clients/algoliasearch-client-java-2/algoliasearch-core/src/main/java/com/algolia/utils/**',

    'tests/output/java/src/test/java/com/algolia/methods/**', // this could be added automatically by the script, but with overhead
    'tests/output/java/src/test/java/com/algolia/client/**',

    // JavaScript
    '!clients/algoliasearch-client-javascript/*',
    '!clients/algoliasearch-client-javascript/.github/**',
    '!clients/algoliasearch-client-javascript/.yarn/**',
    '!clients/algoliasearch-client-javascript/scripts/**',
    '!clients/algoliasearch-client-javascript/packages/algoliasearch/**',
    '!clients/algoliasearch-client-javascript/packages/requester-*/**',
    '!clients/algoliasearch-client-javascript/packages/client-common/**',

    'tests/output/javascript/package.json',
    'tests/output/javascript/src/methods/**',
    'tests/output/javascript/src/client/**',

    // PHP
    '!clients/algoliasearch-client-php/*',
    '!clients/algoliasearch-client-php/.*',
    '!clients/algoliasearch-client-php/lib/*',
    '!clients/algoliasearch-client-php/lib/Cache/**',
    '!clients/algoliasearch-client-php/lib/Exceptions/**',
    '!clients/algoliasearch-client-php/lib/Http/**',
    '!clients/algoliasearch-client-php/lib/Log/**',
    '!clients/algoliasearch-client-php/lib/RequestOptions/**',
    '!clients/algoliasearch-client-php/lib/RetryStrategy/**',
    '!clients/algoliasearch-client-php/lib/Support/**',
    '!clients/algoliasearch-client-php/lib/Configuration/ConfigWithRegion.php',
    'clients/algoliasearch-client-php/lib/Configuration/Configuration.php',

    'tests/output/php/src/methods/**',
    'tests/output/php/src/client/**',
  ],
};
