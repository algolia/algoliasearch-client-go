<p align="center">
  <a href="https://www.algolia.com">
    <img alt="Algolia for Go" src="https://raw.githubusercontent.com/algolia/algoliasearch-client-common/master/banners/go.png" >
  </a>

  <h4 align="center">The perfect starting point to integrate <a href="https://algolia.com" target="_blank">Algolia</a> within your Go project</h4>

  <p align="center">
    <a href="https://github.com/algolia/algoliasearch-client-go/tags"><img src="https://img.shields.io/github/tag/algolia/algoliasearch-client-go.svg?include_prereleases&sort=semver" alt="Github Releases"></img></a>
    <a href="https://pkg.go.dev/github.com/algolia/algoliasearch-client-go/v4"><img src="https://pkg.go.dev/badge/github.com/algolia/algoliasearch-client-go/v4" alt="GoDoc"></img></a>
    <a href="https://goreportcard.com/report/github.com/algolia/algoliasearch-client-go"><img src="https://goreportcard.com/badge/github.com/algolia/algoliasearch-client-go" alt="Go Report Card"></img></a>
    <a href="https://github.com/algolia/algoliasearch-client-go/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></img></a>
    <img src="https://img.shields.io/badge/Go-%3E=1.21-green.svg" alt="Supported version"></img></a>
  </p>
</p>

<p align="center">
  <a href="https://www.algolia.com/doc/libraries/sdk/install#go" target="_blank">Documentation</a>  •
  <a href="https://discourse.algolia.com" target="_blank">Community Forum</a>  •
  <a href="http://stackoverflow.com/questions/tagged/algolia" target="_blank">Stack Overflow</a>  •
  <a href="https://github.com/algolia/algoliasearch-client-go/issues" target="_blank">Report a bug</a>  •
  <a href="https://alg.li/support" target="_blank">Support</a>
</p>

## ✨ Features

* Support Go 1.19 and above
* Typed requests and responses
* First-class support for user-defined structures
* Injectable HTTP client

## 💡 Getting Started

First, install the Algolia API Go Client via the go get command:

```bash
go get github.com/algolia/algoliasearch-client-go/v4
```

You can now import the Algolia API client in your project and play with it.


```go
import "github.com/algolia/algoliasearch-client-go/v4/algolia/search"

client, err := search.NewClient("YOUR_APP_ID", "YOUR_API_KEY")

// Add a new record to your Algolia index
response, err := client.SaveObject(client.NewApiSaveObjectRequest(
  "<YOUR_INDEX_NAME>", map[string]any{"objectID": "id", "test": "val"},
))
if err != nil {
  // handle the eventual error
  panic(err)
}

// use the model directly
print(response)

// Poll the task status to know when it has been indexed
taskResponse, err := searchClient.WaitForTask("<YOUR_INDEX_NAME>", response.TaskID, nil, nil, nil)
if err != nil {
  panic(err)
}

// Fetch search results, with typo tolerance
response, err := client.Search(client.NewApiSearchRequest(

  search.NewEmptySearchMethodParams().SetRequests(
    []search.SearchQuery{*search.SearchForHitsAsSearchQuery(
      search.NewEmptySearchForHits().SetIndexName("<YOUR_INDEX_NAME>").SetQuery("<YOUR_QUERY>").SetHitsPerPage(50))}),
))
if err != nil {
  // handle the eventual error
  panic(err)
}

// use the model directly
print(response)
```

For full documentation, visit the **[Algolia Go API Client](https://www.algolia.com/doc/libraries/sdk/install#go)**.

## ❓ Troubleshooting

Encountering an issue? Before reaching out to support, we recommend heading to our [FAQ](https://support.algolia.com/hc/sections/15061037630609-API-Client-FAQs) where you will find answers for the most common issues and gotchas with the client. You can also open [a GitHub issue](https://github.com/algolia/api-clients-automation/issues/new?assignees=&labels=&projects=&template=Bug_report.md)

## Contributing

This repository hosts the code of the generated Algolia API client for Go, if you'd like to contribute, head over to the [main repository](https://github.com/algolia/api-clients-automation). You can also find contributing guides on [our documentation website](https://api-clients-automation.netlify.app/docs/introduction).

## 📄 License

The Algolia Go API Client is an open-sourced software licensed under the [MIT license](LICENSE).
