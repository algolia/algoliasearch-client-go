package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.abtesting.*;
import com.algolia.utils.*;
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

public class AbtestingClient extends ApiClient {

  public AbtestingClient(String appId, String apiKey) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(null)), null);
  }

  public AbtestingClient(
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

  public AbtestingClient(String appId, String apiKey, String region) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(region)), null);
  }

  public AbtestingClient(
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

  public AbtestingClient(String appId, String apiKey, Requester requester) {
    this(appId, apiKey, requester, null);
  }

  public AbtestingClient(
    String appId,
    String apiKey,
    Requester requester,
    UserAgent.Segment[] userAgentSegments
  ) {
    super(appId, apiKey, requester, "Abtesting", userAgentSegments);
  }

  private static List<StatefulHost> getDefaultHosts(String region) {
    List<StatefulHost> hosts = new ArrayList<StatefulHost>();

    String url = region == null
      ? "analytics.algolia.com"
      : "analytics.{region}.algolia.com".replace("{region}", region);

    hosts.add(
      new StatefulHost(url, "https", EnumSet.of(CallType.READ, CallType.WRITE))
    );
    return hosts;
  }

  /**
   * Creates a new A/B test with provided configuration. You can set an A/B test on two different
   * indices with different settings, or on the same index with different search parameters by
   * providing a customSearchParameters setting on one of the variants.
   *
   * @param addABTestsRequest (required)
   * @return ABTestResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ABTestResponse addABTests(AddABTestsRequest addABTestsRequest)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(addABTestsAsync(addABTestsRequest));
  }

  /**
   * (asynchronously) Creates a new A/B test with provided configuration. You can set an A/B test on
   * two different indices with different settings, or on the same index with different search
   * parameters by providing a customSearchParameters setting on one of the variants.
   *
   * @param addABTestsRequest (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ABTestResponse> addABTestsAsync(
    AddABTestsRequest addABTestsRequest
  ) throws AlgoliaRuntimeException {
    if (addABTestsRequest == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'addABTestsRequest' when calling addABTests(Async)"
      );
    }

    Object bodyObj = addABTestsRequest;

    // create path and map variables
    String requestPath = "/2/abtests";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<ABTestResponse>() {}.getType();
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
   * Delete a test.
   *
   * @param id The A/B test ID. (required)
   * @return ABTestResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ABTestResponse deleteABTest(Integer id)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteABTestAsync(id));
  }

  /**
   * (asynchronously) Delete a test.
   *
   * @param id The A/B test ID. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ABTestResponse> deleteABTestAsync(Integer id)
    throws AlgoliaRuntimeException {
    if (id == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'id' when calling deleteABTest(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/2/abtests/{id}".replaceAll(
          "\\{id\\}",
          this.escapeString(id.toString())
        );

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<ABTestResponse>() {}.getType();
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
   * Returns metadata and metrics for an A/B test.
   *
   * @param id The A/B test ID. (required)
   * @return ABTest
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ABTest getABTest(Integer id) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getABTestAsync(id));
  }

  /**
   * (asynchronously) Returns metadata and metrics for an A/B test.
   *
   * @param id The A/B test ID. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ABTest> getABTestAsync(Integer id)
    throws AlgoliaRuntimeException {
    if (id == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'id' when calling getABTest(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/2/abtests/{id}".replaceAll(
          "\\{id\\}",
          this.escapeString(id.toString())
        );

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<ABTest>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Fetch all existing A/B tests for App that are available for the current API Key. When no data
   * has been processed, the metrics will be returned as null.
   *
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @return ListABTestsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ListABTestsResponse listABTests(Integer offset, Integer limit)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(listABTestsAsync(offset, limit));
  }

  public ListABTestsResponse listABTests() throws AlgoliaRuntimeException {
    return this.listABTests(null, null);
  }

  /**
   * (asynchronously) Fetch all existing A/B tests for App that are available for the current API
   * Key. When no data has been processed, the metrics will be returned as null.
   *
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ListABTestsResponse> listABTestsAsync(
    Integer offset,
    Integer limit
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/abtests";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (offset != null) {
      queryParams.put("offset", parameterToString(offset));
    }

    if (limit != null) {
      queryParams.put("limit", parameterToString(limit));
    }

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<ListABTestsResponse>() {}.getType();
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

  /**
   * Marks the A/B test as stopped. At this point, the test is over and cannot be restarted. As a
   * result, your application is back to normal: index A will perform as usual, receiving 100% of
   * all search requests. Associated metadata and metrics are still stored.
   *
   * @param id The A/B test ID. (required)
   * @return ABTestResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ABTestResponse stopABTest(Integer id) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(stopABTestAsync(id));
  }

  /**
   * (asynchronously) Marks the A/B test as stopped. At this point, the test is over and cannot be
   * restarted. As a result, your application is back to normal: index A will perform as usual,
   * receiving 100% of all search requests. Associated metadata and metrics are still stored.
   *
   * @param id The A/B test ID. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ABTestResponse> stopABTestAsync(Integer id)
    throws AlgoliaRuntimeException {
    if (id == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'id' when calling stopABTest(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/2/abtests/{id}/stop".replaceAll(
          "\\{id\\}",
          this.escapeString(id.toString())
        );

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<ABTestResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }
}
