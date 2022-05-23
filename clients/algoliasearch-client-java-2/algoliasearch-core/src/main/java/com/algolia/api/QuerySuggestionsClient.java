package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.querySuggestions.*;
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

public class QuerySuggestionsClient extends ApiClient {

  public QuerySuggestionsClient(String appId, String apiKey, String region) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(region)), null);
  }

  public QuerySuggestionsClient(
    String appId,
    String apiKey,
    String region,
    AlgoliaAgent.Segment[] algoliaAgentSegments
  ) {
    this(
      appId,
      apiKey,
      new HttpRequester(getDefaultHosts(region)),
      algoliaAgentSegments
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
    AlgoliaAgent.Segment[] algoliaAgentSegments
  ) {
    super(appId, apiKey, requester, "QuerySuggestions", algoliaAgentSegments);
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
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SuccessResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SuccessResponse createConfig(
    QuerySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      createConfigAsync(querySuggestionsIndexWithIndexParam, requestOptions)
    );
  }

  public SuccessResponse createConfig(
    QuerySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam
  ) throws AlgoliaRuntimeException {
    return this.createConfig(querySuggestionsIndexWithIndexParam, null);
  }

  /**
   * (asynchronously) Create a configuration of a Query Suggestions index. There&#39;s a limit of
   * 100 configurations per application.
   *
   * @param querySuggestionsIndexWithIndexParam (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SuccessResponse> createConfigAsync(
    QuerySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam,
    RequestOptions requestOptions
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

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "POST",
          queryParameters,
          bodyObj,
          headers,
          requestOptions,
          false
        );
    Type returnType = new TypeToken<SuccessResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SuccessResponse> createConfigAsync(
    QuerySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam
  ) throws AlgoliaRuntimeException {
    return this.createConfigAsync(querySuggestionsIndexWithIndexParam, null);
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
          requestOptions,
          false
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
   * Delete a configuration of a Query Suggestion's index. By deleting a configuration, you stop all
   * updates to the underlying query suggestion index. Note that when doing this, the underlying
   * index does not change - existing suggestions remain untouched.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SuccessResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SuccessResponse deleteConfig(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteConfigAsync(indexName, requestOptions));
  }

  public SuccessResponse deleteConfig(String indexName)
    throws AlgoliaRuntimeException {
    return this.deleteConfig(indexName, null);
  }

  /**
   * (asynchronously) Delete a configuration of a Query Suggestion&#39;s index. By deleting a
   * configuration, you stop all updates to the underlying query suggestion index. Note that when
   * doing this, the underlying index does not change - existing suggestions remain untouched.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SuccessResponse> deleteConfigAsync(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
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

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "DELETE",
          queryParameters,
          bodyObj,
          headers,
          requestOptions,
          false
        );
    Type returnType = new TypeToken<SuccessResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SuccessResponse> deleteConfigAsync(String indexName)
    throws AlgoliaRuntimeException {
    return this.deleteConfigAsync(indexName, null);
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
          requestOptions,
          false
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
   * Get all the configurations of Query Suggestions. For each index, you get a block of JSON with a
   * list of its configuration settings.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return List&lt;QuerySuggestionsIndex&gt;
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public List<QuerySuggestionsIndex> getAllConfigs(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getAllConfigsAsync(requestOptions));
  }

  public List<QuerySuggestionsIndex> getAllConfigs()
    throws AlgoliaRuntimeException {
    return this.getAllConfigs(null);
  }

  /**
   * (asynchronously) Get all the configurations of Query Suggestions. For each index, you get a
   * block of JSON with a list of its configuration settings.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<List<QuerySuggestionsIndex>> getAllConfigsAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/configs";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions,
          false
        );
    Type returnType = new TypeToken<List<QuerySuggestionsIndex>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<List<QuerySuggestionsIndex>> getAllConfigsAsync()
    throws AlgoliaRuntimeException {
    return this.getAllConfigsAsync(null);
  }

  /**
   * Get the configuration of a single Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return QuerySuggestionsIndex
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public QuerySuggestionsIndex getConfig(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getConfigAsync(indexName, requestOptions));
  }

  public QuerySuggestionsIndex getConfig(String indexName)
    throws AlgoliaRuntimeException {
    return this.getConfig(indexName, null);
  }

  /**
   * (asynchronously) Get the configuration of a single Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<QuerySuggestionsIndex> getConfigAsync(
    String indexName,
    RequestOptions requestOptions
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

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions,
          false
        );
    Type returnType = new TypeToken<QuerySuggestionsIndex>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<QuerySuggestionsIndex> getConfigAsync(
    String indexName
  ) throws AlgoliaRuntimeException {
    return this.getConfigAsync(indexName, null);
  }

  /**
   * Get the status of a Query Suggestion's index. The status includes whether the Query Suggestions
   * index is currently in the process of being built, and the last build time.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Status
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Status getConfigStatus(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getConfigStatusAsync(indexName, requestOptions)
    );
  }

  public Status getConfigStatus(String indexName)
    throws AlgoliaRuntimeException {
    return this.getConfigStatus(indexName, null);
  }

  /**
   * (asynchronously) Get the status of a Query Suggestion&#39;s index. The status includes whether
   * the Query Suggestions index is currently in the process of being built, and the last build
   * time.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Status> getConfigStatusAsync(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
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

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions,
          false
        );
    Type returnType = new TypeToken<Status>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<Status> getConfigStatusAsync(String indexName)
    throws AlgoliaRuntimeException {
    return this.getConfigStatusAsync(indexName, null);
  }

  /**
   * Get the log file of the last build of a single Query Suggestion index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return List&lt;LogFile&gt;
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public List<LogFile> getLogFile(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getLogFileAsync(indexName, requestOptions));
  }

  public List<LogFile> getLogFile(String indexName)
    throws AlgoliaRuntimeException {
    return this.getLogFile(indexName, null);
  }

  /**
   * (asynchronously) Get the log file of the last build of a single Query Suggestion index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<List<LogFile>> getLogFileAsync(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
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

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions,
          false
        );
    Type returnType = new TypeToken<List<LogFile>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<List<LogFile>> getLogFileAsync(String indexName)
    throws AlgoliaRuntimeException {
    return this.getLogFileAsync(indexName, null);
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
          requestOptions,
          false
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
          requestOptions,
          false
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

  /**
   * Update the configuration of a Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param querySuggestionsIndexParam (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SuccessResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SuccessResponse updateConfig(
    String indexName,
    QuerySuggestionsIndexParam querySuggestionsIndexParam,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      updateConfigAsync(indexName, querySuggestionsIndexParam, requestOptions)
    );
  }

  public SuccessResponse updateConfig(
    String indexName,
    QuerySuggestionsIndexParam querySuggestionsIndexParam
  ) throws AlgoliaRuntimeException {
    return this.updateConfig(indexName, querySuggestionsIndexParam, null);
  }

  /**
   * (asynchronously) Update the configuration of a Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param querySuggestionsIndexParam (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SuccessResponse> updateConfigAsync(
    String indexName,
    QuerySuggestionsIndexParam querySuggestionsIndexParam,
    RequestOptions requestOptions
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

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "PUT",
          queryParameters,
          bodyObj,
          headers,
          requestOptions,
          false
        );
    Type returnType = new TypeToken<SuccessResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SuccessResponse> updateConfigAsync(
    String indexName,
    QuerySuggestionsIndexParam querySuggestionsIndexParam
  ) throws AlgoliaRuntimeException {
    return this.updateConfigAsync(indexName, querySuggestionsIndexParam, null);
  }
}
