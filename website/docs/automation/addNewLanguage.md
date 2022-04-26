---
title: Support a new language
---

# Support a new language

:::info

Make sure to first [setup the repository tooling](/docs/automation/setupRepository) to ease your journey!

You will also need to have the [openapi-generator](https://openapi-generator.tech/docs/installation/) installed.

:::

This repository leverages [openapi-generator](https://openapi-generator.tech/) to generate API clients.

## Find a template to start with

Provided templates should be a good starting point to generate a client but make sure to implement the [Algolia requirements](#algolia-requirements) to make it work properly.

You can pick a default template on the [openapi-generator's "generators" page](https://openapi-generator.tech/docs/generators)

### Extract the template locally

```bash
openapi-generator author template -g <templateName> -o templates/<languageName>
```

Example for the `JavaScript` client with the `typescript-node` template:

```bash
openapi-generator author template -g typescript-node -o templates/javascript/
```

## Update the generator config

Add each client in the file [`openapitools.json`](https://github.com/algolia/api-clients-automation/blob/main/openapitools.json), following the others client structure.

> See [`add a new client`](/docs/automation/addNewApiClient) for informations on how to structure your new client.

### Algolia requirements

:::caution

**DX is key.**

**We require the generated API clients to have an up-to-date usage with their ecosystem, make sure to provide correct tooling to lint and format generated code.**

:::

### Strip code

The default generators includes **a lot** of features that won't be used with the Algolia engine:

- Multiple authentication methods: `appId`/`apiKey` are the only authentication methods, located in the requests headers.
- Built-in transporters: A [retry strategy](#retry-strategy) is required to target the DSNs of an `appId`, along with other transporter details listed below.
  - Regional clients are handled differently: [Different client hosts](#different-client-hosts)
- File support, payload format etc.: Requests only require `JSON` support to communicate with the engine.

### Init method

By default, OpenAPI will put the `AppId` and `ApiKey` in every method parameters, but the clients needs to be initialized with those values and put them in the header of every requests, with the right hosts.

The constructor of the client can be edited (from the `.mustache` files) to accept and store those values.

- [First implementation on the JavaScript client](https://github.com/algolia/api-clients-automation/pull/7)
- [Current implementation on the JavaScript client](https://github.com/algolia/api-clients-automation/blob/main/clients/algoliasearch-client-javascript/packages/client-search/src/searchClient.ts#L123)

### Retry strategy

The retry strategy cannot be generated and needs to be implemented outside of the generated client folder. You can achieve this by creating a `utils` (_or any naming that you find relevant_) folder and add a transporter and retry strategy logic to it.

- [First implementation on the JavaScript client](https://github.com/algolia/api-clients-automation/pull/9)
- [Current implementation on the PHP client](https://github.com/algolia/api-clients-automation/tree/main/clients/algoliasearch-client-php/lib/RetryStrategy)

### Different client hosts

Some Algolia clients (search and recommend) targets the default appId host (`${appId}-dsn.algolia.net`, `${appId}.algolia.net`, etc.), while clients like `personalization` have their own regional `host` (`eu` | `us` | `de`).

As the generator does not support reading `servers` in a spec file **yet**, hosts methods and variables are extracted with a custom script and create variables for you to use in the mustache templates, [read more here](/docs/automation/addNewApiClient#generators).

### User Agent

The header 'User-Agent' must respect a strict pattern of a base, client, plus additional user defined segments:

```
base: `Algolia for <language> (<apiVersion>)`
client: `; <clientName> (<clientVersion>)`
segment: `; <Description> ([version])`
```

> The version is optional for segments.

The resulting User Agent is the concatenation of `base`, `client`, and all the `segments`.

For example, if we have:

```
base: `Algolia for Java (5.0.0)`
client: `; Search (5.0.0)`
segment: `; JVM (11.0.14); experimental`
```

Then the resulting User Agent is :

```
Algolia for Java (5.0.0); Search (5.0.0); JVM (11.0.14); experimental
```

You can take a look at the Java implementation [here](https://github.com/algolia/api-clients-automation/pull/347).

### Requesters

> TODO: informations

### Logger

> TODO: informations