package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.search.*;
import com.algolia.utils.*;
import com.algolia.utils.RequestOptions;
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

public class SearchClient extends ApiClient {

  public SearchClient(String appId, String apiKey) {
    this(appId, apiKey, new HttpRequester(getDefaultHosts(appId)), null);
  }

  public SearchClient(
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

  public SearchClient(String appId, String apiKey, Requester requester) {
    this(appId, apiKey, requester, null);
  }

  public SearchClient(
    String appId,
    String apiKey,
    Requester requester,
    UserAgent.Segment[] userAgentSegments
  ) {
    super(appId, apiKey, requester, "Search", userAgentSegments);
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
   * Add a new API Key with specific permissions/restrictions.
   *
   * @param apiKey (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return AddApiKeyResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public AddApiKeyResponse addApiKey(
    ApiKey apiKey,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(addApiKeyAsync(apiKey, requestOptions));
  }

  public AddApiKeyResponse addApiKey(ApiKey apiKey)
    throws AlgoliaRuntimeException {
    return this.addApiKey(apiKey, null);
  }

  /**
   * (asynchronously) Add a new API Key with specific permissions/restrictions.
   *
   * @param apiKey (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<AddApiKeyResponse> addApiKeyAsync(
    ApiKey apiKey,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (apiKey == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'apiKey' when calling addApiKey(Async)"
      );
    }

    Object bodyObj = apiKey;

    // create path and map variables
    String requestPath = "/1/keys";

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
    Type returnType = new TypeToken<AddApiKeyResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<AddApiKeyResponse> addApiKeyAsync(ApiKey apiKey)
    throws AlgoliaRuntimeException {
    return this.addApiKeyAsync(apiKey, null);
  }

  /**
   * Add or replace an object with a given object ID. If the object does not exist, it will be
   * created. If it already exists, it will be replaced.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param body The Algolia object. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtWithObjectIdResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtWithObjectIdResponse addOrUpdateObject(
    String indexName,
    String objectID,
    Object body,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      addOrUpdateObjectAsync(indexName, objectID, body, requestOptions)
    );
  }

  public UpdatedAtWithObjectIdResponse addOrUpdateObject(
    String indexName,
    String objectID,
    Object body
  ) throws AlgoliaRuntimeException {
    return this.addOrUpdateObject(indexName, objectID, body, null);
  }

  /**
   * (asynchronously) Add or replace an object with a given object ID. If the object does not exist,
   * it will be created. If it already exists, it will be replaced.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param body The Algolia object. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtWithObjectIdResponse> addOrUpdateObjectAsync(
    String indexName,
    String objectID,
    Object body,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling addOrUpdateObject(Async)"
      );
    }

    if (objectID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'objectID' when calling addOrUpdateObject(Async)"
      );
    }

    if (body == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'body' when calling addOrUpdateObject(Async)"
      );
    }

    Object bodyObj = body;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/{objectID}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{objectID\\}", this.escapeString(objectID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "PUT",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<UpdatedAtWithObjectIdResponse>() {}
      .getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtWithObjectIdResponse> addOrUpdateObjectAsync(
    String indexName,
    String objectID,
    Object body
  ) throws AlgoliaRuntimeException {
    return this.addOrUpdateObjectAsync(indexName, objectID, body, null);
  }

  /**
   * Add a single source to the list of allowed sources.
   *
   * @param source The source to add. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CreatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public CreatedAtResponse appendSource(
    Source source,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(appendSourceAsync(source, requestOptions));
  }

  public CreatedAtResponse appendSource(Source source)
    throws AlgoliaRuntimeException {
    return this.appendSource(source, null);
  }

  /**
   * (asynchronously) Add a single source to the list of allowed sources.
   *
   * @param source The source to add. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<CreatedAtResponse> appendSourceAsync(
    Source source,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (source == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'source' when calling appendSource(Async)"
      );
    }

    Object bodyObj = source;

    // create path and map variables
    String requestPath = "/1/security/sources/append";

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
    Type returnType = new TypeToken<CreatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<CreatedAtResponse> appendSourceAsync(Source source)
    throws AlgoliaRuntimeException {
    return this.appendSourceAsync(source, null);
  }

  /**
   * Assign or Move a userID to a cluster. The time it takes to migrate (move) a user is
   * proportional to the amount of data linked to the userID. Upon success, the response is 200 OK.
   * A successful response indicates that the operation has been taken into account, and the userID
   * is directly usable.
   *
   * @param xAlgoliaUserID userID to assign. (required)
   * @param assignUserIdParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CreatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public CreatedAtResponse assignUserId(
    String xAlgoliaUserID,
    AssignUserIdParams assignUserIdParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      assignUserIdAsync(xAlgoliaUserID, assignUserIdParams, requestOptions)
    );
  }

  public CreatedAtResponse assignUserId(
    String xAlgoliaUserID,
    AssignUserIdParams assignUserIdParams
  ) throws AlgoliaRuntimeException {
    return this.assignUserId(xAlgoliaUserID, assignUserIdParams, null);
  }

  /**
   * (asynchronously) Assign or Move a userID to a cluster. The time it takes to migrate (move) a
   * user is proportional to the amount of data linked to the userID. Upon success, the response is
   * 200 OK. A successful response indicates that the operation has been taken into account, and the
   * userID is directly usable.
   *
   * @param xAlgoliaUserID userID to assign. (required)
   * @param assignUserIdParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<CreatedAtResponse> assignUserIdAsync(
    String xAlgoliaUserID,
    AssignUserIdParams assignUserIdParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (xAlgoliaUserID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'xAlgoliaUserID' when calling assignUserId(Async)"
      );
    }

    if (assignUserIdParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'assignUserIdParams' when calling assignUserId(Async)"
      );
    }

    Object bodyObj = assignUserIdParams;

    // create path and map variables
    String requestPath = "/1/clusters/mapping";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (xAlgoliaUserID != null) {
      headers.put("X-Algolia-User-ID", this.parameterToString(xAlgoliaUserID));
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
    Type returnType = new TypeToken<CreatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<CreatedAtResponse> assignUserIdAsync(
    String xAlgoliaUserID,
    AssignUserIdParams assignUserIdParams
  ) throws AlgoliaRuntimeException {
    return this.assignUserIdAsync(xAlgoliaUserID, assignUserIdParams, null);
  }

  /**
   * Perform multiple write operations targeting one index, in a single API call.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param batchWriteParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return BatchResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public BatchResponse batch(
    String indexName,
    BatchWriteParams batchWriteParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      batchAsync(indexName, batchWriteParams, requestOptions)
    );
  }

  public BatchResponse batch(
    String indexName,
    BatchWriteParams batchWriteParams
  ) throws AlgoliaRuntimeException {
    return this.batch(indexName, batchWriteParams, null);
  }

  /**
   * (asynchronously) Perform multiple write operations targeting one index, in a single API call.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param batchWriteParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<BatchResponse> batchAsync(
    String indexName,
    BatchWriteParams batchWriteParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling batch(Async)"
      );
    }

    if (batchWriteParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'batchWriteParams' when calling batch(Async)"
      );
    }

    Object bodyObj = batchWriteParams;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/batch".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

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
    Type returnType = new TypeToken<BatchResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<BatchResponse> batchAsync(
    String indexName,
    BatchWriteParams batchWriteParams
  ) throws AlgoliaRuntimeException {
    return this.batchAsync(indexName, batchWriteParams, null);
  }

  /**
   * Assign multiple userIDs to a cluster. Upon success, the response is 200 OK. A successful
   * response indicates that the operation has been taken into account, and the userIDs are directly
   * usable.
   *
   * @param xAlgoliaUserID userID to assign. (required)
   * @param batchAssignUserIdsParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CreatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public CreatedAtResponse batchAssignUserIds(
    String xAlgoliaUserID,
    BatchAssignUserIdsParams batchAssignUserIdsParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      batchAssignUserIdsAsync(
        xAlgoliaUserID,
        batchAssignUserIdsParams,
        requestOptions
      )
    );
  }

  public CreatedAtResponse batchAssignUserIds(
    String xAlgoliaUserID,
    BatchAssignUserIdsParams batchAssignUserIdsParams
  ) throws AlgoliaRuntimeException {
    return this.batchAssignUserIds(
        xAlgoliaUserID,
        batchAssignUserIdsParams,
        null
      );
  }

  /**
   * (asynchronously) Assign multiple userIDs to a cluster. Upon success, the response is 200 OK. A
   * successful response indicates that the operation has been taken into account, and the userIDs
   * are directly usable.
   *
   * @param xAlgoliaUserID userID to assign. (required)
   * @param batchAssignUserIdsParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<CreatedAtResponse> batchAssignUserIdsAsync(
    String xAlgoliaUserID,
    BatchAssignUserIdsParams batchAssignUserIdsParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (xAlgoliaUserID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'xAlgoliaUserID' when calling batchAssignUserIds(Async)"
      );
    }

    if (batchAssignUserIdsParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'batchAssignUserIdsParams' when calling" +
        " batchAssignUserIds(Async)"
      );
    }

    Object bodyObj = batchAssignUserIdsParams;

    // create path and map variables
    String requestPath = "/1/clusters/mapping/batch";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (xAlgoliaUserID != null) {
      headers.put("X-Algolia-User-ID", this.parameterToString(xAlgoliaUserID));
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
    Type returnType = new TypeToken<CreatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<CreatedAtResponse> batchAssignUserIdsAsync(
    String xAlgoliaUserID,
    BatchAssignUserIdsParams batchAssignUserIdsParams
  ) throws AlgoliaRuntimeException {
    return this.batchAssignUserIdsAsync(
        xAlgoliaUserID,
        batchAssignUserIdsParams,
        null
      );
  }

  /**
   * Send a batch of dictionary entries.
   *
   * @param dictionaryName The dictionary to search in. (required)
   * @param batchDictionaryEntriesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse batchDictionaryEntries(
    DictionaryType dictionaryName,
    BatchDictionaryEntriesParams batchDictionaryEntriesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      batchDictionaryEntriesAsync(
        dictionaryName,
        batchDictionaryEntriesParams,
        requestOptions
      )
    );
  }

  public UpdatedAtResponse batchDictionaryEntries(
    DictionaryType dictionaryName,
    BatchDictionaryEntriesParams batchDictionaryEntriesParams
  ) throws AlgoliaRuntimeException {
    return this.batchDictionaryEntries(
        dictionaryName,
        batchDictionaryEntriesParams,
        null
      );
  }

  /**
   * (asynchronously) Send a batch of dictionary entries.
   *
   * @param dictionaryName The dictionary to search in. (required)
   * @param batchDictionaryEntriesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> batchDictionaryEntriesAsync(
    DictionaryType dictionaryName,
    BatchDictionaryEntriesParams batchDictionaryEntriesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (dictionaryName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'dictionaryName' when calling" +
        " batchDictionaryEntries(Async)"
      );
    }

    if (batchDictionaryEntriesParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'batchDictionaryEntriesParams' when calling" +
        " batchDictionaryEntries(Async)"
      );
    }

    Object bodyObj = batchDictionaryEntriesParams;

    // create path and map variables
    String requestPath =
      "/1/dictionaries/{dictionaryName}/batch".replaceAll(
          "\\{dictionaryName\\}",
          this.escapeString(dictionaryName.toString())
        );

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
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> batchDictionaryEntriesAsync(
    DictionaryType dictionaryName,
    BatchDictionaryEntriesParams batchDictionaryEntriesParams
  ) throws AlgoliaRuntimeException {
    return this.batchDictionaryEntriesAsync(
        dictionaryName,
        batchDictionaryEntriesParams,
        null
      );
  }

  /**
   * Create or update a batch of Rules.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param rule (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param clearExistingRules When true, existing Rules are cleared before adding this batch. When
   *     false, existing Rules are kept. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse batchRules(
    String indexName,
    List<Rule> rule,
    Boolean forwardToReplicas,
    Boolean clearExistingRules,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      batchRulesAsync(
        indexName,
        rule,
        forwardToReplicas,
        clearExistingRules,
        requestOptions
      )
    );
  }

  public UpdatedAtResponse batchRules(
    String indexName,
    List<Rule> rule,
    Boolean forwardToReplicas,
    Boolean clearExistingRules
  ) throws AlgoliaRuntimeException {
    return this.batchRules(
        indexName,
        rule,
        forwardToReplicas,
        clearExistingRules,
        null
      );
  }

  public UpdatedAtResponse batchRules(
    String indexName,
    List<Rule> rule,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.batchRules(indexName, rule, null, null, requestOptions);
  }

  public UpdatedAtResponse batchRules(String indexName, List<Rule> rule)
    throws AlgoliaRuntimeException {
    return this.batchRules(indexName, rule, null, null, null);
  }

  /**
   * (asynchronously) Create or update a batch of Rules.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param rule (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param clearExistingRules When true, existing Rules are cleared before adding this batch. When
   *     false, existing Rules are kept. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> batchRulesAsync(
    String indexName,
    List<Rule> rule,
    Boolean forwardToReplicas,
    Boolean clearExistingRules,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling batchRules(Async)"
      );
    }

    if (rule == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'rule' when calling batchRules(Async)"
      );
    }

    Object bodyObj = rule;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/rules/batch".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParameters.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    if (clearExistingRules != null) {
      queryParameters.put(
        "clearExistingRules",
        parameterToString(clearExistingRules)
      );
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
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> batchRulesAsync(
    String indexName,
    List<Rule> rule,
    Boolean forwardToReplicas,
    Boolean clearExistingRules
  ) throws AlgoliaRuntimeException {
    return this.batchRulesAsync(
        indexName,
        rule,
        forwardToReplicas,
        clearExistingRules,
        null
      );
  }

  public CompletableFuture<UpdatedAtResponse> batchRulesAsync(
    String indexName,
    List<Rule> rule,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.batchRulesAsync(indexName, rule, null, null, requestOptions);
  }

  public CompletableFuture<UpdatedAtResponse> batchRulesAsync(
    String indexName,
    List<Rule> rule
  ) throws AlgoliaRuntimeException {
    return this.batchRulesAsync(indexName, rule, null, null, null);
  }

  /**
   * This method allows you to retrieve all index content. It can retrieve up to 1,000 records per
   * call and supports full text search and filters. For performance reasons, some features are not
   * supported, including `distinct`, sorting by `typos`, `words` or `geo distance`. When there is
   * more content to be browsed, the response contains a cursor field. This cursor has to be passed
   * to the subsequent call to browse in order to get the next page of results. When the end of the
   * index has been reached, the cursor field is absent from the response.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param browseRequest (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return BrowseResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public BrowseResponse browse(
    String indexName,
    BrowseRequest browseRequest,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      browseAsync(indexName, browseRequest, requestOptions)
    );
  }

  public BrowseResponse browse(String indexName, BrowseRequest browseRequest)
    throws AlgoliaRuntimeException {
    return this.browse(indexName, browseRequest, null);
  }

  public BrowseResponse browse(String indexName, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.browse(indexName, null, requestOptions);
  }

  public BrowseResponse browse(String indexName)
    throws AlgoliaRuntimeException {
    return this.browse(indexName, null, null);
  }

  /**
   * (asynchronously) This method allows you to retrieve all index content. It can retrieve up to
   * 1,000 records per call and supports full text search and filters. For performance reasons, some
   * features are not supported, including &#x60;distinct&#x60;, sorting by &#x60;typos&#x60;,
   * &#x60;words&#x60; or &#x60;geo distance&#x60;. When there is more content to be browsed, the
   * response contains a cursor field. This cursor has to be passed to the subsequent call to browse
   * in order to get the next page of results. When the end of the index has been reached, the
   * cursor field is absent from the response.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param browseRequest (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<BrowseResponse> browseAsync(
    String indexName,
    BrowseRequest browseRequest,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling browse(Async)"
      );
    }

    Object bodyObj = browseRequest;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/browse".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

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
    Type returnType = new TypeToken<BrowseResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<BrowseResponse> browseAsync(
    String indexName,
    BrowseRequest browseRequest
  ) throws AlgoliaRuntimeException {
    return this.browseAsync(indexName, browseRequest, null);
  }

  public CompletableFuture<BrowseResponse> browseAsync(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.browseAsync(indexName, null, requestOptions);
  }

  public CompletableFuture<BrowseResponse> browseAsync(String indexName)
    throws AlgoliaRuntimeException {
    return this.browseAsync(indexName, null, null);
  }

  /**
   * Remove all synonyms from an index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse clearAllSynonyms(
    String indexName,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      clearAllSynonymsAsync(indexName, forwardToReplicas, requestOptions)
    );
  }

  public UpdatedAtResponse clearAllSynonyms(
    String indexName,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.clearAllSynonyms(indexName, forwardToReplicas, null);
  }

  public UpdatedAtResponse clearAllSynonyms(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.clearAllSynonyms(indexName, null, requestOptions);
  }

  public UpdatedAtResponse clearAllSynonyms(String indexName)
    throws AlgoliaRuntimeException {
    return this.clearAllSynonyms(indexName, null, null);
  }

  /**
   * (asynchronously) Remove all synonyms from an index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> clearAllSynonymsAsync(
    String indexName,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling clearAllSynonyms(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/synonyms/clear".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParameters.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
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
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> clearAllSynonymsAsync(
    String indexName,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.clearAllSynonymsAsync(indexName, forwardToReplicas, null);
  }

  public CompletableFuture<UpdatedAtResponse> clearAllSynonymsAsync(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.clearAllSynonymsAsync(indexName, null, requestOptions);
  }

  public CompletableFuture<UpdatedAtResponse> clearAllSynonymsAsync(
    String indexName
  ) throws AlgoliaRuntimeException {
    return this.clearAllSynonymsAsync(indexName, null, null);
  }

  /**
   * Delete an index's content, but leave settings and index-specific API keys untouched.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse clearObjects(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(clearObjectsAsync(indexName, requestOptions));
  }

  public UpdatedAtResponse clearObjects(String indexName)
    throws AlgoliaRuntimeException {
    return this.clearObjects(indexName, null);
  }

  /**
   * (asynchronously) Delete an index&#39;s content, but leave settings and index-specific API keys
   * untouched.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> clearObjectsAsync(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling clearObjects(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/clear".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

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
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> clearObjectsAsync(
    String indexName
  ) throws AlgoliaRuntimeException {
    return this.clearObjectsAsync(indexName, null);
  }

  /**
   * Delete all Rules in the index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse clearRules(
    String indexName,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      clearRulesAsync(indexName, forwardToReplicas, requestOptions)
    );
  }

  public UpdatedAtResponse clearRules(
    String indexName,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.clearRules(indexName, forwardToReplicas, null);
  }

  public UpdatedAtResponse clearRules(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.clearRules(indexName, null, requestOptions);
  }

  public UpdatedAtResponse clearRules(String indexName)
    throws AlgoliaRuntimeException {
    return this.clearRules(indexName, null, null);
  }

  /**
   * (asynchronously) Delete all Rules in the index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> clearRulesAsync(
    String indexName,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling clearRules(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/rules/clear".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParameters.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
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
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> clearRulesAsync(
    String indexName,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.clearRulesAsync(indexName, forwardToReplicas, null);
  }

  public CompletableFuture<UpdatedAtResponse> clearRulesAsync(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.clearRulesAsync(indexName, null, requestOptions);
  }

  public CompletableFuture<UpdatedAtResponse> clearRulesAsync(String indexName)
    throws AlgoliaRuntimeException {
    return this.clearRulesAsync(indexName, null, null);
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
   * Delete an existing API Key.
   *
   * @param key API Key string. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return DeleteApiKeyResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeleteApiKeyResponse deleteApiKey(
    String key,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteApiKeyAsync(key, requestOptions));
  }

  public DeleteApiKeyResponse deleteApiKey(String key)
    throws AlgoliaRuntimeException {
    return this.deleteApiKey(key, null);
  }

  /**
   * (asynchronously) Delete an existing API Key.
   *
   * @param key API Key string. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeleteApiKeyResponse> deleteApiKeyAsync(
    String key,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (key == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'key' when calling deleteApiKey(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/keys/{key}".replaceAll(
          "\\{key\\}",
          this.escapeString(key.toString())
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
          requestOptions
        );
    Type returnType = new TypeToken<DeleteApiKeyResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<DeleteApiKeyResponse> deleteApiKeyAsync(String key)
    throws AlgoliaRuntimeException {
    return this.deleteApiKeyAsync(key, null);
  }

  /**
   * Remove all objects matching a filter (including geo filters). This method enables you to delete
   * one or more objects based on filters (numeric, facet, tag or geo queries). It doesn't accept
   * empty filters or a query.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return DeletedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeletedAtResponse deleteBy(
    String indexName,
    SearchParams searchParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      deleteByAsync(indexName, searchParams, requestOptions)
    );
  }

  public DeletedAtResponse deleteBy(
    String indexName,
    SearchParams searchParams
  ) throws AlgoliaRuntimeException {
    return this.deleteBy(indexName, searchParams, null);
  }

  /**
   * (asynchronously) Remove all objects matching a filter (including geo filters). This method
   * enables you to delete one or more objects based on filters (numeric, facet, tag or geo
   * queries). It doesn&#39;t accept empty filters or a query.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeletedAtResponse> deleteByAsync(
    String indexName,
    SearchParams searchParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling deleteBy(Async)"
      );
    }

    if (searchParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'searchParams' when calling deleteBy(Async)"
      );
    }

    Object bodyObj = searchParams;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/deleteByQuery".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

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
    Type returnType = new TypeToken<DeletedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<DeletedAtResponse> deleteByAsync(
    String indexName,
    SearchParams searchParams
  ) throws AlgoliaRuntimeException {
    return this.deleteByAsync(indexName, searchParams, null);
  }

  /**
   * Delete an existing index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return DeletedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeletedAtResponse deleteIndex(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteIndexAsync(indexName, requestOptions));
  }

  public DeletedAtResponse deleteIndex(String indexName)
    throws AlgoliaRuntimeException {
    return this.deleteIndex(indexName, null);
  }

  /**
   * (asynchronously) Delete an existing index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeletedAtResponse> deleteIndexAsync(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling deleteIndex(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}".replaceAll(
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
          requestOptions
        );
    Type returnType = new TypeToken<DeletedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<DeletedAtResponse> deleteIndexAsync(
    String indexName
  ) throws AlgoliaRuntimeException {
    return this.deleteIndexAsync(indexName, null);
  }

  /**
   * Delete an existing object.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return DeletedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeletedAtResponse deleteObject(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      deleteObjectAsync(indexName, objectID, requestOptions)
    );
  }

  public DeletedAtResponse deleteObject(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return this.deleteObject(indexName, objectID, null);
  }

  /**
   * (asynchronously) Delete an existing object.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeletedAtResponse> deleteObjectAsync(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling deleteObject(Async)"
      );
    }

    if (objectID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'objectID' when calling deleteObject(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/{objectID}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{objectID\\}", this.escapeString(objectID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "DELETE",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<DeletedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<DeletedAtResponse> deleteObjectAsync(
    String indexName,
    String objectID
  ) throws AlgoliaRuntimeException {
    return this.deleteObjectAsync(indexName, objectID, null);
  }

  /**
   * Delete the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse deleteRule(
    String indexName,
    String objectID,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      deleteRuleAsync(indexName, objectID, forwardToReplicas, requestOptions)
    );
  }

  public UpdatedAtResponse deleteRule(
    String indexName,
    String objectID,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.deleteRule(indexName, objectID, forwardToReplicas, null);
  }

  public UpdatedAtResponse deleteRule(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.deleteRule(indexName, objectID, null, requestOptions);
  }

  public UpdatedAtResponse deleteRule(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return this.deleteRule(indexName, objectID, null, null);
  }

  /**
   * (asynchronously) Delete the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> deleteRuleAsync(
    String indexName,
    String objectID,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling deleteRule(Async)"
      );
    }

    if (objectID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'objectID' when calling deleteRule(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/rules/{objectID}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{objectID\\}", this.escapeString(objectID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParameters.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
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
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> deleteRuleAsync(
    String indexName,
    String objectID,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.deleteRuleAsync(indexName, objectID, forwardToReplicas, null);
  }

  public CompletableFuture<UpdatedAtResponse> deleteRuleAsync(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.deleteRuleAsync(indexName, objectID, null, requestOptions);
  }

  public CompletableFuture<UpdatedAtResponse> deleteRuleAsync(
    String indexName,
    String objectID
  ) throws AlgoliaRuntimeException {
    return this.deleteRuleAsync(indexName, objectID, null, null);
  }

  /**
   * Remove a single source from the list of allowed sources.
   *
   * @param source The IP range of the source. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return DeleteSourceResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeleteSourceResponse deleteSource(
    String source,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteSourceAsync(source, requestOptions));
  }

  public DeleteSourceResponse deleteSource(String source)
    throws AlgoliaRuntimeException {
    return this.deleteSource(source, null);
  }

  /**
   * (asynchronously) Remove a single source from the list of allowed sources.
   *
   * @param source The IP range of the source. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeleteSourceResponse> deleteSourceAsync(
    String source,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (source == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'source' when calling deleteSource(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/security/sources/{source}".replaceAll(
          "\\{source\\}",
          this.escapeString(source.toString())
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
          requestOptions
        );
    Type returnType = new TypeToken<DeleteSourceResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<DeleteSourceResponse> deleteSourceAsync(
    String source
  ) throws AlgoliaRuntimeException {
    return this.deleteSourceAsync(source, null);
  }

  /**
   * Delete a single synonyms set, identified by the given objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return DeletedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeletedAtResponse deleteSynonym(
    String indexName,
    String objectID,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      deleteSynonymAsync(indexName, objectID, forwardToReplicas, requestOptions)
    );
  }

  public DeletedAtResponse deleteSynonym(
    String indexName,
    String objectID,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.deleteSynonym(indexName, objectID, forwardToReplicas, null);
  }

  public DeletedAtResponse deleteSynonym(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.deleteSynonym(indexName, objectID, null, requestOptions);
  }

  public DeletedAtResponse deleteSynonym(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return this.deleteSynonym(indexName, objectID, null, null);
  }

  /**
   * (asynchronously) Delete a single synonyms set, identified by the given objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeletedAtResponse> deleteSynonymAsync(
    String indexName,
    String objectID,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling deleteSynonym(Async)"
      );
    }

    if (objectID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'objectID' when calling deleteSynonym(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/synonyms/{objectID}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{objectID\\}", this.escapeString(objectID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParameters.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
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
    Type returnType = new TypeToken<DeletedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<DeletedAtResponse> deleteSynonymAsync(
    String indexName,
    String objectID,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.deleteSynonymAsync(
        indexName,
        objectID,
        forwardToReplicas,
        null
      );
  }

  public CompletableFuture<DeletedAtResponse> deleteSynonymAsync(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.deleteSynonymAsync(indexName, objectID, null, requestOptions);
  }

  public CompletableFuture<DeletedAtResponse> deleteSynonymAsync(
    String indexName,
    String objectID
  ) throws AlgoliaRuntimeException {
    return this.deleteSynonymAsync(indexName, objectID, null, null);
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
   * Get the permissions of an API key.
   *
   * @param key API Key string. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Key
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Key getApiKey(String key, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getApiKeyAsync(key, requestOptions));
  }

  public Key getApiKey(String key) throws AlgoliaRuntimeException {
    return this.getApiKey(key, null);
  }

  /**
   * (asynchronously) Get the permissions of an API key.
   *
   * @param key API Key string. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Key> getApiKeyAsync(
    String key,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (key == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'key' when calling getApiKey(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/keys/{key}".replaceAll(
          "\\{key\\}",
          this.escapeString(key.toString())
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
          requestOptions
        );
    Type returnType = new TypeToken<Key>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<Key> getApiKeyAsync(String key)
    throws AlgoliaRuntimeException {
    return this.getApiKeyAsync(key, null);
  }

  /**
   * List dictionaries supported per language.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Map&lt;String, Languages&gt;
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Map<String, Languages> getDictionaryLanguages(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getDictionaryLanguagesAsync(requestOptions));
  }

  public Map<String, Languages> getDictionaryLanguages()
    throws AlgoliaRuntimeException {
    return this.getDictionaryLanguages(null);
  }

  /**
   * (asynchronously) List dictionaries supported per language.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Map<String, Languages>> getDictionaryLanguagesAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/dictionaries/*/languages";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<Map<String, Languages>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<Map<String, Languages>> getDictionaryLanguagesAsync()
    throws AlgoliaRuntimeException {
    return this.getDictionaryLanguagesAsync(null);
  }

  /**
   * Retrieve dictionaries settings. The API stores languages whose standard entries are disabled.
   * Fetch settings does not return false values.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetDictionarySettingsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetDictionarySettingsResponse getDictionarySettings(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getDictionarySettingsAsync(requestOptions));
  }

  public GetDictionarySettingsResponse getDictionarySettings()
    throws AlgoliaRuntimeException {
    return this.getDictionarySettings(null);
  }

  /**
   * (asynchronously) Retrieve dictionaries settings. The API stores languages whose standard
   * entries are disabled. Fetch settings does not return false values.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetDictionarySettingsResponse> getDictionarySettingsAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/dictionaries/*/settings";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<GetDictionarySettingsResponse>() {}
      .getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetDictionarySettingsResponse> getDictionarySettingsAsync()
    throws AlgoliaRuntimeException {
    return this.getDictionarySettingsAsync(null);
  }

  /**
   * Return the latest log entries.
   *
   * @param offset First entry to retrieve (zero-based). Log entries are sorted by decreasing date,
   *     therefore 0 designates the most recent log entry. (optional, default to 0)
   * @param length Maximum number of entries to retrieve. The maximum allowed value is 1000.
   *     (optional, default to 10)
   * @param indexName Index for which log entries should be retrieved. When omitted, log entries are
   *     retrieved across all indices. (optional)
   * @param type Type of log entries to retrieve. When omitted, all log entries are retrieved.
   *     (optional, default to all)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetLogsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetLogsResponse getLogs(
    Integer offset,
    Integer length,
    String indexName,
    LogType type,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getLogsAsync(offset, length, indexName, type, requestOptions)
    );
  }

  public GetLogsResponse getLogs(
    Integer offset,
    Integer length,
    String indexName,
    LogType type
  ) throws AlgoliaRuntimeException {
    return this.getLogs(offset, length, indexName, type, null);
  }

  public GetLogsResponse getLogs(RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.getLogs(null, null, null, null, requestOptions);
  }

  public GetLogsResponse getLogs() throws AlgoliaRuntimeException {
    return this.getLogs(null, null, null, null, null);
  }

  /**
   * (asynchronously) Return the latest log entries.
   *
   * @param offset First entry to retrieve (zero-based). Log entries are sorted by decreasing date,
   *     therefore 0 designates the most recent log entry. (optional, default to 0)
   * @param length Maximum number of entries to retrieve. The maximum allowed value is 1000.
   *     (optional, default to 10)
   * @param indexName Index for which log entries should be retrieved. When omitted, log entries are
   *     retrieved across all indices. (optional)
   * @param type Type of log entries to retrieve. When omitted, all log entries are retrieved.
   *     (optional, default to all)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetLogsResponse> getLogsAsync(
    Integer offset,
    Integer length,
    String indexName,
    LogType type,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/logs";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (offset != null) {
      queryParameters.put("offset", parameterToString(offset));
    }

    if (length != null) {
      queryParameters.put("length", parameterToString(length));
    }

    if (indexName != null) {
      queryParameters.put("indexName", parameterToString(indexName));
    }

    if (type != null) {
      queryParameters.put("type", parameterToString(type));
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
    Type returnType = new TypeToken<GetLogsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetLogsResponse> getLogsAsync(
    Integer offset,
    Integer length,
    String indexName,
    LogType type
  ) throws AlgoliaRuntimeException {
    return this.getLogsAsync(offset, length, indexName, type, null);
  }

  public CompletableFuture<GetLogsResponse> getLogsAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.getLogsAsync(null, null, null, null, requestOptions);
  }

  public CompletableFuture<GetLogsResponse> getLogsAsync()
    throws AlgoliaRuntimeException {
    return this.getLogsAsync(null, null, null, null, null);
  }

  /**
   * Retrieve one object from the index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param attributesToRetrieve List of attributes to retrieve. If not specified, all retrievable
   *     attributes are returned. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Map&lt;String, String&gt;
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Map<String, String> getObject(
    String indexName,
    String objectID,
    List<String> attributesToRetrieve,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getObjectAsync(indexName, objectID, attributesToRetrieve, requestOptions)
    );
  }

  public Map<String, String> getObject(
    String indexName,
    String objectID,
    List<String> attributesToRetrieve
  ) throws AlgoliaRuntimeException {
    return this.getObject(indexName, objectID, attributesToRetrieve, null);
  }

  public Map<String, String> getObject(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.getObject(indexName, objectID, null, requestOptions);
  }

  public Map<String, String> getObject(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return this.getObject(indexName, objectID, null, null);
  }

  /**
   * (asynchronously) Retrieve one object from the index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param attributesToRetrieve List of attributes to retrieve. If not specified, all retrievable
   *     attributes are returned. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Map<String, String>> getObjectAsync(
    String indexName,
    String objectID,
    List<String> attributesToRetrieve,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling getObject(Async)"
      );
    }

    if (objectID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'objectID' when calling getObject(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/{objectID}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{objectID\\}", this.escapeString(objectID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (attributesToRetrieve != null) {
      queryParameters.put(
        "attributesToRetrieve",
        parameterToString(attributesToRetrieve)
      );
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
    Type returnType = new TypeToken<Map<String, String>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<Map<String, String>> getObjectAsync(
    String indexName,
    String objectID,
    List<String> attributesToRetrieve
  ) throws AlgoliaRuntimeException {
    return this.getObjectAsync(indexName, objectID, attributesToRetrieve, null);
  }

  public CompletableFuture<Map<String, String>> getObjectAsync(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.getObjectAsync(indexName, objectID, null, requestOptions);
  }

  public CompletableFuture<Map<String, String>> getObjectAsync(
    String indexName,
    String objectID
  ) throws AlgoliaRuntimeException {
    return this.getObjectAsync(indexName, objectID, null, null);
  }

  /**
   * Retrieve one or more objects, potentially from different indices, in a single API call.
   *
   * @param getObjectsParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetObjectsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetObjectsResponse getObjects(
    GetObjectsParams getObjectsParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getObjectsAsync(getObjectsParams, requestOptions)
    );
  }

  public GetObjectsResponse getObjects(GetObjectsParams getObjectsParams)
    throws AlgoliaRuntimeException {
    return this.getObjects(getObjectsParams, null);
  }

  /**
   * (asynchronously) Retrieve one or more objects, potentially from different indices, in a single
   * API call.
   *
   * @param getObjectsParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetObjectsResponse> getObjectsAsync(
    GetObjectsParams getObjectsParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (getObjectsParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'getObjectsParams' when calling getObjects(Async)"
      );
    }

    Object bodyObj = getObjectsParams;

    // create path and map variables
    String requestPath = "/1/indexes/*/objects";

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
    Type returnType = new TypeToken<GetObjectsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetObjectsResponse> getObjectsAsync(
    GetObjectsParams getObjectsParams
  ) throws AlgoliaRuntimeException {
    return this.getObjectsAsync(getObjectsParams, null);
  }

  /**
   * Retrieve the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return Rule
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Rule getRule(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getRuleAsync(indexName, objectID, requestOptions)
    );
  }

  public Rule getRule(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return this.getRule(indexName, objectID, null);
  }

  /**
   * (asynchronously) Retrieve the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Rule> getRuleAsync(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling getRule(Async)"
      );
    }

    if (objectID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'objectID' when calling getRule(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/rules/{objectID}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{objectID\\}", this.escapeString(objectID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<Rule>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<Rule> getRuleAsync(
    String indexName,
    String objectID
  ) throws AlgoliaRuntimeException {
    return this.getRuleAsync(indexName, objectID, null);
  }

  /**
   * Retrieve settings of an index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return IndexSettings
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public IndexSettings getSettings(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getSettingsAsync(indexName, requestOptions));
  }

  public IndexSettings getSettings(String indexName)
    throws AlgoliaRuntimeException {
    return this.getSettings(indexName, null);
  }

  /**
   * (asynchronously) Retrieve settings of an index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<IndexSettings> getSettingsAsync(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling getSettings(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/settings".replaceAll(
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
          requestOptions
        );
    Type returnType = new TypeToken<IndexSettings>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<IndexSettings> getSettingsAsync(String indexName)
    throws AlgoliaRuntimeException {
    return this.getSettingsAsync(indexName, null);
  }

  /**
   * List all allowed sources.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return List&lt;Source&gt;
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public List<Source> getSources(RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getSourcesAsync(requestOptions));
  }

  public List<Source> getSources() throws AlgoliaRuntimeException {
    return this.getSources(null);
  }

  /**
   * (asynchronously) List all allowed sources.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<List<Source>> getSourcesAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/security/sources";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<List<Source>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<List<Source>> getSourcesAsync()
    throws AlgoliaRuntimeException {
    return this.getSourcesAsync(null);
  }

  /**
   * Fetch a synonym object identified by its objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SynonymHit
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SynonymHit getSynonym(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getSynonymAsync(indexName, objectID, requestOptions)
    );
  }

  public SynonymHit getSynonym(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return this.getSynonym(indexName, objectID, null);
  }

  /**
   * (asynchronously) Fetch a synonym object identified by its objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SynonymHit> getSynonymAsync(
    String indexName,
    String objectID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling getSynonym(Async)"
      );
    }

    if (objectID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'objectID' when calling getSynonym(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/synonyms/{objectID}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{objectID\\}", this.escapeString(objectID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<SynonymHit>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SynonymHit> getSynonymAsync(
    String indexName,
    String objectID
  ) throws AlgoliaRuntimeException {
    return this.getSynonymAsync(indexName, objectID, null);
  }

  /**
   * Check the current status of a given task.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param taskID Unique identifier of an task. Numeric value (up to 64bits). (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetTaskResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetTaskResponse getTask(
    String indexName,
    Long taskID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getTaskAsync(indexName, taskID, requestOptions)
    );
  }

  public GetTaskResponse getTask(String indexName, Long taskID)
    throws AlgoliaRuntimeException {
    return this.getTask(indexName, taskID, null);
  }

  /**
   * (asynchronously) Check the current status of a given task.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param taskID Unique identifier of an task. Numeric value (up to 64bits). (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetTaskResponse> getTaskAsync(
    String indexName,
    Long taskID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling getTask(Async)"
      );
    }

    if (taskID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'taskID' when calling getTask(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/task/{taskID}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{taskID\\}", this.escapeString(taskID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<GetTaskResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetTaskResponse> getTaskAsync(
    String indexName,
    Long taskID
  ) throws AlgoliaRuntimeException {
    return this.getTaskAsync(indexName, taskID, null);
  }

  /**
   * Get the top 10 userIDs with the highest number of records per cluster. The data returned will
   * usually be a few seconds behind real time, because userID usage may take up to a few seconds to
   * propagate to the different clusters. Upon success, the response is 200 OK and contains the
   * following array of userIDs and clusters.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return GetTopUserIdsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetTopUserIdsResponse getTopUserIds(RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getTopUserIdsAsync(requestOptions));
  }

  public GetTopUserIdsResponse getTopUserIds() throws AlgoliaRuntimeException {
    return this.getTopUserIds(null);
  }

  /**
   * (asynchronously) Get the top 10 userIDs with the highest number of records per cluster. The
   * data returned will usually be a few seconds behind real time, because userID usage may take up
   * to a few seconds to propagate to the different clusters. Upon success, the response is 200 OK
   * and contains the following array of userIDs and clusters.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetTopUserIdsResponse> getTopUserIdsAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/clusters/mapping/top";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<GetTopUserIdsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<GetTopUserIdsResponse> getTopUserIdsAsync()
    throws AlgoliaRuntimeException {
    return this.getTopUserIdsAsync(null);
  }

  /**
   * Returns the userID data stored in the mapping. The data returned will usually be a few seconds
   * behind real time, because userID usage may take up to a few seconds to propagate to the
   * different clusters. Upon success, the response is 200 OK and contains the following userID
   * data.
   *
   * @param userID userID to assign. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UserId
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UserId getUserId(String userID, RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getUserIdAsync(userID, requestOptions));
  }

  public UserId getUserId(String userID) throws AlgoliaRuntimeException {
    return this.getUserId(userID, null);
  }

  /**
   * (asynchronously) Returns the userID data stored in the mapping. The data returned will usually
   * be a few seconds behind real time, because userID usage may take up to a few seconds to
   * propagate to the different clusters. Upon success, the response is 200 OK and contains the
   * following userID data.
   *
   * @param userID userID to assign. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UserId> getUserIdAsync(
    String userID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (userID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'userID' when calling getUserId(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/clusters/mapping/{userID}".replaceAll(
          "\\{userID\\}",
          this.escapeString(userID.toString())
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
          requestOptions
        );
    Type returnType = new TypeToken<UserId>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UserId> getUserIdAsync(String userID)
    throws AlgoliaRuntimeException {
    return this.getUserIdAsync(userID, null);
  }

  /**
   * Get the status of your clusters' migrations or user creations. Creating a large batch of users
   * or migrating your multi-cluster may take quite some time. This method lets you retrieve the
   * status of the migration, so you can know when it's done. Upon success, the response is 200 OK.
   * A successful response indicates that the operation has been taken into account, and the userIDs
   * are directly usable.
   *
   * @param getClusters Whether to get clusters or not. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return CreatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public CreatedAtResponse hasPendingMappings(
    Boolean getClusters,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      hasPendingMappingsAsync(getClusters, requestOptions)
    );
  }

  public CreatedAtResponse hasPendingMappings(Boolean getClusters)
    throws AlgoliaRuntimeException {
    return this.hasPendingMappings(getClusters, null);
  }

  public CreatedAtResponse hasPendingMappings(RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.hasPendingMappings(null, requestOptions);
  }

  public CreatedAtResponse hasPendingMappings() throws AlgoliaRuntimeException {
    return this.hasPendingMappings(null, null);
  }

  /**
   * (asynchronously) Get the status of your clusters&#39; migrations or user creations. Creating a
   * large batch of users or migrating your multi-cluster may take quite some time. This method lets
   * you retrieve the status of the migration, so you can know when it&#39;s done. Upon success, the
   * response is 200 OK. A successful response indicates that the operation has been taken into
   * account, and the userIDs are directly usable.
   *
   * @param getClusters Whether to get clusters or not. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<CreatedAtResponse> hasPendingMappingsAsync(
    Boolean getClusters,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/clusters/mapping/pending";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (getClusters != null) {
      queryParameters.put("getClusters", parameterToString(getClusters));
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
    Type returnType = new TypeToken<CreatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<CreatedAtResponse> hasPendingMappingsAsync(
    Boolean getClusters
  ) throws AlgoliaRuntimeException {
    return this.hasPendingMappingsAsync(getClusters, null);
  }

  public CompletableFuture<CreatedAtResponse> hasPendingMappingsAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.hasPendingMappingsAsync(null, requestOptions);
  }

  public CompletableFuture<CreatedAtResponse> hasPendingMappingsAsync()
    throws AlgoliaRuntimeException {
    return this.hasPendingMappingsAsync(null, null);
  }

  /**
   * List API keys, along with their associated rights.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return ListApiKeysResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ListApiKeysResponse listApiKeys(RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(listApiKeysAsync(requestOptions));
  }

  public ListApiKeysResponse listApiKeys() throws AlgoliaRuntimeException {
    return this.listApiKeys(null);
  }

  /**
   * (asynchronously) List API keys, along with their associated rights.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ListApiKeysResponse> listApiKeysAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/keys";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<ListApiKeysResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<ListApiKeysResponse> listApiKeysAsync()
    throws AlgoliaRuntimeException {
    return this.listApiKeysAsync(null);
  }

  /**
   * List the clusters available in a multi-clusters setup for a single appID. Upon success, the
   * response is 200 OK and contains the following clusters.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return ListClustersResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ListClustersResponse listClusters(RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(listClustersAsync(requestOptions));
  }

  public ListClustersResponse listClusters() throws AlgoliaRuntimeException {
    return this.listClusters(null);
  }

  /**
   * (asynchronously) List the clusters available in a multi-clusters setup for a single appID. Upon
   * success, the response is 200 OK and contains the following clusters.
   *
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ListClustersResponse> listClustersAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/clusters";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "GET",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<ListClustersResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<ListClustersResponse> listClustersAsync()
    throws AlgoliaRuntimeException {
    return this.listClustersAsync(null);
  }

  /**
   * List existing indexes from an application.
   *
   * @param page Requested page (zero-based). When specified, will retrieve a specific page; the
   *     page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   *     (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return ListIndicesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ListIndicesResponse listIndices(
    Integer page,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(listIndicesAsync(page, requestOptions));
  }

  public ListIndicesResponse listIndices(Integer page)
    throws AlgoliaRuntimeException {
    return this.listIndices(page, null);
  }

  public ListIndicesResponse listIndices(RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.listIndices(null, requestOptions);
  }

  public ListIndicesResponse listIndices() throws AlgoliaRuntimeException {
    return this.listIndices(null, null);
  }

  /**
   * (asynchronously) List existing indexes from an application.
   *
   * @param page Requested page (zero-based). When specified, will retrieve a specific page; the
   *     page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   *     (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ListIndicesResponse> listIndicesAsync(
    Integer page,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/indexes";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (page != null) {
      queryParameters.put("page", parameterToString(page));
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
    Type returnType = new TypeToken<ListIndicesResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<ListIndicesResponse> listIndicesAsync(Integer page)
    throws AlgoliaRuntimeException {
    return this.listIndicesAsync(page, null);
  }

  public CompletableFuture<ListIndicesResponse> listIndicesAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.listIndicesAsync(null, requestOptions);
  }

  public CompletableFuture<ListIndicesResponse> listIndicesAsync()
    throws AlgoliaRuntimeException {
    return this.listIndicesAsync(null, null);
  }

  /**
   * List the userIDs assigned to a multi-clusters appID. The data returned will usually be a few
   * seconds behind real time, because userID usage may take up to a few seconds to propagate to the
   * different clusters. Upon success, the response is 200 OK and contains the following userIDs
   * data.
   *
   * @param page Requested page (zero-based). When specified, will retrieve a specific page; the
   *     page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   *     (optional)
   * @param hitsPerPage Maximum number of objects to retrieve. (optional, default to 100)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return ListUserIdsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ListUserIdsResponse listUserIds(
    Integer page,
    Integer hitsPerPage,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      listUserIdsAsync(page, hitsPerPage, requestOptions)
    );
  }

  public ListUserIdsResponse listUserIds(Integer page, Integer hitsPerPage)
    throws AlgoliaRuntimeException {
    return this.listUserIds(page, hitsPerPage, null);
  }

  public ListUserIdsResponse listUserIds(RequestOptions requestOptions)
    throws AlgoliaRuntimeException {
    return this.listUserIds(null, null, requestOptions);
  }

  public ListUserIdsResponse listUserIds() throws AlgoliaRuntimeException {
    return this.listUserIds(null, null, null);
  }

  /**
   * (asynchronously) List the userIDs assigned to a multi-clusters appID. The data returned will
   * usually be a few seconds behind real time, because userID usage may take up to a few seconds to
   * propagate to the different clusters. Upon success, the response is 200 OK and contains the
   * following userIDs data.
   *
   * @param page Requested page (zero-based). When specified, will retrieve a specific page; the
   *     page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   *     (optional)
   * @param hitsPerPage Maximum number of objects to retrieve. (optional, default to 100)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ListUserIdsResponse> listUserIdsAsync(
    Integer page,
    Integer hitsPerPage,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/clusters/mapping";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (page != null) {
      queryParameters.put("page", parameterToString(page));
    }

    if (hitsPerPage != null) {
      queryParameters.put("hitsPerPage", parameterToString(hitsPerPage));
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
    Type returnType = new TypeToken<ListUserIdsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<ListUserIdsResponse> listUserIdsAsync(
    Integer page,
    Integer hitsPerPage
  ) throws AlgoliaRuntimeException {
    return this.listUserIdsAsync(page, hitsPerPage, null);
  }

  public CompletableFuture<ListUserIdsResponse> listUserIdsAsync(
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.listUserIdsAsync(null, null, requestOptions);
  }

  public CompletableFuture<ListUserIdsResponse> listUserIdsAsync()
    throws AlgoliaRuntimeException {
    return this.listUserIdsAsync(null, null, null);
  }

  /**
   * Perform multiple write operations, potentially targeting multiple indices, in a single API
   * call.
   *
   * @param batchParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return MultipleBatchResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public MultipleBatchResponse multipleBatch(
    BatchParams batchParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      multipleBatchAsync(batchParams, requestOptions)
    );
  }

  public MultipleBatchResponse multipleBatch(BatchParams batchParams)
    throws AlgoliaRuntimeException {
    return this.multipleBatch(batchParams, null);
  }

  /**
   * (asynchronously) Perform multiple write operations, potentially targeting multiple indices, in
   * a single API call.
   *
   * @param batchParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<MultipleBatchResponse> multipleBatchAsync(
    BatchParams batchParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (batchParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'batchParams' when calling multipleBatch(Async)"
      );
    }

    Object bodyObj = batchParams;

    // create path and map variables
    String requestPath = "/1/indexes/*/batch";

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
    Type returnType = new TypeToken<MultipleBatchResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<MultipleBatchResponse> multipleBatchAsync(
    BatchParams batchParams
  ) throws AlgoliaRuntimeException {
    return this.multipleBatchAsync(batchParams, null);
  }

  /**
   * Perform a search operation targeting one or many indices.
   *
   * @param multipleQueriesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return MultipleQueriesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public MultipleQueriesResponse multipleQueries(
    MultipleQueriesParams multipleQueriesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      multipleQueriesAsync(multipleQueriesParams, requestOptions)
    );
  }

  public MultipleQueriesResponse multipleQueries(
    MultipleQueriesParams multipleQueriesParams
  ) throws AlgoliaRuntimeException {
    return this.multipleQueries(multipleQueriesParams, null);
  }

  /**
   * (asynchronously) Perform a search operation targeting one or many indices.
   *
   * @param multipleQueriesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<MultipleQueriesResponse> multipleQueriesAsync(
    MultipleQueriesParams multipleQueriesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (multipleQueriesParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'multipleQueriesParams' when calling" +
        " multipleQueries(Async)"
      );
    }

    Object bodyObj = multipleQueriesParams;

    // create path and map variables
    String requestPath = "/1/indexes/*/queries";

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
    Type returnType = new TypeToken<MultipleQueriesResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<MultipleQueriesResponse> multipleQueriesAsync(
    MultipleQueriesParams multipleQueriesParams
  ) throws AlgoliaRuntimeException {
    return this.multipleQueriesAsync(multipleQueriesParams, null);
  }

  /**
   * Peforms a copy or a move operation on a index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param operationIndexParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse operationIndex(
    String indexName,
    OperationIndexParams operationIndexParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      operationIndexAsync(indexName, operationIndexParams, requestOptions)
    );
  }

  public UpdatedAtResponse operationIndex(
    String indexName,
    OperationIndexParams operationIndexParams
  ) throws AlgoliaRuntimeException {
    return this.operationIndex(indexName, operationIndexParams, null);
  }

  /**
   * (asynchronously) Peforms a copy or a move operation on a index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param operationIndexParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> operationIndexAsync(
    String indexName,
    OperationIndexParams operationIndexParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling operationIndex(Async)"
      );
    }

    if (operationIndexParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'operationIndexParams' when calling" +
        " operationIndex(Async)"
      );
    }

    Object bodyObj = operationIndexParams;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/operation".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

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
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> operationIndexAsync(
    String indexName,
    OperationIndexParams operationIndexParams
  ) throws AlgoliaRuntimeException {
    return this.operationIndexAsync(indexName, operationIndexParams, null);
  }

  /**
   * Update one or more attributes of an existing object. This method lets you update only a part of
   * an existing object, either by adding new attributes or updating existing ones. You can
   * partially update several objects in a single method call. If the index targeted by this
   * operation doesn't exist yet, it's automatically created.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param attributeOrBuiltInOperation List of attributes to update. (required)
   * @param createIfNotExists Creates the record if it does not exist yet. (optional, default to
   *     true)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtWithObjectIdResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtWithObjectIdResponse partialUpdateObject(
    String indexName,
    String objectID,
    List<Map<String, AttributeOrBuiltInOperation>> attributeOrBuiltInOperation,
    Boolean createIfNotExists,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      partialUpdateObjectAsync(
        indexName,
        objectID,
        attributeOrBuiltInOperation,
        createIfNotExists,
        requestOptions
      )
    );
  }

  public UpdatedAtWithObjectIdResponse partialUpdateObject(
    String indexName,
    String objectID,
    List<Map<String, AttributeOrBuiltInOperation>> attributeOrBuiltInOperation,
    Boolean createIfNotExists
  ) throws AlgoliaRuntimeException {
    return this.partialUpdateObject(
        indexName,
        objectID,
        attributeOrBuiltInOperation,
        createIfNotExists,
        null
      );
  }

  public UpdatedAtWithObjectIdResponse partialUpdateObject(
    String indexName,
    String objectID,
    List<Map<String, AttributeOrBuiltInOperation>> attributeOrBuiltInOperation,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.partialUpdateObject(
        indexName,
        objectID,
        attributeOrBuiltInOperation,
        null,
        requestOptions
      );
  }

  public UpdatedAtWithObjectIdResponse partialUpdateObject(
    String indexName,
    String objectID,
    List<Map<String, AttributeOrBuiltInOperation>> attributeOrBuiltInOperation
  ) throws AlgoliaRuntimeException {
    return this.partialUpdateObject(
        indexName,
        objectID,
        attributeOrBuiltInOperation,
        null,
        null
      );
  }

  /**
   * (asynchronously) Update one or more attributes of an existing object. This method lets you
   * update only a part of an existing object, either by adding new attributes or updating existing
   * ones. You can partially update several objects in a single method call. If the index targeted
   * by this operation doesn&#39;t exist yet, it&#39;s automatically created.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param attributeOrBuiltInOperation List of attributes to update. (required)
   * @param createIfNotExists Creates the record if it does not exist yet. (optional, default to
   *     true)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtWithObjectIdResponse> partialUpdateObjectAsync(
    String indexName,
    String objectID,
    List<Map<String, AttributeOrBuiltInOperation>> attributeOrBuiltInOperation,
    Boolean createIfNotExists,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling partialUpdateObject(Async)"
      );
    }

    if (objectID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'objectID' when calling partialUpdateObject(Async)"
      );
    }

    if (attributeOrBuiltInOperation == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'attributeOrBuiltInOperation' when calling" +
        " partialUpdateObject(Async)"
      );
    }

    Object bodyObj = attributeOrBuiltInOperation;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/{objectID}/partial".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{objectID\\}", this.escapeString(objectID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (createIfNotExists != null) {
      queryParameters.put(
        "createIfNotExists",
        parameterToString(createIfNotExists)
      );
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
    Type returnType = new TypeToken<UpdatedAtWithObjectIdResponse>() {}
      .getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtWithObjectIdResponse> partialUpdateObjectAsync(
    String indexName,
    String objectID,
    List<Map<String, AttributeOrBuiltInOperation>> attributeOrBuiltInOperation,
    Boolean createIfNotExists
  ) throws AlgoliaRuntimeException {
    return this.partialUpdateObjectAsync(
        indexName,
        objectID,
        attributeOrBuiltInOperation,
        createIfNotExists,
        null
      );
  }

  public CompletableFuture<UpdatedAtWithObjectIdResponse> partialUpdateObjectAsync(
    String indexName,
    String objectID,
    List<Map<String, AttributeOrBuiltInOperation>> attributeOrBuiltInOperation,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.partialUpdateObjectAsync(
        indexName,
        objectID,
        attributeOrBuiltInOperation,
        null,
        requestOptions
      );
  }

  public CompletableFuture<UpdatedAtWithObjectIdResponse> partialUpdateObjectAsync(
    String indexName,
    String objectID,
    List<Map<String, AttributeOrBuiltInOperation>> attributeOrBuiltInOperation
  ) throws AlgoliaRuntimeException {
    return this.partialUpdateObjectAsync(
        indexName,
        objectID,
        attributeOrBuiltInOperation,
        null,
        null
      );
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

  /**
   * Remove a userID and its associated data from the multi-clusters. Upon success, the response is
   * 200 OK and a task is created to remove the userID data and mapping.
   *
   * @param userID userID to assign. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return RemoveUserIdResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public RemoveUserIdResponse removeUserId(
    String userID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(removeUserIdAsync(userID, requestOptions));
  }

  public RemoveUserIdResponse removeUserId(String userID)
    throws AlgoliaRuntimeException {
    return this.removeUserId(userID, null);
  }

  /**
   * (asynchronously) Remove a userID and its associated data from the multi-clusters. Upon success,
   * the response is 200 OK and a task is created to remove the userID data and mapping.
   *
   * @param userID userID to assign. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<RemoveUserIdResponse> removeUserIdAsync(
    String userID,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (userID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'userID' when calling removeUserId(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/clusters/mapping/{userID}".replaceAll(
          "\\{userID\\}",
          this.escapeString(userID.toString())
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
          requestOptions
        );
    Type returnType = new TypeToken<RemoveUserIdResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<RemoveUserIdResponse> removeUserIdAsync(
    String userID
  ) throws AlgoliaRuntimeException {
    return this.removeUserIdAsync(userID, null);
  }

  /**
   * Replace all allowed sources.
   *
   * @param source The sources to allow. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return ReplaceSourceResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ReplaceSourceResponse replaceSources(
    List<Source> source,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(replaceSourcesAsync(source, requestOptions));
  }

  public ReplaceSourceResponse replaceSources(List<Source> source)
    throws AlgoliaRuntimeException {
    return this.replaceSources(source, null);
  }

  /**
   * (asynchronously) Replace all allowed sources.
   *
   * @param source The sources to allow. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ReplaceSourceResponse> replaceSourcesAsync(
    List<Source> source,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (source == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'source' when calling replaceSources(Async)"
      );
    }

    Object bodyObj = source;

    // create path and map variables
    String requestPath = "/1/security/sources";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "PUT",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<ReplaceSourceResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<ReplaceSourceResponse> replaceSourcesAsync(
    List<Source> source
  ) throws AlgoliaRuntimeException {
    return this.replaceSourcesAsync(source, null);
  }

  /**
   * Restore a deleted API key, along with its associated rights.
   *
   * @param key API Key string. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return AddApiKeyResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public AddApiKeyResponse restoreApiKey(
    String key,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(restoreApiKeyAsync(key, requestOptions));
  }

  public AddApiKeyResponse restoreApiKey(String key)
    throws AlgoliaRuntimeException {
    return this.restoreApiKey(key, null);
  }

  /**
   * (asynchronously) Restore a deleted API key, along with its associated rights.
   *
   * @param key API Key string. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<AddApiKeyResponse> restoreApiKeyAsync(
    String key,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (key == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'key' when calling restoreApiKey(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/keys/{key}/restore".replaceAll(
          "\\{key\\}",
          this.escapeString(key.toString())
        );

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
    Type returnType = new TypeToken<AddApiKeyResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<AddApiKeyResponse> restoreApiKeyAsync(String key)
    throws AlgoliaRuntimeException {
    return this.restoreApiKeyAsync(key, null);
  }

  /**
   * Add an object to the index, automatically assigning it an object ID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param body The Algolia record. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SaveObjectResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SaveObjectResponse saveObject(
    String indexName,
    Object body,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      saveObjectAsync(indexName, body, requestOptions)
    );
  }

  public SaveObjectResponse saveObject(String indexName, Object body)
    throws AlgoliaRuntimeException {
    return this.saveObject(indexName, body, null);
  }

  /**
   * (asynchronously) Add an object to the index, automatically assigning it an object ID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param body The Algolia record. (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SaveObjectResponse> saveObjectAsync(
    String indexName,
    Object body,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling saveObject(Async)"
      );
    }

    if (body == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'body' when calling saveObject(Async)"
      );
    }

    Object bodyObj = body;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

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
    Type returnType = new TypeToken<SaveObjectResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SaveObjectResponse> saveObjectAsync(
    String indexName,
    Object body
  ) throws AlgoliaRuntimeException {
    return this.saveObjectAsync(indexName, body, null);
  }

  /**
   * Create or update the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param rule (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedRuleResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedRuleResponse saveRule(
    String indexName,
    String objectID,
    Rule rule,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      saveRuleAsync(
        indexName,
        objectID,
        rule,
        forwardToReplicas,
        requestOptions
      )
    );
  }

  public UpdatedRuleResponse saveRule(
    String indexName,
    String objectID,
    Rule rule,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.saveRule(indexName, objectID, rule, forwardToReplicas, null);
  }

  public UpdatedRuleResponse saveRule(
    String indexName,
    String objectID,
    Rule rule,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.saveRule(indexName, objectID, rule, null, requestOptions);
  }

  public UpdatedRuleResponse saveRule(
    String indexName,
    String objectID,
    Rule rule
  ) throws AlgoliaRuntimeException {
    return this.saveRule(indexName, objectID, rule, null, null);
  }

  /**
   * (asynchronously) Create or update the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param rule (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedRuleResponse> saveRuleAsync(
    String indexName,
    String objectID,
    Rule rule,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling saveRule(Async)"
      );
    }

    if (objectID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'objectID' when calling saveRule(Async)"
      );
    }

    if (rule == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'rule' when calling saveRule(Async)"
      );
    }

    Object bodyObj = rule;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/rules/{objectID}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{objectID\\}", this.escapeString(objectID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParameters.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
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
    Type returnType = new TypeToken<UpdatedRuleResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedRuleResponse> saveRuleAsync(
    String indexName,
    String objectID,
    Rule rule,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.saveRuleAsync(
        indexName,
        objectID,
        rule,
        forwardToReplicas,
        null
      );
  }

  public CompletableFuture<UpdatedRuleResponse> saveRuleAsync(
    String indexName,
    String objectID,
    Rule rule,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.saveRuleAsync(indexName, objectID, rule, null, requestOptions);
  }

  public CompletableFuture<UpdatedRuleResponse> saveRuleAsync(
    String indexName,
    String objectID,
    Rule rule
  ) throws AlgoliaRuntimeException {
    return this.saveRuleAsync(indexName, objectID, rule, null, null);
  }

  /**
   * Create a new synonym object or update the existing synonym object with the given object ID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param synonymHit (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SaveSynonymResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SaveSynonymResponse saveSynonym(
    String indexName,
    String objectID,
    SynonymHit synonymHit,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      saveSynonymAsync(
        indexName,
        objectID,
        synonymHit,
        forwardToReplicas,
        requestOptions
      )
    );
  }

  public SaveSynonymResponse saveSynonym(
    String indexName,
    String objectID,
    SynonymHit synonymHit,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.saveSynonym(
        indexName,
        objectID,
        synonymHit,
        forwardToReplicas,
        null
      );
  }

  public SaveSynonymResponse saveSynonym(
    String indexName,
    String objectID,
    SynonymHit synonymHit,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.saveSynonym(
        indexName,
        objectID,
        synonymHit,
        null,
        requestOptions
      );
  }

  public SaveSynonymResponse saveSynonym(
    String indexName,
    String objectID,
    SynonymHit synonymHit
  ) throws AlgoliaRuntimeException {
    return this.saveSynonym(indexName, objectID, synonymHit, null, null);
  }

  /**
   * (asynchronously) Create a new synonym object or update the existing synonym object with the
   * given object ID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param synonymHit (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SaveSynonymResponse> saveSynonymAsync(
    String indexName,
    String objectID,
    SynonymHit synonymHit,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling saveSynonym(Async)"
      );
    }

    if (objectID == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'objectID' when calling saveSynonym(Async)"
      );
    }

    if (synonymHit == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'synonymHit' when calling saveSynonym(Async)"
      );
    }

    Object bodyObj = synonymHit;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/synonyms/{objectID}".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{objectID\\}", this.escapeString(objectID.toString()));

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParameters.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
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
    Type returnType = new TypeToken<SaveSynonymResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SaveSynonymResponse> saveSynonymAsync(
    String indexName,
    String objectID,
    SynonymHit synonymHit,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.saveSynonymAsync(
        indexName,
        objectID,
        synonymHit,
        forwardToReplicas,
        null
      );
  }

  public CompletableFuture<SaveSynonymResponse> saveSynonymAsync(
    String indexName,
    String objectID,
    SynonymHit synonymHit,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.saveSynonymAsync(
        indexName,
        objectID,
        synonymHit,
        null,
        requestOptions
      );
  }

  public CompletableFuture<SaveSynonymResponse> saveSynonymAsync(
    String indexName,
    String objectID,
    SynonymHit synonymHit
  ) throws AlgoliaRuntimeException {
    return this.saveSynonymAsync(indexName, objectID, synonymHit, null, null);
  }

  /**
   * Create/update multiple synonym objects at once, potentially replacing the entire list of
   * synonyms if replaceExistingSynonyms is true.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param synonymHit (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param replaceExistingSynonyms Replace all synonyms of the index with the ones sent with this
   *     request. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse saveSynonyms(
    String indexName,
    List<SynonymHit> synonymHit,
    Boolean forwardToReplicas,
    Boolean replaceExistingSynonyms,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      saveSynonymsAsync(
        indexName,
        synonymHit,
        forwardToReplicas,
        replaceExistingSynonyms,
        requestOptions
      )
    );
  }

  public UpdatedAtResponse saveSynonyms(
    String indexName,
    List<SynonymHit> synonymHit,
    Boolean forwardToReplicas,
    Boolean replaceExistingSynonyms
  ) throws AlgoliaRuntimeException {
    return this.saveSynonyms(
        indexName,
        synonymHit,
        forwardToReplicas,
        replaceExistingSynonyms,
        null
      );
  }

  public UpdatedAtResponse saveSynonyms(
    String indexName,
    List<SynonymHit> synonymHit,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.saveSynonyms(indexName, synonymHit, null, null, requestOptions);
  }

  public UpdatedAtResponse saveSynonyms(
    String indexName,
    List<SynonymHit> synonymHit
  ) throws AlgoliaRuntimeException {
    return this.saveSynonyms(indexName, synonymHit, null, null, null);
  }

  /**
   * (asynchronously) Create/update multiple synonym objects at once, potentially replacing the
   * entire list of synonyms if replaceExistingSynonyms is true.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param synonymHit (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param replaceExistingSynonyms Replace all synonyms of the index with the ones sent with this
   *     request. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> saveSynonymsAsync(
    String indexName,
    List<SynonymHit> synonymHit,
    Boolean forwardToReplicas,
    Boolean replaceExistingSynonyms,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling saveSynonyms(Async)"
      );
    }

    if (synonymHit == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'synonymHit' when calling saveSynonyms(Async)"
      );
    }

    Object bodyObj = synonymHit;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/synonyms/batch".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParameters.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    if (replaceExistingSynonyms != null) {
      queryParameters.put(
        "replaceExistingSynonyms",
        parameterToString(replaceExistingSynonyms)
      );
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
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> saveSynonymsAsync(
    String indexName,
    List<SynonymHit> synonymHit,
    Boolean forwardToReplicas,
    Boolean replaceExistingSynonyms
  ) throws AlgoliaRuntimeException {
    return this.saveSynonymsAsync(
        indexName,
        synonymHit,
        forwardToReplicas,
        replaceExistingSynonyms,
        null
      );
  }

  public CompletableFuture<UpdatedAtResponse> saveSynonymsAsync(
    String indexName,
    List<SynonymHit> synonymHit,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.saveSynonymsAsync(
        indexName,
        synonymHit,
        null,
        null,
        requestOptions
      );
  }

  public CompletableFuture<UpdatedAtResponse> saveSynonymsAsync(
    String indexName,
    List<SynonymHit> synonymHit
  ) throws AlgoliaRuntimeException {
    return this.saveSynonymsAsync(indexName, synonymHit, null, null, null);
  }

  /**
   * Perform a search operation targeting one specific index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SearchResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SearchResponse search(
    String indexName,
    SearchParams searchParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      searchAsync(indexName, searchParams, requestOptions)
    );
  }

  public SearchResponse search(String indexName, SearchParams searchParams)
    throws AlgoliaRuntimeException {
    return this.search(indexName, searchParams, null);
  }

  /**
   * (asynchronously) Perform a search operation targeting one specific index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SearchResponse> searchAsync(
    String indexName,
    SearchParams searchParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling search(Async)"
      );
    }

    if (searchParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'searchParams' when calling search(Async)"
      );
    }

    Object bodyObj = searchParams;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/query".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

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
    Type returnType = new TypeToken<SearchResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SearchResponse> searchAsync(
    String indexName,
    SearchParams searchParams
  ) throws AlgoliaRuntimeException {
    return this.searchAsync(indexName, searchParams, null);
  }

  /**
   * Search the dictionary entries.
   *
   * @param dictionaryName The dictionary to search in. (required)
   * @param searchDictionaryEntriesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse searchDictionaryEntries(
    DictionaryType dictionaryName,
    SearchDictionaryEntriesParams searchDictionaryEntriesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      searchDictionaryEntriesAsync(
        dictionaryName,
        searchDictionaryEntriesParams,
        requestOptions
      )
    );
  }

  public UpdatedAtResponse searchDictionaryEntries(
    DictionaryType dictionaryName,
    SearchDictionaryEntriesParams searchDictionaryEntriesParams
  ) throws AlgoliaRuntimeException {
    return this.searchDictionaryEntries(
        dictionaryName,
        searchDictionaryEntriesParams,
        null
      );
  }

  /**
   * (asynchronously) Search the dictionary entries.
   *
   * @param dictionaryName The dictionary to search in. (required)
   * @param searchDictionaryEntriesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> searchDictionaryEntriesAsync(
    DictionaryType dictionaryName,
    SearchDictionaryEntriesParams searchDictionaryEntriesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (dictionaryName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'dictionaryName' when calling" +
        " searchDictionaryEntries(Async)"
      );
    }

    if (searchDictionaryEntriesParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'searchDictionaryEntriesParams' when calling" +
        " searchDictionaryEntries(Async)"
      );
    }

    Object bodyObj = searchDictionaryEntriesParams;

    // create path and map variables
    String requestPath =
      "/1/dictionaries/{dictionaryName}/search".replaceAll(
          "\\{dictionaryName\\}",
          this.escapeString(dictionaryName.toString())
        );

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
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> searchDictionaryEntriesAsync(
    DictionaryType dictionaryName,
    SearchDictionaryEntriesParams searchDictionaryEntriesParams
  ) throws AlgoliaRuntimeException {
    return this.searchDictionaryEntriesAsync(
        dictionaryName,
        searchDictionaryEntriesParams,
        null
      );
  }

  /**
   * Search for values of a given facet, optionally restricting the returned values to those
   * contained in objects matching other search criteria.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param facetName The facet name. (required)
   * @param searchForFacetValuesRequest (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SearchForFacetValuesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SearchForFacetValuesResponse searchForFacetValues(
    String indexName,
    String facetName,
    SearchForFacetValuesRequest searchForFacetValuesRequest,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      searchForFacetValuesAsync(
        indexName,
        facetName,
        searchForFacetValuesRequest,
        requestOptions
      )
    );
  }

  public SearchForFacetValuesResponse searchForFacetValues(
    String indexName,
    String facetName,
    SearchForFacetValuesRequest searchForFacetValuesRequest
  ) throws AlgoliaRuntimeException {
    return this.searchForFacetValues(
        indexName,
        facetName,
        searchForFacetValuesRequest,
        null
      );
  }

  public SearchForFacetValuesResponse searchForFacetValues(
    String indexName,
    String facetName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.searchForFacetValues(
        indexName,
        facetName,
        null,
        requestOptions
      );
  }

  public SearchForFacetValuesResponse searchForFacetValues(
    String indexName,
    String facetName
  ) throws AlgoliaRuntimeException {
    return this.searchForFacetValues(indexName, facetName, null, null);
  }

  /**
   * (asynchronously) Search for values of a given facet, optionally restricting the returned values
   * to those contained in objects matching other search criteria.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param facetName The facet name. (required)
   * @param searchForFacetValuesRequest (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SearchForFacetValuesResponse> searchForFacetValuesAsync(
    String indexName,
    String facetName,
    SearchForFacetValuesRequest searchForFacetValuesRequest,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling searchForFacetValues(Async)"
      );
    }

    if (facetName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'facetName' when calling searchForFacetValues(Async)"
      );
    }

    Object bodyObj = searchForFacetValuesRequest;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/facets/{facetName}/query".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        )
        .replaceAll("\\{facetName\\}", this.escapeString(facetName.toString()));

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
    Type returnType = new TypeToken<SearchForFacetValuesResponse>() {}
      .getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SearchForFacetValuesResponse> searchForFacetValuesAsync(
    String indexName,
    String facetName,
    SearchForFacetValuesRequest searchForFacetValuesRequest
  ) throws AlgoliaRuntimeException {
    return this.searchForFacetValuesAsync(
        indexName,
        facetName,
        searchForFacetValuesRequest,
        null
      );
  }

  public CompletableFuture<SearchForFacetValuesResponse> searchForFacetValuesAsync(
    String indexName,
    String facetName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.searchForFacetValuesAsync(
        indexName,
        facetName,
        null,
        requestOptions
      );
  }

  public CompletableFuture<SearchForFacetValuesResponse> searchForFacetValuesAsync(
    String indexName,
    String facetName
  ) throws AlgoliaRuntimeException {
    return this.searchForFacetValuesAsync(indexName, facetName, null, null);
  }

  /**
   * Search for rules matching various criteria.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchRulesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SearchRulesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SearchRulesResponse searchRules(
    String indexName,
    SearchRulesParams searchRulesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      searchRulesAsync(indexName, searchRulesParams, requestOptions)
    );
  }

  public SearchRulesResponse searchRules(
    String indexName,
    SearchRulesParams searchRulesParams
  ) throws AlgoliaRuntimeException {
    return this.searchRules(indexName, searchRulesParams, null);
  }

  /**
   * (asynchronously) Search for rules matching various criteria.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchRulesParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SearchRulesResponse> searchRulesAsync(
    String indexName,
    SearchRulesParams searchRulesParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling searchRules(Async)"
      );
    }

    if (searchRulesParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'searchRulesParams' when calling searchRules(Async)"
      );
    }

    Object bodyObj = searchRulesParams;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/rules/search".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

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
    Type returnType = new TypeToken<SearchRulesResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SearchRulesResponse> searchRulesAsync(
    String indexName,
    SearchRulesParams searchRulesParams
  ) throws AlgoliaRuntimeException {
    return this.searchRulesAsync(indexName, searchRulesParams, null);
  }

  /**
   * Search or browse all synonyms, optionally filtering them by type.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param query Search for specific synonyms matching this string. (optional, default to )
   * @param type Only search for specific types of synonyms. (optional)
   * @param page Requested page (zero-based). When specified, will retrieve a specific page; the
   *     page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   *     (optional, default to 0)
   * @param hitsPerPage Maximum number of objects to retrieve. (optional, default to 100)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SearchSynonymsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SearchSynonymsResponse searchSynonyms(
    String indexName,
    String query,
    SynonymType type,
    Integer page,
    Integer hitsPerPage,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      searchSynonymsAsync(
        indexName,
        query,
        type,
        page,
        hitsPerPage,
        requestOptions
      )
    );
  }

  public SearchSynonymsResponse searchSynonyms(
    String indexName,
    String query,
    SynonymType type,
    Integer page,
    Integer hitsPerPage
  ) throws AlgoliaRuntimeException {
    return this.searchSynonyms(indexName, query, type, page, hitsPerPage, null);
  }

  public SearchSynonymsResponse searchSynonyms(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.searchSynonyms(
        indexName,
        null,
        null,
        null,
        null,
        requestOptions
      );
  }

  public SearchSynonymsResponse searchSynonyms(String indexName)
    throws AlgoliaRuntimeException {
    return this.searchSynonyms(indexName, null, null, null, null, null);
  }

  /**
   * (asynchronously) Search or browse all synonyms, optionally filtering them by type.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param query Search for specific synonyms matching this string. (optional, default to )
   * @param type Only search for specific types of synonyms. (optional)
   * @param page Requested page (zero-based). When specified, will retrieve a specific page; the
   *     page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   *     (optional, default to 0)
   * @param hitsPerPage Maximum number of objects to retrieve. (optional, default to 100)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SearchSynonymsResponse> searchSynonymsAsync(
    String indexName,
    String query,
    SynonymType type,
    Integer page,
    Integer hitsPerPage,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling searchSynonyms(Async)"
      );
    }

    Object bodyObj = null;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/synonyms/search".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (query != null) {
      queryParameters.put("query", parameterToString(query));
    }

    if (type != null) {
      queryParameters.put("type", parameterToString(type));
    }

    if (page != null) {
      queryParameters.put("page", parameterToString(page));
    }

    if (hitsPerPage != null) {
      queryParameters.put("hitsPerPage", parameterToString(hitsPerPage));
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
    Type returnType = new TypeToken<SearchSynonymsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SearchSynonymsResponse> searchSynonymsAsync(
    String indexName,
    String query,
    SynonymType type,
    Integer page,
    Integer hitsPerPage
  ) throws AlgoliaRuntimeException {
    return this.searchSynonymsAsync(
        indexName,
        query,
        type,
        page,
        hitsPerPage,
        null
      );
  }

  public CompletableFuture<SearchSynonymsResponse> searchSynonymsAsync(
    String indexName,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.searchSynonymsAsync(
        indexName,
        null,
        null,
        null,
        null,
        requestOptions
      );
  }

  public CompletableFuture<SearchSynonymsResponse> searchSynonymsAsync(
    String indexName
  ) throws AlgoliaRuntimeException {
    return this.searchSynonymsAsync(indexName, null, null, null, null, null);
  }

  /**
   * Search for userIDs. The data returned will usually be a few seconds behind real time, because
   * userID usage may take up to a few seconds propagate to the different clusters. To keep updates
   * moving quickly, the index of userIDs isn't built synchronously with the mapping. Instead, the
   * index is built once every 12h, at the same time as the update of userID usage. For example,
   * when you perform a modification like adding or moving a userID, the search will report an
   * outdated value until the next rebuild of the mapping, which takes place every 12h. Upon
   * success, the response is 200 OK and contains the following userIDs data.
   *
   * @param searchUserIdsParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return SearchUserIdsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SearchUserIdsResponse searchUserIds(
    SearchUserIdsParams searchUserIdsParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      searchUserIdsAsync(searchUserIdsParams, requestOptions)
    );
  }

  public SearchUserIdsResponse searchUserIds(
    SearchUserIdsParams searchUserIdsParams
  ) throws AlgoliaRuntimeException {
    return this.searchUserIds(searchUserIdsParams, null);
  }

  /**
   * (asynchronously) Search for userIDs. The data returned will usually be a few seconds behind
   * real time, because userID usage may take up to a few seconds propagate to the different
   * clusters. To keep updates moving quickly, the index of userIDs isn&#39;t built synchronously
   * with the mapping. Instead, the index is built once every 12h, at the same time as the update of
   * userID usage. For example, when you perform a modification like adding or moving a userID, the
   * search will report an outdated value until the next rebuild of the mapping, which takes place
   * every 12h. Upon success, the response is 200 OK and contains the following userIDs data.
   *
   * @param searchUserIdsParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SearchUserIdsResponse> searchUserIdsAsync(
    SearchUserIdsParams searchUserIdsParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (searchUserIdsParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'searchUserIdsParams' when calling searchUserIds(Async)"
      );
    }

    Object bodyObj = searchUserIdsParams;

    // create path and map variables
    String requestPath = "/1/clusters/mapping/search";

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
    Type returnType = new TypeToken<SearchUserIdsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<SearchUserIdsResponse> searchUserIdsAsync(
    SearchUserIdsParams searchUserIdsParams
  ) throws AlgoliaRuntimeException {
    return this.searchUserIdsAsync(searchUserIdsParams, null);
  }

  /**
   * Set dictionaries settings.
   *
   * @param dictionarySettingsParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse setDictionarySettings(
    DictionarySettingsParams dictionarySettingsParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      setDictionarySettingsAsync(dictionarySettingsParams, requestOptions)
    );
  }

  public UpdatedAtResponse setDictionarySettings(
    DictionarySettingsParams dictionarySettingsParams
  ) throws AlgoliaRuntimeException {
    return this.setDictionarySettings(dictionarySettingsParams, null);
  }

  /**
   * (asynchronously) Set dictionaries settings.
   *
   * @param dictionarySettingsParams (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> setDictionarySettingsAsync(
    DictionarySettingsParams dictionarySettingsParams,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (dictionarySettingsParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'dictionarySettingsParams' when calling" +
        " setDictionarySettings(Async)"
      );
    }

    Object bodyObj = dictionarySettingsParams;

    // create path and map variables
    String requestPath = "/1/dictionaries/*/settings";

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(
          requestPath,
          "PUT",
          queryParameters,
          bodyObj,
          headers,
          requestOptions
        );
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> setDictionarySettingsAsync(
    DictionarySettingsParams dictionarySettingsParams
  ) throws AlgoliaRuntimeException {
    return this.setDictionarySettingsAsync(dictionarySettingsParams, null);
  }

  /**
   * Update settings of an index. Only specified settings are overridden; unspecified settings are
   * left unchanged. Specifying null for a setting resets it to its default value.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param indexSettings (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse setSettings(
    String indexName,
    IndexSettings indexSettings,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      setSettingsAsync(
        indexName,
        indexSettings,
        forwardToReplicas,
        requestOptions
      )
    );
  }

  public UpdatedAtResponse setSettings(
    String indexName,
    IndexSettings indexSettings,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.setSettings(indexName, indexSettings, forwardToReplicas, null);
  }

  public UpdatedAtResponse setSettings(
    String indexName,
    IndexSettings indexSettings,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.setSettings(indexName, indexSettings, null, requestOptions);
  }

  public UpdatedAtResponse setSettings(
    String indexName,
    IndexSettings indexSettings
  ) throws AlgoliaRuntimeException {
    return this.setSettings(indexName, indexSettings, null, null);
  }

  /**
   * (asynchronously) Update settings of an index. Only specified settings are overridden;
   * unspecified settings are left unchanged. Specifying null for a setting resets it to its default
   * value.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param indexSettings (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> setSettingsAsync(
    String indexName,
    IndexSettings indexSettings,
    Boolean forwardToReplicas,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (indexName == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexName' when calling setSettings(Async)"
      );
    }

    if (indexSettings == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'indexSettings' when calling setSettings(Async)"
      );
    }

    Object bodyObj = indexSettings;

    // create path and map variables
    String requestPath =
      "/1/indexes/{indexName}/settings".replaceAll(
          "\\{indexName\\}",
          this.escapeString(indexName.toString())
        );

    Map<String, Object> queryParameters = new HashMap<String, Object>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParameters.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
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
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdatedAtResponse> setSettingsAsync(
    String indexName,
    IndexSettings indexSettings,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return this.setSettingsAsync(
        indexName,
        indexSettings,
        forwardToReplicas,
        null
      );
  }

  public CompletableFuture<UpdatedAtResponse> setSettingsAsync(
    String indexName,
    IndexSettings indexSettings,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return this.setSettingsAsync(
        indexName,
        indexSettings,
        null,
        requestOptions
      );
  }

  public CompletableFuture<UpdatedAtResponse> setSettingsAsync(
    String indexName,
    IndexSettings indexSettings
  ) throws AlgoliaRuntimeException {
    return this.setSettingsAsync(indexName, indexSettings, null, null);
  }

  /**
   * Replace every permission of an existing API key.
   *
   * @param key API Key string. (required)
   * @param apiKey (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return UpdateApiKeyResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdateApiKeyResponse updateApiKey(
    String key,
    ApiKey apiKey,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      updateApiKeyAsync(key, apiKey, requestOptions)
    );
  }

  public UpdateApiKeyResponse updateApiKey(String key, ApiKey apiKey)
    throws AlgoliaRuntimeException {
    return this.updateApiKey(key, apiKey, null);
  }

  /**
   * (asynchronously) Replace every permission of an existing API key.
   *
   * @param key API Key string. (required)
   * @param apiKey (required)
   * @param requestOptions The requestOptions to send along with the query, they will be merged with
   *     the transporter requestOptions.
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdateApiKeyResponse> updateApiKeyAsync(
    String key,
    ApiKey apiKey,
    RequestOptions requestOptions
  ) throws AlgoliaRuntimeException {
    if (key == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'key' when calling updateApiKey(Async)"
      );
    }

    if (apiKey == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'apiKey' when calling updateApiKey(Async)"
      );
    }

    Object bodyObj = apiKey;

    // create path and map variables
    String requestPath =
      "/1/keys/{key}".replaceAll(
          "\\{key\\}",
          this.escapeString(key.toString())
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
          requestOptions
        );
    Type returnType = new TypeToken<UpdateApiKeyResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  public CompletableFuture<UpdateApiKeyResponse> updateApiKeyAsync(
    String key,
    ApiKey apiKey
  ) throws AlgoliaRuntimeException {
    return this.updateApiKeyAsync(key, apiKey, null);
  }
}
