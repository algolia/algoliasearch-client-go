package com.algolia.methods.requests;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertEquals;

import com.algolia.EchoRequester;
import com.algolia.EchoResponse;
import com.algolia.api.QuerySuggestionsClient;
import com.algolia.model.querySuggestions.*;
import com.algolia.utils.JSON;
import com.google.gson.reflect.TypeToken;
import java.util.*;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;
import org.skyscreamer.jsonassert.JSONAssert;
import org.skyscreamer.jsonassert.JSONCompareMode;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class QuerySuggestionsClientTests {

  private QuerySuggestionsClient client;
  private EchoRequester requester;

  @BeforeAll
  void init() {
    requester = new EchoRequester();
    client = new QuerySuggestionsClient("appId", "apiKey", requester);
  }

  @Test
  @DisplayName("createConfig")
  void createConfigTest0() {
    QuerySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam0 = new QuerySuggestionsIndexWithIndexParam();
    {
      String indexName1 = "theIndexName";
      querySuggestionsIndexWithIndexParam0.setIndexName(indexName1);
      List<SourceIndex> sourceIndices1 = new ArrayList<>();
      {
        SourceIndex sourceIndices_02 = new SourceIndex();
        {
          String indexName3 = "testIndex";
          sourceIndices_02.setIndexName(indexName3);
          List<Object> facets3 = new ArrayList<>();
          {
            Map<String, String> facets_04 = new HashMap<>();
            {
              String attributes5 = "test";
              facets_04.put("attributes", attributes5);
            }
            facets3.add(facets_04);
          }
          sourceIndices_02.setFacets(facets3);
          List<List<String>> generate3 = new ArrayList<>();
          {
            List<String> generate_04 = new ArrayList<>();
            {
              String generate_0_05 = "facetA";
              generate_04.add(generate_0_05);
              String generate_0_15 = "facetB";
              generate_04.add(generate_0_15);
            }
            generate3.add(generate_04);
            List<String> generate_14 = new ArrayList<>();
            {
              String generate_1_05 = "facetC";
              generate_14.add(generate_1_05);
            }
            generate3.add(generate_14);
          }
          sourceIndices_02.setGenerate(generate3);
        }
        sourceIndices1.add(sourceIndices_02);
      }
      querySuggestionsIndexWithIndexParam0.setSourceIndices(sourceIndices1);
      List<String> languages1 = new ArrayList<>();
      {
        String languages_02 = "french";
        languages1.add(languages_02);
      }
      querySuggestionsIndexWithIndexParam0.setLanguages(languages1);
      List<String> exclude1 = new ArrayList<>();
      {
        String exclude_02 = "test";
        exclude1.add(exclude_02);
      }
      querySuggestionsIndexWithIndexParam0.setExclude(exclude1);
    }

    assertDoesNotThrow(() -> {
      client.createConfig(querySuggestionsIndexWithIndexParam0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/configs");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"indexName\":\"theIndexName\",\"sourceIndices\":[{\"indexName\":\"testIndex\",\"facets\":[{\"attributes\":\"test\"}],\"generate\":[[\"facetA\",\"facetB\"],[\"facetC\"]]}],\"languages\":[\"french\"],\"exclude\":[\"test\"]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
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
    Map<String, String> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, String> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("deleteConfig")
  void deleteConfigTest0() {
    String indexName0 = "theIndexName";

    assertDoesNotThrow(() -> {
      client.deleteConfig(indexName0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/configs/theIndexName");
    assertEquals(req.method, "DELETE");
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
    Map<String, String> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, String> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("getAllConfigs")
  void getAllConfigsTest0() {
    assertDoesNotThrow(() -> {
      client.getAllConfigs();
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/configs");
    assertEquals(req.method, "GET");
  }

  @Test
  @DisplayName("getConfig")
  void getConfigTest0() {
    String indexName0 = "theIndexName";

    assertDoesNotThrow(() -> {
      client.getConfig(indexName0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/configs/theIndexName");
    assertEquals(req.method, "GET");
  }

  @Test
  @DisplayName("getConfigStatus")
  void getConfigStatusTest0() {
    String indexName0 = "theIndexName";

    assertDoesNotThrow(() -> {
      client.getConfigStatus(indexName0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/configs/theIndexName/status");
    assertEquals(req.method, "GET");
  }

  @Test
  @DisplayName("getLogFile")
  void getLogFileTest0() {
    String indexName0 = "theIndexName";

    assertDoesNotThrow(() -> {
      client.getLogFile(indexName0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/logs/theIndexName");
    assertEquals(req.method, "GET");
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
    Map<String, String> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, String> p : actualQuery.entrySet()) {
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
    Map<String, String> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, String> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("updateConfig")
  void updateConfigTest0() {
    String indexName0 = "theIndexName";
    QuerySuggestionsIndexParam querySuggestionsIndexParam0 = new QuerySuggestionsIndexParam();
    {
      List<SourceIndex> sourceIndices1 = new ArrayList<>();
      {
        SourceIndex sourceIndices_02 = new SourceIndex();
        {
          String indexName3 = "testIndex";
          sourceIndices_02.setIndexName(indexName3);
          List<Object> facets3 = new ArrayList<>();
          {
            Map<String, String> facets_04 = new HashMap<>();
            {
              String attributes5 = "test";
              facets_04.put("attributes", attributes5);
            }
            facets3.add(facets_04);
          }
          sourceIndices_02.setFacets(facets3);
          List<List<String>> generate3 = new ArrayList<>();
          {
            List<String> generate_04 = new ArrayList<>();
            {
              String generate_0_05 = "facetA";
              generate_04.add(generate_0_05);
              String generate_0_15 = "facetB";
              generate_04.add(generate_0_15);
            }
            generate3.add(generate_04);
            List<String> generate_14 = new ArrayList<>();
            {
              String generate_1_05 = "facetC";
              generate_14.add(generate_1_05);
            }
            generate3.add(generate_14);
          }
          sourceIndices_02.setGenerate(generate3);
        }
        sourceIndices1.add(sourceIndices_02);
      }
      querySuggestionsIndexParam0.setSourceIndices(sourceIndices1);
      List<String> languages1 = new ArrayList<>();
      {
        String languages_02 = "french";
        languages1.add(languages_02);
      }
      querySuggestionsIndexParam0.setLanguages(languages1);
      List<String> exclude1 = new ArrayList<>();
      {
        String exclude_02 = "test";
        exclude1.add(exclude_02);
      }
      querySuggestionsIndexParam0.setExclude(exclude1);
    }

    assertDoesNotThrow(() -> {
      client.updateConfig(indexName0, querySuggestionsIndexParam0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/configs/theIndexName");
    assertEquals(req.method, "PUT");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"sourceIndices\":[{\"indexName\":\"testIndex\",\"facets\":[{\"attributes\":\"test\"}],\"generate\":[[\"facetA\",\"facetB\"],[\"facetC\"]]}],\"languages\":[\"french\"],\"exclude\":[\"test\"]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
  }
}
