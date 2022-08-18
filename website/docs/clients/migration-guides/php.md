---
title: PHP
---

| Previous                                      | Latest                                             | Description                                                                                                                                                                                                                                          |
|-----------------------------------------------|:---------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `"algolia/algoliasearch-client-php": "^3.2"`  | `"algolia/algoliasearch-client-php": "^4.0@alpha"` | **During the beta phase**, the clients are available under the package 4.x.x-alpha , you can find a full list [here](https://packagist.org/packages/algolia/algoliasearch-client-php).                                                               |
| `Algolia\AlgoliaSearch`                       | `Algolia\AlgoliaSearch\Api`                        | Exported clients have now the namespace suffixed by `Api`.                                                                                                                                                                                           |
| `Algolia\AlgoliaSearch\Support\UserAgent`     | `Algolia\AlgoliaSearch\Support\AlgoliaAgent`       | `UserAgent` class has been renamed to `AlgoliaAgent` for consistency across client languages (`addCustomUserAgent` method also became `addAlgoliaAgent`).                                                                                            |
| `Algolia\AlgoliaSearch\SearchIndex`           | **removed**                                        | Since the method `initIndex` doesn't exist anymore, we decided to merge the `SearchIndex` class inside the `SearchClient` one, now all the methods related to search endpoints are located there.                                                    |
| `Algolia\AlgoliaSearch\Cache\FileCacheDriver` | **removed**                                        | This implementation of the `CacheInterface` is not available anymore in the Client. If you feel the need for it, [please open an issue](https://github.com/algolia/api-clients-automation/issues/new?assignees=&labels=&template=Feature_request.md) |

### Usage

To get started, first uninstall the previously added clients.

```bash
composer remove algolia/algoliasearch-client-php
```

You can now install the `Algoliasearch` clients.

```bash
composer require algolia/algoliasearch-client-php "^4.0@alpha"
```

You can continue this guide on [our installation page](/docs/clients/installation).

### Methods targeting an `indexName`

Prior to the `initIndex` removal stated in the [common breaking changes](/docs/clients/migration-guides/#common-breaking-changes), all methods previously available at the `initIndex` level requires the `indexName` to be sent with the query.

```php
use Algolia\AlgoliaSearch\Api\SearchClient;

$client = SearchClient::create(
    '<YOUR_APP_ID>',
    '<YOUR_API_KEY>'
);

// only query string
$searchResults = $client->search([
    'requests' => [
        ['indexName' => '<YOUR_INDEX_NAME>', 'query' =>'<YOUR_QUERY>'],
    ],
]);

$searchResults2 = $client->search([
    'requests' => [
        [
            'indexName' => '<YOUR_INDEX_NAME>', 
            'query' => '<YOUR_QUERY>',
            'attributesToRetrieve' => ['firstname', 'lastname'],
            'hitsPerPage' => 50,
        ],
    ],
]);
```

### Generate a secured Api Key 

The `SearchClient::generateSecuredApiKey()` from the previous client was removed, but you can still create them:

```php
use Algolia\AlgoliaSearch\Support\Helpers;

// Key will be valid for 25 hours.
$validUntil = time() + (3600 * 25);

$urlEncodedRestrictions = Helpers::buildQuery([
    'restrictIndices' => '<YOUR_INDEX_NAME>',
    'validUntil' => $validUntil,
]);

$content = hash_hmac('sha256', $urlEncodedRestrictions, '<YOUR_SEARCH_KEY>').$urlEncodedRestrictions;
$securedSearchKey =  base64_encode($content);

```