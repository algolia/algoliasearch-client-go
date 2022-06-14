<?php
require '../../../clients/algoliasearch-client-php/vendor/autoload.php';

// Gets the vars from local environment
$env = getenv();

// If the script has been run from docker's playground, fetches the vars from .env file instead
if (isset($env['DOCKER']) && $env['DOCKER'] === "true") {
    $dotenv = Dotenv\Dotenv::createImmutable('../..');
    $dotenv->load();
    $env = $_ENV;
}

return $env;