package com.algolia.playground;

import com.algolia.exceptions.AlgoliaApiException;
import com.algolia.exceptions.AlgoliaRetryException;
import com.algolia.exceptions.AlgoliaRuntimeException;
import com.algolia.model.insights.*;
import com.algolia.api.InsightsClient;
import com.algolia.utils.AlgoliaAgent;
import io.github.cdimascio.dotenv.Dotenv;

public class Insights {
  public static void main(String[] args) {
    Dotenv dotenv = Dotenv.configure().directory("../").load();

    InsightsClient client = new InsightsClient(
      dotenv.get("ALGOLIA_APPLICATION_ID"),
      dotenv.get("ALGOLIA_SEARCH_KEY"),
      new AlgoliaAgent.Segment[] {
        new AlgoliaAgent.Segment("test", "8.0.0"),
        new AlgoliaAgent.Segment("JVM", "11.0.14"),
        new AlgoliaAgent.Segment("no version"),
      }
    );

    String indexName = dotenv.get("SEARCH_INDEX");
    InsightEvents params = new InsightEvents();
    InsightEvent event = new InsightEvent();
    event.setEventType(EventType.CLICK);
    event.setUserToken("user");
    event.setIndex("test_what");
    event.setEventName("test");
    params.addEvents(event);

    try {
      PushEventsResponse result = client.pushEvents(params);
      System.out.println(result);
    } catch (AlgoliaApiException e) {
      // the API failed
      System.err.println("Exception when calling InsightsClient#pushEvents");
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
