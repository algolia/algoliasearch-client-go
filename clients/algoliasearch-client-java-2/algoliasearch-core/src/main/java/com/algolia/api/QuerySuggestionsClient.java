package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.querysuggestions.*;
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

  private static final String[] allowedRegions = { "eu", "us" };

  public QuerySuggestionsClient(String appId, String apiKey, String region) {
    this(appId, apiKey, region, null);
  }

  public QuerySuggestionsClient(String appId, String apiKey, String region, ClientOptions options) {
    super(appId, apiKey, "QuerySuggestions", "4.3.0-SNAPSHOT", options);
    if (options != null && options.getHosts() != null) {
      this.setHosts(options.getHosts());
    } else {
      this.setHosts(getDefaultHosts(region));
    }
    this.setConnectTimeout(2000);
    this.setReadTimeout(5000);
    this.setWriteTimeout(30000);
  }

  private static List<StatefulHost> getDefaultHosts(String region) throws AlgoliaRuntimeException {
    List<StatefulHost> hosts = new ArrayList<StatefulHost>();

    boolean found = false;
    if (region == null) {
      throw new AlgoliaRuntimeException("`region` is missing");
    }
    for (String allowed : allowedRegions) {
      if (allowed.equals(region)) {
        found = true;
        break;
      }
    }
    if (!found) {
      throw new AlgoliaRuntimeException("`region` must be one of the following: eu, us");
    }

    String url = "query-suggestions.{region}.algolia.com".replace("{region}", region);

    hosts.add(new StatefulHost(url, "https", EnumSet.of(CallType.READ, CallType.WRITE)));
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
    return LaunderThrowable.await(createConfigAsync(querySuggestionsIndexWithIndexParam, requestOptions));
  }

  /**
   * Create a configuration of a Query Suggestions index. There's a limit of 100 configurations per
   * application.
   *
   * @param querySuggestionsIndexWithIndexParam (required)
   * @return SuccessResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public SuccessResponse createConfig(QuerySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam)
    throws AlgoliaRuntimeException {
    return this.createConfig(querySuggestionsIndexWithIndexParam, null);
  }

  /**
   * (asynchronously) Create a configuration of a Query Suggestions index. There&#39;s a limit of
   * 100 configurations per application.
   *
   * @param querySuggestionsIndexWithIndexParam (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<SuccessResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<SuccessResponse> createConfigAsync(
    QuerySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (querySuggestionsIndexWithIndexParam == null) {
      throw new AlgoliaRuntimeException("Parameter `querySuggestionsIndexWithIndexParam` is required when calling" + " `createConfig`.");
    }

    Object bodyObj = querySuggestionsIndexWithIndexParam;

    // create path and map variables
    String requestPath = "/1/configs";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "POST", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<SuccessResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) Create a configuration of a Query Suggestions index. There&#39;s a limit of
   * 100 configurations per application.
   *
   * @param querySuggestionsIndexWithIndexParam (required)
   * @return CompletableFuture<SuccessResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<SuccessResponse> createConfigAsync(QuerySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam)
    throws AlgoliaRuntimeException {
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
  public Object del(String path, Map<String, Object> parameters, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(delAsync(path, parameters, requestOptions));
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Object del(String path, Map<String, Object> parameters) throws AlgoliaRuntimeException {
    return this.del(path, parameters, null);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Object del(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.del(path, null, requestOptions);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
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
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> delAsync(String path, Map<String, Object> parameters, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException("Parameter `path` is required when calling `del`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParameters.put(parameter.getKey().toString(), parameterToString(parameter.getValue()));
      }
    }

    Call call = this.buildCall(requestPath, "DELETE", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> delAsync(String path, Map<String, Object> parameters) throws AlgoliaRuntimeException {
    return this.delAsync(path, parameters, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> delAsync(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.delAsync(path, null, requestOptions);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> delAsync(String path) throws AlgoliaRuntimeException {
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
  public SuccessResponse deleteConfig(String indexName, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteConfigAsync(indexName, requestOptions));
  }

  /**
   * Delete a configuration of a Query Suggestion's index. By deleting a configuration, you stop all
   * updates to the underlying query suggestion index. Note that when doing this, the underlying
   * index does not change - existing suggestions remain untouched.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return SuccessResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public SuccessResponse deleteConfig(String indexName) throws AlgoliaRuntimeException {
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
   * @return CompletableFuture<SuccessResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<SuccessResponse> deleteConfigAsync(String indexName, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException("Parameter `indexName` is required when calling `deleteConfig`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/configs/{indexName}".replaceAll("\\{indexName\\}", this.escapeString(indexName.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "DELETE", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<SuccessResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) Delete a configuration of a Query Suggestion&#39;s index. By deleting a
   * configuration, you stop all updates to the underlying query suggestion index. Note that when
   * doing this, the underlying index does not change - existing suggestions remain untouched.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return CompletableFuture<SuccessResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<SuccessResponse> deleteConfigAsync(String indexName) throws AlgoliaRuntimeException {
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
  public Object get(String path, Map<String, Object> parameters, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getAsync(path, parameters, requestOptions));
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Object get(String path, Map<String, Object> parameters) throws AlgoliaRuntimeException {
    return this.get(path, parameters, null);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Object get(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.get(path, null, requestOptions);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
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
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> getAsync(String path, Map<String, Object> parameters, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException("Parameter `path` is required when calling `get`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParameters.put(parameter.getKey().toString(), parameterToString(parameter.getValue()));
      }
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> getAsync(String path, Map<String, Object> parameters) throws AlgoliaRuntimeException {
    return this.getAsync(path, parameters, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> getAsync(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getAsync(path, null, requestOptions);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> getAsync(String path) throws AlgoliaRuntimeException {
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
  public List<QuerySuggestionsIndex> getAllConfigs(RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getAllConfigsAsync(requestOptions));
  }

  /**
   * Get all the configurations of Query Suggestions. For each index, you get a block of JSON with a
   * list of its configuration settings.
   *
   * @return List&lt;QuerySuggestionsIndex&gt;
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<QuerySuggestionsIndex> getAllConfigs() throws AlgoliaRuntimeException {
    return this.getAllConfigs(null);
  }

  /**
   * (asynchronously) Get all the configurations of Query Suggestions. For each index, you get a
   * block of JSON with a list of its configuration settings.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<List<QuerySuggestionsIndex>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<QuerySuggestionsIndex>> getAllConfigsAsync(RequestOptions requestOptions) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/configs";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<List<QuerySuggestionsIndex>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) Get all the configurations of Query Suggestions. For each index, you get a
   * block of JSON with a list of its configuration settings.
   *
   * @return CompletableFuture<List<QuerySuggestionsIndex>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<QuerySuggestionsIndex>> getAllConfigsAsync() throws AlgoliaRuntimeException {
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
  public QuerySuggestionsIndex getConfig(String indexName, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getConfigAsync(indexName, requestOptions));
  }

  /**
   * Get the configuration of a single Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return QuerySuggestionsIndex
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public QuerySuggestionsIndex getConfig(String indexName) throws AlgoliaRuntimeException {
    return this.getConfig(indexName, null);
  }

  /**
   * (asynchronously) Get the configuration of a single Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<QuerySuggestionsIndex> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<QuerySuggestionsIndex> getConfigAsync(String indexName, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException("Parameter `indexName` is required when calling `getConfig`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/configs/{indexName}".replaceAll("\\{indexName\\}", this.escapeString(indexName.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<QuerySuggestionsIndex>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) Get the configuration of a single Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return CompletableFuture<QuerySuggestionsIndex> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<QuerySuggestionsIndex> getConfigAsync(String indexName) throws AlgoliaRuntimeException {
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
  public Status getConfigStatus(String indexName, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getConfigStatusAsync(indexName, requestOptions));
  }

  /**
   * Get the status of a Query Suggestion's index. The status includes whether the Query Suggestions
   * index is currently in the process of being built, and the last build time.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return Status
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Status getConfigStatus(String indexName) throws AlgoliaRuntimeException {
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
   * @return CompletableFuture<Status> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Status> getConfigStatusAsync(String indexName, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException("Parameter `indexName` is required when calling `getConfigStatus`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/configs/{indexName}/status".replaceAll("\\{indexName\\}", this.escapeString(indexName.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<Status>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) Get the status of a Query Suggestion&#39;s index. The status includes whether
   * the Query Suggestions index is currently in the process of being built, and the last build
   * time.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return CompletableFuture<Status> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Status> getConfigStatusAsync(String indexName) throws AlgoliaRuntimeException {
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
  public List<LogFile> getLogFile(String indexName, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getLogFileAsync(indexName, requestOptions));
  }

  /**
   * Get the log file of the last build of a single Query Suggestion index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return List&lt;LogFile&gt;
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<LogFile> getLogFile(String indexName) throws AlgoliaRuntimeException {
    return this.getLogFile(indexName, null);
  }

  /**
   * (asynchronously) Get the log file of the last build of a single Query Suggestion index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<List<LogFile>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<LogFile>> getLogFileAsync(String indexName, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException("Parameter `indexName` is required when calling `getLogFile`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/logs/{indexName}".replaceAll("\\{indexName\\}", this.escapeString(indexName.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<List<LogFile>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) Get the log file of the last build of a single Query Suggestion index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return CompletableFuture<List<LogFile>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<LogFile>> getLogFileAsync(String indexName) throws AlgoliaRuntimeException {
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
  public Object post(String path, Map<String, Object> parameters, Object body, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(postAsync(path, parameters, body, requestOptions));
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Object post(String path, Map<String, Object> parameters, Object body) throws AlgoliaRuntimeException {
    return this.post(path, parameters, body, null);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Object post(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.post(path, null, null, requestOptions);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
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
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> postAsync(String path, Map<String, Object> parameters, Object body, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException("Parameter `path` is required when calling `post`.");
    }

    Object bodyObj = body;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParameters.put(parameter.getKey().toString(), parameterToString(parameter.getValue()));
      }
    }

    Call call = this.buildCall(requestPath, "POST", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> postAsync(String path, Map<String, Object> parameters, Object body) throws AlgoliaRuntimeException {
    return this.postAsync(path, parameters, body, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> postAsync(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.postAsync(path, null, null, requestOptions);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> postAsync(String path) throws AlgoliaRuntimeException {
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
  public Object put(String path, Map<String, Object> parameters, Object body, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(putAsync(path, parameters, body, requestOptions));
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Object put(String path, Map<String, Object> parameters, Object body) throws AlgoliaRuntimeException {
    return this.put(path, parameters, body, null);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Object put(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.put(path, null, null, requestOptions);
  }

  /**
   * This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @return Object
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
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
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> putAsync(String path, Map<String, Object> parameters, Object body, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException("Parameter `path` is required when calling `put`.");
    }

    Object bodyObj = body;

    // create path and map variables
    String requestPath = "/1{path}".replaceAll("\\{path\\}", path.toString());

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (parameters != null) {
      for (Map.Entry<String, Object> parameter : parameters.entrySet()) {
        queryParameters.put(parameter.getKey().toString(), parameterToString(parameter.getValue()));
      }
    }

    Call call = this.buildCall(requestPath, "PUT", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<Object>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param parameters Query parameters to be applied to the current query. (optional)
   * @param body The parameters to send with the custom request. (optional)
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> putAsync(String path, Map<String, Object> parameters, Object body) throws AlgoliaRuntimeException {
    return this.putAsync(path, parameters, body, null);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> putAsync(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.putAsync(path, null, null, requestOptions);
  }

  /**
   * (asynchronously) This method allow you to send requests to the Algolia REST API.
   *
   * @param path The path of the API endpoint to target, anything after the /1 needs to be
   *     specified. (required)
   * @return CompletableFuture<Object> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Object> putAsync(String path) throws AlgoliaRuntimeException {
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
    return LaunderThrowable.await(updateConfigAsync(indexName, querySuggestionsIndexParam, requestOptions));
  }

  /**
   * Update the configuration of a Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param querySuggestionsIndexParam (required)
   * @return SuccessResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public SuccessResponse updateConfig(String indexName, QuerySuggestionsIndexParam querySuggestionsIndexParam)
    throws AlgoliaRuntimeException {
    return this.updateConfig(indexName, querySuggestionsIndexParam, null);
  }

  /**
   * (asynchronously) Update the configuration of a Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param querySuggestionsIndexParam (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<SuccessResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<SuccessResponse> updateConfigAsync(
    String indexName,
    QuerySuggestionsIndexParam querySuggestionsIndexParam,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException("Parameter `indexName` is required when calling `updateConfig`.");
    }

    if (querySuggestionsIndexParam == null) {
      throw new AlgoliaRuntimeException("Parameter `querySuggestionsIndexParam` is required when calling `updateConfig`.");
    }

    Object bodyObj = querySuggestionsIndexParam;

    // create path and map variables
    String requestPath = "/1/configs/{indexName}".replaceAll("\\{indexName\\}", this.escapeString(indexName.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "PUT", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<SuccessResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * (asynchronously) Update the configuration of a Query Suggestions index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param querySuggestionsIndexParam (required)
   * @return CompletableFuture<SuccessResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<SuccessResponse> updateConfigAsync(String indexName, QuerySuggestionsIndexParam querySuggestionsIndexParam)
    throws AlgoliaRuntimeException {
    return this.updateConfigAsync(indexName, querySuggestionsIndexParam, null);
  }
}
