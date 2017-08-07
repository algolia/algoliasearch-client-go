# Changelog

## 2.17.0 (2017-08-07)

### Additions

 - Implement `*WithRequestOptions` methods for both Client and Index implementations

## 2.16.0 (2017-08-02)

### Fixes

 - Correctly stop the browse iteration of `Index.DeleteByQuery`

### Changes

 - Improve performances of `Index.DeleteByQuery`
 - Improve documentation comment of `Index.BrowseAll`
 - `Index.BrowseAll`'s terminal "no more hits" error now has its own variable: `NoMoreHitsErr`

## 2.15.0 (2017-06-30)

### Additions

 - Implement new Index methods for Algolia Query Rules endpoints

## 2.14.0 (2017-06-06)

### Fixes

 - Fix `Client.GetLogs` by correctly passing the given parameters via the URL
 - Accept `indexName` as `Client.GetLogs` parameter as string

## 2.13.0 (2017-04-26)

### Additions

 - Accept percentileComputation query parameter as bool

## 2.12.0 (2017-04-12)

### Additions

 - Add default ProxyFunc to the default transport layer of the HTTP client

### Fixes

 - Improve retry strategy to let it use the provided hosts if any in priority
 - Re-arrange the order of hosts of the retry strategy

## 2.11.0 (2017-03-24)

### Additions

- Add PartialUpdateOp type and helpers to simplify partial update of records

## 2.10.0 (2017-03-24)

### Fixes

- Accept disableExactOnAttributes query and settings parameter as []string
- Accept disablePrefixOnAttributes settings parameter as []string
- Accept paginationLimitedTo settings parameter as int
- Accept length query parameter as int
- Accept offset query parameter as int
- Accept restrictHighlightAndSnippetArrays query and settings parameter as bool
- Accept typoTolerance query parameter as string or bool
- Accept alternativesAsExact settings parameter as []string
- Accept exactOnSingleWordQuery settings parameter as string
- Accept optionalWords query and settings parameter as string or []string
- Accept removeWordsIfNoResults settings parameter as string
- Accupt insidePolygon query parameter as string or [][]float64
- Accept insideBoundingBox query parameter as string or [][]float64
- Accept facetFilters query parameter as string or []string
- Accept facets query parameter as string or []string
- Accept restrictSearchableAttributes query parameter as string or []string
- Accept analyticsTags query parameter as string or []string
- Accept tagFilters query parameter as string or []interface{}
- Accept numericFilters query parameter as string or []interface{}
- Accept getRankingInfo query parameter as int or bool

### Misc

- Run the tests with Go 1.8
- Do not run the tests with Go 1.5 anymore

## 2.9.1 (2017-03-22)

### Fixes

- Fix the `attributesToRetrieve` handling of `GetObject/GetObjects` methods

## 2.9.0 (2017-02-24)

### Additions

- (Get|Add|Update|Delete)APIKey methods

### Changes

- Deprecate (Get|Add|Update|Delete)UserKey methods

## 2.8.0 (2017-02-14)

### Additions

- Allow `maxFacetHits` parameter in queries and parameters

## 2.7.1 (2017-01-25)

### Fixes

- Ensure `net/http.Response.Body.Close()` is called

## 2.7.0 (2016-12-28)

### Changes

- Only accepts the following parameters for `GenerateSecuredAPIKey`:
  + `userToken` (string identifier generally used to rate-limit users per IP)
  + `validUntil` (timestamp of the expiration date)
  + `restrictIndices` (comma-separated string list of the indices to restrict)
  + `referers` (string slice of allowed referers)
  + `restrictSources` (string of the allowed IPv4 network)

## 2.6.0 (2016-12-03)

### Fixes

- Improve the transport layer code
- Keep the last active connection between requests to better handle DNS timeouts

## 2.5.0 (2016-11-29)

### Additions

- Add `Index.SearchForFacetValues` method
  + Same as `Index.SearchFacet`
  + `Index.SearchFacet` is kept for backward-compatibility

## 2.4.0 (2016-11-02)

### Changes

- Accept both boolean and string slice for `Settings.IgnorePlurals`

## 2.3.1 (2016-11-02)

### Fixes

- Handle missing parameters in `Settings.ToMap`:
  + `responseFields`
  + `typoTolerance`

## 2.3.0 (2016-11-02)

### Additions

- Accept `responseFields` parameter in queries and settings

## 2.2.0 (2016-10-19)

### Additions

- Add `Index.SearchFacet` method

## 2.1.2 (2016-10-19)

### Fixes

- Exclude the `testing` package from the build
- Fix typo and type checking for `attributeForDistinct`

## 2.1.1 (2016-10-16)

### Additions

- Accept `facetingAfterDistinct` parameter in queries

## 2.1.0 (2016-10-03)

### Fixes

- Fix `Client.ListIndexes`
- Fix `LogRes` type
- Fix `aroundRadius` type
- Fix `distinct` type
- Fix `removeStopWords` type
- Fix `NewOneWaySynonym` method name (minor typo)
- Fix `NewAltCorrectionSynonym` method name (minor typo)

### Additions

- Add `Client.SetHTTPClient` method
- Add `Index.GetObjectsAttrs` method
- Add `Settings.ToMap` method
- Add `FileSize` field in the `IndexRes` response type
- Add `NumberOfPendingTask` field in the `IndexRes` response type

### Changes

- Implement the new Synonym API
- Rename `fowardToSlaves` parameter into `forwardToReplicas`
- Rename `attributesToIndex` setting into `searchableAttributes`
- Rename `numericAttributesToIndex` setting into `numericAttributesForFiltering`
- Allow `exactOnSingleWordQuery` parameter in queries
- Allow `alternativesAsExact` parameter in queries
- Allow `forwardToReplicas` parameter in queries

### Misc

- Avoid the inclusion of the `syscall` package to be used within Google App Engine

### Tests

- Add tests
- Run tests in parallel
- Only test across the last 3 major versions of Go (currently 1.5, 1.6 and 1.7)

## 2.0.0 (2016-06-06)

- Type every response from the API
- Add tests
- Breaking changes listed here: https://github.com/algolia/algoliasearch-client-go/wiki/Migration-guide-to-version-2.x

## 1.5.0 (2016-06-01)

- Code refactoring ahead of version 2.0.0

## 1.4.0 (2014-10-16)

- Add new secured api keys
- Fix bug with distinct in deleteByQuery

## 1.3.0 (2014-07-07)

- Add new browse methods

## 1.2.1 (2014-05-04)

- Add new methods to add/update api key
- Add batch method to target multiple indices
- Add strategy parameter for the multipleQueries
- Add new method to generate secured api key from query parameters

## 1.2.0 (2015-04-09)

- New implementation for the retry on another DNS

## 1.1.1 (2015-11-23)

- Fix the retry for a network error
- Drop Go 1.2 support
- Add timeout

## 1.1.0 (2014-11-30)

- Bump to 1.1.0.
- Switch to .net.

## 1.0.2 (2014-11-27)

- Bump to 1.0.2.
- Wait more for the ACL.
- Use sandboxes.
- Fix go tip.
- Try to fix go 1.2.
- Ability to set custom hosts.
- Fix missing safeName call.
- Add safeName function.
- Update ChangeLog.

## 1.0.1 (2014-10-22)

- Bump to 1.0.1.
- Fix retry.

## 1.0.0 (2014-10-16)

- Bump to 1.0.0.
- Add setExtraHeader.
- Documentation of new query parameter & index setting.
- Updated default typoTolerance setting & updated removedWordsIfNoResult documentation Add the documentation about the update of an APIKey.
- Added update key.
- Updated default typoTolerance setting & updated removedWordsIfNoResult documentation.
- Remove unsupported version.
- Add new version of Go.
- Improve the encoding of query parametter.
- Add documentation about removeWordsIfNoResult.
- Fixed links.
- Added aroundLatLngViaIP documentation.
- Add notes about the JS API client.
- Add tutorial links + minor enhancements.
- Added documentation of suffix/prefix index name matching in API key.
- Change the cluster.
- Added restrictSearchableAttributes Added multiQueries.
- Documentation: Added deleteByQuery and multipleQueries.
- Added notes on attributesToIndex.
- Update README.md (getObjects)
- Add DeleteByQuery and GetObjects.
- Added disableTypoToleranceOn & altCorrections index settings:
- Add typoTolerance & allowsTyposOnNumericTokens query parameters.
- Increase the sleeping time.
- Documentation: Added words ranking parameter:
- Add missing waitTask Add sleep.
- Added asc(attributeName) & desc(attributeName) documentation in index settings.
- Updated synonyms examples.
- Fix typo.
- Add a note about distinct and the empty queries.
- Added analytics,synonyms,enableSynonymsInHighlight query parameters.
- Update example of generate secured api key.
- Add multipleQueries Add test for multipleQueries Fix test (missing waitTask) Fix url encoding.
- Add test for generate the api key.
- Add generate secured api key.
- Fix the test of index acl.
- New numericFilters documentation.
- Fix Coveralls.
- Add Badges Simplify snippets.
- Change travis.yml.
- Trying to add Travis CI.
- Update README.md.
- Add checking of the length of variadic parametter Change prototype of Query Rename Query to Search Change DeleteObjects to take a array of IDs.
- Add optional parametter for GetObject.
- Add DeleteObjects.
- Change package from main to algoliasearch.
- Add batch.
- Export some functions.
- Fix get/list/delete key for an user Add test on keys.
- Add shuffling of the array of host.
- Add handling of server failure Add handling of server response.
- Fix browse Add user-agent Add content-length in the headers Add content-type in the headers.
- Improve test suite.
- Fix url encoding Fix query index Trying Fix 400.
- Create README.md.
- Impossible to fix urlencoding now fix tests.
- Add make test.
- Add partial test suite.
- Add waitTask.
- Add many functions for an index.
- Fix addKey.
- Add getLogs.
- Add addKey.
- Add sending of body.
- Add parsing of json response.
- Add example of test.
- Add global list/add/delete function.
- Add Index function without body.
- Add listIndexes.
- Add simple Makefile.
- Initial commit.

