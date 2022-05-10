package com.algolia.methods.requests;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertEquals;

import com.algolia.EchoRequester;
import com.algolia.EchoResponse;
import com.algolia.api.PredictClient;
import com.algolia.model.predict.*;
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
class PredictClientTests {

  private PredictClient client;
  private EchoRequester requester;

  @BeforeAll
  void init() {
    requester = new EchoRequester();
    client = new PredictClient("appId", "apiKey", requester);
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
  @DisplayName("fetchUserProfile with minimal parameters for modelsToRetrieve")
  void fetchUserProfileTest0() {
    String userID0 = "user1";
    ModelsToRetrieve params0 = new ModelsToRetrieve();
    {
      List<ModelsToRetrieveEnum> modelsToRetrieve1 = new ArrayList<>();
      {
        ModelsToRetrieveEnum modelsToRetrieve_02 = ModelsToRetrieveEnum.fromValue(
          "funnel_stage"
        );
        modelsToRetrieve1.add(modelsToRetrieve_02);
        ModelsToRetrieveEnum modelsToRetrieve_12 = ModelsToRetrieveEnum.fromValue(
          "order_value"
        );
        modelsToRetrieve1.add(modelsToRetrieve_12);
        ModelsToRetrieveEnum modelsToRetrieve_22 = ModelsToRetrieveEnum.fromValue(
          "affinities"
        );
        modelsToRetrieve1.add(modelsToRetrieve_22);
      }
      params0.setModelsToRetrieve(modelsToRetrieve1);
    }

    assertDoesNotThrow(() -> {
      client.fetchUserProfile(userID0, Params.ofModelsToRetrieve(params0));
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/users/user1/fetch");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"modelsToRetrieve\":[\"funnel_stage\",\"order_value\",\"affinities\"]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
  }

  @Test
  @DisplayName("fetchUserProfile with minimal parameters for typesToRetrieve")
  void fetchUserProfileTest1() {
    String userID0 = "user1";
    TypesToRetrieve params0 = new TypesToRetrieve();
    {
      List<TypesToRetrieveEnum> typesToRetrieve1 = new ArrayList<>();
      {
        TypesToRetrieveEnum typesToRetrieve_02 = TypesToRetrieveEnum.fromValue(
          "properties"
        );
        typesToRetrieve1.add(typesToRetrieve_02);
        TypesToRetrieveEnum typesToRetrieve_12 = TypesToRetrieveEnum.fromValue(
          "segments"
        );
        typesToRetrieve1.add(typesToRetrieve_12);
      }
      params0.setTypesToRetrieve(typesToRetrieve1);
    }

    assertDoesNotThrow(() -> {
      client.fetchUserProfile(userID0, Params.ofTypesToRetrieve(params0));
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/users/user1/fetch");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"typesToRetrieve\":[\"properties\",\"segments\"]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
      );
    });
  }

  @Test
  @DisplayName("fetchUserProfile with all parameters")
  void fetchUserProfileTest2() {
    String userID0 = "user1";
    AllParams params0 = new AllParams();
    {
      List<ModelsToRetrieveEnum> modelsToRetrieve1 = new ArrayList<>();
      {
        ModelsToRetrieveEnum modelsToRetrieve_02 = ModelsToRetrieveEnum.fromValue(
          "funnel_stage"
        );
        modelsToRetrieve1.add(modelsToRetrieve_02);
        ModelsToRetrieveEnum modelsToRetrieve_12 = ModelsToRetrieveEnum.fromValue(
          "order_value"
        );
        modelsToRetrieve1.add(modelsToRetrieve_12);
        ModelsToRetrieveEnum modelsToRetrieve_22 = ModelsToRetrieveEnum.fromValue(
          "affinities"
        );
        modelsToRetrieve1.add(modelsToRetrieve_22);
      }
      params0.setModelsToRetrieve(modelsToRetrieve1);
      List<TypesToRetrieveEnum> typesToRetrieve1 = new ArrayList<>();
      {
        TypesToRetrieveEnum typesToRetrieve_02 = TypesToRetrieveEnum.fromValue(
          "properties"
        );
        typesToRetrieve1.add(typesToRetrieve_02);
        TypesToRetrieveEnum typesToRetrieve_12 = TypesToRetrieveEnum.fromValue(
          "segments"
        );
        typesToRetrieve1.add(typesToRetrieve_12);
      }
      params0.setTypesToRetrieve(typesToRetrieve1);
    }

    assertDoesNotThrow(() -> {
      client.fetchUserProfile(userID0, Params.ofAllParams(params0));
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/1/users/user1/fetch");
    assertEquals(req.method, "POST");

    assertDoesNotThrow(() -> {
      JSONAssert.assertEquals(
        "{\"modelsToRetrieve\":[\"funnel_stage\",\"order_value\",\"affinities\"],\"typesToRetrieve\":[\"properties\",\"segments\"]}",
        req.body,
        JSONCompareMode.STRICT_ORDER
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
}
