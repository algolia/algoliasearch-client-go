package com.algolia.methods.requests;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertEquals;

import com.algolia.EchoRequester;
import com.algolia.EchoResponse;
import com.algolia.api.AnalyticsClient;
import com.algolia.model.analytics.*;
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
class AnalyticsClientTests {

  private AnalyticsClient client;
  private EchoRequester requester;

  @BeforeAll
  void init() {
    requester = new EchoRequester();
    client = new AnalyticsClient("appId", "apiKey", requester);
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
  @DisplayName("get getAverageClickPosition with minimal parameters")
  void getAverageClickPositionTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getAverageClickPosition(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/clicks/averageClickPosition");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getAverageClickPosition with all parameters")
  void getAverageClickPositionTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getAverageClickPosition(index0, startDate0, endDate0, tags0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/clicks/averageClickPosition");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getClickPositions with minimal parameters")
  void getClickPositionsTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getClickPositions(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/clicks/positions");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getClickPositions with all parameters")
  void getClickPositionsTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getClickPositions(index0, startDate0, endDate0, tags0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/clicks/positions");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getClickThroughRate with minimal parameters")
  void getClickThroughRateTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getClickThroughRate(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/clicks/clickThroughRate");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getClickThroughRate with all parameters")
  void getClickThroughRateTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getClickThroughRate(index0, startDate0, endDate0, tags0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/clicks/clickThroughRate");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getConversationRate with minimal parameters")
  void getConversationRateTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getConversationRate(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/conversions/conversionRate");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getConversationRate with all parameters")
  void getConversationRateTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getConversationRate(index0, startDate0, endDate0, tags0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/conversions/conversionRate");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getNoClickRate with minimal parameters")
  void getNoClickRateTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getNoClickRate(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches/noClickRate");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getNoClickRate with all parameters")
  void getNoClickRateTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getNoClickRate(index0, startDate0, endDate0, tags0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches/noClickRate");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getNoResultsRate with minimal parameters")
  void getNoResultsRateTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getNoResultsRate(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches/noResultRate");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getNoResultsRate with all parameters")
  void getNoResultsRateTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getNoResultsRate(index0, startDate0, endDate0, tags0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches/noResultRate");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getSearchesCount with minimal parameters")
  void getSearchesCountTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getSearchesCount(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches/count");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getSearchesCount with all parameters")
  void getSearchesCountTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getSearchesCount(index0, startDate0, endDate0, tags0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches/count");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getSearchesNoClicks with minimal parameters")
  void getSearchesNoClicksTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getSearchesNoClicks(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches/noClicks");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getSearchesNoClicks with all parameters")
  void getSearchesNoClicksTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    int limit0 = 21;
    int offset0 = 42;
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getSearchesNoClicks(
        index0,
        startDate0,
        endDate0,
        limit0,
        offset0,
        tags0
      );
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches/noClicks");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"limit\":\"21\",\"offset\":\"42\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getSearchesNoResults with minimal parameters")
  void getSearchesNoResultsTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getSearchesNoResults(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches/noResults");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getSearchesNoResults with all parameters")
  void getSearchesNoResultsTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    int limit0 = 21;
    int offset0 = 42;
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getSearchesNoResults(
        index0,
        startDate0,
        endDate0,
        limit0,
        offset0,
        tags0
      );
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches/noResults");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"limit\":\"21\",\"offset\":\"42\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getStatus with minimal parameters")
  void getStatusTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getStatus(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/status");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopCountries with minimal parameters")
  void getTopCountriesTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getTopCountries(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/countries");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopCountries with all parameters")
  void getTopCountriesTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    int limit0 = 21;
    int offset0 = 42;
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getTopCountries(
        index0,
        startDate0,
        endDate0,
        limit0,
        offset0,
        tags0
      );
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/countries");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"limit\":\"21\",\"offset\":\"42\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopFilterAttributes with minimal parameters")
  void getTopFilterAttributesTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getTopFilterAttributes(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/filters");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopFilterAttributes with all parameters")
  void getTopFilterAttributesTest1() {
    String index0 = "index";
    String search0 = "mySearch";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    int limit0 = 21;
    int offset0 = 42;
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getTopFilterAttributes(
        index0,
        search0,
        startDate0,
        endDate0,
        limit0,
        offset0,
        tags0
      );
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/filters");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"search\":\"mySearch\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"limit\":\"21\",\"offset\":\"42\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopFilterForAttribute with minimal parameters")
  void getTopFilterForAttributeTest0() {
    String attribute0 = "myAttribute";
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getTopFilterForAttribute(attribute0, index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/filters/myAttribute");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
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
    "get getTopFilterForAttribute with minimal parameters and multiple attributes"
  )
  void getTopFilterForAttributeTest1() {
    String attribute0 = "myAttribute1,myAttribute2";
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getTopFilterForAttribute(attribute0, index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/filters/myAttribute1%2CmyAttribute2");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopFilterForAttribute with all parameters")
  void getTopFilterForAttributeTest2() {
    String attribute0 = "myAttribute";
    String index0 = "index";
    String search0 = "mySearch";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    int limit0 = 21;
    int offset0 = 42;
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getTopFilterForAttribute(
        attribute0,
        index0,
        search0,
        startDate0,
        endDate0,
        limit0,
        offset0,
        tags0
      );
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/filters/myAttribute");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"search\":\"mySearch\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"limit\":\"21\",\"offset\":\"42\",\"tags\":\"tag\"}",
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
    "get getTopFilterForAttribute with all parameters and multiple attributes"
  )
  void getTopFilterForAttributeTest3() {
    String attribute0 = "myAttribute1,myAttribute2";
    String index0 = "index";
    String search0 = "mySearch";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    int limit0 = 21;
    int offset0 = 42;
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getTopFilterForAttribute(
        attribute0,
        index0,
        search0,
        startDate0,
        endDate0,
        limit0,
        offset0,
        tags0
      );
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/filters/myAttribute1%2CmyAttribute2");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"search\":\"mySearch\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"limit\":\"21\",\"offset\":\"42\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopFiltersNoResults with minimal parameters")
  void getTopFiltersNoResultsTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getTopFiltersNoResults(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/filters/noResults");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopFiltersNoResults with all parameters")
  void getTopFiltersNoResultsTest1() {
    String index0 = "index";
    String search0 = "mySearch";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    int limit0 = 21;
    int offset0 = 42;
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getTopFiltersNoResults(
        index0,
        search0,
        startDate0,
        endDate0,
        limit0,
        offset0,
        tags0
      );
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/filters/noResults");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"search\":\"mySearch\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"limit\":\"21\",\"offset\":\"42\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopHits with minimal parameters")
  void getTopHitsTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getTopHits(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/hits");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopHits with all parameters")
  void getTopHitsTest1() {
    String index0 = "index";
    String search0 = "mySearch";
    boolean clickAnalytics0 = true;
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    int limit0 = 21;
    int offset0 = 42;
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getTopHits(
        index0,
        search0,
        clickAnalytics0,
        startDate0,
        endDate0,
        limit0,
        offset0,
        tags0
      );
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/hits");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"search\":\"mySearch\",\"clickAnalytics\":\"true\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"limit\":\"21\",\"offset\":\"42\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopSearches with minimal parameters")
  void getTopSearchesTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getTopSearches(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getTopSearches with all parameters")
  void getTopSearchesTest1() {
    String index0 = "index";
    boolean clickAnalytics0 = true;
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    OrderBy orderBy0 = OrderBy.fromValue("searchCount");
    Direction direction0 = Direction.fromValue("asc");
    int limit0 = 21;
    int offset0 = 42;
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getTopSearches(
        index0,
        clickAnalytics0,
        startDate0,
        endDate0,
        orderBy0,
        direction0,
        limit0,
        offset0,
        tags0
      );
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/searches");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"clickAnalytics\":\"true\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"orderBy\":\"searchCount\",\"direction\":\"asc\",\"limit\":\"21\",\"offset\":\"42\",\"tags\":\"tag\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getUsersCount with minimal parameters")
  void getUsersCountTest0() {
    String index0 = "index";

    assertDoesNotThrow(() -> {
      client.getUsersCount(index0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/users/count");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\"}",
      new TypeToken<HashMap<String, String>>() {}.getType()
    );
    Map<String, Object> actualQuery = req.queryParameters;

    assertEquals(expectedQuery.size(), actualQuery.size());
    for (Map.Entry<String, Object> p : actualQuery.entrySet()) {
      assertEquals(expectedQuery.get(p.getKey()), p.getValue());
    }
  }

  @Test
  @DisplayName("get getUsersCount with all parameters")
  void getUsersCountTest1() {
    String index0 = "index";
    String startDate0 = "1999-09-19";
    String endDate0 = "2001-01-01";
    String tags0 = "tag";

    assertDoesNotThrow(() -> {
      client.getUsersCount(index0, startDate0, endDate0, tags0);
    });
    EchoResponse req = requester.getLastEchoResponse();

    assertEquals(req.path, "/2/users/count");
    assertEquals(req.method, "GET");

    Map<String, String> expectedQuery = JSON.deserialize(
      "{\"index\":\"index\",\"startDate\":\"1999-09-19\",\"endDate\":\"2001-01-01\",\"tags\":\"tag\"}",
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
}
