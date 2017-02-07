# Algolia Search API Client for Go

[Algolia Search](https://www.algolia.com) is a hosted full-text, numerical, and faceted search engine capable of delivering realtime results from the first keystroke.
The **Algolia Search API Client for Go** lets you easily use the [Algolia Search REST API](https://www.algolia.com/doc/rest-api/search) from your Go code.

**Migration note from v1.x to v2.x**

In June 2016, we released the v2 of our Go client. If you were using version 1.x of the client, read the [migration guide to version 2.x](https://github.com/algolia/algoliasearch-client-go/wiki/Migration-guide-to-version-2.x).
Version 1.x are no longer under active development. They are still supported for bug fixes, but will not receive new features.

[![Build Status](https://travis-ci.org/algolia/algoliasearch-client-go.png?branch=master)](https://travis-ci.org/algolia/algoliasearch-client-go) ![Supported version](https://img.shields.io/badge/Go-%3E=1.5-green.svg)


**Note:** An easier-to-read version of this documentation is available on
[Algolia's website](https://www.algolia.com/doc/api-client/go/).

# Table of Contents


**Getting Started**

1. [Install](#install)
1. [Init index - `InitIndex`](#init-index---initindex)
1. [Quick Start](#quick-start)

**Search**

1. [Search an index - `Search`](#search-an-index---search)
1. [Search Response Format](#search-response-format)
1. [Search Parameters](#search-parameters)
1. [Search multiple indices - `MultipleQueries`](#search-multiple-indices---multiplequeries)
1. [Get Objects - `GetObjects`](#get-objects---getobjects)
1. [Search for facet values - `SearchForFacetValues`](#search-for-facet-values---searchforfacetvalues)

**Indexing**

1. [Add Objects - `AddObjects`](#add-objects---addobjects)
1. [Update objects - `UpdateObjects`](#update-objects---updateobjects)
1. [Partial update objects - `PartialUpdateObjects`](#partial-update-objects---partialupdateobjects)
1. [Delete objects - `DeleteObjects`](#delete-objects---deleteobjects)
1. [Delete by query - `DeleteByQuery`](#delete-by-query---deletebyquery)
1. [Wait for operations - `WaitTask`](#wait-for-operations---waittask)

**Settings**

1. [Get settings - `GetSettings`](#get-settings---getsettings)
1. [Set settings - `SetSettings`](#set-settings---setsettings)
1. [Index settings parameters](#index-settings-parameters)

**Parameters**

1. [Overview](#overview)
1. [Search](#search)
1. [Attributes](#attributes)
1. [Ranking](#ranking)
1. [Filtering / Faceting](#filtering--faceting)
1. [Highlighting / Snippeting](#highlighting--snippeting)
1. [Pagination](#pagination)
1. [Typos](#typos)
1. [Geo-Search](#geo-search)
1. [Query Strategy](#query-strategy)
1. [Performance](#performance)
1. [Advanced](#advanced)

**Manage Indices**

1. [Create an index](#create-an-index)
1. [List indices - `ListIndexes`](#list-indices---listindexes)
1. [Delete an index - `DeleteIndex`](#delete-an-index---deleteindex)
1. [Clear an index - `ClearIndex`](#clear-an-index---clearindex)
1. [Copy index - `CopyIndex`](#copy-index---copyindex)
1. [Move index - `MoveIndex`](#move-index---moveindex)

**Api keys**

1. [Overview](#overview)
1. [Generate key - `GenerateSecuredApiKey`](#generate-key---generatesecuredapikey)

**Synonyms**

1. [Save synonym - `AddSynonym`](#save-synonym---addsynonym)
1. [Batch synonyms - `BatchSynonyms`](#batch-synonyms---batchsynonyms)
1. [Editing Synonyms](#editing-synonyms)
1. [Delete synonym - `DeleteSynonym`](#delete-synonym---deletesynonym)
1. [Clear all synonyms - `ClearSynonyms`](#clear-all-synonyms---clearsynonyms)
1. [Get synonym - `GetSynonym`](#get-synonym---getsynonym)
1. [Search synonyms - `SearchSynonyms`](#search-synonyms---searchsynonyms)

**Advanced**

1. [Custom batch - `Batch`](#custom-batch---batch)
1. [Backup / Export an index - `Browse`](#backup--export-an-index---browse)
1. [List api keys - `ListKeys`](#list-api-keys---listkeys)
1. [Add user key - `AddUserKey`](#add-user-key---adduserkey)
1. [Update user key - `UpdateUserKey`](#update-user-key---updateuserkey)
1. [Delete user key - `DeleteUserKey`](#delete-user-key---deleteuserkey)
1. [Get key permissions - `GetUserKey`](#get-key-permissions---getuserkey)
1. [Get latest logs - `GetLogs`](#get-latest-logs---getlogs)
1. [REST API](#rest-api)


# Guides & Tutorials

Check our [online guides](https://www.algolia.com/doc):

* [Data Formatting](https://www.algolia.com/doc/indexing/formatting-your-data)
* [Import and Synchronize data](https://www.algolia.com/doc/indexing/import-synchronize-data/php)
* [Autocomplete](https://www.algolia.com/doc/search/auto-complete)
* [Instant search page](https://www.algolia.com/doc/search/instant-search)
* [Filtering and Faceting](https://www.algolia.com/doc/search/filtering-faceting)
* [Sorting](https://www.algolia.com/doc/relevance/sorting)
* [Ranking Formula](https://www.algolia.com/doc/relevance/ranking)
* [Typo-Tolerance](https://www.algolia.com/doc/relevance/typo-tolerance)
* [Geo-Search](https://www.algolia.com/doc/geo-search/geo-search-overview)
* [Security](https://www.algolia.com/doc/security/best-security-practices)
* [API-Keys](https://www.algolia.com/doc/security/api-keys)
* [REST API](https://www.algolia.com/doc/rest)


# Getting Started



## Install

Download AlgoliaSearch using

```bash
go get github.com/algolia/algoliasearch-client-go/algoliasearch
```

## Init index - `InitIndex` 

To initialize the client, you need your **Application ID** and **API Key**. You can find both of them on [your Algolia account](https://www.algolia.com/api-keys).

```go
import "github.com/algolia/algoliasearch-client-go/algoliasearch"

client := algoliasearch.NewClient("YourApplicationID", "YourAPIKey")
```

## Quick Start

In 30 seconds, this quick start tutorial will show you how to index and search objects.

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

Settings can be customized to tune the search behavior. For example, you can add a custom sort by number of followers to the already great built-in relevance:

```go
settings := algoliasearch.Map{
  "customRanking": []string{"desc(followers)"},
}

res, err := index.SetSettings(settings)
```

You can also configure the list of attributes you want to index by order of importance (first = most important):

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

Since the engine is designed to suggest results as you type, you'll generally search by prefix. In this case the order of attributes is very important to decide which hit is the best:

```go
res, err := index.Search("or", nil)
res, err = index.Search("jim", nil)
```

**Note:** **Note:** If you are building a web application, you may be more interested in using our [JavaScript client](https://github.com/algolia/algoliasearch-client-javascript) to perform queries.

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


# Search



## Search an index - `Search` 

**Notes:** If you are building a web application, you may be more interested in using our [JavaScript client](https://github.com/algolia/algoliasearch-client-javascript) to perform queries. It brings two benefits:
  * Your users get a better response time by not going through your servers
  * It will offload unnecessary tasks from your servers.

To perform a search, you only need to initialize the index and perform a call to the search function.

The search query allows only to retrieve 1000 hits. If you need to retrieve more than 1000 hits (e.g. for SEO), you can use [Backup / Export an index](#backup--export-an-index).

```go
params := algoliasearch.Map{
  "attributesToRetrieve": []string{"firstname", "lastname"},
  "hitsPerPage":          50,
}

res, err = index.Search("jimmie paint", params)
```

## Search Response Format

### Sample

The server response will look like:

```json
{
  "hits": [
    {
      "firstname": "Jimmie",
      "lastname": "Barninger",
      "objectID": "433",
      "_highlightResult": {
        "firstname": {
          "value": "<em>Jimmie</em>",
          "matchLevel": "partial"
        },
        "lastname": {
          "value": "Barninger",
          "matchLevel": "none"
        },
        "company": {
          "value": "California <em>Paint</em> & Wlpaper Str",
          "matchLevel": "partial"
        }
      }
    }
  ],
  "page": 0,
  "nbHits": 1,
  "nbPages": 1,
  "hitsPerPage": 20,
  "processingTimeMS": 1,
  "query": "jimmie paint",
  "params": "query=jimmie+paint&attributesToRetrieve=firstname,lastname&hitsPerPage=50"
}
```

### Fields

- `hits` (array): The hits returned by the search, sorted according to the ranking formula.

    Hits are made of the JSON objects that you stored in the index; therefore, they are mostly schema-less. However, Algolia does enrich them with a few additional fields:

    - `_highlightResult` (object, optional): Highlighted attributes. *Note: Only returned when [attributesToHighlight](#attributestohighlight) is non-empty.*

        - `${attribute_name}` (object): Highlighting for one attribute.

            - `value` (string): Markup text with occurrences highlighted. The tags used for highlighting are specified via [highlightPreTag](#highlightpretag) and [highlightPostTag](#highlightposttag).

            - `matchLevel` (string, enum) = {`none` \| `partial` \| `full`}: Indicates how well the attribute matched the search query.

            - `matchedWords` (array): List of words *from the query* that matched the object.

            - `fullyHighlighted` (boolean): Whether the entire attribute value is highlighted.

    - `_snippetResult` (object, optional): Snippeted attributes. *Note: Only returned when [attributesToSnippet](#attributestosnippet) is non-empty.*

        - `${attribute_name}` (object): Snippeting for the corresponding attribute.

            - `value` (string): Markup text with occurrences highlighted and optional ellipsis indicators. The tags used for highlighting are specified via [highlightPreTag](#highlightpretag) and [highlightPostTag](#highlightposttag). The text used to indicate ellipsis is specified via [snippetEllipsisText](#snippetellipsistext).

            - `matchLevel` (string, enum) = {`none` \| `partial` \| `full`}: Indicates how well the attribute matched the search query.

    - `_rankingInfo` (object, optional): Ranking information. *Note: Only returned when [getRankingInfo](#getrankinginfo) is `true`.*

        - `nbTypos` (integer): Number of typos encountered when matching the record. Corresponds to the `typos` ranking criterion in the ranking formula.

        - `firstMatchedWord` (integer): Position of the most important matched attribute in the attributes to index list. Corresponds to the `attribute` ranking criterion in the ranking formula.

        - `proximityDistance` (integer): When the query contains more than one word, the sum of the distances between matched words. Corresponds to the `proximity` criterion in the ranking formula.

        - `userScore` (integer): Custom ranking for the object, expressed as a single numerical value. Conceptually, it's what the position of the object would be in the list of all objects sorted by custom ranking. Corresponds to the `custom` criterion in the ranking formula.

        - `geoDistance` (integer): Distance between the geo location in the search query and the best matching geo location in the record, divided by the geo precision.

        - `geoPrecision` (integer): Precision used when computed the geo distance, in meters. All distances will be floored to a multiple of this precision.

        - `nbExactWords` (integer): Number of exactly matched words. If `alternativeAsExact` is set, it may include plurals and/or synonyms.

        - `words` (integer): Number of matched words, including prefixes and typos.

        - `filters` (integer): *This field is reserved for advanced usage.* It will be zero in most cases.

        - `matchedGeoLocation` (object): Geo location that matched the query. *Note: Only returned for a geo search.*

            - `lat` (float): Latitude of the matched location.

            - `lng` (float): Longitude of the matched location.

            - `distance` (integer): Distance between the matched location and the search location (in meters). **Caution:** Contrary to `geoDistance`, this value is *not* divided by the geo precision.

    - `_distinctSeqID` (integer): *Note: Only returned when [distinct](#distinct) is non-zero.* When two consecutive results have the same value for the attribute used for "distinct", this field is used to distinguish between them.

- `nbHits` (integer): Number of hits that the search query matched.

- `page` (integer): Index of the current page (zero-based). See the [page](#page) search parameter. *Note: Not returned if you use `offset`/`length` for pagination.*

- `hitsPerPage` (integer): Maximum number of hits returned per page. See the [hitsPerPage](#hitsperpage) search parameter. *Note: Not returned if you use `offset`/`length` for pagination.*

- `nbPages` (integer): Number of pages corresponding to the number of hits. Basically, `ceil(nbHits / hitsPerPage)`. *Note: Not returned if you use `offset`/`length` for pagination.*

- `processingTimeMS` (integer): Time that the server took to process the request, in milliseconds. *Note: This does not include network time.*

- `query` (string): An echo of the query text. See the [query](#query) search parameter.

- `queryAfterRemoval` (string, optional): *Note: Only returned when [removeWordsIfNoResults](#removewordsifnoresults) is set to `lastWords` or `firstWords`.* A markup text indicating which parts of the original query have been removed in order to retrieve a non-empty result set. The removed parts are surrounded by `<em>` tags.

- `params` (string, URL-encoded): An echo of all search parameters.

- `message` (string, optional): Used to return warnings about the query.

- `aroundLatLng` (string, optional): *Note: Only returned when [aroundLatLngViaIP](#aroundlatlngviaip) is set.* The computed geo location. **Warning: for legacy reasons, this parameter is a string and not an object.** Format: `${lat},${lng}`, where the latitude and longitude are expressed as decimal floating point numbers.

- `automaticRadius` (integer, optional): *Note: Only returned for geo queries without an explicitly specified radius (see `aroundRadius`).* The automatically computed radius. **Warning: for legacy reasons, this parameter is a string and not an integer.**

When [getRankingInfo](#getrankinginfo) is set to `true`, the following additional fields are returned:

- `serverUsed` (string): Actual host name of the server that processed the request. (Our DNS supports automatic failover and load balancing, so this may differ from the host name used in the request.)

- `parsedQuery` (string): The query string that will be searched, after normalization. Normalization includes removing stop words (if [removeStopWords](#removestopwords) is enabled), and transforming portions of the query string into phrase queries (see [advancedSyntax](#advancedsyntax)).

- `timeoutCounts` (boolean) - DEPRECATED: Please use `exhaustiveFacetsCount` in remplacement.

- `timeoutHits` (boolean) - DEPRECATED: Please use `exhaustiveFacetsCount` in remplacement.

... and ranking information is also added to each of the hits (see above).

When [facets](#facets) is non-empty, the following additional fields are returned:

- `facets` (object): Maps each facet name to the corresponding facet counts:

    - `${facet_name}` (object): Facet counts for the corresponding facet name:

        - `${facet_value}` (integer): Count for this facet value.

- `facets_stats` (object, optional): *Note: Only returned when at least one of the returned facets contains numerical values.* Statistics for numerical facets:

    - `${facet_name}` (object): The statistics for a given facet:

        - `min` (integer | float): The minimum value in the result set.

        - `max` (integer | float): The maximum value in the result set.

        - `avg` (integer | float): The average facet value in the result set.

        - `sum` (integer | float): The sum of all values in the result set.

- `exhaustiveFacetsCount` (boolean): Whether the counts are exhaustive (`true`) or approximate (`false`). *Note: In some conditions when [distinct](#distinct) is greater than 1 and an empty query without refinement is sent, the facet counts may not always be exhaustive.*

## Search Parameters

You can see the full list of search parameters here:
[https://www.algolia.com/doc/api-client/go/parameters/](https://www.algolia.com/doc/api-client/go/parameters/)

## Search multiple indices - `MultipleQueries` 

You can send multiple queries with a single API call using a batch of queries:

```go
// Perform 3 queries in a single API call:
//  - 1st query targets the `categories` index
//  - 2nd and 3rd queries target the `products` index
queries := []algoliasearch.IndexedQuery{
  {
    IndexName: "categories",
    Params:    algoliasearch.Map{"query": "computer", "hitsPerPage": 3},
  },
  {
    IndexName: "products",
    Params:    algoliasearch.Map{"query": "computer", "hitsPerPage": 3, "filters": "_tags:promotion"},
  },
  {
    IndexName: "products",
    Params:    algoliasearch.Map{"query": "computer", "hitsPerPage": 10},
  },
}

res, err := client.MultipleQueries(queries, "")
```

You can specify a `strategy` parameter to optimize your multiple queries:

- `none`: Execute the sequence of queries until the end.
- `stopIfEnoughMatches`: Execute the sequence of queries until the number of hits is reached by the sum of hits.

### Response

The resulting JSON contains the following fields:

- `results` (array): The results for each request, in the order they were submitted. The contents are the same as in [Search an index](#search-an-index).
    Each result also includes the following additional fields:

    - `index` (string): The name of the targeted index.
    - `processed` (boolean, optional): *Note: Only returned when `strategy` is `stopIfEnoughmatches`.* Whether the query was processed.

## Get Objects - `GetObjects` 

You can easily retrieve an object using its `objectID` and optionally specify a comma separated list of attributes you want:

```go
// Retrieves the object with all its attributes
object, err := index.GetObject("myID", nil)

// Retrieves the object with only its `firstname` attribute
object, err = index.GetObject("myID", []string{"firstname"})

// Retrieves the object with only its `firstname` and `lastname` attributes
object, err = index.GetObject("myID", []string{"firstname", "lastname"})
```

You can also retrieve a set of objects:

```go
objects, err = index.GetObjects([]string{"myID1", "myID2"})
```

## Search for facet values - `SearchForFacetValues` 

When there are many facet values for a given facet, it may be useful to search within them. For example, you may have dozens of 'brands' for a given index of 'clothes'. Rather than displaying all of the brands, it is often best to only display the most popular and add a search input to allow the user to search for the brand that they are looking for.

Searching on facet values is different than a regular search because you are searching only on *facet values*, not *objects*.

The results are sorted by decreasing count. By default, maximum 10 results are returned. This can be adjusted via [maxFacetHits](#maxfacethits). No pagination is possible.

The facet search can optionally take regular search query parameters.
In that case, it will return only facet values that both:

1. match the facet query
2. are contained in objects matching the regular search query.

**Warning:** For a facet to be searchable, it must have been declared with the `searchable()` modifier in the [attributesForFaceting](#attributesforfaceting) index setting.

#### Example

go


# Indexing



## Add Objects - `AddObjects` 

Each entry in an index has a unique identifier called `objectID`. There are two ways to add an entry to the index:

 1. Supplying your own `objectID`.
 2. Using automatic `objectID` assignment. You will be able to access it in the response.

Using your own unique IDs when creating records is a good way to make future updates easier without having to keep track of Algolia's generated IDs.
The value you provide for objectIDs can be an integer or a string.

You don't need to explicitly create an index, it will be automatically created the first time you add an object.
Objects are schema less so you don't need any configuration to start indexing.
If you wish to configure things, the settings section provides details about advanced settings.

Example with automatic `objectID` assignments:

```go
objects := []algoliasearch.Object{
  {"firstname": "Jimmie", "lastname": "Barninger"},
  {"firstname": "Ray", "lastname": "Charles"},
}

res, err := index.AddObjects(objects)
```

Example with manual `objectID` assignments:

```go
objects := []algoliasearch.Object{
  {"objectID": "1", "firstname": "Jimmie", "lastname": "Barninger"},
  {"objectID": "2", "firstname": "Ray", "lastname": "Charles"},
}

res, err := index.AddObjects(objects)
```

To add a single object, use the following method:

```go
object := Object{
  "objectID":  "myID",
  "firstname": "Jimmie",
  "lastname":  "Barninger",
}
res, err := index.AddObject(object)
```

## Update objects - `UpdateObjects` 

You have three options when updating an existing object:

 1. Replace all its attributes.
 2. Replace only some attributes.
 3. Apply an operation to some attributes.

Example on how to replace all attributes existing objects:

```go
objects := []algoliasearch.Object{
  {"objectID": "myID1", "firstname": "Jimmie", "lastname": "Barninger"},
  {"objectID": "myID2", "firstname": "Ray", "lastname": "Charles"},
}

res, err := index.UpdateObjects(objects)
```

To update a single object, you can use the following method:

```go
object := algoliasearch.Object{
  "objectID":  "myID",
  "firstname": "Jimmie",
  "lastname":  "Barninger",
  "city":      "New York",
}
res, err := index.UpdateObject(object)
```

## Partial update objects - `PartialUpdateObjects` 

You have many ways to update an object's attributes:

 1. Set the attribute value
 2. Add a string or number element to an array
 3. Remove an element from an array
 4. Add a string or number element to an array if it doesn't exist
 5. Increment an attribute
 6. Decrement an attribute

Example to update only the city attribute of an existing object:

```go
object := algoliasearch.Object{
  "objectID": "myID"
  "city":     "San Francisco"
}
res, err := index.PartialUpdateObject(object)
```

Example to add a tag:

```go
object := algoliasearch.Object{
  "objectID": "myID",
  "_tags":    algoliasearch.Map{
    "_operation": "Add",
    "value":      "MyTags",
  }
}

res, err := index.PartialUpdateObject(object)
```

Example to remove a tag:

```go
object := algoliasearch.Object{
  "objectID": "myID",
  "_tags":    algoliasearch.Map{
    "_operation": "Remove",
    "value":      "MyTags",
  }
}

res, err := index.PartialUpdateObject(object)
```

Example to add a tag if it doesn't exist:

```go
object := algoliasearch.Object{
  "objectID": "myID",
  "_tags":    algoliasearch.Map{
    "_operation": "AddUnique",
    "value":      "MyTags",
  }
}

res, err := index.PartialUpdateObject(object)
```

Example to increment a numeric value:

```go
object := algoliasearch.Object{
  "objectID": "myID",
  "price":    algoliasearch.Map{
    "_operation": "Increment",
    "value":      42,
  }
}

res, err := index.PartialUpdateObject(object)
```

Note: Here we are incrementing the value by `42`. To increment just by one, put
`value:1`.

Example to decrement a numeric value:

```go
object := algoliasearch.Object{
  "objectID": "myID",
  "price":    algoliasearch.Map{
    "_operation": "Decrement",
    "value":      42,
  }
}

res, err := index.PartialUpdateObject(object)
```

Note: Here we are decrementing the value by `42`. To decrement just by one, put
`value:1`.

To partial update multiple objects using one API call, you can use the following method:

```go
objects := []algoliasearch.Object{
  {"objectID": "myID1", "lastname": "Barninger"},
  {"objectID": "myID2", "firstname": "Ray"},
}

res, err := index.PartialUpdateObjects(objects)
```

## Delete objects - `DeleteObjects` 

You can delete objects using their `objectID`:

```go
objectIDs := []string{"myID1", "myID2"}

res, err := index.DeleteObjects(objectIDs)
```

To delete a single object, you can use the following method:

```go
res, err := index.DeleteObject("myID")
```

## Delete by query - `DeleteByQuery` 

You can delete all objects matching a single query with the following code. Internally, the API client performs the query, deletes all matching hits, and waits until the deletions have been applied.

Take your precautions when using this method. Calling it with an empty query will result in cleaning the index of all its records.

```go
params := algoliasearch.Map{
  // Set your query parameters here
}

err := index.DeleteByQuery("john", params)
```

## Wait for operations - `WaitTask` 

All write operations in Algolia are asynchronous by design.

It means that when you add or update an object to your index, our servers will
reply to your request with a `taskID` as soon as they understood the write
operation.

The actual insert and indexing will be done after replying to your code.

You can wait for a task to complete using the `waitTask` method on the `taskID` returned by a write operation.

For example, to wait for indexing of a new object:

```go
object := algoliasearch.Object{
  "firstname": "Jimmie",
  "lastname":  "Barninger",
}

res, err := index.AddObject(object)
err = index.WaitTask(res.TaskID)
```

If you want to ensure multiple objects have been indexed, you only need to check
the biggest `taskID`.


# Settings



## Get settings - `GetSettings` 

You can retrieve settings:

```go
// Retrieves all the settings of the index
settings, err := index.GetSettings()
```

## Set settings - `SetSettings` 

```go
// Updates the settings
settings := algoliasearch.Map{
  "customRanking": []string{"desc(followers)"},
}
res, err := index.SetSettings(settings)
```

You can find the list of parameters you can set in the [Settings Parameters](#index-settings-parameters) section

**Warning**

Performance wise, it's better to do a `SetSettings` before pushing the data

### Replica settings

You can forward all settings updates to the replicas of an index by using the `forwardToReplicas` option:

```go
settings := algoliasearch.Map{
  "customRanking":   []string{"asc(followers)"},
  "forwardToReplicas": true,
}
res, err := index.SetSettings(settings)
```

## Index settings parameters

You can see the full list of settings parameters here:
[https://www.algolia.com/doc/api-client/go/parameters/](https://www.algolia.com/doc/api-client/go/parameters/)


# Manage Indices



## Create an index

To create an index, you need to perform any indexing operation like:
- set settings
- add object

## List indices - `ListIndexes` 

You can list all your indices along with their associated information (number of entries, disk size, etc.) with the `ListIndexes` method:

```go
indexes, err := client.ListIndexes()
```

## Delete an index - `DeleteIndex` 

You can delete an index using its name:

```go
res, err := index.Delete()
```

## Clear an index - `ClearIndex` 

You can delete the index contents without removing settings and index specific API keys by using the `clearIndex` command:

```go
res, err := index.Clear()
```

## Copy index - `CopyIndex` 

You can copy an existing index using the `copy` command.

**Warning**: The copy command will overwrite the destination index.

```go
index := client.InitIndex("MyNewIndex")

// Copy `MyNewIndex` to `MyIndex`
res, err := index.Copy("MyIndex")
```

## Move index - `MoveIndex` 

In some cases, you may want to totally reindex all your data. In order to keep your existing service
running while re-importing your data we recommend the usage of a temporary index plus an atomical
move using the `MoveIndex` method.

```go
// Rename MyNewIndex in MyIndex (and overwrite it)
res, err := client.MoveIndex("MyNewIndex", "MyIndex")
```

**Note:** The MoveIndex method overrides the destination index, and deletes the temporary one.
  In other words, there is no need to call the `ClearIndex` or `DeleteIndex` methods to clean the temporary index.
It also overrides all the settings of the destination index (except the [replicas](#replicas) parameter that need to not be part of the temporary index settings).

**Recommended steps**
If you want to fully update your index `MyIndex` every night, we recommend the following process:

 1. Get settings and synonyms from the old index using [Get settings](#get-settings)
  and [Get synonym](#get-synonym).
 1. Apply settings and synonyms to the temporary index `MyTmpIndex`, (this will create the `MyTmpIndex` index)
  using [Set settings](#set-settings) and [Batch synonyms](#batch-synonyms) ([!] Make sure to remove the [replicas](#replicas) parameter from the settings if it exists.
 1. Import your records into a new index using [Add Objects](#add-objects)).
 1. Atomically replace the index `MyIndex` with the content and settings of the index `MyTmpIndex`
 using the [Move index](#move-index) method.
 This will automatically override the old index without any downtime on the search.
 
 You'll end up with only one index called `MyIndex`, that contains the records and settings pushed to `MyTmpIndex`
 and the replica-indices that were initially attached to `MyIndex` will be in sync with the new data.


# Api keys



## Overview

When creating your Algolia Account, you'll notice there are 3 different API Keys:

- **Admin API Key** - it provides full control of all your indices.
*The admin API key should always be kept secure;
do NOT give it to anybody; do NOT use it from outside your back-end as it will
allow the person who has it to query/change/delete data*

- **Search-Only API Key** - It allows you to search on every indices.

- **Monitoring API Key** - It allows you to access the [Monitoring API](https://www.algolia.com/doc/rest-api/monitoring)

### Other types of API keys

The *Admin API Key* and *Search-Only API Key* both have really large scope and sometimes you want to give a key to
someone that have restricted permissions, can it be an index, a rate limit, a validity limit, ...

To address those use-cases we have two different type of keys:

- **Secured API Keys**

When you need to restrict the scope of the *Search Key*, we recommend to use *Secured API Key*.
You can generate them on the fly (without any call to the API)
from the *Search Only API Key* or any search *User Key* using the [Generate key](#generate-key) method

- **User API Keys**

If *Secured API Keys* does not meet your requirements, you can make use of *User keys*.
Managing and especially creating those keys requires a call to the API.

We have several methods to manage them:

- [Add user key](#add-user-key)
- [Update user key](#update-user-key)
- [Delete user key](#delete-user-key)
- [List api keys](#list-api-keys)
- [Get key permissions](#get-key-permissions)

## Generate key - `GenerateSecuredApiKey` 

When you need to restrict the scope of the *Search Key*, we recommend to use *Secured API Key*.
You can generate a *Secured API Key* from the *Search Only API Key* or any search *User API Key*

There is a few things to know about *Secured API Keys*
- They always need to be generated **on your backend** using one of our API Client
- You can generate them on the fly (without any call to the API)
- They will not appear on the dashboard as they are generated without any call to the API
- The key you use to generate it **needs to become private** and you should not use it in your frontend.
- The generated secured API key **will inherit any restriction from the search key it has been generated from**

You can then use the key in your frontend code

```js
var client = algoliasearch('YourApplicationID', 'YourPublicAPIKey');

var index = client.initIndex('indexName')

index.search('something', function(err, content) {
  if (err) {
    console.error(err);
    return;
  }

  console.log(content);
});
```

#### Filters

Every filter set in the API key will always be applied. On top of that [filters](#filters) can be applied
in the query parameters.

```go
params := algoliasearch.Map{
  "filters": "_tags:user_42",
}

key, err := client.GenerateSecuredAPIKey("YourSearchOnlyApiKey", params)
```

**Warning**:

If you set filters in the key `groups:admin`, and `groups:press OR groups:visitors` in the query parameters,
this will be equivalent to `groups:admin AND (groups:press OR groups:visitors)`

##### Having one API Key per User

One of the usage of secured API keys, is to have allow users to see only part of an index, when this index
contains the data of all users.
In that case, you can tag all records with their associated `user_id` in order to add a `user_id=42` filter when
generating the *Secured API Key* to retrieve only what a user is tagged in.

**Warning**

If you're generating *Secured API Keys* using the [JavaScript client](http://github.com/algolia/algoliasearch-client-javascript) in your frontend,
it will result in a security breach since the user is able to modify the filters you've set
by modifying the code from the browser.

#### Valid Until

You can set a Unix timestamp used to define the expiration date of the API key

```go
// Generate a public API key that is valid for 1 hour

valid_until := time.Now().Unix() + 3600

params := algoliasearch.Map{
  "validUntil": valid_until,
}

key, err := client.GenerateSecuredAPIKey("YourSearchOnlyApiKey", params)
```

#### Index Restriction

You can restrict the key to a list of index names allowed for the secured API key

```go
// Generate a public API key that is restricted to `index1` and `index2`

params := algoliasearch.Map{
  "restrictIndices": "index1,index2",
}

key, err := client.GenerateSecuredAPIKey("YourSearchOnlyApiKey", params)
```

#### Rate Limiting

If you want to rate limit a secured API Key, the API key you generate the secured api key from need to be rate-limited.
You can do that either via the dashboard or via the API using the
[Add user key](#add-user-key) or [Update user key](#update-user-key) method

##### User Rate Limiting

By default the rate limits will only use the `IP`.

This can be an issue when several of your end users are using the same IP.
To avoid that, you can set a `userToken` query parameter when generating the key.

When set, a unique user will be identified by his `IP + user_token` instead of only by his `IP`.

This allows you to restrict a single user to performing a maximum of `N` API calls per hour,
even if he shares his `IP` with another user.

```go
params := algoliasearch.Map{
  "filters":   "_tags:user_42",
  "userToken": "user_42",
}

key, err := client.GenerateSecuredAPIKey("YourSearchOnlyApiKey", params)
```

#### Network restriction

For more protection against API key leaking and reuse you can restrict the key to be valid only from specific IPv4 networks

```go
// Generate a public API key that is restricted to `192.168.1.0/24`

params := algoliasearch.Map{
  "restrictSources": "192.168.1.0/24",
}

key, err := client.GenerateSecuredAPIKey("YourSearchOnlyApiKey", params)
```


# Synonyms



## Save synonym - `AddSynonym` 

This method saves a single synonym record into the index.

In this example, we specify true to forward the creation to replica indices.
By default the behavior is to save only on the specified index.

```go
uniqueID := "a-unique-identifier"
synonym := algoliasearch.NewSynonym(uniqueID, []string{
    "car", "vehicle", "auto",
})
res, err := index.AddSynonym(uniqueID, synonym, true)
```

## Batch synonyms - `BatchSynonyms` 

Use the batch method to create a large number of synonyms at once,
forward them to replica indices if desired,
and optionally replace all existing synonyms
on the index with the content of the batch using the replaceExistingSynonyms parameter.

You should always use replaceExistingSynonyms to atomically replace all synonyms
on a production index. This is the only way to ensure the index always
has a full list of synonyms to use during the indexing of the new list.

```go
// Batch synonyms, with replica forwarding and atomic replacement of existing synonyms
res, err := index.BatchSynonyms([]algoliasearch.Synonym{
    algoliasearch.NewSynonym("a-unique-identifier", []string{"car", "vehicle", "auto"}),
    algoliasearch.NewSynonym("another-unique-identifier", []string{"street", "st"}),
}, true, true)
```

## Editing Synonyms

Updating the value of a specific synonym record is the same as creating one.
Make sure you specify the same objectID used to create the record and the synonyms
will be updated.
When updating multiple synonyms in a batch call (but not all synonyms),
make sure you set replaceExistingSynonyms to false (or leave it out,
false is the default value).
Otherwise, the entire synonym list will be replaced only partially with the records
in the batch update.

## Delete synonym - `DeleteSynonym` 

Use the normal index delete method to delete synonyms,
specifying the objectID of the synonym record you want to delete.
Forward the deletion to replica indices by setting the forwardToReplicas parameter to true.

```go
// Delete and forward to replicas
res, err := index.DeleteSynonym("a-unique-identifier", true)
```

## Clear all synonyms - `ClearSynonyms` 

This is a convenience method to delete all synonyms at once.
It should not be used on a production index to then push a new list of synonyms:
there would be a short period of time during which the index would have no synonyms
at all.

To atomically replace all synonyms of an index,
use the batch method with the replaceExistingSynonyms parameter set to true.

```go
// Clear synonyms and forward to replicas
res, err := index.ClearSynonyms(true)
```

## Get synonym - `GetSynonym` 

Search for synonym records by their objectID or by the text they contain.
Both methods are covered here.

```go
synonym, err := index.GetSynonym("a-unique-identifier")
```

## Search synonyms - `SearchSynonyms` 

Search for synonym records similar to how you’d search normally.

Accepted search parameters:
- query: the actual search query to find synonyms. Use an empty query to browse all the synonyms of an index.
- type: restrict the search to a specific type of synonym. Use an empty string to search all types (default behavior). Multiple types can be specified using a comma-separated list or an array.
- page: the page to fetch when browsing through several pages of results. This value is zero-based.
hitsPerPage: the number of synonyms to return for each call. The default value is 100.

```go
// Searching for "street" in synonyms and one-way synonyms; fetch the second page with 10 hits per page
synonyms, err := index.SearchSynonyms("street", []string{"synonym", "oneWaySynonym"}, 1, 10)
```


# Advanced



## Custom batch - `Batch` 

You may want to perform multiple operations with one API call to reduce latency.

If you have one index per user, you may want to perform a batch operations across several indices.
We expose a method to perform this type of batch:

```go
person := algoliasearch.Map{
  "firstname": "Jimmie",
  "lastname":  "Barninger",
}

operation := algoliasearch.BatchOperation{
  Action: "addObject",
  Body:   person,
}

operations := []algoliasearch.BatchOperationIndexed{
  {IndexName: "prodIndex", BatchOperation: operation},
  {IndexName: "devIndex", BatchOperation: operation},
}

res, err := client.Batch(operations)
```

The attribute **action** can have these values:

- addObject
- updateObject
- partialUpdateObject
- partialUpdateObjectNoCreate
- deleteObject

## Backup / Export an index - `Browse` 

The `search` method cannot return more than 1,000 results. If you need to
retrieve all the content of your index (for backup, SEO purposes or for running
a script on it), you should use the `browse` method instead. This method lets
you retrieve objects beyond the 1,000 limit.

This method is optimized for speed. To make it fast, distinct, typo-tolerance,
word proximity, geo distance and number of matched words are disabled. Results
are still returned ranked by attributes and custom ranking.

#### Response Format

##### Sample

```json
{
  "hits": [
    {
      "firstname": "Jimmie",
      "lastname": "Barninger",
      "objectID": "433"
    }
  ],
  "processingTimeMS": 7,
  "query": "",
  "params": "filters=level%3D20",
  "cursor": "ARJmaWx0ZXJzPWxldmVsJTNEMjABARoGODA4OTIzvwgAgICAgICAgICAAQ=="
}
```

##### Fields

- `cursor` (string, optional): A cursor to retrieve the next chunk of data. If absent, it means that the end of the index has been reached.
- `query` (string): Query text used to filter the results.
- `params` (string, URL-encoded): Search parameters used to filter the results.
- `processingTimeMS` (integer): Time that the server took to process the request, in milliseconds. *Note: This does not include network time.*

The following fields are provided for convenience purposes, and **only when the browse is not filtered**:

- `nbHits` (integer): Number of objects in the index.
- `page` (integer): Index of the current page (zero-based).
- `hitsPerPage` (integer): Maximum number of hits returned per page.
- `nbPages` (integer): Number of pages corresponding to the number of hits. Basically, `ceil(nbHits / hitsPerPage)`.

#### Example

```go
// Iterate with a filter over the index
params := algoliasearch.Map{
  "filters": "i<42",
  "query":   "text",
}
it, err := index.BrowseAll(params)

var hit Map
for {
  if hit, err = it.Next(); err != nil {
  if err.Error() == "No more hits" {
    // End of results
  } else {
    // Error while browsing
  }
  break
  }
}

// Retrieve the next cursor from the `Browse` method
res, err := index.Browse(params, "")
fmt.Println(res.Cursor)
```

## List api keys - `ListKeys` 

To list existing keys, you can use:

```go
// Lists global API Keys
keys, err := client.ListKeys()

// Lists API Keys that can access only to this index
keys, err = index.ListKeys()
```

Each key is defined by a set of permissions that specify the authorized actions. The different permissions are:

* **search**: Allowed to search.
* **browse**: Allowed to retrieve all index contents via the browse API.
* **addObject**: Allowed to add/update an object in the index.
* **deleteObject**: Allowed to delete an existing object.
* **deleteIndex**: Allowed to delete index content.
* **settings**: allows to get index settings.
* **editSettings**: Allowed to change index settings.
* **analytics**: Allowed to retrieve analytics through the analytics API.
* **listIndexes**: Allowed to list all accessible indexes.

## Add user key - `AddUserKey` 

To create API keys:

```go
// Creates a new index specific API key valid for 300 seconds, with a rate
// limit of 100 calls per hour per IP and a maximum of 20 hits
acl := []string{"search"}
params := algoliasearch.Map{
  "validity":               300,
  "maxQueriesPerIPPerHour": 100,
  "maxHitsPerQuery":        20,
}

res, err := index.AddKey(acl, params)
fmt.Println(res.Key)
```

You can also create an API Key with advanced settings:

##### validity

Add a validity period. The key will be valid for a specific period of time (in seconds).

##### maxQueriesPerIPPerHour

Specify the maximum number of API calls allowed from an IP address per hour. Each time an API call is performed with this key, a check is performed. If the IP at the source of the call did more than this number of calls in the last hour, a 403 code is returned. Defaults to 0 (no rate limit). This parameter can be used to protect you from attempts at retrieving your entire index contents by massively querying the index.

  

Note: If you are sending the query through your servers, you must use the `Algolia.with_rate_limits("EndUserIP", "APIKeyWithRateLimit") do ... end` block to enable rate-limit.

##### maxHitsPerQuery

Specify the maximum number of hits this API key can retrieve in one call. Defaults to 0 (unlimited). This parameter can be used to protect you from attempts at retrieving your entire index contents by massively querying the index.

##### indexes

Specify the list of targeted indices. You can target all indices starting with a prefix or ending with a suffix using the '\*' character. For example, "dev\_\*" matches all indices starting with "dev\_" and "\*\_dev" matches all indices ending with "\_dev". Defaults to all indices if empty or blank.

##### referers

Specify the list of referers. You can target all referers starting with a prefix, ending with a suffix using the '\*' character. For example, "https://algolia.com/\*" matches all referers starting with "https://algolia.com/" and "\*.algolia.com" matches all referers ending with ".algolia.com". If you want to allow the domain algolia.com you can use "\*algolia.com/\*". Defaults to all referers if empty or blank.

##### queryParameters

Specify the list of query parameters. You can force the query parameters for a query using the url string format (param1=X&param2=Y...).

##### description

Specify a description to describe where the key is used.

```go
// Creates a new index specific API key valid for 300 seconds, with a rate
// limit of 100 calls per hour per IP and a maximum of 20 hits
acl := []string{"search"}
params := algoliasearch.Map{
  "indexes":                "myIndex",
  "validity":               300,
  "maxQueriesPerIPPerHour": 100,
  "maxHitsPerQuery":        20,
}

res, err := client.AddKey(acl, params)
fmt.Println(res.Key)
```

## Update user key - `UpdateUserKey` 

To update the permissions of an existing key:

```go
// Update an existing index specific API key valid for 300 seconds, with a rate
// limit of 100 calls per hour per IP and a maximum of 20 hits
params := algoliasearch.Map{
  "acl":                    []string{"search"},
  "indexes":                "myIndex",
  "validity":               300,
  "maxQueriesPerIPPerHour": 100,
  "maxHitsPerQuery":        20,
}

res, err := client.UpdateKey("f420238212c54dcfad07ea0aa6d5c45f", params)
fmt.Println(res.Key)
```

To get the permissions of a given key:

```go
// Gets the rights of a global key
key, err := client.GetKey("f420238212c54dcfad07ea0aa6d5c45f")

// Gets the rights of an index specific key
key, err = index.GetKey("71671c38001bf3ac857bc82052485107")
```

## Delete user key - `DeleteUserKey` 

To delete an existing key:

```go
// Deletes a global key
res, err := client.DeleteKey("f420238212c54dcfad07ea0aa6d5c45f")

// Deletes an index specific key
res, err := index.DeleteKey("71671c38001bf3ac857bc82052485107")
```

## Get key permissions - `GetUserKey` 

To get the permissions of a given key:

```go
// Gets the rights of a global key
key, err := client.GetKey("f420238212c54dcfad07ea0aa6d5c45f")

// Gets the rights of an index specific key
key, err = index.GetKey("71671c38001bf3ac857bc82052485107")
```

## Get latest logs - `GetLogs` 

You can retrieve the latest logs via this API. Each log entry contains:

* Timestamp in ISO-8601 format
* Client IP
* Request Headers (API Key is obfuscated)
* Request URL
* Request method
* Request body
* Answer HTTP code
* Answer body
* SHA1 ID of entry

You can retrieve the logs of your last 1,000 API calls and browse them using the offset/length parameters:

#### offset

Specify the first entry to retrieve (0-based, 0 is the most recent log entry). Defaults to 0.

#### length

Specify the maximum number of entries to retrieve starting at the offset. Defaults to 10. Maximum allowed value: 1,000.

#### onlyErrors

Retrieve only logs with an HTTP code different than 200 or 201. (deprecated)

#### type

Specify the type of logs to retrieve:

* `query`: Retrieve only the queries.
* `build`: Retrieve only the build operations.
* `error`: Retrieve only the errors (same as `onlyErrors` parameters).

```go
// Gets last 10 log entries
params := algoliasearch.Map{
  "offset": 0,
  "length": 10,
  "type":   "all",
}
res, err := client.GetLogs(params)

// Gets last 100 log entries
params["length"] = 100
res, err = client.GetLogs(params)

// Gets last 100 log errors
params["type"] = "error"
res, err = client.GetLogs(params)
```

## REST API

We've developed API clients for the most common programming languages and platforms.
These clients are advanced wrappers on top of our REST API itself and have been made
in order to help you integrating the service within your apps:
for both indexing and search.

Everything that can be done using the REST API can be done using those clients.

The REST API lets your interact directly with Algolia platforms from anything that can send an HTTP request
[Go to the REST API doc](https://algolia.com/doc/rest)


