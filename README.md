# Algolia Search API Client for Go

[Algolia Search](https://www.algolia.com) is a hosted full-text, numerical,
and faceted search engine capable of delivering realtime results from the first keystroke.

The **Algolia Search API Client for Go** lets
you easily use the [Algolia Search REST API](https://www.algolia.com/doc/rest-api/search) from
your Go code.

[![Build Status](https://travis-ci.org/algolia/algoliasearch-client-go.png?branch=master)](https://travis-ci.org/algolia/algoliasearch-client-go) ![Supported version](https://img.shields.io/badge/Go-%3E=1.7-green.svg)


**Migration note from v1.x to v2.x**

In June 2016, we released the v2 of our Go client. If you were using version 1.x of the client, read the [migration guide to version 2.x](https://github.com/algolia/algoliasearch-client-go/wiki/Migration-guide-to-version-2.x).
Version 1.x are no longer under active development. They are still supported for bug fixes, but will not receive new features.




## API Documentation

You can find the full reference on [Algolia's website](https://www.algolia.com/doc/api-client/go/).



1. **[Supported platforms](#supported-platforms)**


1. **[Install](#install)**


1. **[Quick Start](#quick-start)**


1. **[Push data](#push-data)**


1. **[Configure](#configure)**


1. **[Search](#search)**


1. **[Search UI](#search-ui)**


1. **[List of available methods](#list-of-available-methods)**


# Getting Started



## Supported platforms

This API client is compatible with Go 1.7 and above.

## Install

Download the Go client using:

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
// Search for a first name
res, err := index.Search("jimmie", nil)

// Search for a first name with typo
res, err = index.Search("jimie", nil)

// Search for a company
res, err = index.Search("california paint", nil)

// Search for a first name and a company
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
  routing: true,
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





### Personalization





### Search

- [Search index](https://algolia.com/doc/api-reference/api-methods/search/?language=go)
- [Search for facet values](https://algolia.com/doc/api-reference/api-methods/search-for-facet-values/?language=go)
- [Search multiple indices](https://algolia.com/doc/api-reference/api-methods/multiple-queries/?language=go)
- [Browse index](https://algolia.com/doc/api-reference/api-methods/browse/?language=go)




### Indexing

- [Add objects](https://algolia.com/doc/api-reference/api-methods/add-objects/?language=go)
- [Save objects](https://algolia.com/doc/api-reference/api-methods/save-objects/?language=go)
- [Partial update objects](https://algolia.com/doc/api-reference/api-methods/partial-update-objects/?language=go)
- [Delete objects](https://algolia.com/doc/api-reference/api-methods/delete-objects/?language=go)
- [Replace all objects](https://algolia.com/doc/api-reference/api-methods/replace-all-objects/?language=go)
- [Delete by](https://algolia.com/doc/api-reference/api-methods/delete-by/?language=go)
- [Clear objects](https://algolia.com/doc/api-reference/api-methods/clear-objects/?language=go)
- [Get objects](https://algolia.com/doc/api-reference/api-methods/get-objects/?language=go)
- [Custom batch](https://algolia.com/doc/api-reference/api-methods/batch/?language=go)




### Settings

- [Get settings](https://algolia.com/doc/api-reference/api-methods/get-settings/?language=go)
- [Set settings](https://algolia.com/doc/api-reference/api-methods/set-settings/?language=go)
- [Copy settings](https://algolia.com/doc/api-reference/api-methods/copy-settings/?language=go)




### Manage indices

- [List indices](https://algolia.com/doc/api-reference/api-methods/list-indices/?language=go)
- [Delete index](https://algolia.com/doc/api-reference/api-methods/delete-index/?language=go)
- [Copy index](https://algolia.com/doc/api-reference/api-methods/copy-index/?language=go)
- [Move index](https://algolia.com/doc/api-reference/api-methods/move-index/?language=go)




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
- [Replace all synonyms](https://algolia.com/doc/api-reference/api-methods/replace-all-synonyms/?language=go)
- [Copy synonyms](https://algolia.com/doc/api-reference/api-methods/copy-synonyms/?language=go)
- [Export Synonyms](https://algolia.com/doc/api-reference/api-methods/export-synonyms/?language=go)




### Query rules

- [Save rule](https://algolia.com/doc/api-reference/api-methods/save-rule/?language=go)
- [Batch rules](https://algolia.com/doc/api-reference/api-methods/batch-rules/?language=go)
- [Get rule](https://algolia.com/doc/api-reference/api-methods/get-rule/?language=go)
- [Delete rule](https://algolia.com/doc/api-reference/api-methods/delete-rule/?language=go)
- [Clear rules](https://algolia.com/doc/api-reference/api-methods/clear-rules/?language=go)
- [Search rules](https://algolia.com/doc/api-reference/api-methods/search-rules/?language=go)
- [Replace all rules](https://algolia.com/doc/api-reference/api-methods/replace-all-rules/?language=go)
- [Copy rules](https://algolia.com/doc/api-reference/api-methods/copy-rules/?language=go)
- [Export rules](https://algolia.com/doc/api-reference/api-methods/export-rules/?language=go)




### A/B Test

- [Add A/B test](https://algolia.com/doc/api-reference/api-methods/add-ab-test/?language=go)
- [Get A/B test](https://algolia.com/doc/api-reference/api-methods/get-ab-test/?language=go)
- [List A/B tests](https://algolia.com/doc/api-reference/api-methods/list-ab-tests/?language=go)
- [Stop A/B test](https://algolia.com/doc/api-reference/api-methods/stop-ab-test/?language=go)
- [Delete A/B test](https://algolia.com/doc/api-reference/api-methods/delete-ab-test/?language=go)




### MultiClusters

- [Assign or Move userID](https://algolia.com/doc/api-reference/api-methods/assign-user-id/?language=go)
- [Get top userID](https://algolia.com/doc/api-reference/api-methods/get-top-user-id/?language=go)
- [Get userID](https://algolia.com/doc/api-reference/api-methods/get-user-id/?language=go)
- [List clusters](https://algolia.com/doc/api-reference/api-methods/list-clusters/?language=go)
- [List userIDs](https://algolia.com/doc/api-reference/api-methods/list-user-id/?language=go)
- [Remove userID](https://algolia.com/doc/api-reference/api-methods/remove-user-id/?language=go)
- [Search userID](https://algolia.com/doc/api-reference/api-methods/search-user-id/?language=go)




### Advanced

- [Get logs](https://algolia.com/doc/api-reference/api-methods/get-logs/?language=go)
- [Configuring timeouts](https://algolia.com/doc/api-reference/api-methods/configuring-timeouts/?language=go)
- [Set extra header](https://algolia.com/doc/api-reference/api-methods/set-extra-header/?language=go)
- [Wait for operations](https://algolia.com/doc/api-reference/api-methods/wait-task/?language=go)





## Getting Help

- **Need help**? Ask a question to the [Algolia Community](https://discourse.algolia.com/) or on [Stack Overflow](http://stackoverflow.com/questions/tagged/algolia).
- **Found a bug?** You can open a [GitHub issue](https://github.com/algolia/algoliasearch-client-go/issues).

