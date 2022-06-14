<?php

$env = require_once('../loadEnv.php');

use Algolia\AlgoliaSearch\Api\AbtestingClient;

$client = AbtestingClient::create($env['ALGOLIA_APPLICATION_ID'], $env['ALGOLIA_ANALYTICS_KEY']);

$abTest = [
    'name' => 'testing',
    'variants' => [
        [
            'index' => 'test1',
            'trafficPercentage' => 30,
            'description' => 'a description',
        ],
        [
            'index' => 'test2',
            'trafficPercentage' => 50,
        ],
    ],
    'endAt' => '2022-02-01',
];

var_dump(
    $client->addABTests($abTest)
);
