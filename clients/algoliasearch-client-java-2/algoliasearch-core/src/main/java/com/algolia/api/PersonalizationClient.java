package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.personalization.*;
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

public class PersonalizationClient extends ApiClient {

  public PersonalizationClient(String appId, String apiKey, String region) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(region)), null);
  }

  public PersonalizationClient(
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

  public PersonalizationClient(
    String appId,
    String apiKey,
    Requester requester
  ) {
    this(appId, apiKey, requester, null);
  }

  public PersonalizationClient(
    String appId,
    String apiKey,
    Requester requester,
    UserAgent.Segment[] userAgentSegments
  ) {
    super(appId, apiKey, requester, "Personalization", userAgentSegments);
  }

  private static List<StatefulHost> getDefaultHosts(String region) {
    List<StatefulHost> hosts = new ArrayList<StatefulHost>();

    String url =
      "personalization.{region}.algolia.com".replace("{region}", region);

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
   * Delete the user profile and all its associated data. Returns, as part of the response, a date
   * until which the data can safely be considered as deleted for the given user. This means if you
   * send events for the given user before this date, they will be ignored. Any data received after
   * the deletedUntil date will start building a new user profile. It might take a couple hours for
   * the deletion request to be fully processed.
   *
   * @param userToken userToken representing the user for which to fetch the Personalization
   *     profile. (required)
   * @return DeleteUserProfileResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeleteUserProfileResponse deleteUserProfile(String userToken)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteUserProfileAsync(userToken));
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeleteUserProfileResponse> deleteUserProfileAsync(
    String userToken
  ) throws AlgoliaRuntimeException {
    if (userToken == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'userToken' when calling deleteUserProfile(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/profiles/{userToken}".replaceAll(
          "\\{userToken\\}",
          this.escapeString(userToken.toString())
        );

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<DeleteUserProfileResponse>() {}.getType();
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
   * The strategy contains information on the events and facets that impact user profiles and
   * personalized search results.
   *
   * @return PersonalizationStrategyParams
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public PersonalizationStrategyParams getPersonalizationStrategy()
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getPersonalizationStrategyAsync());
  }

  /**
   * (asynchronously) The strategy contains information on the events and facets that impact user
   * profiles and personalized search results.
   *
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<PersonalizationStrategyParams> getPersonalizationStrategyAsync()
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/strategies/personalization";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<PersonalizationStrategyParams>() {}
      .getType();
    return this.executeAsync(call, returnType);
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
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetUserTokenResponse getUserTokenProfile(String userToken)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getUserTokenProfileAsync(userToken));
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetUserTokenResponse> getUserTokenProfileAsync(
    String userToken
  ) throws AlgoliaRuntimeException {
    if (userToken == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'userToken' when calling getUserTokenProfile(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/profiles/personalization/{userToken}".replaceAll(
          "\\{userToken\\}",
          this.escapeString(userToken.toString())
        );

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<GetUserTokenResponse>() {}.getType();
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
   * A strategy defines the events and facets that impact user profiles and personalized search
   * results.
   *
   * @param personalizationStrategyParams (required)
   * @return SetPersonalizationStrategyResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SetPersonalizationStrategyResponse setPersonalizationStrategy(
    PersonalizationStrategyParams personalizationStrategyParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      setPersonalizationStrategyAsync(personalizationStrategyParams)
    );
  }

  /**
   * (asynchronously) A strategy defines the events and facets that impact user profiles and
   * personalized search results.
   *
   * @param personalizationStrategyParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SetPersonalizationStrategyResponse> setPersonalizationStrategyAsync(
    PersonalizationStrategyParams personalizationStrategyParams
  ) throws AlgoliaRuntimeException {
    if (personalizationStrategyParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'personalizationStrategyParams' when calling" +
        " setPersonalizationStrategy(Async)"
      );
    }

    Object bodyObj = personalizationStrategyParams;

    // create path and map variables
    String requestPath = "/1/strategies/personalization";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SetPersonalizationStrategyResponse>() {}
      .getType();
    return this.executeAsync(call, returnType);
  }
}
