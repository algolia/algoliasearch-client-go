package com.algolia.methods.requests;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertEquals;

import com.algolia.EchoRequester;
import com.algolia.EchoResponse;
import com.algolia.api.AbtestingClient;
import com.algolia.model.abtesting.*;
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
class AbtestingClientTests {

  private AbtestingClient client;
  private EchoRequester requester;

  @BeforeAll
  void init() {
    requester = new EchoRequester();
    client = new AbtestingClient("appId", "apiKey", requester);
  }

  @Test
  @DisplayName("addABTests with minimal parameters")
  void addABTestsTest0() {
    AddABTestsRequest addABTestsRequest0 = new AddABTestsRequest();
    {
      String endAt1 = "2022-12-31T00:00:00.000Z";
      addABTestsRequest0.setEndAt(endAt1);
      String name1 = "myABTest";
      addABTestsRequest0.setName(name1);
      List<AddABTestsVariant> variant1 = new ArrayList<>();
      {
        AbTestsVariant variant_02 = new AbTestsVariant();
        {
          String index3 = "AB_TEST_1";
          variant_02.setIndex(index3);
          int trafficPercentage3 = 30;
          variant_02.setTrafficPercentage(trafficPercentage3);
        }
        variant1.add(AddABTestsVariant.ofAbTestsVariant(variant_02));
        AbTestsVariant variant_12 = new AbTestsVariant();
        {
          String index3 = "AB_TEST_2";
          variant_12.setIndex(index3);
          int trafficPercentage3 = 50;
          variant_12.setTrafficPercentage(trafficPercentage3);
        }
        variant1.add(AddABTestsVariant.ofAbTestsVariant(variant_12));
      }
      addABTestsRequest0.setVariant(variant1);
    }

    assertDoesNotThrow(() -> {
      client.addABTests(addABTestsRequest0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/abtests");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"endAt\":\"2022-12-31T00:00:00.000Z\",\"name\":\"myABTest\",\"variant\":[{\"index\":\"AB_TEST_1\",\"trafficPercentage\":30},{\"index\":\"AB_TEST_2\",\"trafficPercentage\":50}]}",
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
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("deleteABTest")
  void deleteABTestTest0() {
    int id0 = 42;

    assertDoesNotThrow(() -> {
      client.deleteABTest(id0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/abtests/42");
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
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("getABTest")
  void getABTestTest0() {
    int id0 = 42;

    assertDoesNotThrow(() -> {
      client.getABTest(id0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/abtests/42");
    assertEquals(req.method, "GET");
  }

  @Test
  @DisplayName("listABTests with minimal parameters")
  void listABTestsTest0() {
    int offset0 = 42;
    int limit0 = 21;

    assertDoesNotThrow(() -> {
      client.listABTests(offset0, limit0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/abtests");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"offset\":\"42\",\"limit\":\"21\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
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

  @Test
  @DisplayName("stopABTest")
  void stopABTestTest0() {
    int id0 = 42;

    assertDoesNotThrow(() -> {
      client.stopABTest(id0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/abtests/42/stop");
    assertEquals(req.method, "POST");
  }
}
