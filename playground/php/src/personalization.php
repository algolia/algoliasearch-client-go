<?php

$env = require_once('../loadEnv.php');

use Algolia\AlgoliaSearch\Api\PersonalizationClient;

$client = PersonalizationClient::create($env['ALGOLIA_APPLICATION_ID'], $env['ALGOLIA_RECOMMENDATION_KEY']);

var_dump(
    $client->deleteUserProfile('userToken')
);
