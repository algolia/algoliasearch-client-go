<?php

require '../../../clients/algoliasearch-client-php/vendor/autoload.php';

use Algolia\AlgoliaSearch\Api\SearchClient;

$client = SearchClient::create(
    getenv('ALGOLIA_APPLICATION_ID'),
    getenv('ALGOLIA_ADMIN_KEY')
);
$indexName = getenv('SEARCH_INDEX');


$response = $client->saveObject(
    $indexName,
    ['objectID' => "111", 'name' => getenv('SEARCH_QUERY')],
);

var_dump($response);

$client->waitForTask($indexName, $response['taskID']);

var_dump(
    $client->search([
        'requests' => [
            ['indexName' => $indexName, 'query' => getenv('SEARCH_QUERY')],
        ],
    ])
);
