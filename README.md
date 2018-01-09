# Algolia Search API Client for Go

[Algolia Search](https://www.algolia.com) is a hosted full-text, numerical,
and faceted search engine capable of delivering realtime results from the first keystroke.

The **Algolia Search API Client for Go** lets
you easily use the [Algolia Search REST API](https://www.algolia.com/doc/rest-api/search) from
your Go code.

[![Build Status](https://travis-ci.org/algolia/algoliasearch-client-go.png?branch=master)](https://travis-ci.org/algolia/algoliasearch-client-go) ![Supported version](https://img.shields.io/badge/Go-%3E=1.5-green.svg)


**Migration note from v1.x to v2.x**

In June 2016, we released the v2 of our Go client. If you were using version 1.x of the client, read the [migration guide to version 2.x](https://github.com/algolia/algoliasearch-client-go/wiki/Migration-guide-to-version-2.x).
Version 1.x are no longer under active development. They are still supported for bug fixes, but will not receive new features.




## API Documentation

You can find the full reference on [Algolia's website](https://www.algolia.com/doc/api-client/go/).


## Table of Contents



1. **[Install](#install)**


1. **[Quick Start](#quick-start)**


1. **[Push data](#push-data)**


1. **[Configure](#configure)**


1. **[Search](#search)**


1. **[Search UI](#search-ui)**


1. **[List of available methods](#list-of-available-methods)**


# Getting Started



## Install

Download AlgoliaSearch using

```bash
go get github.com/algolia/algoliasearch-client-go/algoliasearch
```

## Quick Start

In 30 seconds, this quick start tutorial will show you how to index and search objects.

### Initialize the client

To begin, you will need to initialize the client. In order to do this you will need your **Application ID** and **API Key**.
You can find both on [your Algolia account](https://www.algolia.com/api-keys).

```go
import "github.com/algolia/algoliasearch-client-go/algoliasearch"

client := algoliasearch.NewClient("YourApplicationID", "YourAPIKey")
index := client.InitIndex("your_index_name")
```

## Push data

Without any prior configuration, you can start indexing [500 contacts](https://github.com/algolia/datasets/blob/master/contacts/contacts.json) in the ```contacts``` index using the following code:
```go
index := client.InitIndex("contacts")
content, _ := ioutil.ReadFile("contacts.json")

var objects []algoliasearch.Object
if err := json.Unmarshal(content, &objects); err != nil {
  return
}

res, err := index.AddObjects(objects)
```

## Configure

Settings can be customized to fine tune the search behavior. For example, you can add a custom sort by number of followers to further enhance the built-in relevance:

```go
settings := algoliasearch.Map{
  "customRanking": []string{"desc(followers)"},
}

res, err := index.SetSettings(settings)
```

You can also configure the list of attributes you want to index by order of importance (most important first).

**Note:** The Algolia engine is designed to suggest results as you type, which means you'll generally search by prefix.
In this case, the order of attributes is very important to decide which hit is the best:

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

## Search

You can now search for contacts using `firstname`, `lastname`, `company`, etc. (even with typos):

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

## Search UI

**Warning:** If you are building a web application, you may be more interested in using one of our
[frontend search UI libraries](https://www.algolia.com/doc/guides/search-ui/search-libraries/)

The following example shows how to build a front-end search quickly using
[InstantSearch.js](https://community.algolia.com/instantsearch.js/)

### index.html

```html
<!doctype html>
<head>
  <meta charset="UTF-8">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/instantsearch.js@2.3/dist/instantsearch.min.css">
  <!-- Always use `2.x` versions in production rather than `2` to mitigate any side effects on your website,
  Find the latest version on InstantSearch.js website: https://community.algolia.com/instantsearch.js/v2/guides/usage.html -->
</head>
<body>
  <header>
    <div>
       <input id="search-input" placeholder="Search for products">
       <!-- We use a specific placeholder in the input to guides users in their search. -->
    
  </header>
  <main>
      
      
  </main>

  <script type="text/html" id="hit-template">
    
      <p class="hit-name">{{{_highlightResult.firstname.value}}} {{{_highlightResult.lastname.value}}}</p>
    
  </script>

  <script src="https://cdn.jsdelivr.net/npm/instantsearch.js@2.3/dist/instantsearch.min.js"></script>
  <script src="app.js"></script>
</body>
```

### app.js

```js
var search = instantsearch({
  // Replace with your own values
  appId: 'YourApplicationID',
  apiKey: 'YourSearchOnlyAPIKey', // search only API key, no ADMIN key
  indexName: 'contacts',
  urlSync: true,
  searchParameters: {
    hitsPerPage: 10
  }
});

search.addWidget(
  instantsearch.widgets.searchBox({
    container: '#search-input'
  })
);

search.addWidget(
  instantsearch.widgets.hits({
    container: '#hits',
    templates: {
      item: document.getElementById('hit-template').innerHTML,
      empty: "We didn't find any results for the search <em>\"{{query}}\"</em>"
    }
  })
);

search.start();
```




## List of available methods







### Search

- [Search an index](https://algolia.com/doc/api-reference/api-methods/search/?language=go)
- [Search for facet values](https://algolia.com/doc/api-reference/api-methods/search-for-facet-values/?language=go)
- [Search multiple indexes](https://algolia.com/doc/api-reference/api-methods/multiple-queries/?language=go)
- [Browse an index](https://algolia.com/doc/api-reference/api-methods/browse/?language=go)





### Indexing

- [Add objects](https://algolia.com/doc/api-reference/api-methods/add-objects/?language=go)
- [Update objects](https://algolia.com/doc/api-reference/api-methods/update-objects/?language=go)
- [Partial update objects](https://algolia.com/doc/api-reference/api-methods/partial-update-objects/?language=go)
- [Delete objects](https://algolia.com/doc/api-reference/api-methods/delete-objects/?language=go)
- [Delete by query](https://algolia.com/doc/api-reference/api-methods/delete-by-query/?language=go)
- [Get objects](https://algolia.com/doc/api-reference/api-methods/get-objects/?language=go)
- [Custom batch](https://algolia.com/doc/api-reference/api-methods/batch/?language=go)
- [Wait for operations](https://algolia.com/doc/api-reference/api-methods/wait-task/?language=go)





### Settings

- [Get settings](https://algolia.com/doc/api-reference/api-methods/get-settings/?language=go)
- [Set settings](https://algolia.com/doc/api-reference/api-methods/set-settings/?language=go)





### Manage indices

- [List indexes](https://algolia.com/doc/api-reference/api-methods/list-indices/?language=go)
- [Delete index](https://algolia.com/doc/api-reference/api-methods/delete-index/?language=go)
- [Copy index](https://algolia.com/doc/api-reference/api-methods/copy-index/?language=go)
- [Move index](https://algolia.com/doc/api-reference/api-methods/move-index/?language=go)
- [Clear index](https://algolia.com/doc/api-reference/api-methods/clear-index/?language=go)





### API Keys

- [Create secured API Key](https://algolia.com/doc/api-reference/api-methods/generate-secured-api-key/?language=go)
- [Add API Key](https://algolia.com/doc/api-reference/api-methods/add-api-key/?language=go)
- [Update API Key](https://algolia.com/doc/api-reference/api-methods/update-api-key/?language=go)
- [Delete API Key](https://algolia.com/doc/api-reference/api-methods/delete-api-key/?language=go)
- [Get API Key permissions](https://algolia.com/doc/api-reference/api-methods/get-api-key/?language=go)
- [List API Keys](https://algolia.com/doc/api-reference/api-methods/list-api-keys/?language=go)





### Synonyms

- [Save synonym](https://algolia.com/doc/api-reference/api-methods/save-synonym/?language=go)
- [Batch synonyms](https://algolia.com/doc/api-reference/api-methods/batch-synonyms/?language=go)
- [Delete synonym](https://algolia.com/doc/api-reference/api-methods/delete-synonym/?language=go)
- [Clear all synonyms](https://algolia.com/doc/api-reference/api-methods/clear-synonyms/?language=go)
- [Get synonym](https://algolia.com/doc/api-reference/api-methods/get-synonym/?language=go)
- [Search synonyms](https://algolia.com/doc/api-reference/api-methods/search-synonyms/?language=go)
- [Export Synonyms](https://algolia.com/doc/api-reference/api-methods/export-synonyms/?language=go)





### Query rules

- [Save a single rule](https://algolia.com/doc/api-reference/api-methods/rules-save/?language=go)
- [Batch save multiple rules](https://algolia.com/doc/api-reference/api-methods/rules-save-batch/?language=go)
- [Read a rule](https://algolia.com/doc/api-reference/api-methods/rules-read/?language=go)
- [Delete a single rule](https://algolia.com/doc/api-reference/api-methods/rules-delete/?language=go)
- [Clear all rules](https://algolia.com/doc/api-reference/api-methods/rules-clear/?language=go)
- [Search for rules](https://algolia.com/doc/api-reference/api-methods/rules-search/?language=go)
- [Export rules](https://algolia.com/doc/api-reference/api-methods/rules-export/?language=go)








### Advanced

- [Get latest logs](https://algolia.com/doc/api-reference/api-methods/get-logs/?language=go)



## Getting Help

- **Need help**? Ask a question to the [Algolia Community](https://discourse.algolia.com/) or on [Stack Overflow](http://stackoverflow.com/questions/tagged/algolia).
- **Found a bug?** You can open a [GitHub issue](https://github.com/algolia/algoliasearch-client-go/issues).

