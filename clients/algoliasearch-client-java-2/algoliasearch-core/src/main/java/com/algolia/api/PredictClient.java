// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.predict.*;
import com.algolia.utils.*;
import com.algolia.utils.retry.CallType;
import com.algolia.utils.retry.StatefulHost;
import com.fasterxml.jackson.core.type.TypeReference;
import java.util.ArrayList;
import java.util.EnumSet;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.CompletableFuture;
import okhttp3.Call;

public class PredictClient extends ApiClient {

  private static final String[] allowedRegions = { "eu", "us" };

  public PredictClient(String appId, String apiKey, String region) {
    this(appId, apiKey, region, null);
  }

  public PredictClient(String appId, String apiKey, String region, ClientOptions options) {
    super(appId, apiKey, "Predict", "4.0.0-SNAPSHOT", options);
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
    if (region != null) {
      for (String allowed : allowedRegions) {
        if (allowed.equals(region)) {
          found = true;
          break;
        }
      }
    }

    if (region == null || !found) {
      throw new AlgoliaRuntimeException("`region` is required and must be one of the following: eu, us");
    }

    String url = "predict.{region}.algolia.com".replace("{region}", region);

    hosts.add(new StatefulHost(url, "https", EnumSet.of(CallType.READ, CallType.WRITE)));
    return hosts;
  }

  /**
   * Activate an existing model template. This action triggers the training and inference pipelines
   * for the selected model. The model is added with `modelStatus=pending`. If a model with the
   * exact same source & index already exists, the API endpoint returns an error.
   *
   * @param activateModelParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return ActivateModelInstanceResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public ActivateModelInstanceResponse activateModelInstance(ActivateModelParams activateModelParams, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(activateModelInstanceAsync(activateModelParams, requestOptions));
  }

  /**
   * Activate an existing model template. This action triggers the training and inference pipelines
   * for the selected model. The model is added with `modelStatus=pending`. If a model with the
   * exact same source & index already exists, the API endpoint returns an error.
   *
   * @param activateModelParams (required)
   * @return ActivateModelInstanceResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public ActivateModelInstanceResponse activateModelInstance(ActivateModelParams activateModelParams) throws AlgoliaRuntimeException {
    return this.activateModelInstance(activateModelParams, null);
  }

  /**
   * (asynchronously) Activate an existing model template. This action triggers the training and
   * inference pipelines for the selected model. The model is added with
   * &#x60;modelStatus&#x3D;pending&#x60;. If a model with the exact same source &amp; index already
   * exists, the API endpoint returns an error.
   *
   * @param activateModelParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<ActivateModelInstanceResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<ActivateModelInstanceResponse> activateModelInstanceAsync(
    ActivateModelParams activateModelParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (activateModelParams == null) {
      throw new AlgoliaRuntimeException("Parameter `activateModelParams` is required when calling `activateModelInstance`.");
    }

    Object bodyObj = activateModelParams;

    // create path and map variables
    String requestPath = "/1/predict/models";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "POST", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<ActivateModelInstanceResponse>() {});
  }

  /**
   * (asynchronously) Activate an existing model template. This action triggers the training and
   * inference pipelines for the selected model. The model is added with
   * &#x60;modelStatus&#x3D;pending&#x60;. If a model with the exact same source &amp; index already
   * exists, the API endpoint returns an error.
   *
   * @param activateModelParams (required)
   * @return CompletableFuture<ActivateModelInstanceResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<ActivateModelInstanceResponse> activateModelInstanceAsync(ActivateModelParams activateModelParams)
    throws AlgoliaRuntimeException {
    return this.activateModelInstanceAsync(activateModelParams, null);
  }

  /**
   * Create a new segment. All segments added by this endpoint will have a computed type. The
   * endpoint receives a filters parameter, with a syntax similar to filters for Rules.
   *
   * @param createSegmentParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CreateSegmentResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CreateSegmentResponse createSegment(CreateSegmentParams createSegmentParams, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(createSegmentAsync(createSegmentParams, requestOptions));
  }

  /**
   * Create a new segment. All segments added by this endpoint will have a computed type. The
   * endpoint receives a filters parameter, with a syntax similar to filters for Rules.
   *
   * @param createSegmentParams (required)
   * @return CreateSegmentResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CreateSegmentResponse createSegment(CreateSegmentParams createSegmentParams) throws AlgoliaRuntimeException {
    return this.createSegment(createSegmentParams, null);
  }

  /**
   * (asynchronously) Create a new segment. All segments added by this endpoint will have a computed
   * type. The endpoint receives a filters parameter, with a syntax similar to filters for Rules.
   *
   * @param createSegmentParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<CreateSegmentResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<CreateSegmentResponse> createSegmentAsync(
    CreateSegmentParams createSegmentParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (createSegmentParams == null) {
      throw new AlgoliaRuntimeException("Parameter `createSegmentParams` is required when calling `createSegment`.");
    }

    Object bodyObj = createSegmentParams;

    // create path and map variables
    String requestPath = "/1/segments";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "POST", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<CreateSegmentResponse>() {});
  }

  /**
   * (asynchronously) Create a new segment. All segments added by this endpoint will have a computed
   * type. The endpoint receives a filters parameter, with a syntax similar to filters for Rules.
   *
   * @param createSegmentParams (required)
   * @return CompletableFuture<CreateSegmentResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<CreateSegmentResponse> createSegmentAsync(CreateSegmentParams createSegmentParams)
    throws AlgoliaRuntimeException {
    return this.createSegmentAsync(createSegmentParams, null);
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
   * @throws AlgoliaRuntimeException If it fails to process the API call
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
    return this.executeAsync(call, new TypeReference<Object>() {});
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
   * Delete the model’s configuration, pipelines and generated predictions.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return DeleteModelInstanceResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public DeleteModelInstanceResponse deleteModelInstance(String modelID, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteModelInstanceAsync(modelID, requestOptions));
  }

  /**
   * Delete the model’s configuration, pipelines and generated predictions.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @return DeleteModelInstanceResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public DeleteModelInstanceResponse deleteModelInstance(String modelID) throws AlgoliaRuntimeException {
    return this.deleteModelInstance(modelID, null);
  }

  /**
   * (asynchronously) Delete the model’s configuration, pipelines and generated predictions.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<DeleteModelInstanceResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<DeleteModelInstanceResponse> deleteModelInstanceAsync(String modelID, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (modelID == null) {
      throw new AlgoliaRuntimeException("Parameter `modelID` is required when calling `deleteModelInstance`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/predict/models/{modelID}".replaceAll("\\{modelID\\}", this.escapeString(modelID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "DELETE", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<DeleteModelInstanceResponse>() {});
  }

  /**
   * (asynchronously) Delete the model’s configuration, pipelines and generated predictions.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @return CompletableFuture<DeleteModelInstanceResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<DeleteModelInstanceResponse> deleteModelInstanceAsync(String modelID) throws AlgoliaRuntimeException {
    return this.deleteModelInstanceAsync(modelID, null);
  }

  /**
   * Delete the segment’s configuration. User intents (predictions) from the segment are not
   * deleted. All segment types (computed or custom) can be deleted. When the query is successful,
   * the HTTP response is 200 OK and returns the date until which you can safely consider the data
   * as being deleted.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return DeleteSegmentResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public DeleteSegmentResponse deleteSegment(String segmentID, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteSegmentAsync(segmentID, requestOptions));
  }

  /**
   * Delete the segment’s configuration. User intents (predictions) from the segment are not
   * deleted. All segment types (computed or custom) can be deleted. When the query is successful,
   * the HTTP response is 200 OK and returns the date until which you can safely consider the data
   * as being deleted.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @return DeleteSegmentResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public DeleteSegmentResponse deleteSegment(String segmentID) throws AlgoliaRuntimeException {
    return this.deleteSegment(segmentID, null);
  }

  /**
   * (asynchronously) Delete the segment’s configuration. User intents (predictions) from the
   * segment are not deleted. All segment types (computed or custom) can be deleted. When the query
   * is successful, the HTTP response is 200 OK and returns the date until which you can safely
   * consider the data as being deleted.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<DeleteSegmentResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<DeleteSegmentResponse> deleteSegmentAsync(String segmentID, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (segmentID == null) {
      throw new AlgoliaRuntimeException("Parameter `segmentID` is required when calling `deleteSegment`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/segments/{segmentID}".replaceAll("\\{segmentID\\}", this.escapeString(segmentID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "DELETE", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<DeleteSegmentResponse>() {});
  }

  /**
   * (asynchronously) Delete the segment’s configuration. User intents (predictions) from the
   * segment are not deleted. All segment types (computed or custom) can be deleted. When the query
   * is successful, the HTTP response is 200 OK and returns the date until which you can safely
   * consider the data as being deleted.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @return CompletableFuture<DeleteSegmentResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<DeleteSegmentResponse> deleteSegmentAsync(String segmentID) throws AlgoliaRuntimeException {
    return this.deleteSegmentAsync(segmentID, null);
  }

  /**
   * Delete all data and predictions associated with an authenticated user (userID) or an anonymous
   * user (cookieID, sessionID).
   *
   * @param userID User ID for authenticated users or cookie ID for non-authenticated repeated users
   *     (visitors). (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return DeleteUserProfileResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public DeleteUserProfileResponse deleteUserProfile(String userID, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteUserProfileAsync(userID, requestOptions));
  }

  /**
   * Delete all data and predictions associated with an authenticated user (userID) or an anonymous
   * user (cookieID, sessionID).
   *
   * @param userID User ID for authenticated users or cookie ID for non-authenticated repeated users
   *     (visitors). (required)
   * @return DeleteUserProfileResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public DeleteUserProfileResponse deleteUserProfile(String userID) throws AlgoliaRuntimeException {
    return this.deleteUserProfile(userID, null);
  }

  /**
   * (asynchronously) Delete all data and predictions associated with an authenticated user (userID)
   * or an anonymous user (cookieID, sessionID).
   *
   * @param userID User ID for authenticated users or cookie ID for non-authenticated repeated users
   *     (visitors). (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<DeleteUserProfileResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<DeleteUserProfileResponse> deleteUserProfileAsync(String userID, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (userID == null) {
      throw new AlgoliaRuntimeException("Parameter `userID` is required when calling `deleteUserProfile`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/users/{userID}".replaceAll("\\{userID\\}", this.escapeString(userID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "DELETE", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<DeleteUserProfileResponse>() {});
  }

  /**
   * (asynchronously) Delete all data and predictions associated with an authenticated user (userID)
   * or an anonymous user (cookieID, sessionID).
   *
   * @param userID User ID for authenticated users or cookie ID for non-authenticated repeated users
   *     (visitors). (required)
   * @return CompletableFuture<DeleteUserProfileResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<DeleteUserProfileResponse> deleteUserProfileAsync(String userID) throws AlgoliaRuntimeException {
    return this.deleteUserProfileAsync(userID, null);
  }

  /**
   * Get the list of segments with their configuration.
   *
   * @param type The type of segments to fetch. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return List<Segment>
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<Segment> fetchAllSegments(SegmentType type, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(fetchAllSegmentsAsync(type, requestOptions));
  }

  /**
   * Get the list of segments with their configuration.
   *
   * @param type The type of segments to fetch. (optional)
   * @return List<Segment>
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<Segment> fetchAllSegments(SegmentType type) throws AlgoliaRuntimeException {
    return this.fetchAllSegments(type, null);
  }

  /**
   * Get the list of segments with their configuration.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return List<Segment>
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<Segment> fetchAllSegments(RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.fetchAllSegments(null, requestOptions);
  }

  /**
   * Get the list of segments with their configuration.
   *
   * @return List<Segment>
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<Segment> fetchAllSegments() throws AlgoliaRuntimeException {
    return this.fetchAllSegments(null, null);
  }

  /**
   * (asynchronously) Get the list of segments with their configuration.
   *
   * @param type The type of segments to fetch. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<List<Segment>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<Segment>> fetchAllSegmentsAsync(SegmentType type, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/segments";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (type != null) {
      queryParameters.put("type", parameterToString(type));
    }

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<List<Segment>>() {});
  }

  /**
   * (asynchronously) Get the list of segments with their configuration.
   *
   * @param type The type of segments to fetch. (optional)
   * @return CompletableFuture<List<Segment>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<Segment>> fetchAllSegmentsAsync(SegmentType type) throws AlgoliaRuntimeException {
    return this.fetchAllSegmentsAsync(type, null);
  }

  /**
   * (asynchronously) Get the list of segments with their configuration.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<List<Segment>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<Segment>> fetchAllSegmentsAsync(RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return this.fetchAllSegmentsAsync(null, requestOptions);
  }

  /**
   * (asynchronously) Get the list of segments with their configuration.
   *
   * @return CompletableFuture<List<Segment>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<Segment>> fetchAllSegmentsAsync() throws AlgoliaRuntimeException {
    return this.fetchAllSegmentsAsync(null, null);
  }

  /**
   * Get all users with predictions in the provided application.
   *
   * @param fetchAllUserProfilesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return FetchAllUserProfilesResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public FetchAllUserProfilesResponse fetchAllUserProfiles(
    FetchAllUserProfilesParams fetchAllUserProfilesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(fetchAllUserProfilesAsync(fetchAllUserProfilesParams, requestOptions));
  }

  /**
   * Get all users with predictions in the provided application.
   *
   * @param fetchAllUserProfilesParams (required)
   * @return FetchAllUserProfilesResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public FetchAllUserProfilesResponse fetchAllUserProfiles(FetchAllUserProfilesParams fetchAllUserProfilesParams)
    throws AlgoliaRuntimeException {
    return this.fetchAllUserProfiles(fetchAllUserProfilesParams, null);
  }

  /**
   * (asynchronously) Get all users with predictions in the provided application.
   *
   * @param fetchAllUserProfilesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<FetchAllUserProfilesResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<FetchAllUserProfilesResponse> fetchAllUserProfilesAsync(
    FetchAllUserProfilesParams fetchAllUserProfilesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (fetchAllUserProfilesParams == null) {
      throw new AlgoliaRuntimeException("Parameter `fetchAllUserProfilesParams` is required when calling" + " `fetchAllUserProfiles`.");
    }

    Object bodyObj = fetchAllUserProfilesParams;

    // create path and map variables
    String requestPath = "/1/users";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "POST", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<FetchAllUserProfilesResponse>() {});
  }

  /**
   * (asynchronously) Get all users with predictions in the provided application.
   *
   * @param fetchAllUserProfilesParams (required)
   * @return CompletableFuture<FetchAllUserProfilesResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<FetchAllUserProfilesResponse> fetchAllUserProfilesAsync(FetchAllUserProfilesParams fetchAllUserProfilesParams)
    throws AlgoliaRuntimeException {
    return this.fetchAllUserProfilesAsync(fetchAllUserProfilesParams, null);
  }

  /**
   * Get the segment configuration.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Segment
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Segment fetchSegment(String segmentID, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(fetchSegmentAsync(segmentID, requestOptions));
  }

  /**
   * Get the segment configuration.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @return Segment
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public Segment fetchSegment(String segmentID) throws AlgoliaRuntimeException {
    return this.fetchSegment(segmentID, null);
  }

  /**
   * (asynchronously) Get the segment configuration.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<Segment> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Segment> fetchSegmentAsync(String segmentID, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    if (segmentID == null) {
      throw new AlgoliaRuntimeException("Parameter `segmentID` is required when calling `fetchSegment`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/segments/{segmentID}".replaceAll("\\{segmentID\\}", this.escapeString(segmentID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<Segment>() {});
  }

  /**
   * (asynchronously) Get the segment configuration.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @return CompletableFuture<Segment> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<Segment> fetchSegmentAsync(String segmentID) throws AlgoliaRuntimeException {
    return this.fetchSegmentAsync(segmentID, null);
  }

  /**
   * Get predictions, properties (raw, computed or custom) and segments (computed or custom) for a
   * user profile.
   *
   * @param userID User ID for authenticated users or cookie ID for non-authenticated repeated users
   *     (visitors). (required)
   * @param params (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UserProfile
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public UserProfile fetchUserProfile(String userID, Params params, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(fetchUserProfileAsync(userID, params, requestOptions));
  }

  /**
   * Get predictions, properties (raw, computed or custom) and segments (computed or custom) for a
   * user profile.
   *
   * @param userID User ID for authenticated users or cookie ID for non-authenticated repeated users
   *     (visitors). (required)
   * @param params (required)
   * @return UserProfile
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public UserProfile fetchUserProfile(String userID, Params params) throws AlgoliaRuntimeException {
    return this.fetchUserProfile(userID, params, null);
  }

  /**
   * (asynchronously) Get predictions, properties (raw, computed or custom) and segments (computed
   * or custom) for a user profile.
   *
   * @param userID User ID for authenticated users or cookie ID for non-authenticated repeated users
   *     (visitors). (required)
   * @param params (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<UserProfile> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<UserProfile> fetchUserProfileAsync(String userID, Params params, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (userID == null) {
      throw new AlgoliaRuntimeException("Parameter `userID` is required when calling `fetchUserProfile`.");
    }

    if (params == null) {
      throw new AlgoliaRuntimeException("Parameter `params` is required when calling `fetchUserProfile`.");
    }

    Object bodyObj = params;

    // create path and map variables
    String requestPath = "/1/users/{userID}/fetch".replaceAll("\\{userID\\}", this.escapeString(userID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "POST", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<UserProfile>() {});
  }

  /**
   * (asynchronously) Get predictions, properties (raw, computed or custom) and segments (computed
   * or custom) for a user profile.
   *
   * @param userID User ID for authenticated users or cookie ID for non-authenticated repeated users
   *     (visitors). (required)
   * @param params (required)
   * @return CompletableFuture<UserProfile> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<UserProfile> fetchUserProfileAsync(String userID, Params params) throws AlgoliaRuntimeException {
    return this.fetchUserProfileAsync(userID, params, null);
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
   * @throws AlgoliaRuntimeException If it fails to process the API call
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
    return this.executeAsync(call, new TypeReference<Object>() {});
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
   * Get a list of all available model types. Each model type can be activated more than once, by
   * selecting a different data source.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return List<GetAvailableModelTypesResponseInner>
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<GetAvailableModelTypesResponseInner> getAvailableModelTypes(RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getAvailableModelTypesAsync(requestOptions));
  }

  /**
   * Get a list of all available model types. Each model type can be activated more than once, by
   * selecting a different data source.
   *
   * @return List<GetAvailableModelTypesResponseInner>
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<GetAvailableModelTypesResponseInner> getAvailableModelTypes() throws AlgoliaRuntimeException {
    return this.getAvailableModelTypes(null);
  }

  /**
   * (asynchronously) Get a list of all available model types. Each model type can be activated more
   * than once, by selecting a different data source.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<List<GetAvailableModelTypesResponseInner>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<GetAvailableModelTypesResponseInner>> getAvailableModelTypesAsync(RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/predict/modeltypes";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<List<GetAvailableModelTypesResponseInner>>() {});
  }

  /**
   * (asynchronously) Get a list of all available model types. Each model type can be activated more
   * than once, by selecting a different data source.
   *
   * @return CompletableFuture<List<GetAvailableModelTypesResponseInner>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<GetAvailableModelTypesResponseInner>> getAvailableModelTypesAsync() throws AlgoliaRuntimeException {
    return this.getAvailableModelTypesAsync(null);
  }

  /**
   * Get the configuration for a model that was activated.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return ModelInstance
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public ModelInstance getModelInstanceConfig(String modelID, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getModelInstanceConfigAsync(modelID, requestOptions));
  }

  /**
   * Get the configuration for a model that was activated.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @return ModelInstance
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public ModelInstance getModelInstanceConfig(String modelID) throws AlgoliaRuntimeException {
    return this.getModelInstanceConfig(modelID, null);
  }

  /**
   * (asynchronously) Get the configuration for a model that was activated.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<ModelInstance> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<ModelInstance> getModelInstanceConfigAsync(String modelID, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (modelID == null) {
      throw new AlgoliaRuntimeException("Parameter `modelID` is required when calling `getModelInstanceConfig`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/predict/models/{modelID}".replaceAll("\\{modelID\\}", this.escapeString(modelID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<ModelInstance>() {});
  }

  /**
   * (asynchronously) Get the configuration for a model that was activated.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @return CompletableFuture<ModelInstance> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<ModelInstance> getModelInstanceConfigAsync(String modelID) throws AlgoliaRuntimeException {
    return this.getModelInstanceConfigAsync(modelID, null);
  }

  /**
   * Get a list of all model instances.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return List<ModelInstance>
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<ModelInstance> getModelInstances(RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getModelInstancesAsync(requestOptions));
  }

  /**
   * Get a list of all model instances.
   *
   * @return List<ModelInstance>
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<ModelInstance> getModelInstances() throws AlgoliaRuntimeException {
    return this.getModelInstances(null);
  }

  /**
   * (asynchronously) Get a list of all model instances.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<List<ModelInstance>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<ModelInstance>> getModelInstancesAsync(RequestOptions requestOptions) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/predict/models";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<List<ModelInstance>>() {});
  }

  /**
   * (asynchronously) Get a list of all model instances.
   *
   * @return CompletableFuture<List<ModelInstance>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<ModelInstance>> getModelInstancesAsync() throws AlgoliaRuntimeException {
    return this.getModelInstancesAsync(null);
  }

  /**
   * Get the model instance’ training metrics.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return List<ModelMetrics>
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<ModelMetrics> getModelMetrics(String modelID, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getModelMetricsAsync(modelID, requestOptions));
  }

  /**
   * Get the model instance’ training metrics.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @return List<ModelMetrics>
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public List<ModelMetrics> getModelMetrics(String modelID) throws AlgoliaRuntimeException {
    return this.getModelMetrics(modelID, null);
  }

  /**
   * (asynchronously) Get the model instance’ training metrics.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<List<ModelMetrics>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<ModelMetrics>> getModelMetricsAsync(String modelID, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (modelID == null) {
      throw new AlgoliaRuntimeException("Parameter `modelID` is required when calling `getModelMetrics`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/predict/models/{modelID}/metrics".replaceAll("\\{modelID\\}", this.escapeString(modelID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<List<ModelMetrics>>() {});
  }

  /**
   * (asynchronously) Get the model instance’ training metrics.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @return CompletableFuture<List<ModelMetrics>> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<List<ModelMetrics>> getModelMetricsAsync(String modelID) throws AlgoliaRuntimeException {
    return this.getModelMetricsAsync(modelID, null);
  }

  /**
   * Get the profiles of users that belong to a segment.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param fetchAllUserProfilesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetSegmentUsersResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public GetSegmentUsersResponse getSegmentUsers(
    String segmentID,
    FetchAllUserProfilesParams fetchAllUserProfilesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getSegmentUsersAsync(segmentID, fetchAllUserProfilesParams, requestOptions));
  }

  /**
   * Get the profiles of users that belong to a segment.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param fetchAllUserProfilesParams (required)
   * @return GetSegmentUsersResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public GetSegmentUsersResponse getSegmentUsers(String segmentID, FetchAllUserProfilesParams fetchAllUserProfilesParams)
    throws AlgoliaRuntimeException {
    return this.getSegmentUsers(segmentID, fetchAllUserProfilesParams, null);
  }

  /**
   * (asynchronously) Get the profiles of users that belong to a segment.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param fetchAllUserProfilesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<GetSegmentUsersResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<GetSegmentUsersResponse> getSegmentUsersAsync(
    String segmentID,
    FetchAllUserProfilesParams fetchAllUserProfilesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (segmentID == null) {
      throw new AlgoliaRuntimeException("Parameter `segmentID` is required when calling `getSegmentUsers`.");
    }

    if (fetchAllUserProfilesParams == null) {
      throw new AlgoliaRuntimeException("Parameter `fetchAllUserProfilesParams` is required when calling `getSegmentUsers`.");
    }

    Object bodyObj = fetchAllUserProfilesParams;

    // create path and map variables
    String requestPath = "/1/segments/{segmentID}/users".replaceAll("\\{segmentID\\}", this.escapeString(segmentID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "POST", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<GetSegmentUsersResponse>() {});
  }

  /**
   * (asynchronously) Get the profiles of users that belong to a segment.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param fetchAllUserProfilesParams (required)
   * @return CompletableFuture<GetSegmentUsersResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<GetSegmentUsersResponse> getSegmentUsersAsync(
    String segmentID,
    FetchAllUserProfilesParams fetchAllUserProfilesParams
  ) throws AlgoliaRuntimeException {
    return this.getSegmentUsersAsync(segmentID, fetchAllUserProfilesParams, null);
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
   * @throws AlgoliaRuntimeException If it fails to process the API call
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

    Object bodyObj = body != null ? body : new Object();

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
    return this.executeAsync(call, new TypeReference<Object>() {});
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
   * @throws AlgoliaRuntimeException If it fails to process the API call
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

    Object bodyObj = body != null ? body : new Object();

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
    return this.executeAsync(call, new TypeReference<Object>() {});
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
   * Update a model’s configuration.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @param updateModelParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdateModelInstanceResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public UpdateModelInstanceResponse updateModelInstance(
    String modelID,
    UpdateModelParams updateModelParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(updateModelInstanceAsync(modelID, updateModelParams, requestOptions));
  }

  /**
   * Update a model’s configuration.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @param updateModelParams (required)
   * @return UpdateModelInstanceResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public UpdateModelInstanceResponse updateModelInstance(String modelID, UpdateModelParams updateModelParams)
    throws AlgoliaRuntimeException {
    return this.updateModelInstance(modelID, updateModelParams, null);
  }

  /**
   * (asynchronously) Update a model’s configuration.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @param updateModelParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<UpdateModelInstanceResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<UpdateModelInstanceResponse> updateModelInstanceAsync(
    String modelID,
    UpdateModelParams updateModelParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (modelID == null) {
      throw new AlgoliaRuntimeException("Parameter `modelID` is required when calling `updateModelInstance`.");
    }

    if (updateModelParams == null) {
      throw new AlgoliaRuntimeException("Parameter `updateModelParams` is required when calling `updateModelInstance`.");
    }

    Object bodyObj = updateModelParams;

    // create path and map variables
    String requestPath = "/1/predict/models/{modelID}".replaceAll("\\{modelID\\}", this.escapeString(modelID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "POST", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<UpdateModelInstanceResponse>() {});
  }

  /**
   * (asynchronously) Update a model’s configuration.
   *
   * @param modelID The ID of the model to retrieve. (required)
   * @param updateModelParams (required)
   * @return CompletableFuture<UpdateModelInstanceResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<UpdateModelInstanceResponse> updateModelInstanceAsync(String modelID, UpdateModelParams updateModelParams)
    throws AlgoliaRuntimeException {
    return this.updateModelInstanceAsync(modelID, updateModelParams, null);
  }

  /**
   * Update a segment’s configuration.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param updateSegmentParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdateSegmentResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public UpdateSegmentResponse updateSegment(String segmentID, UpdateSegmentParams updateSegmentParams, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(updateSegmentAsync(segmentID, updateSegmentParams, requestOptions));
  }

  /**
   * Update a segment’s configuration.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param updateSegmentParams (required)
   * @return UpdateSegmentResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public UpdateSegmentResponse updateSegment(String segmentID, UpdateSegmentParams updateSegmentParams) throws AlgoliaRuntimeException {
    return this.updateSegment(segmentID, updateSegmentParams, null);
  }

  /**
   * (asynchronously) Update a segment’s configuration.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param updateSegmentParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<UpdateSegmentResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<UpdateSegmentResponse> updateSegmentAsync(
    String segmentID,
    UpdateSegmentParams updateSegmentParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (segmentID == null) {
      throw new AlgoliaRuntimeException("Parameter `segmentID` is required when calling `updateSegment`.");
    }

    if (updateSegmentParams == null) {
      throw new AlgoliaRuntimeException("Parameter `updateSegmentParams` is required when calling `updateSegment`.");
    }

    Object bodyObj = updateSegmentParams;

    // create path and map variables
    String requestPath = "/1/segments/{segmentID}".replaceAll("\\{segmentID\\}", this.escapeString(segmentID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "POST", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<UpdateSegmentResponse>() {});
  }

  /**
   * (asynchronously) Update a segment’s configuration.
   *
   * @param segmentID The ID of the Segment to fetch. (required)
   * @param updateSegmentParams (required)
   * @return CompletableFuture<UpdateSegmentResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<UpdateSegmentResponse> updateSegmentAsync(String segmentID, UpdateSegmentParams updateSegmentParams)
    throws AlgoliaRuntimeException {
    return this.updateSegmentAsync(segmentID, updateSegmentParams, null);
  }
}
