package com.algolia.methods.requests;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertEquals;

import com.algolia.EchoRequester;
import com.algolia.EchoResponse;
import com.algolia.api.RecommendClient;
import com.algolia.model.recommend.*;
import com.algolia.utils.JSON;
import com.algolia.utils.RequestOptions;
import com.google.gson.reflect.TypeToken;
import java.util.*;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;
import org.skyscreamer.jsonassert.JSONAssert;
import org.skyscreamer.jsonassert.JSONCompareMode;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class RecommendClientTests {

  private RecommendClient client;
  private EchoRequester requester;

  @BeforeAll
  void init() {
    requester = new EchoRequester();
    client = new RecommendClient("appId", "apiKey", requester);
  }

  @Test
  @DisplayName("allow del method for a custom path with minimal parameters")
  void delTest0() {
    String path0 = "/test/minimal";

    assertDoesNotThrow(() -> {
      client.del(path0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/minimal");
    assertEquals(req.method, "DELETE");
  }

  @Test
  @DisplayName("allow del method for a custom path with all parameters")
  void delTest1() {
    String path0 = "/test/all";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }

    assertDoesNotThrow(() -> {
      client.del(path0, parameters0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/all");
    assertEquals(req.method, "DELETE");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("allow get method for a custom path with minimal parameters")
  void getTest0() {
    String path0 = "/test/minimal";

    assertDoesNotThrow(() -> {
      client.get(path0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/minimal");
    assertEquals(req.method, "GET");
  }

  @Test
  @DisplayName("allow get method for a custom path with all parameters")
  void getTest1() {
    String path0 = "/test/all";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }

    assertDoesNotThrow(() -> {
      client.get(path0, parameters0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/all");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName(
    "get recommendations for recommend model with minimal parameters"
  )
  void getRecommendationsTest0() {
    GetRecommendationsParams getRecommendationsParams0 = new GetRecommendationsParams();
    {
      List<RecommendationsRequest> requests1 = new ArrayList<>();
      {
        RecommendationRequest requests_02 = new RecommendationRequest();
        {
          String indexName3 = "indexName";
          requests_02.setIndexName(indexName3);
          String objectID3 = "objectID";
          requests_02.setObjectID(objectID3);
          RecommendationModels model3 = RecommendationModels.fromValue(
            "related-products"
          );
          requests_02.setModel(model3);
          int threshold3 = 42;
          requests_02.setThreshold(threshold3);
        }
        requests1.add(
          RecommendationsRequest.ofRecommendationRequest(requests_02)
        );
      }
      getRecommendationsParams0.setRequests(requests1);
    }

    assertDoesNotThrow(() -> {
      client.getRecommendations(getRecommendationsParams0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/indexes/*/recommendations");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"requests\":[{\"indexName\":\"indexName\",\"objectID\":\"objectID\",\"model\":\"related-products\",\"threshold\":42}]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
  }

  @Test
  @DisplayName("get recommendations for recommend model with all parameters")
  void getRecommendationsTest1() {
    GetRecommendationsParams getRecommendationsParams0 = new GetRecommendationsParams();
    {
      List<RecommendationsRequest> requests1 = new ArrayList<>();
      {
        RecommendationRequest requests_02 = new RecommendationRequest();
        {
          String indexName3 = "indexName";
          requests_02.setIndexName(indexName3);
          String objectID3 = "objectID";
          requests_02.setObjectID(objectID3);
          RecommendationModels model3 = RecommendationModels.fromValue(
            "related-products"
          );
          requests_02.setModel(model3);
          int threshold3 = 42;
          requests_02.setThreshold(threshold3);
          int maxRecommendations3 = 10;
          requests_02.setMaxRecommendations(maxRecommendations3);
          SearchParamsObject queryParameters3 = new SearchParamsObject();
          {
            String query4 = "myQuery";
            queryParameters3.setQuery(query4);
            List<String> facetFilters4 = new ArrayList<>();
            {
              String facetFilters_05 = "query";
              facetFilters4.add(facetFilters_05);
            }
            queryParameters3.setFacetFilters(
              FacetFilters.ofListString(facetFilters4)
            );
          }
          requests_02.setQueryParameters(queryParameters3);
          SearchParamsObject fallbackParameters3 = new SearchParamsObject();
          {
            String query4 = "myQuery";
            fallbackParameters3.setQuery(query4);
            List<String> facetFilters4 = new ArrayList<>();
            {
              String facetFilters_05 = "fallback";
              facetFilters4.add(facetFilters_05);
            }
            fallbackParameters3.setFacetFilters(
              FacetFilters.ofListString(facetFilters4)
            );
          }
          requests_02.setFallbackParameters(fallbackParameters3);
        }
        requests1.add(
          RecommendationsRequest.ofRecommendationRequest(requests_02)
        );
      }
      getRecommendationsParams0.setRequests(requests1);
    }

    assertDoesNotThrow(() -> {
      client.getRecommendations(getRecommendationsParams0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/indexes/*/recommendations");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"requests\":[{\"indexName\":\"indexName\",\"objectID\":\"objectID\",\"model\":\"related-products\",\"threshold\":42,\"maxRecommendations\":10,\"queryParameters\":{\"query\":\"myQuery\",\"facetFilters\":[\"query\"]},\"fallbackParameters\":{\"query\":\"myQuery\",\"facetFilters\":[\"fallback\"]}}]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
  }

  @Test
  @DisplayName("get recommendations for trending model with minimal parameters")
  void getRecommendationsTest2() {
    GetRecommendationsParams getRecommendationsParams0 = new GetRecommendationsParams();
    {
      List<RecommendationsRequest> requests1 = new ArrayList<>();
      {
        TrendingRequest requests_02 = new TrendingRequest();
        {
          String indexName3 = "indexName";
          requests_02.setIndexName(indexName3);
          TrendingModels model3 = TrendingModels.fromValue("trending-items");
          requests_02.setModel(model3);
          int threshold3 = 42;
          requests_02.setThreshold(threshold3);
        }
        requests1.add(RecommendationsRequest.ofTrendingRequest(requests_02));
      }
      getRecommendationsParams0.setRequests(requests1);
    }

    assertDoesNotThrow(() -> {
      client.getRecommendations(getRecommendationsParams0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/indexes/*/recommendations");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"requests\":[{\"indexName\":\"indexName\",\"model\":\"trending-items\",\"threshold\":42}]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
  }

  @Test
  @DisplayName("get recommendations for trending model with all parameters")
  void getRecommendationsTest3() {
    GetRecommendationsParams getRecommendationsParams0 = new GetRecommendationsParams();
    {
      List<RecommendationsRequest> requests1 = new ArrayList<>();
      {
        TrendingRequest requests_02 = new TrendingRequest();
        {
          String indexName3 = "indexName";
          requests_02.setIndexName(indexName3);
          TrendingModels model3 = TrendingModels.fromValue("trending-items");
          requests_02.setModel(model3);
          int threshold3 = 42;
          requests_02.setThreshold(threshold3);
          int maxRecommendations3 = 10;
          requests_02.setMaxRecommendations(maxRecommendations3);
          String facetName3 = "myFacetName";
          requests_02.setFacetName(facetName3);
          String facetValue3 = "myFacetValue";
          requests_02.setFacetValue(facetValue3);
          SearchParamsObject queryParameters3 = new SearchParamsObject();
          {
            String query4 = "myQuery";
            queryParameters3.setQuery(query4);
            List<String> facetFilters4 = new ArrayList<>();
            {
              String facetFilters_05 = "query";
              facetFilters4.add(facetFilters_05);
            }
            queryParameters3.setFacetFilters(
              FacetFilters.ofListString(facetFilters4)
            );
          }
          requests_02.setQueryParameters(queryParameters3);
          SearchParamsObject fallbackParameters3 = new SearchParamsObject();
          {
            String query4 = "myQuery";
            fallbackParameters3.setQuery(query4);
            List<String> facetFilters4 = new ArrayList<>();
            {
              String facetFilters_05 = "fallback";
              facetFilters4.add(facetFilters_05);
            }
            fallbackParameters3.setFacetFilters(
              FacetFilters.ofListString(facetFilters4)
            );
          }
          requests_02.setFallbackParameters(fallbackParameters3);
        }
        requests1.add(RecommendationsRequest.ofTrendingRequest(requests_02));
      }
      getRecommendationsParams0.setRequests(requests1);
    }

    assertDoesNotThrow(() -> {
      client.getRecommendations(getRecommendationsParams0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/indexes/*/recommendations");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"requests\":[{\"indexName\":\"indexName\",\"model\":\"trending-items\",\"threshold\":42,\"maxRecommendations\":10,\"facetName\":\"myFacetName\",\"facetValue\":\"myFacetValue\",\"queryParameters\":{\"query\":\"myQuery\",\"facetFilters\":[\"query\"]},\"fallbackParameters\":{\"query\":\"myQuery\",\"facetFilters\":[\"fallback\"]}}]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
  }

  @Test
  @DisplayName("get multiple recommendations with minimal parameters")
  void getRecommendationsTest4() {
    GetRecommendationsParams getRecommendationsParams0 = new GetRecommendationsParams();
    {
      List<RecommendationsRequest> requests1 = new ArrayList<>();
      {
        RecommendationRequest requests_02 = new RecommendationRequest();
        {
          String indexName3 = "indexName1";
          requests_02.setIndexName(indexName3);
          String objectID3 = "objectID1";
          requests_02.setObjectID(objectID3);
          RecommendationModels model3 = RecommendationModels.fromValue(
            "related-products"
          );
          requests_02.setModel(model3);
          int threshold3 = 21;
          requests_02.setThreshold(threshold3);
        }
        requests1.add(
          RecommendationsRequest.ofRecommendationRequest(requests_02)
        );
        RecommendationRequest requests_12 = new RecommendationRequest();
        {
          String indexName3 = "indexName2";
          requests_12.setIndexName(indexName3);
          String objectID3 = "objectID2";
          requests_12.setObjectID(objectID3);
          RecommendationModels model3 = RecommendationModels.fromValue(
            "related-products"
          );
          requests_12.setModel(model3);
          int threshold3 = 21;
          requests_12.setThreshold(threshold3);
        }
        requests1.add(
          RecommendationsRequest.ofRecommendationRequest(requests_12)
        );
      }
      getRecommendationsParams0.setRequests(requests1);
    }

    assertDoesNotThrow(() -> {
      client.getRecommendations(getRecommendationsParams0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/indexes/*/recommendations");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"requests\":[{\"indexName\":\"indexName1\",\"objectID\":\"objectID1\",\"model\":\"related-products\",\"threshold\":21},{\"indexName\":\"indexName2\",\"objectID\":\"objectID2\",\"model\":\"related-products\",\"threshold\":21}]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
  }

  @Test
  @DisplayName("get multiple recommendations with all parameters")
  void getRecommendationsTest5() {
    GetRecommendationsParams getRecommendationsParams0 = new GetRecommendationsParams();
    {
      List<RecommendationsRequest> requests1 = new ArrayList<>();
      {
        RecommendationRequest requests_02 = new RecommendationRequest();
        {
          String indexName3 = "indexName1";
          requests_02.setIndexName(indexName3);
          String objectID3 = "objectID1";
          requests_02.setObjectID(objectID3);
          RecommendationModels model3 = RecommendationModels.fromValue(
            "related-products"
          );
          requests_02.setModel(model3);
          int threshold3 = 21;
          requests_02.setThreshold(threshold3);
          int maxRecommendations3 = 10;
          requests_02.setMaxRecommendations(maxRecommendations3);
          SearchParamsObject queryParameters3 = new SearchParamsObject();
          {
            String query4 = "myQuery";
            queryParameters3.setQuery(query4);
            List<String> facetFilters4 = new ArrayList<>();
            {
              String facetFilters_05 = "query1";
              facetFilters4.add(facetFilters_05);
            }
            queryParameters3.setFacetFilters(
              FacetFilters.ofListString(facetFilters4)
            );
          }
          requests_02.setQueryParameters(queryParameters3);
          SearchParamsObject fallbackParameters3 = new SearchParamsObject();
          {
            String query4 = "myQuery";
            fallbackParameters3.setQuery(query4);
            List<String> facetFilters4 = new ArrayList<>();
            {
              String facetFilters_05 = "fallback1";
              facetFilters4.add(facetFilters_05);
            }
            fallbackParameters3.setFacetFilters(
              FacetFilters.ofListString(facetFilters4)
            );
          }
          requests_02.setFallbackParameters(fallbackParameters3);
        }
        requests1.add(
          RecommendationsRequest.ofRecommendationRequest(requests_02)
        );
        RecommendationRequest requests_12 = new RecommendationRequest();
        {
          String indexName3 = "indexName2";
          requests_12.setIndexName(indexName3);
          String objectID3 = "objectID2";
          requests_12.setObjectID(objectID3);
          RecommendationModels model3 = RecommendationModels.fromValue(
            "related-products"
          );
          requests_12.setModel(model3);
          int threshold3 = 21;
          requests_12.setThreshold(threshold3);
          int maxRecommendations3 = 10;
          requests_12.setMaxRecommendations(maxRecommendations3);
          SearchParamsObject queryParameters3 = new SearchParamsObject();
          {
            String query4 = "myQuery";
            queryParameters3.setQuery(query4);
            List<String> facetFilters4 = new ArrayList<>();
            {
              String facetFilters_05 = "query2";
              facetFilters4.add(facetFilters_05);
            }
            queryParameters3.setFacetFilters(
              FacetFilters.ofListString(facetFilters4)
            );
          }
          requests_12.setQueryParameters(queryParameters3);
          SearchParamsObject fallbackParameters3 = new SearchParamsObject();
          {
            String query4 = "myQuery";
            fallbackParameters3.setQuery(query4);
            List<String> facetFilters4 = new ArrayList<>();
            {
              String facetFilters_05 = "fallback2";
              facetFilters4.add(facetFilters_05);
            }
            fallbackParameters3.setFacetFilters(
              FacetFilters.ofListString(facetFilters4)
            );
          }
          requests_12.setFallbackParameters(fallbackParameters3);
        }
        requests1.add(
          RecommendationsRequest.ofRecommendationRequest(requests_12)
        );
      }
      getRecommendationsParams0.setRequests(requests1);
    }

    assertDoesNotThrow(() -> {
      client.getRecommendations(getRecommendationsParams0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/indexes/*/recommendations");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"requests\":[{\"indexName\":\"indexName1\",\"objectID\":\"objectID1\",\"model\":\"related-products\",\"threshold\":21,\"maxRecommendations\":10,\"queryParameters\":{\"query\":\"myQuery\",\"facetFilters\":[\"query1\"]},\"fallbackParameters\":{\"query\":\"myQuery\",\"facetFilters\":[\"fallback1\"]}},{\"indexName\":\"indexName2\",\"objectID\":\"objectID2\",\"model\":\"related-products\",\"threshold\":21,\"maxRecommendations\":10,\"queryParameters\":{\"query\":\"myQuery\",\"facetFilters\":[\"query2\"]},\"fallbackParameters\":{\"query\":\"myQuery\",\"facetFilters\":[\"fallback2\"]}}]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
  }

  @Test
  @DisplayName("get frequently bought together recommendations")
  void getRecommendationsTest6() {
    GetRecommendationsParams getRecommendationsParams0 = new GetRecommendationsParams();
    {
      List<RecommendationsRequest> requests1 = new ArrayList<>();
      {
        RecommendationRequest requests_02 = new RecommendationRequest();
        {
          String indexName3 = "indexName1";
          requests_02.setIndexName(indexName3);
          String objectID3 = "objectID1";
          requests_02.setObjectID(objectID3);
          RecommendationModels model3 = RecommendationModels.fromValue(
            "bought-together"
          );
          requests_02.setModel(model3);
          int threshold3 = 42;
          requests_02.setThreshold(threshold3);
        }
        requests1.add(
          RecommendationsRequest.ofRecommendationRequest(requests_02)
        );
      }
      getRecommendationsParams0.setRequests(requests1);
    }

    assertDoesNotThrow(() -> {
      client.getRecommendations(getRecommendationsParams0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/indexes/*/recommendations");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"requests\":[{\"indexName\":\"indexName1\",\"objectID\":\"objectID1\",\"model\":\"bought-together\",\"threshold\":42}]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
  }

  @Test
  @DisplayName("allow post method for a custom path with minimal parameters")
  void postTest0() {
    String path0 = "/test/minimal";

    assertDoesNotThrow(() -> {
      client.post(path0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/minimal");
    assertEquals(req.method, "POST");
  }

  @Test
  @DisplayName("allow post method for a custom path with all parameters")
  void postTest1() {
    String path0 = "/test/all";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String body1 = "parameters";
      body0.put("body", body1);
    }

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/all");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"body\":\"parameters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("requestOptions can override default query parameters")
  void postTest2() {
    String path0 = "/test/requestOptions";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String facet1 = "filters";
      body0.put("facet", facet1);
    }

    RequestOptions requestOptions = new RequestOptions();
    requestOptions.addExtraQueryParameters("query", "myQueryParameter");

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"facet\":\"filters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"myQueryParameter\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("requestOptions merges query parameters with default ones")
  void postTest3() {
    String path0 = "/test/requestOptions";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String facet1 = "filters";
      body0.put("facet", facet1);
    }

    RequestOptions requestOptions = new RequestOptions();
    requestOptions.addExtraQueryParameters("query2", "myQueryParameter");

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"facet\":\"filters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\",\"query2\":\"myQueryParameter\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("requestOptions can override default headers")
  void postTest4() {
    String path0 = "/test/requestOptions";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String facet1 = "filters";
      body0.put("facet", facet1);
    }

    RequestOptions requestOptions = new RequestOptions();
    requestOptions.addExtraHeader("x-algolia-api-key", "myApiKey");

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"facet\":\"filters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }

    Map<String, String> expectedHeaders = JSON.deserialize(
      "{\"x-algolia-api-key\":\"myApiKey\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, String> actualHeaders = req.headers;

    for (Map.Entry<String, String> p : expectedHeaders.entrySet()) {
      assertEquals(actualHeaders.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("requestOptions merges headers with default ones")
  void postTest5() {
    String path0 = "/test/requestOptions";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String facet1 = "filters";
      body0.put("facet", facet1);
    }

    RequestOptions requestOptions = new RequestOptions();
    requestOptions.addExtraHeader("x-algolia-api-key", "myApiKey");

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"facet\":\"filters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }

    Map<String, String> expectedHeaders = JSON.deserialize(
      "{\"x-algolia-api-key\":\"myApiKey\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, String> actualHeaders = req.headers;

    for (Map.Entry<String, String> p : expectedHeaders.entrySet()) {
      assertEquals(actualHeaders.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("requestOptions queryParameters accepts booleans")
  void postTest6() {
    String path0 = "/test/requestOptions";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String facet1 = "filters";
      body0.put("facet", facet1);
    }

    RequestOptions requestOptions = new RequestOptions();
    requestOptions.addExtraQueryParameters("isItWorking", true);

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"facet\":\"filters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\",\"isItWorking\":\"true\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("requestOptions queryParameters accepts integers")
  void postTest7() {
    String path0 = "/test/requestOptions";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String facet1 = "filters";
      body0.put("facet", facet1);
    }

    RequestOptions requestOptions = new RequestOptions();
    requestOptions.addExtraQueryParameters("myParam", 2);

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"facet\":\"filters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\",\"myParam\":\"2\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("requestOptions queryParameters accepts list of string")
  void postTest8() {
    String path0 = "/test/requestOptions";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String facet1 = "filters";
      body0.put("facet", facet1);
    }

    RequestOptions requestOptions = new RequestOptions();
    List<Object> requestOptionsQueryParameters = new ArrayList<>();
    requestOptionsQueryParameters.add("c");
    requestOptionsQueryParameters.add("d");
    requestOptions.addExtraQueryParameters(
      "myParam",
      requestOptionsQueryParameters
    );

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"facet\":\"filters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\",\"myParam\":\"c,d\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("requestOptions queryParameters accepts list of booleans")
  void postTest9() {
    String path0 = "/test/requestOptions";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String facet1 = "filters";
      body0.put("facet", facet1);
    }

    RequestOptions requestOptions = new RequestOptions();
    List<Object> requestOptionsQueryParameters = new ArrayList<>();
    requestOptionsQueryParameters.add(true);
    requestOptionsQueryParameters.add(true);
    requestOptionsQueryParameters.add(false);
    requestOptions.addExtraQueryParameters(
      "myParam",
      requestOptionsQueryParameters
    );

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"facet\":\"filters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\",\"myParam\":\"true,true,false\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("requestOptions queryParameters accepts list of integers")
  void postTest10() {
    String path0 = "/test/requestOptions";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String facet1 = "filters";
      body0.put("facet", facet1);
    }

    RequestOptions requestOptions = new RequestOptions();
    List<Object> requestOptionsQueryParameters = new ArrayList<>();
    requestOptionsQueryParameters.add(1);
    requestOptionsQueryParameters.add(2);
    requestOptions.addExtraQueryParameters(
      "myParam",
      requestOptionsQueryParameters
    );

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"facet\":\"filters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\",\"myParam\":\"1,2\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("allow put method for a custom path with minimal parameters")
  void putTest0() {
    String path0 = "/test/minimal";

    assertDoesNotThrow(() -> {
      client.put(path0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/minimal");
    assertEquals(req.method, "PUT");
  }

  @Test
  @DisplayName("allow put method for a custom path with all parameters")
  void putTest1() {
    String path0 = "/test/all";
    Map<String, Object> parameters0 = new HashMap<>();
    {
      String query1 = "parameters";
      parameters0.put("query", query1);
    }
    Map<String, String> body0 = new HashMap<>();
    {
      String body1 = "parameters";
      body0.put("body", body1);
    }

    assertDoesNotThrow(() -> {
      client.put(path0, parameters0, body0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/test/all");
    assertEquals(req.method, "PUT");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"body\":\"parameters\"}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"query\":\"parameters\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }
}
