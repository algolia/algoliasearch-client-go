<?php

$env = require_once('../loadEnv.php');

use Algolia\AlgoliaSearch\Api\RecommendClient;

$client = RecommendClient::create($env['ALGOLIA_APPLICATION_ID'], $env['ALGOLIA_ADMIN_KEY']);
$indexName = $env['SEARCH_INDEX'];
$query = $env['SEARCH_QUERY'];

var_dump($client->getRecommendations(
    [
        'requests' => [
            [
                'indexName' => $indexName,
                'model' => 'bought-together',
                'objectID' => $query,
                'threshold' => 0
            ]
        ]
    ]
));
