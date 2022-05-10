package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.querySuggestions.*;
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

public class QuerySuggestionsClient extends ApiClient {

  public QuerySuggestionsClient(String appId, String apiKey, String region) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(region)), null);
  }

  public QuerySuggestionsClient(
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

  public QuerySuggestionsClient(
    String appId,
    String apiKey,
    Requester requester
  ) {
    this(appId, apiKey, requester, null);
  }

  public QuerySuggestionsClient(
    String appId,
    String apiKey,
    Requester requester,
    UserAgent.Segment[] userAgentSegments
  ) {
    super(appId, apiKey, requester, "QuerySuggestions", userAgentSegments);
  }

  private static List<StatefulHost> getDefaultHosts(String region) {
    List<StatefulHost> hosts = new ArrayList<StatefulHost>();

    String url =
      "query-suggestions.{region}.algolia.com".replace("{region}", region);

    hosts.add(
      new StatefulHost(url, "https", EnumSet.of(CallType.READ, CallType.WRITE))
    );
    return hosts;
  }

  /**
   * Create a configuration of a Query Suggestions index. There's a limit of 100 configurations per
   * application.
   *
   * @param querySuggestionsIndexWithIndexParam (required)
   * @return SucessResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SucessResponse createConfig(
    QuerySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      createConfigAsync(querySuggestionsIndexWithIndexParam)
    );
  }

  /**
   * (asynchronously) Create a configuration of a Query Suggestions index. There&#39;s a limit of
   * 100 configurations per application.
   *
   * @param querySuggestionsIndexWithIndexParam (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SucessResponse> createConfigAsync(
    QuerySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam
  ) throws AlgoliaRuntimeException {
    if (querySuggestionsIndexWithIndexParam == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'querySuggestionsIndexWithIndexParam' when calling" +
        " createConfig(Async)"
      );
    }

    Object bodyObj = querySuggestionsIndexWithIndexParam;

    // create path and map variables
    String requestPath = "/1/configs";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SucessResponse>() {}.getType();
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
   * Delete a configuration of a Query Suggestion's index. By deleting a configuraton, you stop all
   * updates to the underlying query suggestion index. Note that when doing this, the underlying
   * index does not change - existing suggestions remain untouched.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return SucessResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SucessResponse deleteConfig(String indexName)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteConfigAsync(indexName));
  }

  /**
   * (asynchronously) Delete a configuration of a Query Suggestion&#39;s index. By deleting a
   * configuraton, you stop all updates to the underlying query suggestion index. Note that when
   * doing this, the underlying index does not change - existing suggestions remain untouched.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SucessResponse> deleteConfigAsync(String indexName)
    throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling deleteConfig(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/configs/{indexName}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SucessResponse>() {}.getType();
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
   * Get all the configurations of Query Suggestions. For each index, you get a block of JSON with a
   * list of its configuration settings.
   *
   * @return List&lt;QuerySuggestionsIndex&gt;
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public List<QuerySuggestionsIndex> getAllConfigs()
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getAllConfigsAsync());
  }

  /**
   * (asynchronously) Get all the configurations of Query Suggestions. For each index, you get a
   * block of JSON with a list of its configuration settings.
   *
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<List<QuerySuggestionsIndex>> getAllConfigsAsync()
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/configs";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<List<QuerySuggestionsIndex>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Get the configuration of a single Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return QuerySuggestionsIndex
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public QuerySuggestionsIndex getConfig(String indexName)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getConfigAsync(indexName));
  }

  /**
   * (asynchronously) Get the configuration of a single Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<QuerySuggestionsIndex> getConfigAsync(
    String indexName
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling getConfig(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/configs/{indexName}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<QuerySuggestionsIndex>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Get the status of a Query Suggestion's index. The status includes whether the Query Suggestions
   * index is currently in the process of being built, and the last build time.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return Status
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Status getConfigStatus(String indexName)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getConfigStatusAsync(indexName));
  }

  /**
   * (asynchronously) Get the status of a Query Suggestion&#39;s index. The status includes whether
   * the Query Suggestions index is currently in the process of being built, and the last build
   * time.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Status> getConfigStatusAsync(String indexName)
    throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling getConfigStatus(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/configs/{indexName}/status".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<Status>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Get the log file of the last build of a single Query Suggestion index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return List&lt;LogFile&gt;
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public List<LogFile> getLogFile(String indexName)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getLogFileAsync(indexName));
  }

  /**
   * (asynchronously) Get the log file of the last build of a single Query Suggestion index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<List<LogFile>> getLogFileAsync(String indexName)
    throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling getLogFile(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/logs/{indexName}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<List<LogFile>>() {}.getType();
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
   * Update the configuration of a Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param querySuggestionsIndexParam (required)
   * @return SucessResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SucessResponse updateConfig(
    String indexName,
    QuerySuggestionsIndexParam querySuggestionsIndexParam
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      updateConfigAsync(indexName, querySuggestionsIndexParam)
    );
  }

  /**
   * (asynchronously) Update the configuration of a Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param querySuggestionsIndexParam (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SucessResponse> updateConfigAsync(
    String indexName,
    QuerySuggestionsIndexParam querySuggestionsIndexParam
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling updateConfig(Async)"
      );
    }

    if (querySuggestionsIndexParam == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'querySuggestionsIndexParam' when calling" +
        " updateConfig(Async)"
      );
    }

    Object bodyObj = querySuggestionsIndexParam;

    // create path and map variables
    String requestPath =
      "/1/configs/{indexName}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "PUT", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SucessResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }
}
