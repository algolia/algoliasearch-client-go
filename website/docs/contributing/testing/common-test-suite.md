---
title: Common Test Suite
---

# Common Test Suite

The CTS aims at ensuring minimal working operation for the API clients, by comparing the request formed by sample parameters.
It is automatically generated for all languages, from a JSON entry point.

:::info

Common Test Suite requires all clients to be built.

[CLI commands for the Common Test Suite](/docs/contributing/CLI/cts-commands)

:::

## How to add test

The test generation script requires a JSON file name from the `operationId` (e.g. `search.json`), located in the `CTS/<client>/requests/` folder (e.g. `CTS/search/requests/`).

> See the [browse test file for the search client](https://github.com/algolia/api-clients-automation/blob/main/tests/CTS/methods/requests/search/browse.json)

```json
[
  {
    "testName": "The name of the test (e.g. 'searches to testIndex with query parameters')",
    // The parameters of you'd like to pass to the method
    "parameters": {
      "indexName": "testIndex",
      "searchParams": {
        "offset": 42,
        "limit": 21,
        "query": "the string to search"
      },
      "facets": ["*"]
    },
    // Extra options sent with your method
    "requestOptions": {
      // Merged with transporter query parameters
      "queryParameters": {
        "anOtherParam": true
      },
      // Merged with transporter headers
      "headers": {
        "x-header": "test"
      }
    },
    // The payload of the request
    "request": {
      "path": "/1/indexes/testIndex/query",
      "method": "POST",
      "body": { "query": "the string to search" },
      "queryParameters": {
        "otherParam": "22",
        "anOtherParam": "true"
      },
      "headers": {
        "x-header": "test"
      }
    }
  }
]
```

And that's it! If the name of the file matches an `operationId` in the spec, a test will be generated and will be calling the method name `operationId`.

The list of `queryParameters` must match exactly the actual value, the CTS has to check the number of query parameters and the value of each.

## How to add a new language

Create a template in [`tests/CTS/methods/requests/templates/<languageName>/requests.mustache`](https://github.com/algolia/api-clients-automation/tree/main/tests/CTS/methods/requests/templates) that parses an array of tests into the test framework of choice.

> See [implementation of the JavaScript tests](https://github.com/algolia/api-clients-automation/blob/main/tests/CTS/methods/requests/templates/javascript/requests.mustache)

When writing your template, here is a list of variables accessible from `mustache`:

```json
{
  "import": "the name of the package or library to import",
  "client": "the name of the API Client object to instanciate and import",
  "blocks": [
    {
      // The list of test to implement
      "operationID": "the name of the endpoint and the cts file to test",
      "tests": [
        {
          "testName": "the descriptive name test (default to `method`)",
          "method": "the method to call on the API Client",
          "parameters": {
            // Object of all parameters with their name, to be used for languages that require the parameter name
            "parameterName": "value"
            // ...
          },
          "parametersWithDataType": [
            {
              "key": "key",
              "value": "value",
              // booleans indicating the data type
              "isArray": false,
              "isObject": true,
              "isFreeFormObject": false,
              "isString": false,
              "isInteger": false,
              "isLong": false,
              "isDouble": false,
              "isEnum": false,
              "isBoolean": false,
              "objectName": "SearchParams",
              // oneOfModel empty if there is no oneOf
              "oneOfModel": {
                "parentClassName": "SearchParams",
                "type": "SearchParamsObject"
              },
              // properties used to have unique name and link to parent
              "parent": "theParentObject",
              "suffix": 7,
              "parentSuffix": 6,
              // boolean indicating if it is the last parameter
              "-last": false
            }
          ],
          "requestOptionsWithDataType": [
            {
              "key": "key",
              "value": "value",
              // booleans indicating the data type
              "isArray": false,
              "isObject": true,
              "isFreeFormObject": false,
              "isString": false,
              "isInteger": false,
              "isLong": false,
              "isDouble": false,
              "isEnum": false,
              "isBoolean": false,
              "objectName": "SearchParams",
              // oneOfModel empty if there is no oneOf
              "oneOfModel": {
                "parentClassName": "SearchParams",
                "type": "SearchParamsObject"
              },
              // properties used to have unique name and link to parent
              "parent": "theParentObject",
              "suffix": 7,
              "parentSuffix": 6,
              // boolean indicating if it is the last parameter
              "-last": false
            }
          ],
          // boolean indicating if the method has parameters, useful for `GET` requests
          "hasParameters": true,
          // boolean indicating if the method has requestOptions, useful to shape your template only if provided
          "hasRequestOptions": true,
          "request": {
            "path": "the expected path of the request",
            "method": "the expected method: GET, POST, PUT, DELETE or PATCH",
            "body": {
              // The expected body of the request
            },
            "queryParameters": {
              // key: string map
              "parameterName": "stringify version of the value"
            },
            "headers": {
              // key: string map
              "headerName": "stringify version of the value"
            }
          }
        }
      ]
    }
  ]
}
```

## Add common tests to every clients

You might want to test how every clients behaves, without having to duplicate the same tests. We provide 4 methods on every clients, common to all languages.

You can find [the common folder](https://github.com/algolia/api-clients-automation/tree/main/tests/CTS/methods/requests/common) in the CTS too. [Adding a test](#how-to-add-test) in this folder will generate tests for all the clients.

## Get the list of remaining CTS to implement

To get the list of `operationId` not yet in the CTS but in the spec, run this command:

```bash
rm -rf ./specs/bundled
comm -3 <(grep -r operationId ./specs | awk -F: '{gsub(/ /,""); print $NF}' | sort) <(find ./tests/CTS/clients -type f -name '*.json' | awk -F/ '{gsub(/.json/,"");print $NF}' | sort)
```
