package com.algolia.client;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import com.algolia.EchoInterceptor;
import com.algolia.EchoResponse;
import com.algolia.api.PredictClient;
import com.algolia.model.predict.*;
import com.algolia.utils.ClientOptions;
import com.algolia.utils.HttpRequester;
import java.util.*;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class PredictClientClientTests {

  private HttpRequester requester;
  private EchoInterceptor echo;

  @BeforeAll
  void init() {
    requester = new HttpRequester();
    echo = new EchoInterceptor();
    requester.addInterceptor(echo.getEchoInterceptor());
  }

  PredictClient createClient() {
    return new PredictClient("appId", "apiKey", "ew", new ClientOptions().setRequester(requester));
  }

  @Test
  @DisplayName("calls api with default read timeouts")
  void commonApiTest0() {
    PredictClient $client = createClient();

    String path0 = "/test";

    $client.get(path0);
    EchoResponse result = echo.getLastResponse();

    assertEquals(2000, result.connectTimeout);
    assertEquals(5000, result.responseTimeout);
  }

  @Test
  @DisplayName("calls api with default write timeouts")
  void commonApiTest1() {
    PredictClient $client = createClient();

    String path0 = "/test";

    $client.post(path0);
    EchoResponse result = echo.getLastResponse();

    assertEquals(2000, result.connectTimeout);
    assertEquals(30000, result.responseTimeout);
  }

  @Test
  @DisplayName("throws when region is not given")
  void parametersTest0() {
    {
      Exception exception = assertThrows(
        Exception.class,
        () -> {
          PredictClient $client = new PredictClient("my-app-id", "my-api-key", "", new ClientOptions().setRequester(requester));
        }
      );
      assertEquals("`region` is required and must be one of the following: ue, ew", exception.getMessage());
    }
  }

  @Test
  @DisplayName("throws when incorrect region is given")
  void parametersTest1() {
    {
      Exception exception = assertThrows(
        Exception.class,
        () -> {
          PredictClient $client = new PredictClient("my-app-id", "my-api-key", "not_a_region", new ClientOptions().setRequester(requester));
        }
      );
      assertEquals("`region` is required and must be one of the following: ue, ew", exception.getMessage());
    }
  }

  @Test
  @DisplayName("does not throw when region is given")
  void parametersTest2() {
    PredictClient $client = new PredictClient("my-app-id", "my-api-key", "ew", new ClientOptions().setRequester(requester));
  }
}
