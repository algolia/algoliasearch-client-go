package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.recommend.*;
import com.algolia.utils.*;
import com.algolia.utils.retry.CallType;
import com.algolia.utils.retry.StatefulHost;
import com.google.gson.reflect.TypeToken;
import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.Collections;
import java.util.EnumSet;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Random;
import java.util.concurrent.CompletableFuture;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import okhttp3.Call;

public class RecommendClient extends ApiClient {

  public RecommendClient(String appId, String apiKey) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(appId)), null);
  }

  public RecommendClient(
    String appId,
    String apiKey,
    UserAgent.Segment[] userAgentSegments
  ) {
    this(
      appId,
      apiKey,
      new HttpRequester(getDefaultHosts(appId)),
      userAgentSegments
    );
  }

  public RecommendClient(String appId, String apiKey, Requester requester) {
    this(appId, apiKey, requester, null);
  }

  public RecommendClient(
    String appId,
    String apiKey,
    Requester requester,
    UserAgent.Segment[] userAgentSegments
  ) {
    super(appId, apiKey, requester, "Recommend", userAgentSegments);
  }

  private static List<StatefulHost> getDefaultHosts(String appId) {
    List<StatefulHost> hosts = new ArrayList<StatefulHost>();
    hosts.add(
      new StatefulHost(
        appId + "-dsn.algolia.net",
        "https",
        EnumSet.of(CallType.READ)
      )
    );
    hosts.add(
      new StatefulHost(
        appId + ".algolia.net",
        "https",
        EnumSet.of(CallType.WRITE)
      )
    );

    List<StatefulHost> commonHosts = new ArrayList<StatefulHost>();
    hosts.add(
      new StatefulHost(
        appId + "-1.algolianet.net",
        "https",
        EnumSet.of(CallType.READ, CallType.WRITE)
      )
    );
    hosts.add(
      new StatefulHost(
        appId + "-2.algolianet.net",
        "https",
        EnumSet.of(CallType.READ, CallType.WRITE)
      )
    );
    hosts.add(
      new StatefulHost(
        appId + "-3.algolianet.net",
        "https",
        EnumSet.of(CallType.READ, CallType.WRITE)
      )
    );

    Collections.shuffle(commonHosts, new Random());

    return Stream
      .concat(hosts.stream(), commonHosts.stream())
      .collect(Collectors.toList());
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @return Object
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Object del(String path, Map<String, Object> parameters)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(delAsync(path, parameters));
  }

  public Object del(String path) throws AlgoliaRuntimeException {
    return this.del(path, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Object> delAsync(
    String path,
    Map<String, Object> parameters
  ) throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'path' when calling del(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParams.put(
          parameter.getKey().toString(),
          parameterToString(parameter.getValue())
        );
      }
    }

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @return Object
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Object get(String path, Map<String, Object> parameters)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getAsync(path, parameters));
  }

  public Object get(String path) throws AlgoliaRuntimeException {
    return this.get(path, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Object> getAsync(
    String path,
    Map<String, Object> parameters
  ) throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'path' when calling get(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParams.put(
          parameter.getKey().toString(),
          parameterToString(parameter.getValue())
        );
      }
    }

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Returns recommendations or trending results, for a specific model and `objectID`.
   *
   * @param getRecommendationsParams (required)
   * @return GetRecommendationsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetRecommendationsResponse getRecommendations(
    GetRecommendationsParams getRecommendationsParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getRecommendationsAsync(getRecommendationsParams)
    );
  }

  /**
   * (asynchronously) Returns recommendations or trending results, for a specific model and
   * &#x60;objectID&#x60;.
   *
   * @param getRecommendationsParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetRecommendationsResponse> getRecommendationsAsync(
    GetRecommendationsParams getRecommendationsParams
  ) throws AlgoliaRuntimeException {
    if (getRecommendationsParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'getRecommendationsParams' when calling" +
        " getRecommendations(Async)"
      );
    }

    Object bodyObj = getRecommendationsParams;

    // create path and map variables
    String requestPath = "/1/indexes/*/recommendations";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<GetRecommendationsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @return Object
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Object post(String path, Map<String, Object> parameters, Object body)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(postAsync(path, parameters, body));
  }

  public Object post(String path) throws AlgoliaRuntimeException {
    return this.post(path, null, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Object> postAsync(
    String path,
    Map<String, Object> parameters,
    Object body
  ) throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'path' when calling post(Async)"
      );
    }

    Object bodyObj = body;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParams.put(
          parameter.getKey().toString(),
          parameterToString(parameter.getValue())
        );
      }
    }

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @return Object
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Object put(String path, Map<String, Object> parameters, Object body)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(putAsync(path, parameters, body));
  }

  public Object put(String path) throws AlgoliaRuntimeException {
    return this.put(path, null, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Object> putAsync(
    String path,
    Map<String, Object> parameters,
    Object body
  ) throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'path' when calling put(Async)"
      );
    }

    Object bodyObj = body;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParams.put(
          parameter.getKey().toString(),
          parameterToString(parameter.getValue())
        );
      }
    }

    Call call =
      this.buildCall(requestPath, "PUT", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }
}
