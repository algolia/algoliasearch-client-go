package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.analytics.*;
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

public class AnalyticsClient extends ApiClient {

  public AnalyticsClient(String appId, String apiKey) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(null)), null);
  }

  public AnalyticsClient(String appId, String apiKey, AlgoliaAgent.Segment[] algoliaAgentSegments) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(null)), algoliaAgentSegments);
  }

  public AnalyticsClient(String appId, String apiKey, String region) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(region)), null);
  }

  public AnalyticsClient(String appId, String apiKey, String region, AlgoliaAgent.Segment[] algoliaAgentSegments) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(region)), algoliaAgentSegments);
  }

  public AnalyticsClient(String appId, String apiKey, Requester requester) {
    this(appId, apiKey, requester, null);
  }

  public AnalyticsClient(String appId, String apiKey, Requester requester, AlgoliaAgent.Segment[] algoliaAgentSegments) {
    super(appId, apiKey, requester, "Analytics", algoliaAgentSegments);
  }

  private static List<StatefulHost> getDefaultHosts(String region) {
    List<StatefulHost> hosts = new ArrayList<StatefulHost>();

    String url = region == null ? "analytics.algolia.com" : "analytics.{region}.algolia.com".replace("{region}", region);

    hosts.add(new StatefulHost(url, "https", EnumSet.of(CallType.READ, CallType.WRITE)));
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
  public Object del(String path, Map<String, Object> parameters, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(delAsync(path, parameters, requestOptions));
  }

  public Object del(String path, Map<String, Object> parameters) throws AlgoliaRuntimeException {
    return this.del(path, parameters, null);
  }

  public Object del(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
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
  public CompletableFuture<Object> delAsync(String path, Map<String, Object> parameters, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'path' when calling del(Async)");
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

  public CompletableFuture<Object> delAsync(String path, Map<String, Object> parameters) throws AlgoliaRuntimeException {
    return this.delAsync(path, parameters, null);
  }

  public CompletableFuture<Object> delAsync(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.delAsync(path, null, requestOptions);
  }

  public CompletableFuture<Object> delAsync(String path) throws AlgoliaRuntimeException {
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
  public Object get(String path, Map<String, Object> parameters, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getAsync(path, parameters, requestOptions));
  }

  public Object get(String path, Map<String, Object> parameters) throws AlgoliaRuntimeException {
    return this.get(path, parameters, null);
  }

  public Object get(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
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
  public CompletableFuture<Object> getAsync(String path, Map<String, Object> parameters, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'path' when calling get(Async)");
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

  public CompletableFuture<Object> getAsync(String path, Map<String, Object> parameters) throws AlgoliaRuntimeException {
    return this.getAsync(path, parameters, null);
  }

  public CompletableFuture<Object> getAsync(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getAsync(path, null, requestOptions);
  }

  public CompletableFuture<Object> getAsync(String path) throws AlgoliaRuntimeException {
    return this.getAsync(path, null, null);
  }

  /**
   * Returns the average click position. The endpoint returns a value for the complete given time
   * range, as well as a value per day.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetAverageClickPositionResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetAverageClickPositionResponse getAverageClickPosition(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getAverageClickPositionAsync(index, startDate, endDate, tags, requestOptions));
  }

  public GetAverageClickPositionResponse getAverageClickPosition(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getAverageClickPosition(index, startDate, endDate, tags, null);
  }

  public GetAverageClickPositionResponse getAverageClickPosition(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getAverageClickPosition(index, null, null, null, requestOptions);
  }

  public GetAverageClickPositionResponse getAverageClickPosition(String index) throws AlgoliaRuntimeException {
    return this.getAverageClickPosition(index, null, null, null, null);
  }

  /**
   * (asynchronously) Returns the average click position. The endpoint returns a value for the
   * complete given time range, as well as a value per day.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetAverageClickPositionResponse> getAverageClickPositionAsync(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getAverageClickPosition(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/clicks/averageClickPosition";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetAverageClickPositionResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetAverageClickPositionResponse> getAverageClickPositionAsync(
    String index,
    String startDate,
    String endDate,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getAverageClickPositionAsync(index, startDate, endDate, tags, null);
  }

  public CompletableFuture<GetAverageClickPositionResponse> getAverageClickPositionAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getAverageClickPositionAsync(index, null, null, null, requestOptions);
  }

  public CompletableFuture<GetAverageClickPositionResponse> getAverageClickPositionAsync(String index) throws AlgoliaRuntimeException {
    return this.getAverageClickPositionAsync(index, null, null, null, null);
  }

  /**
   * Returns the distribution of clicks per range of positions. If the groups all have a count of 0,
   * it means Algolia didn’t receive any click events for the queries with the clickAnalytics search
   * parameter set to true. The count is 0 until Algolia receives at least one click event.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetClickPositionsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetClickPositionsResponse getClickPositions(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getClickPositionsAsync(index, startDate, endDate, tags, requestOptions));
  }

  public GetClickPositionsResponse getClickPositions(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getClickPositions(index, startDate, endDate, tags, null);
  }

  public GetClickPositionsResponse getClickPositions(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getClickPositions(index, null, null, null, requestOptions);
  }

  public GetClickPositionsResponse getClickPositions(String index) throws AlgoliaRuntimeException {
    return this.getClickPositions(index, null, null, null, null);
  }

  /**
   * (asynchronously) Returns the distribution of clicks per range of positions. If the groups all
   * have a count of 0, it means Algolia didn’t receive any click events for the queries with the
   * clickAnalytics search parameter set to true. The count is 0 until Algolia receives at least one
   * click event.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetClickPositionsResponse> getClickPositionsAsync(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getClickPositions(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/clicks/positions";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetClickPositionsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetClickPositionsResponse> getClickPositionsAsync(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getClickPositionsAsync(index, startDate, endDate, tags, null);
  }

  public CompletableFuture<GetClickPositionsResponse> getClickPositionsAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getClickPositionsAsync(index, null, null, null, requestOptions);
  }

  public CompletableFuture<GetClickPositionsResponse> getClickPositionsAsync(String index) throws AlgoliaRuntimeException {
    return this.getClickPositionsAsync(index, null, null, null, null);
  }

  /**
   * Returns a click-through rate (CTR). The endpoint returns a value for the complete given time
   * range, as well as a value per day. It also returns the count of clicks and searches used to
   * compute the rates.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetClickThroughRateResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetClickThroughRateResponse getClickThroughRate(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getClickThroughRateAsync(index, startDate, endDate, tags, requestOptions));
  }

  public GetClickThroughRateResponse getClickThroughRate(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getClickThroughRate(index, startDate, endDate, tags, null);
  }

  public GetClickThroughRateResponse getClickThroughRate(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getClickThroughRate(index, null, null, null, requestOptions);
  }

  public GetClickThroughRateResponse getClickThroughRate(String index) throws AlgoliaRuntimeException {
    return this.getClickThroughRate(index, null, null, null, null);
  }

  /**
   * (asynchronously) Returns a click-through rate (CTR). The endpoint returns a value for the
   * complete given time range, as well as a value per day. It also returns the count of clicks and
   * searches used to compute the rates.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetClickThroughRateResponse> getClickThroughRateAsync(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getClickThroughRate(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/clicks/clickThroughRate";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetClickThroughRateResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetClickThroughRateResponse> getClickThroughRateAsync(
    String index,
    String startDate,
    String endDate,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getClickThroughRateAsync(index, startDate, endDate, tags, null);
  }

  public CompletableFuture<GetClickThroughRateResponse> getClickThroughRateAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getClickThroughRateAsync(index, null, null, null, requestOptions);
  }

  public CompletableFuture<GetClickThroughRateResponse> getClickThroughRateAsync(String index) throws AlgoliaRuntimeException {
    return this.getClickThroughRateAsync(index, null, null, null, null);
  }

  /**
   * Returns a conversion rate (CR). The endpoint returns a value for the complete given time range,
   * as well as a value per day. It also returns the count of conversion and searches used to
   * compute the rates.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetConversationRateResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetConversationRateResponse getConversationRate(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getConversationRateAsync(index, startDate, endDate, tags, requestOptions));
  }

  public GetConversationRateResponse getConversationRate(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getConversationRate(index, startDate, endDate, tags, null);
  }

  public GetConversationRateResponse getConversationRate(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getConversationRate(index, null, null, null, requestOptions);
  }

  public GetConversationRateResponse getConversationRate(String index) throws AlgoliaRuntimeException {
    return this.getConversationRate(index, null, null, null, null);
  }

  /**
   * (asynchronously) Returns a conversion rate (CR). The endpoint returns a value for the complete
   * given time range, as well as a value per day. It also returns the count of conversion and
   * searches used to compute the rates.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetConversationRateResponse> getConversationRateAsync(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getConversationRate(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/conversions/conversionRate";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetConversationRateResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetConversationRateResponse> getConversationRateAsync(
    String index,
    String startDate,
    String endDate,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getConversationRateAsync(index, startDate, endDate, tags, null);
  }

  public CompletableFuture<GetConversationRateResponse> getConversationRateAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getConversationRateAsync(index, null, null, null, requestOptions);
  }

  public CompletableFuture<GetConversationRateResponse> getConversationRateAsync(String index) throws AlgoliaRuntimeException {
    return this.getConversationRateAsync(index, null, null, null, null);
  }

  /**
   * Returns the rate at which searches didn't lead to any clicks. The endpoint returns a value for
   * the complete given time range, as well as a value per day. It also returns the count of
   * searches and searches without clicks.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetNoClickRateResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetNoClickRateResponse getNoClickRate(String index, String startDate, String endDate, String tags, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getNoClickRateAsync(index, startDate, endDate, tags, requestOptions));
  }

  public GetNoClickRateResponse getNoClickRate(String index, String startDate, String endDate, String tags) throws AlgoliaRuntimeException {
    return this.getNoClickRate(index, startDate, endDate, tags, null);
  }

  public GetNoClickRateResponse getNoClickRate(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getNoClickRate(index, null, null, null, requestOptions);
  }

  public GetNoClickRateResponse getNoClickRate(String index) throws AlgoliaRuntimeException {
    return this.getNoClickRate(index, null, null, null, null);
  }

  /**
   * (asynchronously) Returns the rate at which searches didn&#39;t lead to any clicks. The endpoint
   * returns a value for the complete given time range, as well as a value per day. It also returns
   * the count of searches and searches without clicks.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetNoClickRateResponse> getNoClickRateAsync(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getNoClickRate(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/searches/noClickRate";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetNoClickRateResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetNoClickRateResponse> getNoClickRateAsync(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getNoClickRateAsync(index, startDate, endDate, tags, null);
  }

  public CompletableFuture<GetNoClickRateResponse> getNoClickRateAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getNoClickRateAsync(index, null, null, null, requestOptions);
  }

  public CompletableFuture<GetNoClickRateResponse> getNoClickRateAsync(String index) throws AlgoliaRuntimeException {
    return this.getNoClickRateAsync(index, null, null, null, null);
  }

  /**
   * Returns the rate at which searches didn't return any results. The endpoint returns a value for
   * the complete given time range, as well as a value per day. It also returns the count of
   * searches and searches without results used to compute the rates.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetNoResultsRateResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetNoResultsRateResponse getNoResultsRate(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getNoResultsRateAsync(index, startDate, endDate, tags, requestOptions));
  }

  public GetNoResultsRateResponse getNoResultsRate(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getNoResultsRate(index, startDate, endDate, tags, null);
  }

  public GetNoResultsRateResponse getNoResultsRate(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getNoResultsRate(index, null, null, null, requestOptions);
  }

  public GetNoResultsRateResponse getNoResultsRate(String index) throws AlgoliaRuntimeException {
    return this.getNoResultsRate(index, null, null, null, null);
  }

  /**
   * (asynchronously) Returns the rate at which searches didn&#39;t return any results. The endpoint
   * returns a value for the complete given time range, as well as a value per day. It also returns
   * the count of searches and searches without results used to compute the rates.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetNoResultsRateResponse> getNoResultsRateAsync(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getNoResultsRate(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/searches/noResultRate";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetNoResultsRateResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetNoResultsRateResponse> getNoResultsRateAsync(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getNoResultsRateAsync(index, startDate, endDate, tags, null);
  }

  public CompletableFuture<GetNoResultsRateResponse> getNoResultsRateAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getNoResultsRateAsync(index, null, null, null, requestOptions);
  }

  public CompletableFuture<GetNoResultsRateResponse> getNoResultsRateAsync(String index) throws AlgoliaRuntimeException {
    return this.getNoResultsRateAsync(index, null, null, null, null);
  }

  /**
   * Returns the number of searches across the given time range. The endpoint returns a value for
   * the complete given time range, as well as a value per day.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetSearchesCountResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetSearchesCountResponse getSearchesCount(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getSearchesCountAsync(index, startDate, endDate, tags, requestOptions));
  }

  public GetSearchesCountResponse getSearchesCount(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getSearchesCount(index, startDate, endDate, tags, null);
  }

  public GetSearchesCountResponse getSearchesCount(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getSearchesCount(index, null, null, null, requestOptions);
  }

  public GetSearchesCountResponse getSearchesCount(String index) throws AlgoliaRuntimeException {
    return this.getSearchesCount(index, null, null, null, null);
  }

  /**
   * (asynchronously) Returns the number of searches across the given time range. The endpoint
   * returns a value for the complete given time range, as well as a value per day.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetSearchesCountResponse> getSearchesCountAsync(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getSearchesCount(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/searches/count";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetSearchesCountResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetSearchesCountResponse> getSearchesCountAsync(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getSearchesCountAsync(index, startDate, endDate, tags, null);
  }

  public CompletableFuture<GetSearchesCountResponse> getSearchesCountAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getSearchesCountAsync(index, null, null, null, requestOptions);
  }

  public CompletableFuture<GetSearchesCountResponse> getSearchesCountAsync(String index) throws AlgoliaRuntimeException {
    return this.getSearchesCountAsync(index, null, null, null, null);
  }

  /**
   * Returns top searches that didn't lead to any clicks. Limited to the 1000 most frequent ones.
   * For each search, also returns the average number of found hits.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetSearchesNoClicksResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetSearchesNoClicksResponse getSearchesNoClicks(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getSearchesNoClicksAsync(index, startDate, endDate, limit, offset, tags, requestOptions));
  }

  public GetSearchesNoClicksResponse getSearchesNoClicks(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getSearchesNoClicks(index, startDate, endDate, limit, offset, tags, null);
  }

  public GetSearchesNoClicksResponse getSearchesNoClicks(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getSearchesNoClicks(index, null, null, null, null, null, requestOptions);
  }

  public GetSearchesNoClicksResponse getSearchesNoClicks(String index) throws AlgoliaRuntimeException {
    return this.getSearchesNoClicks(index, null, null, null, null, null, null);
  }

  /**
   * (asynchronously) Returns top searches that didn&#39;t lead to any clicks. Limited to the 1000
   * most frequent ones. For each search, also returns the average number of found hits.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetSearchesNoClicksResponse> getSearchesNoClicksAsync(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getSearchesNoClicks(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/searches/noClicks";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (limit != null) {
      queryParameters.put("limit", parameterToString(limit));
    }

    if (offset != null) {
      queryParameters.put("offset", parameterToString(offset));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetSearchesNoClicksResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetSearchesNoClicksResponse> getSearchesNoClicksAsync(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getSearchesNoClicksAsync(index, startDate, endDate, limit, offset, tags, null);
  }

  public CompletableFuture<GetSearchesNoClicksResponse> getSearchesNoClicksAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getSearchesNoClicksAsync(index, null, null, null, null, null, requestOptions);
  }

  public CompletableFuture<GetSearchesNoClicksResponse> getSearchesNoClicksAsync(String index) throws AlgoliaRuntimeException {
    return this.getSearchesNoClicksAsync(index, null, null, null, null, null, null);
  }

  /**
   * Returns top searches that didn't return any results. Limited to the 1000 most frequent ones.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetSearchesNoResultsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetSearchesNoResultsResponse getSearchesNoResults(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getSearchesNoResultsAsync(index, startDate, endDate, limit, offset, tags, requestOptions));
  }

  public GetSearchesNoResultsResponse getSearchesNoResults(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getSearchesNoResults(index, startDate, endDate, limit, offset, tags, null);
  }

  public GetSearchesNoResultsResponse getSearchesNoResults(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getSearchesNoResults(index, null, null, null, null, null, requestOptions);
  }

  public GetSearchesNoResultsResponse getSearchesNoResults(String index) throws AlgoliaRuntimeException {
    return this.getSearchesNoResults(index, null, null, null, null, null, null);
  }

  /**
   * (asynchronously) Returns top searches that didn&#39;t return any results. Limited to the 1000
   * most frequent ones.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetSearchesNoResultsResponse> getSearchesNoResultsAsync(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getSearchesNoResults(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/searches/noResults";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (limit != null) {
      queryParameters.put("limit", parameterToString(limit));
    }

    if (offset != null) {
      queryParameters.put("offset", parameterToString(offset));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetSearchesNoResultsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetSearchesNoResultsResponse> getSearchesNoResultsAsync(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getSearchesNoResultsAsync(index, startDate, endDate, limit, offset, tags, null);
  }

  public CompletableFuture<GetSearchesNoResultsResponse> getSearchesNoResultsAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getSearchesNoResultsAsync(index, null, null, null, null, null, requestOptions);
  }

  public CompletableFuture<GetSearchesNoResultsResponse> getSearchesNoResultsAsync(String index) throws AlgoliaRuntimeException {
    return this.getSearchesNoResultsAsync(index, null, null, null, null, null, null);
  }

  /**
   * Returns the latest update time of the analytics API for a given index. If the index has been
   * recently created and/or no search has been performed yet the updated time will be null.
   *
   * @param index The index name to target. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetStatusResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetStatusResponse getStatus(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getStatusAsync(index, requestOptions));
  }

  public GetStatusResponse getStatus(String index) throws AlgoliaRuntimeException {
    return this.getStatus(index, null);
  }

  /**
   * (asynchronously) Returns the latest update time of the analytics API for a given index. If the
   * index has been recently created and/or no search has been performed yet the updated time will
   * be null.
   *
   * @param index The index name to target. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetStatusResponse> getStatusAsync(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getStatus(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/status";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetStatusResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetStatusResponse> getStatusAsync(String index) throws AlgoliaRuntimeException {
    return this.getStatusAsync(index, null);
  }

  /**
   * Returns top countries. Limited to the 1000 most frequent ones.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetTopCountriesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetTopCountriesResponse getTopCountries(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getTopCountriesAsync(index, startDate, endDate, limit, offset, tags, requestOptions));
  }

  public GetTopCountriesResponse getTopCountries(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopCountries(index, startDate, endDate, limit, offset, tags, null);
  }

  public GetTopCountriesResponse getTopCountries(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getTopCountries(index, null, null, null, null, null, requestOptions);
  }

  public GetTopCountriesResponse getTopCountries(String index) throws AlgoliaRuntimeException {
    return this.getTopCountries(index, null, null, null, null, null, null);
  }

  /**
   * (asynchronously) Returns top countries. Limited to the 1000 most frequent ones.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetTopCountriesResponse> getTopCountriesAsync(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getTopCountries(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/countries";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (limit != null) {
      queryParameters.put("limit", parameterToString(limit));
    }

    if (offset != null) {
      queryParameters.put("offset", parameterToString(offset));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetTopCountriesResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetTopCountriesResponse> getTopCountriesAsync(
    String index,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopCountriesAsync(index, startDate, endDate, limit, offset, tags, null);
  }

  public CompletableFuture<GetTopCountriesResponse> getTopCountriesAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getTopCountriesAsync(index, null, null, null, null, null, requestOptions);
  }

  public CompletableFuture<GetTopCountriesResponse> getTopCountriesAsync(String index) throws AlgoliaRuntimeException {
    return this.getTopCountriesAsync(index, null, null, null, null, null, null);
  }

  /**
   * Returns top filter attributes. Limited to the 1000 most used filters.
   *
   * @param index The index name to target. (required)
   * @param search The query term to search for. Must match the exact user input. (optional)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetTopFilterAttributesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetTopFilterAttributesResponse getTopFilterAttributes(
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getTopFilterAttributesAsync(index, search, startDate, endDate, limit, offset, tags, requestOptions));
  }

  public GetTopFilterAttributesResponse getTopFilterAttributes(
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopFilterAttributes(index, search, startDate, endDate, limit, offset, tags, null);
  }

  public GetTopFilterAttributesResponse getTopFilterAttributes(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getTopFilterAttributes(index, null, null, null, null, null, null, requestOptions);
  }

  public GetTopFilterAttributesResponse getTopFilterAttributes(String index) throws AlgoliaRuntimeException {
    return this.getTopFilterAttributes(index, null, null, null, null, null, null, null);
  }

  /**
   * (asynchronously) Returns top filter attributes. Limited to the 1000 most used filters.
   *
   * @param index The index name to target. (required)
   * @param search The query term to search for. Must match the exact user input. (optional)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetTopFilterAttributesResponse> getTopFilterAttributesAsync(
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getTopFilterAttributes(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/filters";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (search != null) {
      queryParameters.put("search", parameterToString(search));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (limit != null) {
      queryParameters.put("limit", parameterToString(limit));
    }

    if (offset != null) {
      queryParameters.put("offset", parameterToString(offset));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetTopFilterAttributesResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetTopFilterAttributesResponse> getTopFilterAttributesAsync(
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopFilterAttributesAsync(index, search, startDate, endDate, limit, offset, tags, null);
  }

  public CompletableFuture<GetTopFilterAttributesResponse> getTopFilterAttributesAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getTopFilterAttributesAsync(index, null, null, null, null, null, null, requestOptions);
  }

  public CompletableFuture<GetTopFilterAttributesResponse> getTopFilterAttributesAsync(String index) throws AlgoliaRuntimeException {
    return this.getTopFilterAttributesAsync(index, null, null, null, null, null, null, null);
  }

  /**
   * Returns top filters for the given attribute. Limited to the 1000 most used filters.
   *
   * @param attribute The exact name of the attribute. (required)
   * @param index The index name to target. (required)
   * @param search The query term to search for. Must match the exact user input. (optional)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetTopFilterForAttributeResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetTopFilterForAttributeResponse getTopFilterForAttribute(
    String attribute,
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getTopFilterForAttributeAsync(attribute, index, search, startDate, endDate, limit, offset, tags, requestOptions)
    );
  }

  public GetTopFilterForAttributeResponse getTopFilterForAttribute(
    String attribute,
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopFilterForAttribute(attribute, index, search, startDate, endDate, limit, offset, tags, null);
  }

  public GetTopFilterForAttributeResponse getTopFilterForAttribute(String attribute, String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getTopFilterForAttribute(attribute, index, null, null, null, null, null, null, requestOptions);
  }

  public GetTopFilterForAttributeResponse getTopFilterForAttribute(String attribute, String index) throws AlgoliaRuntimeException {
    return this.getTopFilterForAttribute(attribute, index, null, null, null, null, null, null, null);
  }

  /**
   * (asynchronously) Returns top filters for the given attribute. Limited to the 1000 most used
   * filters.
   *
   * @param attribute The exact name of the attribute. (required)
   * @param index The index name to target. (required)
   * @param search The query term to search for. Must match the exact user input. (optional)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetTopFilterForAttributeResponse> getTopFilterForAttributeAsync(
    String attribute,
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (attribute == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'attribute' when calling" + " getTopFilterForAttribute(Async)");
    }

    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getTopFilterForAttribute(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/filters/{attribute}".replaceAll("\\{attribute\\}", this.escapeString(attribute.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (search != null) {
      queryParameters.put("search", parameterToString(search));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (limit != null) {
      queryParameters.put("limit", parameterToString(limit));
    }

    if (offset != null) {
      queryParameters.put("offset", parameterToString(offset));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetTopFilterForAttributeResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetTopFilterForAttributeResponse> getTopFilterForAttributeAsync(
    String attribute,
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopFilterForAttributeAsync(attribute, index, search, startDate, endDate, limit, offset, tags, null);
  }

  public CompletableFuture<GetTopFilterForAttributeResponse> getTopFilterForAttributeAsync(
    String attribute,
    String index,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.getTopFilterForAttributeAsync(attribute, index, null, null, null, null, null, null, requestOptions);
  }

  public CompletableFuture<GetTopFilterForAttributeResponse> getTopFilterForAttributeAsync(String attribute, String index)
    throws AlgoliaRuntimeException {
    return this.getTopFilterForAttributeAsync(attribute, index, null, null, null, null, null, null, null);
  }

  /**
   * Returns top filters with no results. Limited to the 1000 most used filters.
   *
   * @param index The index name to target. (required)
   * @param search The query term to search for. Must match the exact user input. (optional)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetTopFiltersNoResultsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetTopFiltersNoResultsResponse getTopFiltersNoResults(
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getTopFiltersNoResultsAsync(index, search, startDate, endDate, limit, offset, tags, requestOptions));
  }

  public GetTopFiltersNoResultsResponse getTopFiltersNoResults(
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopFiltersNoResults(index, search, startDate, endDate, limit, offset, tags, null);
  }

  public GetTopFiltersNoResultsResponse getTopFiltersNoResults(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getTopFiltersNoResults(index, null, null, null, null, null, null, requestOptions);
  }

  public GetTopFiltersNoResultsResponse getTopFiltersNoResults(String index) throws AlgoliaRuntimeException {
    return this.getTopFiltersNoResults(index, null, null, null, null, null, null, null);
  }

  /**
   * (asynchronously) Returns top filters with no results. Limited to the 1000 most used filters.
   *
   * @param index The index name to target. (required)
   * @param search The query term to search for. Must match the exact user input. (optional)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetTopFiltersNoResultsResponse> getTopFiltersNoResultsAsync(
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getTopFiltersNoResults(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/filters/noResults";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (search != null) {
      queryParameters.put("search", parameterToString(search));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (limit != null) {
      queryParameters.put("limit", parameterToString(limit));
    }

    if (offset != null) {
      queryParameters.put("offset", parameterToString(offset));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetTopFiltersNoResultsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetTopFiltersNoResultsResponse> getTopFiltersNoResultsAsync(
    String index,
    String search,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopFiltersNoResultsAsync(index, search, startDate, endDate, limit, offset, tags, null);
  }

  public CompletableFuture<GetTopFiltersNoResultsResponse> getTopFiltersNoResultsAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getTopFiltersNoResultsAsync(index, null, null, null, null, null, null, requestOptions);
  }

  public CompletableFuture<GetTopFiltersNoResultsResponse> getTopFiltersNoResultsAsync(String index) throws AlgoliaRuntimeException {
    return this.getTopFiltersNoResultsAsync(index, null, null, null, null, null, null, null);
  }

  /**
   * Returns top hits. Limited to the 1000 most frequent ones.
   *
   * @param index The index name to target. (required)
   * @param search The query term to search for. Must match the exact user input. (optional)
   * @param clickAnalytics Whether to include the click-through and conversion rates for a search.
   *     (optional, default to false)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetTopHitsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetTopHitsResponse getTopHits(
    String index,
    String search,
    Boolean clickAnalytics,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getTopHitsAsync(index, search, clickAnalytics, startDate, endDate, limit, offset, tags, requestOptions));
  }

  public GetTopHitsResponse getTopHits(
    String index,
    String search,
    Boolean clickAnalytics,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopHits(index, search, clickAnalytics, startDate, endDate, limit, offset, tags, null);
  }

  public GetTopHitsResponse getTopHits(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getTopHits(index, null, null, null, null, null, null, null, requestOptions);
  }

  public GetTopHitsResponse getTopHits(String index) throws AlgoliaRuntimeException {
    return this.getTopHits(index, null, null, null, null, null, null, null, null);
  }

  /**
   * (asynchronously) Returns top hits. Limited to the 1000 most frequent ones.
   *
   * @param index The index name to target. (required)
   * @param search The query term to search for. Must match the exact user input. (optional)
   * @param clickAnalytics Whether to include the click-through and conversion rates for a search.
   *     (optional, default to false)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetTopHitsResponse> getTopHitsAsync(
    String index,
    String search,
    Boolean clickAnalytics,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getTopHits(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/hits";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (search != null) {
      queryParameters.put("search", parameterToString(search));
    }

    if (clickAnalytics != null) {
      queryParameters.put("clickAnalytics", parameterToString(clickAnalytics));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (limit != null) {
      queryParameters.put("limit", parameterToString(limit));
    }

    if (offset != null) {
      queryParameters.put("offset", parameterToString(offset));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetTopHitsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetTopHitsResponse> getTopHitsAsync(
    String index,
    String search,
    Boolean clickAnalytics,
    String startDate,
    String endDate,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopHitsAsync(index, search, clickAnalytics, startDate, endDate, limit, offset, tags, null);
  }

  public CompletableFuture<GetTopHitsResponse> getTopHitsAsync(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getTopHitsAsync(index, null, null, null, null, null, null, null, requestOptions);
  }

  public CompletableFuture<GetTopHitsResponse> getTopHitsAsync(String index) throws AlgoliaRuntimeException {
    return this.getTopHitsAsync(index, null, null, null, null, null, null, null, null);
  }

  /**
   * Returns top searches. Limited to the 1000 most frequent ones. For each search, also returns the
   * average number of hits returned.
   *
   * @param index The index name to target. (required)
   * @param clickAnalytics Whether to include the click-through and conversion rates for a search.
   *     (optional, default to false)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param orderBy Reorder the results. (optional, default to searchCount)
   * @param direction The sorting of the result. (optional, default to asc)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetTopSearchesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetTopSearchesResponse getTopSearches(
    String index,
    Boolean clickAnalytics,
    String startDate,
    String endDate,
    OrderBy orderBy,
    Direction direction,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getTopSearchesAsync(index, clickAnalytics, startDate, endDate, orderBy, direction, limit, offset, tags, requestOptions)
    );
  }

  public GetTopSearchesResponse getTopSearches(
    String index,
    Boolean clickAnalytics,
    String startDate,
    String endDate,
    OrderBy orderBy,
    Direction direction,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopSearches(index, clickAnalytics, startDate, endDate, orderBy, direction, limit, offset, tags, null);
  }

  public GetTopSearchesResponse getTopSearches(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getTopSearches(index, null, null, null, null, null, null, null, null, requestOptions);
  }

  public GetTopSearchesResponse getTopSearches(String index) throws AlgoliaRuntimeException {
    return this.getTopSearches(index, null, null, null, null, null, null, null, null, null);
  }

  /**
   * (asynchronously) Returns top searches. Limited to the 1000 most frequent ones. For each search,
   * also returns the average number of hits returned.
   *
   * @param index The index name to target. (required)
   * @param clickAnalytics Whether to include the click-through and conversion rates for a search.
   *     (optional, default to false)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param orderBy Reorder the results. (optional, default to searchCount)
   * @param direction The sorting of the result. (optional, default to asc)
   * @param limit Number of records to return. Limit is the size of the page. (optional, default to
   *     10)
   * @param offset Position of the starting record. Used for paging. 0 is the first record.
   *     (optional, default to 0)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetTopSearchesResponse> getTopSearchesAsync(
    String index,
    Boolean clickAnalytics,
    String startDate,
    String endDate,
    OrderBy orderBy,
    Direction direction,
    Integer limit,
    Integer offset,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getTopSearches(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/searches";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (clickAnalytics != null) {
      queryParameters.put("clickAnalytics", parameterToString(clickAnalytics));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (orderBy != null) {
      queryParameters.put("orderBy", parameterToString(orderBy));
    }

    if (direction != null) {
      queryParameters.put("direction", parameterToString(direction));
    }

    if (limit != null) {
      queryParameters.put("limit", parameterToString(limit));
    }

    if (offset != null) {
      queryParameters.put("offset", parameterToString(offset));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetTopSearchesResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetTopSearchesResponse> getTopSearchesAsync(
    String index,
    Boolean clickAnalytics,
    String startDate,
    String endDate,
    OrderBy orderBy,
    Direction direction,
    Integer limit,
    Integer offset,
    String tags
  ) throws AlgoliaRuntimeException {
    return this.getTopSearchesAsync(index, clickAnalytics, startDate, endDate, orderBy, direction, limit, offset, tags, null);
  }

  public CompletableFuture<GetTopSearchesResponse> getTopSearchesAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getTopSearchesAsync(index, null, null, null, null, null, null, null, null, requestOptions);
  }

  public CompletableFuture<GetTopSearchesResponse> getTopSearchesAsync(String index) throws AlgoliaRuntimeException {
    return this.getTopSearchesAsync(index, null, null, null, null, null, null, null, null, null);
  }

  /**
   * Returns the distinct count of users across the given time range. The endpoint returns a value
   * for the complete given time range, as well as a value per day.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetUsersCountResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetUsersCountResponse getUsersCount(String index, String startDate, String endDate, String tags, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getUsersCountAsync(index, startDate, endDate, tags, requestOptions));
  }

  public GetUsersCountResponse getUsersCount(String index, String startDate, String endDate, String tags) throws AlgoliaRuntimeException {
    return this.getUsersCount(index, startDate, endDate, tags, null);
  }

  public GetUsersCountResponse getUsersCount(String index, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.getUsersCount(index, null, null, null, requestOptions);
  }

  public GetUsersCountResponse getUsersCount(String index) throws AlgoliaRuntimeException {
    return this.getUsersCount(index, null, null, null, null);
  }

  /**
   * (asynchronously) Returns the distinct count of users across the given time range. The endpoint
   * returns a value for the complete given time range, as well as a value per day.
   *
   * @param index The index name to target. (required)
   * @param startDate The lower bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param endDate The upper bound timestamp (a date, a string like \"2006-01-02\") of the period
   *     to analyze. (optional)
   * @param tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags
   *     set at search time. Multiple tags can be combined with the operators OR and AND. If a tag
   *     contains characters like spaces or parentheses, it should be URL encoded. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetUsersCountResponse> getUsersCountAsync(
    String index,
    String startDate,
    String endDate,
    String tags,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (index == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'index' when calling getUsersCount(Async)");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/2/users/count";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (index != null) {
      queryParameters.put("index", parameterToString(index));
    }

    if (startDate != null) {
      queryParameters.put("startDate", parameterToString(startDate));
    }

    if (endDate != null) {
      queryParameters.put("endDate", parameterToString(endDate));
    }

    if (tags != null) {
      queryParameters.put("tags", parameterToString(tags));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    Type returnType = new TypeToken<GetUsersCountResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetUsersCountResponse> getUsersCountAsync(String index, String startDate, String endDate, String tags)
    throws AlgoliaRuntimeException {
    return this.getUsersCountAsync(index, startDate, endDate, tags, null);
  }

  public CompletableFuture<GetUsersCountResponse> getUsersCountAsync(String index, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getUsersCountAsync(index, null, null, null, requestOptions);
  }

  public CompletableFuture<GetUsersCountResponse> getUsersCountAsync(String index) throws AlgoliaRuntimeException {
    return this.getUsersCountAsync(index, null, null, null, null);
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

  public Object post(String path, Map<String, Object> parameters, Object body) throws AlgoliaRuntimeException {
    return this.post(path, parameters, body, null);
  }

  public Object post(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
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
  public CompletableFuture<Object> postAsync(String path, Map<String, Object> parameters, Object body, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'path' when calling post(Async)");
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

  public CompletableFuture<Object> postAsync(String path, Map<String, Object> parameters, Object body) throws AlgoliaRuntimeException {
    return this.postAsync(path, parameters, body, null);
  }

  public CompletableFuture<Object> postAsync(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.postAsync(path, null, null, requestOptions);
  }

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

  public Object put(String path, Map<String, Object> parameters, Object body) throws AlgoliaRuntimeException {
    return this.put(path, parameters, body, null);
  }

  public Object put(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
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
  public CompletableFuture<Object> putAsync(String path, Map<String, Object> parameters, Object body, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (path == null) {
      throw new AlgoliaRuntimeException("Missing the required parameter 'path' when calling put(Async)");
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

  public CompletableFuture<Object> putAsync(String path, Map<String, Object> parameters, Object body) throws AlgoliaRuntimeException {
    return this.putAsync(path, parameters, body, null);
  }

  public CompletableFuture<Object> putAsync(String path, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.putAsync(path, null, null, requestOptions);
  }

  public CompletableFuture<Object> putAsync(String path) throws AlgoliaRuntimeException {
    return this.putAsync(path, null, null, null);
  }
}
