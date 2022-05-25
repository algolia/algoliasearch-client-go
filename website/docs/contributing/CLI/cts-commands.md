---
title: Common Test Suite commands
---

# Common Test Suite commands

:::info

Common Test Suite requires all clients to be built.

:::

The Common Test Suite commands are used to [`generate`](#generate) and [`run`](#run) tests.

## Usage

> `language` and `client` defaults to `all`

> The `client` parameter is variadic, you can pass multiple `client` name

```bash
yarn docker cts generate <language | all> <client... | all>
```

### Available options

| Option      | Command           | Description                                                   |
| ----------- | :---------------- | :------------------------------------------------------------ |
| verbose     | -v, --verbose     | Make the process verbose, display logs from third party tools |
| interactive | -i, --interactive | Open prompt to query parameters                               |

## Generate

### Generate CTS for all clients for all supported languages

```bash
yarn docker cts generate
```

### Generate CTS for specific client for specific language

```bash
yarn docker cts generate java sources
```

### Generate CTS for many client for specific language

```bash
yarn docker cts generate php sources recommend search
```

## Run

### Run CTS for all supported languages

```bash
yarn docker cts run
```

### Run CTS for a specific languages

```bash
yarn docker cts run javascript
```
