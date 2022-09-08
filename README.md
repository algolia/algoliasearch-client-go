<div align="center">

# API Clients Automation

The Algolia API clients are generated from [OpenAPI specs](https://swagger.io/specification/), leveraging the open-source [openapi-generator](https://openapi-generator.tech/) tool.

[![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat-square)](./LICENSE.md)

<p align="center">
  <strong>
  <a href="https://api-clients-automation.netlify.app/">API clients automation documentation</a> •
  <a href="https://www.algolia.com/doc/">Algolia documentation</a>
  </strong>
</p>

</div>

**Migration note from current API clients**

> In July 2022, we released an alpha version generated API clients for the JavaScript, Java and PHP languages. If you are using the latest stable of those clients and looking to upgrade, read the [migration guide](https://api-clients-automation.netlify.app/docs/clients/migration-guides/). You can still browse the documentation of the stable clients on [the Algolia documentation](https://www.algolia.com/doc/).

## 💡 Getting Started with the clients

You can read `getting started` guides and how to use the API clients on [our documentation](https://api-clients-automation.netlify.app/docs/clients/installation).

## ✨ Contributing

> Looking to add a new client, or fix a bug? Make sure to take a look at [our contribution guides](https://api-clients-automation.netlify.app/docs/contributing/introduction).

### Setup repository tooling

```bash
nvm use && yarn
```

### Setup dev environment

> **Make sure to have Docker installed so you don't have to install the tooling for every API clients.**

```bash
yarn docker:setup
```

[Read more on our documentation](https://api-clients-automation.netlify.app/docs/contributing/setup-repository)

### CLI

The CLI allows you to make changes locally and run commands through the docker container.

- [Specs CLI commands](https://api-clients-automation.netlify.app/docs/contributing/CLI/specs-commands)
- [Clients CLI commands](https://api-clients-automation.netlify.app/docs/contributing/CLI/clients-commands)
- [CTS CLI commands](https://api-clients-automation.netlify.app/docs/contributing/CLI/cts-commands)

### Guides and requirements

Read the guides and requirements to:

- [Add a new client](https://api-clients-automation.netlify.app/docs/contributing/add-new-api-client)
- [Add a new language](https://api-clients-automation.netlify.app/docs/contributing/add-new-api-language)

### Tests

Test the generated clients by running:

- The [`playground`](./playground) (see [documentation](https://api-clients-automation.netlify.app/docs/contributing/testing/playground))
- The [`Common Test Suite`](./tests/) (see [documentation](https://api-clients-automation.netlify.app/docs/contributing/testing/common-test-suite)).

For full documentation, visit the **[online documentation](https://api-clients-automation.netlify.app/docs/contributing/introduction)**.

## ❓ Troubleshooting

Encountering an issue with the API clients? Before reaching out to support, we recommend heading to our [FAQ](https://www.algolia.com/doc/api-client/troubleshooting/faq/javascript/) where you will find answers to the most common issues and gotchas with the client.

You can also [open an issue on GitHub](https://github.com/algolia/api-clients-automation/issues/new/choose).

## 📄 License

Algolia API clients automation is an open-sourced software licensed under the [MIT license](LICENSE.md).
