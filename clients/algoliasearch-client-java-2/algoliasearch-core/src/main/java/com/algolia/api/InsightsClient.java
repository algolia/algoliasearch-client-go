package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.insights.*;
import com.algolia.utils.*;
import com.algolia.utils.RequestOptions;
import com.algolia.utils.retry.CallType;
import com.algolia.utils.retry.StatefulHost;
import com.google.gson.reflect.TypeToken;
import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.EnumSet;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.CompletableFuture;
import okhttp3.Call;

public class InsightsClient extends ApiClient {

  public InsightsClient(String appId, String apiKey) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(null)), null);
  }

  public InsightsClient(
    String appId,
    String apiKey,
    UserAgent.Segment[] userAgentSegments
  ) {
    this(
      appId,
      apiKey,
      new HttpRequester(getDefaultHosts(null)),
      userAgentSegments
    );
  }

  public InsightsClient(String appId, String apiKey, String region) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(region)), null);
  }

  public InsightsClient(
    String appId,
    String apiKey,
    String region,
    UserAgent.Segment[] userAgentSegments
  ) {
    this(
      appId,
      apiKey,
      new HttpRequester(getDefaultHosts(region)),
      userAgentSegments
    );
  }

  public InsightsClient(String appId, String apiKey, Requester requester) {
    this(appId, apiKey, requester, null);
  }

  public InsightsClient(
    String appId,
    String apiKey,
    Requester requester,
    UserAgent.Segment[] userAgentSegments
  ) {
    super(appId, apiKey, requester, "Insights", userAgentSegments);
  }

  private static List<StatefulHost> getDefaultHosts(String region) {
    List<StatefulHost> hosts = new ArrayList<StatefulHost>();

    String url = region == null
      ? "insights.algolia.io"
      : "insights.{region}.algolia.io".replace("{region}", region);

    hosts.add(
      new StatefulHost(url, "https", EnumSet.of(CallType.READ, CallType.WRITE))
    );
    return hosts;
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Object
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Object del(
    String path,
    Map<String, Object> parameters,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(delAsync(path, parameters, requestOptions));
  }

  public Object del(String path, Map<String, Object> parameters)
    throws AlgoliaRuntimeException {
    return this.del(path, parameters, null);
  }

  public Object del(String path, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.del(path, null, requestOptions);
  }

  public Object del(String path) throws AlgoliaRuntimeException {
    return this.del(path, null, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Object> delAsync(
    String path,
    Map<String, Object> parameters,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'path' when calling del(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParameters.put(
          parameter.getKey().toString(),
          parameterToString(parameter.getValue())
        );
      }
    }

    Call call =
      this.buildCall(
          requestPath,
          "DELETE",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<Object> delAsync(
    String path,
    Map<String, Object> parameters
  ) throws AlgoliaRuntimeException {
    return this.delAsync(path, parameters, null);
  }

  public CompletableFuture<Object> delAsync(
    String path,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.delAsync(path, null, requestOptions);
  }

  public CompletableFuture<Object> delAsync(String path)
    throws AlgoliaRuntimeException {
    return this.delAsync(path, null, null);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Object
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Object get(
    String path,
    Map<String, Object> parameters,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getAsync(path, parameters, requestOptions));
  }

  public Object get(String path, Map<String, Object> parameters)
    throws AlgoliaRuntimeException {
    return this.get(path, parameters, null);
  }

  public Object get(String path, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.get(path, null, requestOptions);
  }

  public Object get(String path) throws AlgoliaRuntimeException {
    return this.get(path, null, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Object> getAsync(
    String path,
    Map<String, Object> parameters,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'path' when calling get(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParameters.put(
          parameter.getKey().toString(),
          parameterToString(parameter.getValue())
        );
      }
    }

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<Object> getAsync(
    String path,
    Map<String, Object> parameters
  ) throws AlgoliaRuntimeException {
    return this.getAsync(path, parameters, null);
  }

  public CompletableFuture<Object> getAsync(
    String path,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.getAsync(path, null, requestOptions);
  }

  public CompletableFuture<Object> getAsync(String path)
    throws AlgoliaRuntimeException {
    return this.getAsync(path, null, null);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Object
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Object post(
    String path,
    Map<String, Object> parameters,
    Object body,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      postAsync(path, parameters, body, requestOptions)
    );
  }

  public Object post(String path, Map<String, Object> parameters, Object body)
    throws AlgoliaRuntimeException {
    return this.post(path, parameters, body, null);
  }

  public Object post(String path, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.post(path, null, null, requestOptions);
  }

  public Object post(String path) throws AlgoliaRuntimeException {
    return this.post(path, null, null, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Object> postAsync(
    String path,
    Map<String, Object> parameters,
    Object body,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'path' when calling post(Async)"
      );
    }

    Object bodyObj = body;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParameters.put(
          parameter.getKey().toString(),
          parameterToString(parameter.getValue())
        );
      }
    }

    Call call =
      this.buildCall(
          requestPath,
          "POST",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<Object> postAsync(
    String path,
    Map<String, Object> parameters,
    Object body
  ) throws AlgoliaRuntimeException {
    return this.postAsync(path, parameters, body, null);
  }

  public CompletableFuture<Object> postAsync(
    String path,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.postAsync(path, null, null, requestOptions);
  }

  public CompletableFuture<Object> postAsync(String path)
    throws AlgoliaRuntimeException {
    return this.postAsync(path, null, null, null);
  }

  /**
   * This command pushes an array of events. An event is - an action: `eventName` - performed in a
   * context: `eventType` - at some point in time provided: `timestamp` - by an end user:
   * `userToken` - on something: `index` Notes: - To be accepted, all events sent must be valid. -
   * The size of the body must be *less than 2 MB*. - When an event is tied to an Algolia search, it
   * must also provide a `queryID`. If that event is a `click`, their absolute `positions` should
   * also be passed. - We consider that an `index` provides access to 2 resources: objects and
   * filters. An event can only interact with a single resource type, but not necessarily on a
   * single item. As such an event will accept an array of `objectIDs` or `filters`.
   *
   * @param insightEvents (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return PushEventsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public PushEventsResponse pushEvents(
    InsightEvents insightEvents,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      pushEventsAsync(insightEvents, requestOptions)
    );
  }

  public PushEventsResponse pushEvents(InsightEvents insightEvents)
    throws AlgoliaRuntimeException {
    return this.pushEvents(insightEvents, null);
  }

  /**
   * (asynchronously) This command pushes an array of events. An event is - an action:
   * &#x60;eventName&#x60; - performed in a context: &#x60;eventType&#x60; - at some point in time
   * provided: &#x60;timestamp&#x60; - by an end user: &#x60;userToken&#x60; - on something:
   * &#x60;index&#x60; Notes: - To be accepted, all events sent must be valid. - The size of the
   * body must be *less than 2 MB*. - When an event is tied to an Algolia search, it must also
   * provide a &#x60;queryID&#x60;. If that event is a &#x60;click&#x60;, their absolute
   * &#x60;positions&#x60; should also be passed. - We consider that an &#x60;index&#x60; provides
   * access to 2 resources: objects and filters. An event can only interact with a single resource
   * type, but not necessarily on a single item. As such an event will accept an array of
   * &#x60;objectIDs&#x60; or &#x60;filters&#x60;.
   *
   * @param insightEvents (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<PushEventsResponse> pushEventsAsync(
    InsightEvents insightEvents,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (insightEvents == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'insightEvents' when calling pushEvents(Async)"
      );
    }

    Object bodyObj = insightEvents;

    // create path and map variables
    String requestPath = "/1/events";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "POST",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<PushEventsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<PushEventsResponse> pushEventsAsync(
    InsightEvents insightEvents
  ) throws AlgoliaRuntimeException {
    return this.pushEventsAsync(insightEvents, null);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Object
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Object put(
    String path,
    Map<String, Object> parameters,
    Object body,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      putAsync(path, parameters, body, requestOptions)
    );
  }

  public Object put(String path, Map<String, Object> parameters, Object body)
    throws AlgoliaRuntimeException {
    return this.put(path, parameters, body, null);
  }

  public Object put(String path, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.put(path, null, null, requestOptions);
  }

  public Object put(String path) throws AlgoliaRuntimeException {
    return this.put(path, null, null, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Object> putAsync(
    String path,
    Map<String, Object> parameters,
    Object body,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'path' when calling put(Async)"
      );
    }

    Object bodyObj = body;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParameters.put(
          parameter.getKey().toString(),
          parameterToString(parameter.getValue())
        );
      }
    }

    Call call =
      this.buildCall(
          requestPath,
          "PUT",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<Object> putAsync(
    String path,
    Map<String, Object> parameters,
    Object body
  ) throws AlgoliaRuntimeException {
    return this.putAsync(path, parameters, body, null);
  }

  public CompletableFuture<Object> putAsync(
    String path,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.putAsync(path, null, null, requestOptions);
  }

  public CompletableFuture<Object> putAsync(String path)
    throws AlgoliaRuntimeException {
    return this.putAsync(path, null, null, null);
  }
}
