package com.algolia.playground;

import com.algolia.api.SearchClient;
import com.algolia.exceptions.*;
import com.algolia.model.search.*;
import com.algolia.utils.AlgoliaAgent;
import com.algolia.utils.ClientOptions;
import io.github.cdimascio.dotenv.Dotenv;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;

public class Search {

  public static void main(String[] args) {
    Dotenv dotenv = Dotenv.configure().directory("../").load();

    SearchClient client = new SearchClient(
      dotenv.get("ALGOLIA_APPLICATION_ID"),
      dotenv.get("ALGOLIA_SEARCH_KEY"),
      ClientOptions.build()
        .addAlgoliaAgentSegment("test", "8.0.0")
        .addAlgoliaAgentSegment("JVM", "11.0.14")
        .addAlgoliaAgentSegment("no version")
    );

    String indexName = dotenv.get("SEARCH_INDEX");
    String query = dotenv.get("SEARCH_QUERY");

    try {
      List<Map<String, Object>> records = Arrays.asList(Collections.singletonMap("name", "Tom Cruise"), Collections.singletonMap("name", "Scarlett Johansson"));

      for (Map<String, Object> record : records) {
        client.saveObject(
          indexName,
          record
        );
      }

      SearchMethodParams searchMethodParams = new SearchMethodParams();
      List<SearchQuery> requests = new ArrayList<>();
      SearchForHits request = new SearchForHits();
      request.setIndexName(indexName);
      request.setQuery(query);
      requests.add(SearchQuery.ofSearchForHits(request));
      searchMethodParams.setRequests(requests);

      CompletableFuture<SearchResponses> result = client.searchAsync(searchMethodParams);

      SearchResponses sr = result.get();
      System.out.println(sr);
    } catch (InterruptedException e) {
      System.err.println("InterrupedException" + e.getMessage());
      e.printStackTrace();
    } catch (ExecutionException e) {
      System.err.println("ExecutionException" + e.getMessage());
      e.printStackTrace();
    } catch (AlgoliaApiException e) {
      // the API failed
      System.err.println("Exception when calling SearchClient#search");
      System.err.println("Status code: " + e.getHttpErrorCode());
      System.err.println("Reason: " + e.getMessage());
      e.printStackTrace();
    } catch (AlgoliaRetryException e) {
      // the retry failed
      System.err.println("Exception in the retry strategy");
      e.printStackTrace();
    } catch (AlgoliaRuntimeException e) {
      // the serialization or something else failed
      e.printStackTrace();
    }
  }
}
