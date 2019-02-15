# ChangeLog

# [2.27.0](https://github.com/algolia/algoliasearch-client-go/compare/2.26.1...2.27.0) (2018-12-17)

### Summary

Hello everyone,

Quick release of the new Insights Client which will let you interact with the
Insights API, responsible for handling the interactions with Algolia
Insights services. This new version also adds support for the setting and
retrieving the Personalization strategy of an Algolia application.

### Changes

- **added:** Implement Personalization-related methods ([0ab7e0b](https://github.com/algolia/algoliasearch-client-go/commit/0ab7e0b))
- **added:** Implement Insights client and related methods ([0d8fdd3](https://github.com/algolia/algoliasearch-client-go/commit/0d8fdd3))
- **changed:** Handle calls to https://insights.algolia.io in transport layer ([0ef2010](https://github.com/algolia/algoliasearch-client-go/commit/0ef2010))
- **removed:** Delete useless Secured API Key tests (better covered in CTS in the v3) ([713245f](https://github.com/algolia/algoliasearch-client-go/commit/713245f))
- **chore:** Remove useless Glide configuration ([a8eb049](https://github.com/algolia/algoliasearch-client-go/commit/a8eb049))

# [2.26.1](https://github.com/algolia/algoliasearch-client-go/compare/2.26.0...2.26.1) (2018-12-07)

### Summary

Quick fix release removing a wrongly inserted deprecation note on
`Client.MoveIndex{WithRequestOptions}`.

### Changes

- **fix:** Remove wrong deprecation of `Client.MoveIndex{WithRequestOptions}` ([c5465a5](https://github.com/algolia/algoliasearch-client-go/commit/c5465a5))

# [2.26.0](https://github.com/algolia/algoliasearch-client-go/compare/2.25.0...2.26.0) (2018-12-07)

### Summary

Hello everyone,

Big release this december fixing a lot of issues but mainly adding important
new features:

 * Support of new features of Query Rules
 * Add Multi-Cluster Management (MCM) capabilities
 * Add CopySettings, CopyRules, CopySynonyms functions
 * Add ReplaceAllObjects, ReplaceAllRules, ReplaceAllSynonyms functions
 * Add support for cross-application index copy via `Account.CopyIndex`

### Changes

- **fix:** change AddAPIKey type from read to write ([8465764](https://github.com/algolia/algoliasearch-client-go/commit/8465764))
- **added:** Implement AccountClient.CopyIndex ([ebe51f5](https://github.com/algolia/algoliasearch-client-go/commit/ebe51f5))
- **changed:** remove extra empty line ([ddc2dbd](https://github.com/algolia/algoliasearch-client-go/commit/ddc2dbd))
- **changed:** Use the environment variables from our Common Test Suite for testing ([21bf674](https://github.com/algolia/algoliasearch-client-go/commit/21bf674))
- **added:** Implement Index.GetAppID and Client.GetAppID methods ([04db81f](https://github.com/algolia/algoliasearch-client-go/commit/04db81f))
- **added:** Implement Index.ReplaceAll{Objects,Rules,Synonyms} ([1c2e61a](https://github.com/algolia/algoliasearch-client-go/commit/1c2e61a))
- **deprecated:** Index.{Delete,Clear} replaced by Client.{Delete,Clear}Index ([799eff0](https://github.com/algolia/algoliasearch-client-go/commit/799eff0))
- **added:** Implement Client.Copy{Settings,Synonyms,Rules} ([70ffd7d](https://github.com/algolia/algoliasearch-client-go/commit/70ffd7d))
- **deprecated:** Index.{Copy,Move} replaced by Client.{Copy,Move}Index ([e1d3d27](https://github.com/algolia/algoliasearch-client-go/commit/e1d3d27))
- **changed:** Deprecate Index.Move method in favor of Index.MoveTo ([0c7497a](https://github.com/algolia/algoliasearch-client-go/commit/0c7497a))
- **deprecated:** Client.MoveIndex in favor of Index.Move ([31d9996](https://github.com/algolia/algoliasearch-client-go/commit/31d9996))
- **added:** Add tests for MCM ([52eecbe](https://github.com/algolia/algoliasearch-client-go/commit/52eecbe))
- **added:** Introduce Multi Cluster Management (MCM) ([53ff241](https://github.com/algolia/algoliasearch-client-go/commit/53ff241))
- **fix:** objectID is required to deleteObject ([efff124](https://github.com/algolia/algoliasearch-client-go/commit/efff124))
- **fix:** Add missing types for `facetFilters` ([f71d328](https://github.com/algolia/algoliasearch-client-go/commit/f71d328))
- **test:** Add tests for Query Rules v2 ([b76d198](https://github.com/algolia/algoliasearch-client-go/commit/b76d198))
- **feat:** Implement Query Rules v2 ([9bf1d31](https://github.com/algolia/algoliasearch-client-go/commit/9bf1d31))
- **chore:** Add mitchellh/mapstructure for tests ([47339c0](https://github.com/algolia/algoliasearch-client-go/commit/47339c0))
- **chore(md):** Update contribution-related files ([83eef94](https://github.com/algolia/algoliasearch-client-go/commit/83eef94))

# [2.25.0](https://github.com/algolia/algoliasearch-client-go/compare/2.24.0...2.25.0) (2018-09-06)

### Summary

Hello everyone,

Following up after some vacations, here is the first release of September. It
only includes minor additions: new search and settings parameters are now
exposed. The list of changes speaks for itself this time. Feel free to consult
the details of each parameter in [the official Algolia documentation](https://www.algolia.com/doc/api-reference/api-parameters/).

### Changes

- **chore:** Fix release.sh srcipt to properly update client version in transport.go ([fba4b19](https://github.com/algolia/algoliasearch-client-go/commit/fba4b19))
- **feat:** Enable type-checking on restrictSources query parameter ([a3afd63](https://github.com/algolia/algoliasearch-client-go/commit/a3afd63))
- **feat:** Expose decompoundedAttributes as a setting parameter ([d00989d](https://github.com/algolia/algoliasearch-client-go/commit/d00989d))
- **feat:** Expose camelCaseAttributes as a setting parameter ([815489f](https://github.com/algolia/algoliasearch-client-go/commit/815489f))
- **feat:** Expose queryLanguages as a setting and search parameter ([6c7e3b6](https://github.com/algolia/algoliasearch-client-go/commit/6c7e3b6))
- **feat:** Expose keepDiacriticsOnCharacters as a setting parameter ([d9cf86a](https://github.com/algolia/algoliasearch-client-go/commit/d9cf86a))
- **doc:** Add missing date in ChangeLog.md ([776db89](https://github.com/algolia/algoliasearch-client-go/commit/776db89))

# [2.24.0](https://github.com/algolia/algoliasearch-client-go/compare/2.23.1...2.24.0) (2018-08-24)

### Summary

Hello everyone,

After recent investigations, we found out that the implementation of our retry
strategy in the Go client was missing a reseting feature. The retry strategy is
responsible for determining which hosts should be choosen for each call and to
try them sequentially if some are down or unreachable. However, the previous
implementation was lacking a reset of the hosts when all were marked down or
have been marked down for too long.

Because the new retry strategy implementation relies on the `context` package
from the standard library that was introduced in Go 1.7, Go 1.6 (released in
February 2017) and previous versions are no longer supported.

To improve the configurability of the client, the following methods are now
replacing the deprecated `Client.SetTimeout` method:

 - `Client.SetReadTimeout`
 - `Client.SetWriteTimeout`
 - `Client.SetAnalyticsTimeout`

Finally, for debugging purpose, we introduced a debugging output, controlled
by the `ALGOLIA_DEBUG` environment variable. Set it to `1` and extra
information will be displayed. Those informations being truncated, you can
access the unbounded information by setting `ALGOLIA_DEBUG` to anything number
greater than 1.

### Changes

- **feat:** Enable debugging output if ALGOLIA_DEBUG is set ([0cec615](https://github.com/algolia/algoliasearch-client-go/commit/0cec615))
- **refactor:** Deprecate Client.SetTimeout in favor of Client.Set{Read,Write,Analytics}Timeout methods ([198d9a1](https://github.com/algolia/algoliasearch-client-go/commit/198d9a1))
- **feat:** Add Client.Set{Read,Write,Analytics}Timeout methods ([c5dd60b](https://github.com/algolia/algoliasearch-client-go/commit/c5dd60b))
- **chore:** Remove intermediate Go versions to test in Travis ([e9f203e](https://github.com/algolia/algoliasearch-client-go/commit/e9f203e))
- **test:** Add tests for the RetryStrategy implementation ([642a11a](https://github.com/algolia/algoliasearch-client-go/commit/642a11a))
- **fix:** Improve retry strategy by reseting default hosts after some time ([dd317c3](https://github.com/algolia/algoliasearch-client-go/commit/dd317c3))
- **chore:** Upgrade github.com/stretchr/testify dependency from 1.1.4 to 1.2.2 ([65735eb](https://github.com/algolia/algoliasearch-client-go/commit/65735eb))
- **chore:** Drop support for Go 1.6 ([7f9e66e](https://github.com/algolia/algoliasearch-client-go/commit/7f9e66e))
- **doc:** Fix list typo in the ChangeLog.md ([eaa9f97](https://github.com/algolia/algoliasearch-client-go/commit/eaa9f97))

# [2.23.1](https://github.com/algolia/algoliasearch-client-go/compare/2.23.0...2.23.1) (2018-06-29)

### Summary

Hello everyone,

Today's patch release should make `Settings.ToMap` users happy, as some missing
fields that were not exported are now correctly made available in the result
`Map`. For more details, check the commit below. Other changes only touch chore
files.

### Changes

- **chore:** Add release.sh script ([dd734e0](https://github.com/algolia/algoliasearch-client-go/commit/dd734e0))
- **fix:** Export missing attributes in Settings.ToMap ([a0ec60a](https://github.com/algolia/algoliasearch-client-go/commit/a0ec60a))
- **misc:** Update README ([9c969af](https://github.com/algolia/algoliasearch-client-go/commit/9c969af))

# [2.23.0](https://github.com/algolia/algoliasearch-client-go/compare/2.22.0...2.23.0) (2018-06-19)
Hello everyone,

Big release today and as you may see, new format for the ChangeLog that we are
trying to improve. A lot of additions since 2.22.0, mainly for [A/B
testing](https://www.algolia.com/doc/guides/analytics/abtest-overview/)
features that are now available. After a careful audit of the client, we have
also started to deprecate some methods in favor of new ones, mainly for
consistency reasons. And as usual, a few fixes regarding some input/response
types that were wrong.

As usual, feel free to report any issue of question you may have in our [Github
issue tracker](https://github.com/algolia/algoliasearch-client-go/issues) and
to contribute by submitting your Pull Requests directly to [our Github
repository](https://github.com/algolia/algoliasearch-client-go/pulls). And for
a more detailed assistance regarding Algolia and its features, you may also
contact us directly at support@algolia.com.

Have a nice day.

- **feat:** Implement AB Testing via the new Analytics handler ([f1e8432](https://github.com/algolia/algoliasearch-client-go/commit/f1e8432))
- **chore:** Use testify to shorten tests and replace glide with dep ([32bd096](https://github.com/algolia/algoliasearch-client-go/commit/32bd096))
- **fix:** Expose missing fields related to Query Rules ([def6f71](https://github.com/algolia/algoliasearch-client-go/commit/def6f71))
- **test:** Look for missing objectID error for BatchOperation ([0e93e1e](https://github.com/algolia/algoliasearch-client-go/commit/0e93e1e))
- **test:** Minor logs added ([fcd7eef](https://github.com/algolia/algoliasearch-client-go/commit/fcd7eef))
- **refactor:** Better describe object errors ([cb025d7](https://github.com/algolia/algoliasearch-client-go/commit/cb025d7))
- **fix:** Prevent invalid batch requests correctly ([355486b](https://github.com/algolia/algoliasearch-client-go/commit/355486b))
- **feat:** Enforce Rule.ObjectID not to be empty (closes #397) ([22e2592](https://github.com/algolia/algoliasearch-client-go/commit/22e2592))
- **test:** Use *APIKey methods instead of deprecated *UserKey methods for tests ([5ff7df4](https://github.com/algolia/algoliasearch-client-go/commit/5ff7df4))
- **test:** Prevent tests from deleting the search-only API key of the application ([16fc6ef](https://github.com/algolia/algoliasearch-client-go/commit/16fc6ef))
- **refactor:** Deprecate all API key related Index methods in favor of Client ones (closes #413) ([53b957f](https://github.com/algolia/algoliasearch-client-go/commit/53b957f))
- **feat:** Type-check search parameter sumOrFiltersScores (closes #402) ([0d07917](https://github.com/algolia/algoliasearch-client-go/commit/0d07917))
- **fix:** Rename mistyped RulePatternAnchoring StarstWith into StartsWith (closes #399) ([1aa19ad](https://github.com/algolia/algoliasearch-client-go/commit/1aa19ad))
- **fix:** Expose missing ID field in SaveSynonym response (closes #393) ([a83e4dc](https://github.com/algolia/algoliasearch-client-go/commit/a83e4dc))
- **feat:** Deprecate Index.AddSynonym in favor of Index.SaveSynonym (closes #391) ([d718682](https://github.com/algolia/algoliasearch-client-go/commit/d718682))
- **feat:** Deprecate Client.ListKeys in favor of Client.ListAPIKeys (closes #390) ([39bd6c1](https://github.com/algolia/algoliasearch-client-go/commit/39bd6c1))
- **test:** Add integration test for GenerateSecuredAPIKey ([0eae8c2](https://github.com/algolia/algoliasearch-client-go/commit/0eae8c2))
- **fix:** Remove 'referers' field check for GenerateSecuredAPIKey (closes #388) ([98fb88c](https://github.com/algolia/algoliasearch-client-go/commit/98fb88c))
- **fix:** Expose missing NbPages field to listIndexesRes (closes #387) ([fe19eb4](https://github.com/algolia/algoliasearch-client-go/commit/fe19eb4))
- **fix:** Add missing trailing S to IndexRes' NumberOfPendingTasks (closes #386) ([9c72938](https://github.com/algolia/algoliasearch-client-go/commit/9c72938))
- **fix:** Use correct return type for DeleteBy (closes #383) ([bddeb63](https://github.com/algolia/algoliasearch-client-go/commit/bddeb63))
- **feat:** Expose processed boolean for MultipleQueriesRes when using stopIfEnoughMatchesStrategy (closes #379) ([58a1cd1](https://github.com/algolia/algoliasearch-client-go/commit/58a1cd1))
- **feat:** Expose exhaustiveFacetsCount boolean in SearchFacetRes (closes #377) ([fb4ba25](https://github.com/algolia/algoliasearch-client-go/commit/fb4ba25))

# [2.22.0](https://github.com/algolia/algoliasearch-client-go/compare/2.21.3...2.22.0) (2018-03-26)
- **feat:** Add clickAnalytics to search option parameters ([7b8d667](https://github.com/algolia/algoliasearch-client-go/commit/7b8d667))
- **chore:** Update contribution-related files ([af85276](https://github.com/algolia/algoliasearch-client-go/commit/af85276))
- **chore:** use gotest to have colored tests output ([f3d66ff](https://github.com/algolia/algoliasearch-client-go/commit/f3d66ff))
- **chore:** reformat .travis.yml ([1290509](https://github.com/algolia/algoliasearch-client-go/commit/1290509))
- **chore:** test against Go 1.10 in Travis ([4c21ba0](https://github.com/algolia/algoliasearch-client-go/commit/4c21ba0))
- **test:** Fix typos and printf formats ([05560b2](https://github.com/algolia/algoliasearch-client-go/commit/05560b2))

# [2.21.3](https://github.com/algolia/algoliasearch-client-go/compare/2.21.2...2.21.3) (2018-02-13)
- **fix:** Expose missing indexes in Key ([df4fedb](https://github.com/algolia/algoliasearch-client-go/commit/df4fedb))

# [2.21.2](https://github.com/algolia/algoliasearch-client-go/compare/2.21.1...2.21.2) (2018-02-05)
- **doc:** Bump to 2.21.2 ([ce90912](https://github.com/algolia/algoliasearch-client-go/commit/ce90912))
- **fix:** Expose missing maxFacetHits as setting ([1c59801](https://github.com/algolia/algoliasearch-client-go/commit/1c59801))
- **fix:** Expose missing paginationLimitedTo as setting ([5fab01d](https://github.com/algolia/algoliasearch-client-go/commit/5fab01d))
- **fix:** Expose missing restrictHighlightAndSnippetArrays as setting ([5d02c93](https://github.com/algolia/algoliasearch-client-go/commit/5d02c93))
- **test:** add 1.9 to travis ([bb6ab50](https://github.com/algolia/algoliasearch-client-go/commit/bb6ab50))
- **refactor:** use time.Since for passed time ([b718cf5](https://github.com/algolia/algoliasearch-client-go/commit/b718cf5))
- **refactor:** run gofmt with simplify flag ([1d6924d](https://github.com/algolia/algoliasearch-client-go/commit/1d6924d))

# [2.21.1](https://github.com/algolia/algoliasearch-client-go/compare/2.21.0...2.21.1) (2017-11-29)
- **fix:** Add missing removeWordsIfNoResults to settings type ([4628a86](https://github.com/algolia/algoliasearch-client-go/commit/4628a86))

# [2.21.0](https://github.com/algolia/algoliasearch-client-go/compare/2.20.0...2.21.0) (2017-11-14)
- **test:** Add tests for scoped copy ([4f5d0cd](https://github.com/algolia/algoliasearch-client-go/commit/4f5d0cd))
- **feat:** Implement scoped copy methods ([3c8e3b1](https://github.com/algolia/algoliasearch-client-go/commit/3c8e3b1))
- **feat:** Expose SortFacetValuesBy setting ([e9d1386](https://github.com/algolia/algoliasearch-client-go/commit/e9d1386))

# [2.20.0](https://github.com/algolia/algoliasearch-client-go/compare/2.19.1...2.20.0) (2017-10-11)
- **doc:** Reformat ChangeLog.md ([68257924](https://github.com/algolia/algoliasearch-client-go/commit/68257924))
- **doc:** Add missing description of 2.19.1 release to the ChangeLog.md ([8114f6e](https://github.com/algolia/algoliasearch-client-go/commit/8114f6e))
- **feat:** Implement RuleIterator to browse over all the query rules of a given index ([40e35c2](https://github.com/algolia/algoliasearch-client-go/commit/40e35c2))
- **feat:** Implement SynonymIterator to browse over all the synonyms of a given index ([3554426](https://github.com/algolia/algoliasearch-client-go/commit/3554426))
- **test:** Remove unused Client parameter from addOneObject function ([13edced](https://github.com/algolia/algoliasearch-client-go/commit/13edced))
- **fix:** Expose correct highlighting field (_highlightResult not _highlightedResult) for Synonym answers ([688e76a](https://github.com/algolia/algoliasearch-client-go/commit/688e76a))
- **fix:** Add missing _highlightResult field for Query Rules answers ([666987b](https://github.com/algolia/algoliasearch-client-go/commit/666987b))
- **test:** Move shared testing code to dedicated functions ([76e0896](https://github.com/algolia/algoliasearch-client-go/commit/76e0896))
- **refactor:** Remove commented dead code ([b7988cf](https://github.com/algolia/algoliasearch-client-go/commit/b7988cf))
- **refactor:** Move NoMoreHitsErr definition to a dedicated file ([6e3dff3](https://github.com/algolia/algoliasearch-client-go/commit/6e3dff3))

# 2.19.1 (2017-09-06)
## Changes
- Expose the `ExhaustiveNbHits` boolean in `QueryRes`
- Remove useless `omitempty` tags from `BrowseRes` and `QueryRes`

# 2.19.0 (2017-08-28)
## Additions
- Implement `Index.DeleteBy`
## Changes
- Deprecate `Index.DeleteByQuery`

# 2.18.0 (2017-08-23)
## Additions
- Implement `Client.SetMaxIdleConnsPerHosts` to let the user override `MaxIdleConnsPerHost`
## Changes
- Change the default value of `MaxIdleConnsPerHost` from 2 to 64

# 2.17.0 (2017-08-07)
## Additions
- Implement `*WithRequestOptions` methods for both Client and Index implementations

# 2.16.0 (2017-08-02)
## Fixes
- Correctly stop the browse iteration of `Index.DeleteByQuery`
## Changes
- Improve performances of `Index.DeleteByQuery`
- Improve documentation comment of `Index.BrowseAll`
- `Index.BrowseAll`'s terminal "no more hits" error now has its own variable: `NoMoreHitsErr`

# 2.15.0 (2017-06-30)
## Additions
- Implement new Index methods for Algolia Query Rules endpoints

# 2.14.0 (2017-06-06)
## Fixes
- Fix `Client.GetLogs` by correctly passing the given parameters via the URL
- Accept `indexName` as `Client.GetLogs` parameter as string

# 2.13.0 (2017-04-26)
## Additions
- Accept percentileComputation query parameter as bool

# 2.12.0 (2017-04-12)
## Additions
- Add default ProxyFunc to the default transport layer of the HTTP client
## Fixes
- Improve retry strategy to let it use the provided hosts if any in priority
- Re-arrange the order of hosts of the retry strategy

# 2.11.0 (2017-03-24)
## Additions
- Add PartialUpdateOp type and helpers to simplify partial update of records

# 2.10.0 (2017-03-24)
## Fixes
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
## Misc
- Run the tests with Go 1.8
- Do not run the tests with Go 1.5 anymore

# 2.9.1 (2017-03-22)
## Fixes
- Fix the `attributesToRetrieve` handling of `GetObject/GetObjects` methods

# 2.9.0 (2017-02-24)
## Additions
- (Get|Add|Update|Delete)APIKey methods
## Changes
- Deprecate (Get|Add|Update|Delete)UserKey methods

# 2.8.0 (2017-02-14)
## Additions
- Allow `maxFacetHits` parameter in queries and parameters

# 2.7.1 (2017-01-25)
## Fixes
- Ensure `net/http.Response.Body.Close()` is called

# 2.7.0 (2016-12-28)
## Changes
- Only accepts the following parameters for `GenerateSecuredAPIKey`:
  + `userToken` (string identifier generally used to rate-limit users per IP)
  + `validUntil` (timestamp of the expiration date)
  + `restrictIndices` (comma-separated string list of the indices to restrict)
  + `referers` (string slice of allowed referers)
  + `restrictSources` (string of the allowed IPv4 network)

# 2.6.0 (2016-12-03)
## Fixes
- Improve the transport layer code
- Keep the last active connection between requests to better handle DNS timeouts

# 2.5.0 (2016-11-29)
## Additions
- Add `Index.SearchForFacetValues` method
  + Same as `Index.SearchFacet`
  + `Index.SearchFacet` is kept for backward-compatibility

# 2.4.0 (2016-11-02)
## Changes
- Accept both boolean and string slice for `Settings.IgnorePlurals`

# 2.3.1 (2016-11-02)
## Fixes
- Handle missing parameters in `Settings.ToMap`:
  + `responseFields`
  + `typoTolerance`

# 2.3.0 (2016-11-02)
## Additions
- Accept `responseFields` parameter in queries and settings

# 2.2.0 (2016-10-19)
## Additions
- Add `Index.SearchFacet` method

# 2.1.2 (2016-10-19)
## Fixes
- Exclude the `testing` package from the build
- Fix typo and type checking for `attributeForDistinct`

# 2.1.1 (2016-10-16)
## Additions
- Accept `facetingAfterDistinct` parameter in queries

# 2.1.0 (2016-10-03)
## Fixes
- Fix `Client.ListIndexes`
- Fix `LogRes` type
- Fix `aroundRadius` type
- Fix `distinct` type
- Fix `removeStopWords` type
- Fix `NewOneWaySynonym` method name (minor typo)
- Fix `NewAltCorrectionSynonym` method name (minor typo)
## Additions
- Add `Client.SetHTTPClient` method
- Add `Index.GetObjectsAttrs` method
- Add `Settings.ToMap` method
- Add `FileSize` field in the `IndexRes` response type
- Add `NumberOfPendingTask` field in the `IndexRes` response type
## Changes
- Implement the new Synonym API
- Rename `fowardToSlaves` parameter into `forwardToReplicas`
- Rename `attributesToIndex` setting into `searchableAttributes`
- Rename `numericAttributesToIndex` setting into `numericAttributesForFiltering`
- Allow `exactOnSingleWordQuery` parameter in queries
- Allow `alternativesAsExact` parameter in queries
- Allow `forwardToReplicas` parameter in queries
## Misc
- Avoid the inclusion of the `syscall` package to be used within Google App Engine
## Tests
- Add tests
- Run tests in parallel
- Only test across the last 3 major versions of Go (currently 1.5, 1.6 and 1.7)

# 2.0.0 (2016-06-06)
- Type every response from the API
- Add tests
- Breaking changes listed here: https://github.com/algolia/algoliasearch-client-go/wiki/Migration-guide-to-version-2.x

# 1.5.0 (2016-06-01)
- Code refactoring ahead of version 2.0.0

# 1.4.0 (2014-10-16)
- Add new secured api keys
- Fix bug with distinct in deleteByQuery

# 1.3.0 (2014-07-07)
- Add new browse methods

# 1.2.1 (2014-05-04)
- Add new methods to add/update api key
- Add batch method to target multiple indices
- Add strategy parameter for the multipleQueries
- Add new method to generate secured api key from query parameters

# 1.2.0 (2015-04-09)
- New implementation for the retry on another DNS

# 1.1.1 (2015-11-23)
- Fix the retry for a network error
- Drop Go 1.2 support
- Add timeout

# 1.1.0 (2014-11-30)
- Bump to 1.1.0.
- Switch to .net.

# 1.0.2 (2014-11-27)
- Bump to 1.0.2.
- Wait more for the ACL.
- Use sandboxes.
- Fix go tip.
- Try to fix go 1.2.
- Ability to set custom hosts.
- Fix missing safeName call.
- Add safeName function.
- Update ChangeLog.

# 1.0.1 (2014-10-22)
- Bump to 1.0.1.
- Fix retry.

# 1.0.0 (2014-10-16)
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
