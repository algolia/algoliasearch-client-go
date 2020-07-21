<p align="center">
  <a href="https://www.algolia.com">
    <img alt="Algolia for Go" src="https://raw.githubusercontent.com/algolia/algoliasearch-client-common/master/banners/go.png" >
  </a>

  <h4 align="center">The perfect starting point to integrate <a href="https://algolia.com" target="_blank">Algolia</a> within your Go project</h4>

  <p align="center">
    <a href="https://circleci.com/gh/algolia/algoliasearch-client-go"><img src="https://circleci.com/gh/algolia/algoliasearch-client-go.svg?style=shield" alt="CircleCI" /></a>
    <a href="https://github.com/algolia/algoliasearch-client-go/releases"><img src="https://img.shields.io/github/tag/algolia/algoliasearch-client-go.svg" alt="Github Releases"></img></a>
    <a href="https://godoc.org/github.com/algolia/algoliasearch-client-go"><img src="https://godoc.org/github.com/algolia/algoliasearch-client-go?status.svg" alt="GoDoc"></img></a>
    <a href="https://goreportcard.com/report/github.com/algolia/algoliasearch-client-go"><img src="https://goreportcard.com/badge/github.com/algolia/algoliasearch-client-go" alt="Go Report Card"></img></a>
    <a href="https://github.com/algolia/algoliasearch-client-go/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></img></a>
    <img src="https://img.shields.io/badge/Go-%3E=1.11-green.svg" alt="Supported version"></img></a>
  </p>
</p>

<p align="center">
  <a href="https://www.algolia.com/doc/api-client/getting-started/install/go/" target="_blank">Documentation</a>  •
  <a href="https://discourse.algolia.com" target="_blank">Community Forum</a>  •
  <a href="http://stackoverflow.com/questions/tagged/algolia" target="_blank">Stack Overflow</a>  •
  <a href="https://github.com/algolia/algoliasearch-client-go/issues" target="_blank">Report a bug</a>  •
  <a href="https://www.algolia.com/doc/api-client/troubleshooting/faq/go/" target="_blank">FAQ</a>  •
  <a href="https://www.algolia.com/support" target="_blank">Support</a>
</p>

## ✨ Features

* Support Go 1.11 and above
* Typed requests and responses
* First-class support for user-defined structures
* Injectable HTTP client

**Migration note from v2.x to v3.x**

> In June 2019, we released v3 of our Go client. If you are using version 2.x
> of the client, read the [migration guide to version 3.x](https://www.algolia.com/doc/api-client/getting-started/upgrade-guides/go/).
> Version 2.x will **no longer** be under active development.

## 💡 Getting Started

First, add the Algolia Go API Client as a new module to your Go project:

```bash
# First line is optional if your project is already defined as a Go module
go mod init <YOUR_PROJECT_NAME>
go get github.com/algolia/algoliasearch-client-go/v3@v3.4.0
```

Then, create objects on your index:

```go
package main

import "github.com/algolia/algoliasearch-client-go/v3/algolia/search"

type Contact struct {
	ObjectID string `json:"objectID"`
	Name     string `json:"name"`
}

func main() {
	client := search.NewClient("YourApplicationID", "YourAPIKey")
	index := client.InitIndex("your_index_name")

	res, err := index.SaveObjects([]Contact{
		{ObjectID: "1", Name: "Foo"},
	})
}
```

Finally, you may begin searching a object using the `Search` method:

```go
package main

import (
	"fmt"
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Contact struct {
	ObjectID string `json:"objectID"`
	Name     string `json:"name"`
}

func main() {
	client := search.NewClient("YourApplicationID", "YourAPIKey")
	index := client.InitIndex("your_index_name")

	res, err := index.Search("Foo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var contacts []Contact

	err = res.UnmarshalHits(&contacts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(contacts)
}
```

For full documentation, visit the **[Algolia Go API Client](https://www.algolia.com/doc/api-client/getting-started/install/go/)**.

## ❓ Troubleshooting

Encountering an issue? Before reaching out to support, we recommend heading to our [FAQ](https://www.algolia.com/doc/api-client/troubleshooting/faq/go/) where you will find answers for the most common issues and gotchas with the client.

## 📄 License

Algolia Go API Client is an open-sourced software licensed under the [MIT license](LICENSE).
