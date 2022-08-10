// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.personalization.*;
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

public class PersonalizationClient extends ApiClient {

  private static final String[] allowedRegions = { "eu", "us" };

  public PersonalizationClient(String appId, String apiKey, String region) {
    this(appId, apiKey, region, null);
  }

  public PersonalizationClient(String appId, String apiKey, String region, ClientOptions options) {
    super(appId, apiKey, "Personalization", "4.4.7-SNAPSHOT", options);
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

    String url = "personalization.{region}.algolia.com".replace("{region}", region);

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
   * Delete the user profile and all its associated data. Returns, as part of the response, a date
   * until which the data can safely be considered as deleted for the given user. This means if you
   * send events for the given user before this date, they will be ignored. Any data received after
   * the deletedUntil date will start building a new user profile. It might take a couple hours for
   * the deletion request to be fully processed.
   *
   * @param userToken userToken representing the user for which to fetch the Personalization
   *     profile. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return DeleteUserProfileResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public DeleteUserProfileResponse deleteUserProfile(String userToken, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteUserProfileAsync(userToken, requestOptions));
  }

  /**
   * Delete the user profile and all its associated data. Returns, as part of the response, a date
   * until which the data can safely be considered as deleted for the given user. This means if you
   * send events for the given user before this date, they will be ignored. Any data received after
   * the deletedUntil date will start building a new user profile. It might take a couple hours for
   * the deletion request to be fully processed.
   *
   * @param userToken userToken representing the user for which to fetch the Personalization
   *     profile. (required)
   * @return DeleteUserProfileResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public DeleteUserProfileResponse deleteUserProfile(String userToken) throws AlgoliaRuntimeException {
    return this.deleteUserProfile(userToken, null);
  }

  /**
   * (asynchronously) Delete the user profile and all its associated data. Returns, as part of the
   * response, a date until which the data can safely be considered as deleted for the given user.
   * This means if you send events for the given user before this date, they will be ignored. Any
   * data received after the deletedUntil date will start building a new user profile. It might take
   * a couple hours for the deletion request to be fully processed.
   *
   * @param userToken userToken representing the user for which to fetch the Personalization
   *     profile. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<DeleteUserProfileResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<DeleteUserProfileResponse> deleteUserProfileAsync(String userToken, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (userToken == null) {
      throw new AlgoliaRuntimeException("Parameter `userToken` is required when calling `deleteUserProfile`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/profiles/{userToken}".replaceAll("\\{userToken\\}", this.escapeString(userToken.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "DELETE", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<DeleteUserProfileResponse>() {});
  }

  /**
   * (asynchronously) Delete the user profile and all its associated data. Returns, as part of the
   * response, a date until which the data can safely be considered as deleted for the given user.
   * This means if you send events for the given user before this date, they will be ignored. Any
   * data received after the deletedUntil date will start building a new user profile. It might take
   * a couple hours for the deletion request to be fully processed.
   *
   * @param userToken userToken representing the user for which to fetch the Personalization
   *     profile. (required)
   * @return CompletableFuture<DeleteUserProfileResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<DeleteUserProfileResponse> deleteUserProfileAsync(String userToken) throws AlgoliaRuntimeException {
    return this.deleteUserProfileAsync(userToken, null);
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
   * The strategy contains information on the events and facets that impact user profiles and
   * personalized search results.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return PersonalizationStrategyParams
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public PersonalizationStrategyParams getPersonalizationStrategy(RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getPersonalizationStrategyAsync(requestOptions));
  }

  /**
   * The strategy contains information on the events and facets that impact user profiles and
   * personalized search results.
   *
   * @return PersonalizationStrategyParams
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public PersonalizationStrategyParams getPersonalizationStrategy() throws AlgoliaRuntimeException {
    return this.getPersonalizationStrategy(null);
  }

  /**
   * (asynchronously) The strategy contains information on the events and facets that impact user
   * profiles and personalized search results.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<PersonalizationStrategyParams> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<PersonalizationStrategyParams> getPersonalizationStrategyAsync(RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/strategies/personalization";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<PersonalizationStrategyParams>() {});
  }

  /**
   * (asynchronously) The strategy contains information on the events and facets that impact user
   * profiles and personalized search results.
   *
   * @return CompletableFuture<PersonalizationStrategyParams> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<PersonalizationStrategyParams> getPersonalizationStrategyAsync() throws AlgoliaRuntimeException {
    return this.getPersonalizationStrategyAsync(null);
  }

  /**
   * Get the user profile built from Personalization strategy. The profile is structured by facet
   * name used in the strategy. Each facet value is mapped to its score. Each score represents the
   * user affinity for a specific facet value given the userToken past events and the
   * Personalization strategy defined. Scores are bounded to 20. The last processed event timestamp
   * is provided using the ISO 8601 format for debugging purposes.
   *
   * @param userToken userToken representing the user for which to fetch the Personalization
   *     profile. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetUserTokenResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public GetUserTokenResponse getUserTokenProfile(String userToken, RequestOptions requestOptions) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getUserTokenProfileAsync(userToken, requestOptions));
  }

  /**
   * Get the user profile built from Personalization strategy. The profile is structured by facet
   * name used in the strategy. Each facet value is mapped to its score. Each score represents the
   * user affinity for a specific facet value given the userToken past events and the
   * Personalization strategy defined. Scores are bounded to 20. The last processed event timestamp
   * is provided using the ISO 8601 format for debugging purposes.
   *
   * @param userToken userToken representing the user for which to fetch the Personalization
   *     profile. (required)
   * @return GetUserTokenResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public GetUserTokenResponse getUserTokenProfile(String userToken) throws AlgoliaRuntimeException {
    return this.getUserTokenProfile(userToken, null);
  }

  /**
   * (asynchronously) Get the user profile built from Personalization strategy. The profile is
   * structured by facet name used in the strategy. Each facet value is mapped to its score. Each
   * score represents the user affinity for a specific facet value given the userToken past events
   * and the Personalization strategy defined. Scores are bounded to 20. The last processed event
   * timestamp is provided using the ISO 8601 format for debugging purposes.
   *
   * @param userToken userToken representing the user for which to fetch the Personalization
   *     profile. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<GetUserTokenResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<GetUserTokenResponse> getUserTokenProfileAsync(String userToken, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    if (userToken == null) {
      throw new AlgoliaRuntimeException("Parameter `userToken` is required when calling `getUserTokenProfile`.");
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/profiles/personalization/{userToken}".replaceAll("\\{userToken\\}", this.escapeString(userToken.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "GET", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<GetUserTokenResponse>() {});
  }

  /**
   * (asynchronously) Get the user profile built from Personalization strategy. The profile is
   * structured by facet name used in the strategy. Each facet value is mapped to its score. Each
   * score represents the user affinity for a specific facet value given the userToken past events
   * and the Personalization strategy defined. Scores are bounded to 20. The last processed event
   * timestamp is provided using the ISO 8601 format for debugging purposes.
   *
   * @param userToken userToken representing the user for which to fetch the Personalization
   *     profile. (required)
   * @return CompletableFuture<GetUserTokenResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<GetUserTokenResponse> getUserTokenProfileAsync(String userToken) throws AlgoliaRuntimeException {
    return this.getUserTokenProfileAsync(userToken, null);
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
   * A strategy defines the events and facets that impact user profiles and personalized search
   * results.
   *
   * @param personalizationStrategyParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SetPersonalizationStrategyResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public SetPersonalizationStrategyResponse setPersonalizationStrategy(
    PersonalizationStrategyParams personalizationStrategyParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(setPersonalizationStrategyAsync(personalizationStrategyParams, requestOptions));
  }

  /**
   * A strategy defines the events and facets that impact user profiles and personalized search
   * results.
   *
   * @param personalizationStrategyParams (required)
   * @return SetPersonalizationStrategyResponse
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public SetPersonalizationStrategyResponse setPersonalizationStrategy(PersonalizationStrategyParams personalizationStrategyParams)
    throws AlgoliaRuntimeException {
    return this.setPersonalizationStrategy(personalizationStrategyParams, null);
  }

  /**
   * (asynchronously) A strategy defines the events and facets that impact user profiles and
   * personalized search results.
   *
   * @param personalizationStrategyParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CompletableFuture<SetPersonalizationStrategyResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<SetPersonalizationStrategyResponse> setPersonalizationStrategyAsync(
    PersonalizationStrategyParams personalizationStrategyParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (personalizationStrategyParams == null) {
      throw new AlgoliaRuntimeException(
        "Parameter `personalizationStrategyParams` is required when calling" + " `setPersonalizationStrategy`."
      );
    }

    Object bodyObj = personalizationStrategyParams;

    // create path and map variables
    String requestPath = "/1/strategies/personalization";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call = this.buildCall(requestPath, "POST", queryParameters, bodyObj, headers, requestOptions, false);
    return this.executeAsync(call, new TypeReference<SetPersonalizationStrategyResponse>() {});
  }

  /**
   * (asynchronously) A strategy defines the events and facets that impact user profiles and
   * personalized search results.
   *
   * @param personalizationStrategyParams (required)
   * @return CompletableFuture<SetPersonalizationStrategyResponse> The awaitable future
   * @throws AlgoliaRuntimeException If it fails to process the API call
   */
  public CompletableFuture<SetPersonalizationStrategyResponse> setPersonalizationStrategyAsync(
    PersonalizationStrategyParams personalizationStrategyParams
  ) throws AlgoliaRuntimeException {
    return this.setPersonalizationStrategyAsync(personalizationStrategyParams, null);
  }
}
