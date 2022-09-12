package com.algolia.methods.requests;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.junit.jupiter.api.Assertions.fail;

import com.algolia.EchoInterceptor;
import com.algolia.EchoResponse;
import com.algolia.api.PredictClient;
import com.algolia.model.predict.*;
import com.algolia.utils.ClientOptions;
import com.algolia.utils.HttpRequester;
import com.algolia.utils.JSONBuilder;
import com.algolia.utils.RequestOptions;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.util.*;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;
import org.skyscreamer.jsonassert.JSONAssert;
import org.skyscreamer.jsonassert.JSONCompareMode;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class PredictClientRequestsTests {

  private PredictClient client;
  private EchoInterceptor echo;
  private ObjectMapper json;

  @BeforeAll
  void init() {
    json = new JSONBuilder().failOnUnknown(true).build();
    HttpRequester requester = new HttpRequester();
    echo = new EchoInterceptor();
    requester.addInterceptor(echo.getEchoInterceptor());
    client = new PredictClient("appId", "apiKey", "eu", new ClientOptions().setRequester(requester));
  }

  @Test
  @DisplayName("allow del method for a custom path with minimal parameters")
  void delTest0() {
    String path0 = "/test/minimal";

    assertDoesNotThrow(() -> {
      client.del(path0);
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/minimal");
    assertEquals(req.method, "DELETE");
    assertNull(req.body);
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
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/all");
    assertEquals(req.method, "DELETE");
    assertNull(req.body);

    try {
      Map<String, String> expectedQuery = json.readValue("{\"query\":\"parameters\"}", new TypeReference<HashMap<String, String>>() {});
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
    }
  }

  @Test
  @DisplayName("deleteUserProfile")
  void deleteUserProfileTest0() {
    String userID0 = "user1";

    assertDoesNotThrow(() -> {
      client.deleteUserProfile(userID0);
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/users/user1");
    assertEquals(req.method, "DELETE");
    assertNull(req.body);
  }

  @Test
  @DisplayName("fetchAllUserProfiles with minimal parameters for modelsToRetrieve")
  void fetchAllUserProfilesTest0() {
    ModelsToRetrieveParam fetchAllUserProfilesParams0 = new ModelsToRetrieveParam();
    {
      List<ModelsToRetrieve> modelsToRetrieve1 = new ArrayList<>();
      {
        ModelsToRetrieve modelsToRetrieve_02 = ModelsToRetrieve.fromValue("funnel_stage");
        modelsToRetrieve1.add(modelsToRetrieve_02);
        ModelsToRetrieve modelsToRetrieve_12 = ModelsToRetrieve.fromValue("order_value");
        modelsToRetrieve1.add(modelsToRetrieve_12);
        ModelsToRetrieve modelsToRetrieve_22 = ModelsToRetrieve.fromValue("affinities");
        modelsToRetrieve1.add(modelsToRetrieve_22);
      }
      fetchAllUserProfilesParams0.setModelsToRetrieve(modelsToRetrieve1);
    }

    assertDoesNotThrow(() -> {
      client.fetchAllUserProfiles(FetchAllUserProfilesParams.of(fetchAllUserProfilesParams0));
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/users");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"modelsToRetrieve\":[\"funnel_stage\",\"order_value\",\"affinities\"]}", req.body, JSONCompareMode.STRICT);
    });
  }

  @Test
  @DisplayName("fetchAllUserProfiles with minimal parameters for typesToRetrieve")
  void fetchAllUserProfilesTest1() {
    TypesToRetrieveParam fetchAllUserProfilesParams0 = new TypesToRetrieveParam();
    {
      List<TypesToRetrieve> typesToRetrieve1 = new ArrayList<>();
      {
        TypesToRetrieve typesToRetrieve_02 = TypesToRetrieve.fromValue("properties");
        typesToRetrieve1.add(typesToRetrieve_02);
        TypesToRetrieve typesToRetrieve_12 = TypesToRetrieve.fromValue("segments");
        typesToRetrieve1.add(typesToRetrieve_12);
      }
      fetchAllUserProfilesParams0.setTypesToRetrieve(typesToRetrieve1);
    }

    assertDoesNotThrow(() -> {
      client.fetchAllUserProfiles(FetchAllUserProfilesParams.of(fetchAllUserProfilesParams0));
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/users");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"typesToRetrieve\":[\"properties\",\"segments\"]}", req.body, JSONCompareMode.STRICT);
    });
  }

  @Test
  @DisplayName("fetchAllUserProfiles with a limit")
  void fetchAllUserProfilesTest2() {
    LimitParam fetchAllUserProfilesParams0 = new LimitParam();
    {
      int limit1 = 10;
      fetchAllUserProfilesParams0.setLimit(limit1);
    }

    assertDoesNotThrow(() -> {
      client.fetchAllUserProfiles(FetchAllUserProfilesParams.of(fetchAllUserProfilesParams0));
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/users");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"limit\":10}", req.body, JSONCompareMode.STRICT);
    });
  }

  @Test
  @DisplayName("fetchAllUserProfiles with a nextPageToken")
  void fetchAllUserProfilesTest3() {
    NextPageTokenParam fetchAllUserProfilesParams0 = new NextPageTokenParam();
    {
      String nextPageToken1 = "nextPageTokenExample123";
      fetchAllUserProfilesParams0.setNextPageToken(nextPageToken1);
    }

    assertDoesNotThrow(() -> {
      client.fetchAllUserProfiles(FetchAllUserProfilesParams.of(fetchAllUserProfilesParams0));
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/users");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"nextPageToken\":\"nextPageTokenExample123\"}", req.body, JSONCompareMode.STRICT);
    });
  }

  @Test
  @DisplayName("fetchAllUserProfiles with a previousPageToken")
  void fetchAllUserProfilesTest4() {
    PreviousPageTokenParam fetchAllUserProfilesParams0 = new PreviousPageTokenParam();
    {
      String previousPageToken1 = "previousPageTokenExample123";
      fetchAllUserProfilesParams0.setPreviousPageToken(previousPageToken1);
    }

    assertDoesNotThrow(() -> {
      client.fetchAllUserProfiles(FetchAllUserProfilesParams.of(fetchAllUserProfilesParams0));
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/users");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"previousPageToken\":\"previousPageTokenExample123\"}", req.body, JSONCompareMode.STRICT);
    });
  }

  @Test
  @DisplayName("fetchUserProfile with minimal parameters for modelsToRetrieve")
  void fetchUserProfileTest0() {
    String userID0 = "user1";
    ModelsToRetrieveParam params0 = new ModelsToRetrieveParam();
    {
      List<ModelsToRetrieve> modelsToRetrieve1 = new ArrayList<>();
      {
        ModelsToRetrieve modelsToRetrieve_02 = ModelsToRetrieve.fromValue("funnel_stage");
        modelsToRetrieve1.add(modelsToRetrieve_02);
        ModelsToRetrieve modelsToRetrieve_12 = ModelsToRetrieve.fromValue("order_value");
        modelsToRetrieve1.add(modelsToRetrieve_12);
        ModelsToRetrieve modelsToRetrieve_22 = ModelsToRetrieve.fromValue("affinities");
        modelsToRetrieve1.add(modelsToRetrieve_22);
      }
      params0.setModelsToRetrieve(modelsToRetrieve1);
    }

    assertDoesNotThrow(() -> {
      client.fetchUserProfile(userID0, Params.of(params0));
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/users/user1/fetch");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"modelsToRetrieve\":[\"funnel_stage\",\"order_value\",\"affinities\"]}", req.body, JSONCompareMode.STRICT);
    });
  }

  @Test
  @DisplayName("fetchUserProfile with minimal parameters for typesToRetrieve")
  void fetchUserProfileTest1() {
    String userID0 = "user1";
    TypesToRetrieveParam params0 = new TypesToRetrieveParam();
    {
      List<TypesToRetrieve> typesToRetrieve1 = new ArrayList<>();
      {
        TypesToRetrieve typesToRetrieve_02 = TypesToRetrieve.fromValue("properties");
        typesToRetrieve1.add(typesToRetrieve_02);
        TypesToRetrieve typesToRetrieve_12 = TypesToRetrieve.fromValue("segments");
        typesToRetrieve1.add(typesToRetrieve_12);
      }
      params0.setTypesToRetrieve(typesToRetrieve1);
    }

    assertDoesNotThrow(() -> {
      client.fetchUserProfile(userID0, Params.of(params0));
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/users/user1/fetch");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"typesToRetrieve\":[\"properties\",\"segments\"]}", req.body, JSONCompareMode.STRICT);
    });
  }

  @Test
  @DisplayName("fetchUserProfile with all parameters")
  void fetchUserProfileTest2() {
    String userID0 = "user1";
    AllParams params0 = new AllParams();
    {
      List<ModelsToRetrieve> modelsToRetrieve1 = new ArrayList<>();
      {
        ModelsToRetrieve modelsToRetrieve_02 = ModelsToRetrieve.fromValue("funnel_stage");
        modelsToRetrieve1.add(modelsToRetrieve_02);
        ModelsToRetrieve modelsToRetrieve_12 = ModelsToRetrieve.fromValue("order_value");
        modelsToRetrieve1.add(modelsToRetrieve_12);
        ModelsToRetrieve modelsToRetrieve_22 = ModelsToRetrieve.fromValue("affinities");
        modelsToRetrieve1.add(modelsToRetrieve_22);
      }
      params0.setModelsToRetrieve(modelsToRetrieve1);
      List<TypesToRetrieve> typesToRetrieve1 = new ArrayList<>();
      {
        TypesToRetrieve typesToRetrieve_02 = TypesToRetrieve.fromValue("properties");
        typesToRetrieve1.add(typesToRetrieve_02);
        TypesToRetrieve typesToRetrieve_12 = TypesToRetrieve.fromValue("segments");
        typesToRetrieve1.add(typesToRetrieve_12);
      }
      params0.setTypesToRetrieve(typesToRetrieve1);
    }

    assertDoesNotThrow(() -> {
      client.fetchUserProfile(userID0, Params.of(params0));
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/users/user1/fetch");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"modelsToRetrieve\":[\"funnel_stage\",\"order_value\",\"affinities\"],\"typesToRetrieve\":[\"properties\",\"segments\"]}",
        req.body,
        JSONCompareMode.STRICT
      );
    });
  }

  @Test
  @DisplayName("allow get method for a custom path with minimal parameters")
  void getTest0() {
    String path0 = "/test/minimal";

    assertDoesNotThrow(() -> {
      client.get(path0);
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/minimal");
    assertEquals(req.method, "GET");
    assertNull(req.body);
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
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/all");
    assertEquals(req.method, "GET");
    assertNull(req.body);

    try {
      Map<String, String> expectedQuery = json.readValue("{\"query\":\"parameters\"}", new TypeReference<HashMap<String, String>>() {});
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
    }
  }

  @Test
  @DisplayName("allow post method for a custom path with minimal parameters")
  void postTest0() {
    String path0 = "/test/minimal";

    assertDoesNotThrow(() -> {
      client.post(path0);
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/minimal");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{}", req.body, JSONCompareMode.STRICT);
    });
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
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/all");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"body\":\"parameters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue("{\"query\":\"parameters\"}", new TypeReference<HashMap<String, String>>() {});
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
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
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"facet\":\"filters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue(
        "{\"query\":\"myQueryParameter\"}",
        new TypeReference<HashMap<String, String>>() {}
      );
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
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
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"facet\":\"filters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue(
        "{\"query\":\"parameters\",\"query2\":\"myQueryParameter\"}",
        new TypeReference<HashMap<String, String>>() {}
      );
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
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
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"facet\":\"filters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue("{\"query\":\"parameters\"}", new TypeReference<HashMap<String, String>>() {});
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
    }

    try {
      Map<String, String> expectedHeaders = json.readValue(
        "{\"x-algolia-api-key\":\"myApiKey\"}",
        new TypeReference<HashMap<String, String>>() {}
      );
      Map<String, String> actualHeaders = req.headers;

      for (Map.Entry<String, String> p : expectedHeaders.entrySet()) {
        assertEquals(actualHeaders.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse headers json");
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
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"facet\":\"filters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue("{\"query\":\"parameters\"}", new TypeReference<HashMap<String, String>>() {});
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
    }

    try {
      Map<String, String> expectedHeaders = json.readValue(
        "{\"x-algolia-api-key\":\"myApiKey\"}",
        new TypeReference<HashMap<String, String>>() {}
      );
      Map<String, String> actualHeaders = req.headers;

      for (Map.Entry<String, String> p : expectedHeaders.entrySet()) {
        assertEquals(actualHeaders.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse headers json");
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
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"facet\":\"filters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue(
        "{\"query\":\"parameters\",\"isItWorking\":\"true\"}",
        new TypeReference<HashMap<String, String>>() {}
      );
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
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
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"facet\":\"filters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue(
        "{\"query\":\"parameters\",\"myParam\":\"2\"}",
        new TypeReference<HashMap<String, String>>() {}
      );
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
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
    requestOptions.addExtraQueryParameters("myParam", Arrays.asList("c", "d"));

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"facet\":\"filters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue(
        "{\"query\":\"parameters\",\"myParam\":\"c,d\"}",
        new TypeReference<HashMap<String, String>>() {}
      );
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
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
    requestOptions.addExtraQueryParameters("myParam", Arrays.asList(true, true, false));

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"facet\":\"filters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue(
        "{\"query\":\"parameters\",\"myParam\":\"true,true,false\"}",
        new TypeReference<HashMap<String, String>>() {}
      );
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
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
    requestOptions.addExtraQueryParameters("myParam", Arrays.asList(1, 2));

    assertDoesNotThrow(() -> {
      client.post(path0, parameters0, body0, requestOptions);
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/requestOptions");
    assertEquals(req.method, "POST");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"facet\":\"filters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue(
        "{\"query\":\"parameters\",\"myParam\":\"1,2\"}",
        new TypeReference<HashMap<String, String>>() {}
      );
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
    }
  }

  @Test
  @DisplayName("allow put method for a custom path with minimal parameters")
  void putTest0() {
    String path0 = "/test/minimal";

    assertDoesNotThrow(() -> {
      client.put(path0);
    });
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/minimal");
    assertEquals(req.method, "PUT");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{}", req.body, JSONCompareMode.STRICT);
    });
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
    EchoResponse req = echo.getLastResponse();

    assertEquals(req.path, "/1/test/all");
    assertEquals(req.method, "PUT");
    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals("{\"body\":\"parameters\"}", req.body, JSONCompareMode.STRICT);
    });

    try {
      Map<String, String> expectedQuery = json.readValue("{\"query\":\"parameters\"}", new TypeReference<HashMap<String, String>>() {});
      Map<String, Object> actualQuery = req.queryParameters;

      assertEquals(expectedQuery.size(), actualQuery.size());
      for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
        assertEquals(expectedQuery.get(p.getKey()), p.getValue());
      }
    } catch (JsonProcessingException e) {
      fail("failed to parse queryParameters json");
    }
  }
}
