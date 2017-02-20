# Algolia Search API Client for Go

[Algolia Search](https://www.algolia.com) is a hosted full-text, numerical, and faceted search engine capable of delivering realtime results from the first keystroke.
The **Algolia Search API Client for Go** lets you easily use the [Algolia Search REST API](https://www.algolia.com/doc/rest-api/search) from your Go code.

[![Build Status](https://travis-ci.org/algolia/algoliasearch-client-go.png?branch=master)](https://travis-ci.org/algolia/algoliasearch-client-go) ![Supported version](https://img.shields.io/badge/Go-%3E=1.5-green.svg)


**Migration note from v1.x to v2.x**

In June 2016, we released the v2 of our Go client. If you were using version 1.x of the client, read the [migration guide to version 2.x](https://github.com/algolia/algoliasearch-client-go/wiki/Migration-guide-to-version-2.x).
Version 1.x are no longer under active development. They are still supported for bug fixes, but will not receive new features.




## API Documentation

You can find the full reference on [Algolia's website](https://www.algolia.com/doc/api-client/go/).


## Table of Contents


1. **[Install](#install)**


1. **[Quick Start](#quick-start)**

    * [Initialize the client](#initialize-the-client)
    * [Push data](#push-data)
    * [Search](#search)
    * [Configure](#configure)
    * [Frontend search](#frontend-search)

1. **[Getting Help](#getting-help)**





# Getting Started



## Install

Download AlgoliaSearch using

```bash
go get github.com/algolia/algoliasearch-client-go/algoliasearch
```

## Quick Start

In 30 seconds, this quick start tutorial will show you how to index and search objects.

### Initialize the client

You first need to initialize the client. For that you need your **Application ID** and **API Key**.
You can find both of them on [your Algolia account](https://www.algolia.com/api-keys).

```go
import "github.com/algolia/algoliasearch-client-go/algoliasearch"

client := algoliasearch.NewClient("YourApplicationID", "YourAPIKey")
```

### Push data

Without any prior configuration, you can start indexing [500 contacts](https://github.com/algolia/algoliasearch-client-csharp/blob/master/contacts.json) in the ```contacts``` index using the following code:
```go
index := client.InitIndex("contacts")
content, _ := ioutil.ReadFile("contacts.json")

var objects []algoliasearch.Object
if err := json.Unmarshal(content, &objects); err != nil {
  return
}

res, err := index.AddObjects(objects)
```

### Search

You can now search for contacts using firstname, lastname, company, etc. (even with typos):

```go
// Search by firstname
res, err := index.Search("jimmie", nil)

// Search by firstname with typo
res, err = index.Search("jimie", nil)

// Search for a company
res, err = index.Search("california paint", nil)

// Search for a firstname & company
res, err = index.Search("jimmie paint", nil)
```

### Configure

Settings can be customized to tune the search behavior. For example, you can add a custom sort by number of followers to the already great built-in relevance:

```go
settings := algoliasearch.Map{
  "customRanking": []string{"desc(followers)"},
}

res, err := index.SetSettings(settings)
```

You can also configure the list of attributes you want to index by order of importance (first = most important):

**Note:** Since the engine is designed to suggest results as you type, you'll generally search by prefix.
In this case the order of attributes is very important to decide which hit is the best:

```go
settings := algoliasearch.Map{
  "searchableAttributes": []string{
	"firstname",
	"lastname",
	"company",
	"email",
	"city",
	"address",
  },
}

res, err := index.SetSettings(settings)
```

### Frontend search

**Note:** If you are building a web application, you may be more interested in using our [JavaScript client](https://github.com/algolia/algoliasearch-client-javascript) to perform queries.

It brings two benefits:
  * Your users get a better response time by not going through your servers
  * It will offload unnecessary tasks from your servers

```html
<script src="https://cdn.jsdelivr.net/algoliasearch/3/algoliasearch.min.js"></script>
<script>
var client = algoliasearch('ApplicationID', 'apiKey');
var index = client.initIndex('indexName');

// perform query "jim"
index.search('jim', searchCallback);

// the last optional argument can be used to add search parameters
index.search(
  'jim', {
    hitsPerPage: 5,
    facets: '*',
    maxValuesPerFacet: 10
  },
  searchCallback
);

function searchCallback(err, content) {
  if (err) {
    console.error(err);
    return;
  }

  console.log(content);
}
</script>
```

## Getting Help

- **Need help**? Ask a question to the [Algolia Community](https://discourse.algolia.com/) or on [Stack Overflow](http://stackoverflow.com/questions/tagged/algolia).
- **Found a bug?** You can open a [GitHub issue](https://github.com/algolia/algoliasearch-client-go/issues).



