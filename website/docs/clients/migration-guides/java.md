---
title: Java
---

### Methods targeting an `indexName`

Prior to the `initIndex` removal stated in the [common breaking changes](/docs/clients/migration-guides/#common-breaking-changes), all methods previously available at the `initIndex` level requires the `indexName` to be sent with the query.

That also mean you need to explicit the type you want to be returned from your queries, when it applies.

```java
import com.algolia.api.SearchClient;
import com.algolia.model.search.*;

SearchClient client = new SearchClient("<YOUR_APP_ID>", "<YOUR_API_KEY>");

client.search(
  new SearchMethodParams()
    .addRequests(SearchQuery.of(
      new SearchForHits()
        .setIndexName("<YOUR_INDEX_NAME>")
        .setQuery("<YOUR_QUERY>")
    )
  ),
  MyObject.class
);
```




