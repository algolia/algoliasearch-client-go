package com.algolia.client;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

import com.algolia.EchoInterceptor;
import com.algolia.EchoResponse;
import com.algolia.api.PersonalizationClient;
import com.algolia.model.personalization.*;
import com.algolia.utils.ClientOptions;
import com.algolia.utils.HttpRequester;
import java.util.*;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class PersonalizationClientClientTests {

  private HttpRequester requester;
  private EchoInterceptor echo;

  @BeforeAll
  void init() {
    requester = new HttpRequester();
    echo = new EchoInterceptor();
    requester.addInterceptor(echo.getEchoInterceptor());
  }

  PersonalizationClient createClient() {
    return new PersonalizationClient("appId", "apiKey", "us", ClientOptions.build().setRequester(requester));
  }

  @Test
  @DisplayName("calls api with correct user agent")
  void commonApiTest0() {
    PersonalizationClient $client = createClient();

    String path0 = "/test";

    $client.post(path0);
    EchoResponse result = echo.getLastResponse();

    {
      String regexp =
        "^Algolia for Java \\(\\d+\\.\\d+\\.\\d+(-.*)?\\)(; [a-zA-Z. ]+" +
        " (\\(\\d+\\.\\d+\\.\\d+(-.*)?\\))?)*(; Personalization" +
        " (\\(\\d+\\.\\d+\\.\\d+(-.*)?\\)))(; [a-zA-Z. ]+" +
        " (\\(\\d+\\.\\d+\\.\\d+(-.*)?\\))?)*$";
      assertTrue(
        result.headers.get("user-agent").matches(regexp),
        "Expected " + result.headers.get("user-agent") + " to match the following regex: " + regexp
      );
    }
  }

  @Test
  @DisplayName("calls api with correct timeouts")
  void commonApiTest1() {
    PersonalizationClient $client = createClient();

    String path0 = "/test";

    $client.post(path0);
    EchoResponse result = echo.getLastResponse();

    assertEquals(2000, result.connectTimeout);
    assertEquals(30000, result.responseTimeout);
  }

  @Test
  @DisplayName("throws when incorrect region is given")
  void parametersTest0() {
    {
      Exception exception = assertThrows(
        Exception.class,
        () -> {
          PersonalizationClient $client = new PersonalizationClient(
            "my-app-id",
            "my-api-key",
            "not_a_region",
            ClientOptions.build().setRequester(requester)
          );
        }
      );
      assertEquals("`region` must be one of the following: eu, us", exception.getMessage());
    }
  }

  @Test
  @DisplayName("does not throw when region is given")
  void parametersTest1() {
    PersonalizationClient $client = new PersonalizationClient(
      "my-app-id",
      "my-api-key",
      "us",
      ClientOptions.build().setRequester(requester)
    );
  }
}
