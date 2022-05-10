package com.algolia.api;

import com.algolia.ApiClient;
import com.algolia.exceptions.*;
import com.algolia.model.search.*;
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
   * @return AddApiKeyResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public AddApiKeyResponse addApiKey(ApiKey apiKey)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(addApiKeyAsync(apiKey));
  }

  /**
   * (asynchronously) Add a new API Key with specific permissions/restrictions.
   *
   * @param apiKey (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<AddApiKeyResponse> addApiKeyAsync(ApiKey apiKey)
    throws AlgoliaRuntimeException {
    if (apiKey == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'apiKey' when calling addApiKey(Async)"
      );
    }

    Object bodyObj = apiKey;

    // create path and map variables
    String requestPath = "/1/keys";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<AddApiKeyResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Add or replace an object with a given object ID. If the object does not exist, it will be
   * created. If it already exists, it will be replaced.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param body The Algolia object. (required)
   * @return UpdatedAtWithObjectIdResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtWithObjectIdResponse addOrUpdateObject(
    String indexName,
    String objectID,
    Object body
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      addOrUpdateObjectAsync(indexName, objectID, body)
    );
  }

  /**
   * (asynchronously) Add or replace an object with a given object ID. If the object does not exist,
   * it will be created. If it already exists, it will be replaced.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param body The Algolia object. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtWithObjectIdResponse> addOrUpdateObjectAsync(
    String indexName,
    String objectID,
    Object body
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "PUT", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtWithObjectIdResponse>() {}
      .getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Add a single source to the list of allowed sources.
   *
   * @param source The source to add. (required)
   * @return CreatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public CreatedAtResponse appendSource(Source source)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(appendSourceAsync(source));
  }

  /**
   * (asynchronously) Add a single source to the list of allowed sources.
   *
   * @param source The source to add. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<CreatedAtResponse> appendSourceAsync(Source source)
    throws AlgoliaRuntimeException {
    if (source == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'source' when calling appendSource(Async)"
      );
    }

    Object bodyObj = source;

    // create path and map variables
    String requestPath = "/1/security/sources/append";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<CreatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Assign or Move a userID to a cluster. The time it takes to migrate (move) a user is
   * proportional to the amount of data linked to the userID. Upon success, the response is 200 OK.
   * A successful response indicates that the operation has been taken into account, and the userID
   * is directly usable.
   *
   * @param xAlgoliaUserID userID to assign. (required)
   * @param assignUserIdParams (required)
   * @return CreatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public CreatedAtResponse assignUserId(
    String xAlgoliaUserID,
    AssignUserIdParams assignUserIdParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      assignUserIdAsync(xAlgoliaUserID, assignUserIdParams)
    );
  }

  /**
   * (asynchronously) Assign or Move a userID to a cluster. The time it takes to migrate (move) a
   * user is proportional to the amount of data linked to the userID. Upon success, the response is
   * 200 OK. A successful response indicates that the operation has been taken into account, and the
   * userID is directly usable.
   *
   * @param xAlgoliaUserID userID to assign. (required)
   * @param assignUserIdParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<CreatedAtResponse> assignUserIdAsync(
    String xAlgoliaUserID,
    AssignUserIdParams assignUserIdParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (xAlgoliaUserID != null) {
      queryParams.put("X-Algolia-User-ID", parameterToString(xAlgoliaUserID));
    }

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<CreatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Perform multiple write operations targeting one index, in a single API call.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param batchWriteParams (required)
   * @return BatchResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public BatchResponse batch(
    String indexName,
    BatchWriteParams batchWriteParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(batchAsync(indexName, batchWriteParams));
  }

  /**
   * (asynchronously) Perform multiple write operations targeting one index, in a single API call.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param batchWriteParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<BatchResponse> batchAsync(
    String indexName,
    BatchWriteParams batchWriteParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<BatchResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Assign multiple userIDs to a cluster. Upon success, the response is 200 OK. A successful
   * response indicates that the operation has been taken into account, and the userIDs are directly
   * usable.
   *
   * @param xAlgoliaUserID userID to assign. (required)
   * @param batchAssignUserIdsParams (required)
   * @return CreatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public CreatedAtResponse batchAssignUserIds(
    String xAlgoliaUserID,
    BatchAssignUserIdsParams batchAssignUserIdsParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      batchAssignUserIdsAsync(xAlgoliaUserID, batchAssignUserIdsParams)
    );
  }

  /**
   * (asynchronously) Assign multiple userIDs to a cluster. Upon success, the response is 200 OK. A
   * successful response indicates that the operation has been taken into account, and the userIDs
   * are directly usable.
   *
   * @param xAlgoliaUserID userID to assign. (required)
   * @param batchAssignUserIdsParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<CreatedAtResponse> batchAssignUserIdsAsync(
    String xAlgoliaUserID,
    BatchAssignUserIdsParams batchAssignUserIdsParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (xAlgoliaUserID != null) {
      queryParams.put("X-Algolia-User-ID", parameterToString(xAlgoliaUserID));
    }

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<CreatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Send a batch of dictionary entries.
   *
   * @param dictionaryName The dictionary to search in. (required)
   * @param batchDictionaryEntriesParams (required)
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse batchDictionaryEntries(
    DictionaryType dictionaryName,
    BatchDictionaryEntriesParams batchDictionaryEntriesParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      batchDictionaryEntriesAsync(dictionaryName, batchDictionaryEntriesParams)
    );
  }

  /**
   * (asynchronously) Send a batch of dictionary entries.
   *
   * @param dictionaryName The dictionary to search in. (required)
   * @param batchDictionaryEntriesParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> batchDictionaryEntriesAsync(
    DictionaryType dictionaryName,
    BatchDictionaryEntriesParams batchDictionaryEntriesParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
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
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse batchRules(
    String indexName,
    List<Rule> rule,
    Boolean forwardToReplicas,
    Boolean clearExistingRules
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      batchRulesAsync(indexName, rule, forwardToReplicas, clearExistingRules)
    );
  }

  public UpdatedAtResponse batchRules(String indexName, List<Rule> rule)
    throws AlgoliaRuntimeException {
    return this.batchRules(indexName, rule, null, null);
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> batchRulesAsync(
    String indexName,
    List<Rule> rule,
    Boolean forwardToReplicas,
    Boolean clearExistingRules
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParams.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    if (clearExistingRules != null) {
      queryParams.put(
        "clearExistingRules",
        parameterToString(clearExistingRules)
      );
    }

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
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
   * @return BrowseResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public BrowseResponse browse(String indexName, BrowseRequest browseRequest)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(browseAsync(indexName, browseRequest));
  }

  public BrowseResponse browse(String indexName)
    throws AlgoliaRuntimeException {
    return this.browse(indexName, null);
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<BrowseResponse> browseAsync(
    String indexName,
    BrowseRequest browseRequest
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<BrowseResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Remove all synonyms from an index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse clearAllSynonyms(
    String indexName,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      clearAllSynonymsAsync(indexName, forwardToReplicas)
    );
  }

  public UpdatedAtResponse clearAllSynonyms(String indexName)
    throws AlgoliaRuntimeException {
    return this.clearAllSynonyms(indexName, null);
  }

  /**
   * (asynchronously) Remove all synonyms from an index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> clearAllSynonymsAsync(
    String indexName,
    Boolean forwardToReplicas
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParams.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Delete an index's content, but leave settings and index-specific API keys untouched.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse clearObjects(String indexName)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(clearObjectsAsync(indexName));
  }

  /**
   * (asynchronously) Delete an index&#39;s content, but leave settings and index-specific API keys
   * untouched.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> clearObjectsAsync(
    String indexName
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Delete all Rules in the index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse clearRules(
    String indexName,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      clearRulesAsync(indexName, forwardToReplicas)
    );
  }

  public UpdatedAtResponse clearRules(String indexName)
    throws AlgoliaRuntimeException {
    return this.clearRules(indexName, null);
  }

  /**
   * (asynchronously) Delete all Rules in the index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> clearRulesAsync(
    String indexName,
    Boolean forwardToReplicas
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParams.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
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
   * Delete an existing API Key.
   *
   * @param key API Key string. (required)
   * @return DeleteApiKeyResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeleteApiKeyResponse deleteApiKey(String key)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteApiKeyAsync(key));
  }

  /**
   * (asynchronously) Delete an existing API Key.
   *
   * @param key API Key string. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeleteApiKeyResponse> deleteApiKeyAsync(String key)
    throws AlgoliaRuntimeException {
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<DeleteApiKeyResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Remove all objects matching a filter (including geo filters). This method enables you to delete
   * one or more objects based on filters (numeric, facet, tag or geo queries). It doesn't accept
   * empty filters or a query.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchParams (required)
   * @return DeletedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeletedAtResponse deleteBy(
    String indexName,
    SearchParams searchParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteByAsync(indexName, searchParams));
  }

  /**
   * (asynchronously) Remove all objects matching a filter (including geo filters). This method
   * enables you to delete one or more objects based on filters (numeric, facet, tag or geo
   * queries). It doesn&#39;t accept empty filters or a query.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeletedAtResponse> deleteByAsync(
    String indexName,
    SearchParams searchParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<DeletedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Delete an existing index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return DeletedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeletedAtResponse deleteIndex(String indexName)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteIndexAsync(indexName));
  }

  /**
   * (asynchronously) Delete an existing index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeletedAtResponse> deleteIndexAsync(
    String indexName
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<DeletedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Delete an existing object.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @return DeletedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeletedAtResponse deleteObject(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteObjectAsync(indexName, objectID));
  }

  /**
   * (asynchronously) Delete an existing object.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeletedAtResponse> deleteObjectAsync(
    String indexName,
    String objectID
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<DeletedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Delete the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse deleteRule(
    String indexName,
    String objectID,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      deleteRuleAsync(indexName, objectID, forwardToReplicas)
    );
  }

  public UpdatedAtResponse deleteRule(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return this.deleteRule(indexName, objectID, null);
  }

  /**
   * (asynchronously) Delete the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> deleteRuleAsync(
    String indexName,
    String objectID,
    Boolean forwardToReplicas
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParams.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Remove a single source from the list of allowed sources.
   *
   * @param source The IP range of the source. (required)
   * @return DeleteSourceResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeleteSourceResponse deleteSource(String source)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(deleteSourceAsync(source));
  }

  /**
   * (asynchronously) Remove a single source from the list of allowed sources.
   *
   * @param source The IP range of the source. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeleteSourceResponse> deleteSourceAsync(
    String source
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<DeleteSourceResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Delete a single synonyms set, identified by the given objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return DeletedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public DeletedAtResponse deleteSynonym(
    String indexName,
    String objectID,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      deleteSynonymAsync(indexName, objectID, forwardToReplicas)
    );
  }

  public DeletedAtResponse deleteSynonym(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return this.deleteSynonym(indexName, objectID, null);
  }

  /**
   * (asynchronously) Delete a single synonyms set, identified by the given objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<DeletedAtResponse> deleteSynonymAsync(
    String indexName,
    String objectID,
    Boolean forwardToReplicas
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParams.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<DeletedAtResponse>() {}.getType();
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
   * Get the permissions of an API key.
   *
   * @param key API Key string. (required)
   * @return Key
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Key getApiKey(String key) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getApiKeyAsync(key));
  }

  /**
   * (asynchronously) Get the permissions of an API key.
   *
   * @param key API Key string. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Key> getApiKeyAsync(String key)
    throws AlgoliaRuntimeException {
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<Key>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * List dictionaries supported per language.
   *
   * @return Map&lt;String, Languages&gt;
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Map<String, Languages> getDictionaryLanguages()
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getDictionaryLanguagesAsync());
  }

  /**
   * (asynchronously) List dictionaries supported per language.
   *
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Map<String, Languages>> getDictionaryLanguagesAsync()
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/dictionaries/*/languages";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<Map<String, Languages>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Retrieve dictionaries settings. The API stores languages whose standard entries are disabled.
   * Fetch settings does not return false values.
   *
   * @return GetDictionarySettingsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetDictionarySettingsResponse getDictionarySettings()
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getDictionarySettingsAsync());
  }

  /**
   * (asynchronously) Retrieve dictionaries settings. The API stores languages whose standard
   * entries are disabled. Fetch settings does not return false values.
   *
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetDictionarySettingsResponse> getDictionarySettingsAsync()
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/dictionaries/*/settings";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<GetDictionarySettingsResponse>() {}
      .getType();
    return this.executeAsync(call, returnType);
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
   * @return GetLogsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetLogsResponse getLogs(
    Integer offset,
    Integer length,
    String indexName,
    LogType type
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getLogsAsync(offset, length, indexName, type)
    );
  }

  public GetLogsResponse getLogs() throws AlgoliaRuntimeException {
    return this.getLogs(null, null, null, null);
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetLogsResponse> getLogsAsync(
    Integer offset,
    Integer length,
    String indexName,
    LogType type
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/logs";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (offset != null) {
      queryParams.put("offset", parameterToString(offset));
    }

    if (length != null) {
      queryParams.put("length", parameterToString(length));
    }

    if (indexName != null) {
      queryParams.put("indexName", parameterToString(indexName));
    }

    if (type != null) {
      queryParams.put("type", parameterToString(type));
    }

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<GetLogsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Retrieve one object from the index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param attributesToRetrieve List of attributes to retrieve. If not specified, all retrievable
   *     attributes are returned. (optional)
   * @return Map&lt;String, String&gt;
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Map<String, String> getObject(
    String indexName,
    String objectID,
    List<String> attributesToRetrieve
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      getObjectAsync(indexName, objectID, attributesToRetrieve)
    );
  }

  public Map<String, String> getObject(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return this.getObject(indexName, objectID, null);
  }

  /**
   * (asynchronously) Retrieve one object from the index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param attributesToRetrieve List of attributes to retrieve. If not specified, all retrievable
   *     attributes are returned. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Map<String, String>> getObjectAsync(
    String indexName,
    String objectID,
    List<String> attributesToRetrieve
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (attributesToRetrieve != null) {
      queryParams.put(
        "attributesToRetrieve",
        parameterToString(attributesToRetrieve)
      );
    }

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<Map<String, String>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Retrieve one or more objects, potentially from different indices, in a single API call.
   *
   * @param getObjectsParams (required)
   * @return GetObjectsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetObjectsResponse getObjects(GetObjectsParams getObjectsParams)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getObjectsAsync(getObjectsParams));
  }

  /**
   * (asynchronously) Retrieve one or more objects, potentially from different indices, in a single
   * API call.
   *
   * @param getObjectsParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetObjectsResponse> getObjectsAsync(
    GetObjectsParams getObjectsParams
  ) throws AlgoliaRuntimeException {
    if (getObjectsParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'getObjectsParams' when calling getObjects(Async)"
      );
    }

    Object bodyObj = getObjectsParams;

    // create path and map variables
    String requestPath = "/1/indexes/*/objects";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<GetObjectsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Retrieve the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @return Rule
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public Rule getRule(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getRuleAsync(indexName, objectID));
  }

  /**
   * (asynchronously) Retrieve the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<Rule> getRuleAsync(
    String indexName,
    String objectID
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<Rule>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Retrieve settings of an index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return IndexSettings
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public IndexSettings getSettings(String indexName)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getSettingsAsync(indexName));
  }

  /**
   * (asynchronously) Retrieve settings of an index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<IndexSettings> getSettingsAsync(String indexName)
    throws AlgoliaRuntimeException {
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<IndexSettings>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * List all allowed sources.
   *
   * @return List&lt;Source&gt;
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public List<Source> getSources() throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getSourcesAsync());
  }

  /**
   * (asynchronously) List all allowed sources.
   *
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<List<Source>> getSourcesAsync()
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/security/sources";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<List<Source>>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Fetch a synonym object identified by its objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @return SynonymHit
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SynonymHit getSynonym(String indexName, String objectID)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getSynonymAsync(indexName, objectID));
  }

  /**
   * (asynchronously) Fetch a synonym object identified by its objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SynonymHit> getSynonymAsync(
    String indexName,
    String objectID
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SynonymHit>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Check the current status of a given task.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param taskID Unique identifier of an task. Numeric value (up to 64bits). (required)
   * @return GetTaskResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetTaskResponse getTask(String indexName, Integer taskID)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getTaskAsync(indexName, taskID));
  }

  /**
   * (asynchronously) Check the current status of a given task.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param taskID Unique identifier of an task. Numeric value (up to 64bits). (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetTaskResponse> getTaskAsync(
    String indexName,
    Integer taskID
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<GetTaskResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Get the top 10 userIDs with the highest number of records per cluster. The data returned will
   * usually be a few seconds behind real time, because userID usage may take up to a few seconds to
   * propagate to the different clusters. Upon success, the response is 200 OK and contains the
   * following array of userIDs and clusters.
   *
   * @return GetTopUserIdsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public GetTopUserIdsResponse getTopUserIds() throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getTopUserIdsAsync());
  }

  /**
   * (asynchronously) Get the top 10 userIDs with the highest number of records per cluster. The
   * data returned will usually be a few seconds behind real time, because userID usage may take up
   * to a few seconds to propagate to the different clusters. Upon success, the response is 200 OK
   * and contains the following array of userIDs and clusters.
   *
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<GetTopUserIdsResponse> getTopUserIdsAsync()
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/clusters/mapping/top";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<GetTopUserIdsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Returns the userID data stored in the mapping. The data returned will usually be a few seconds
   * behind real time, because userID usage may take up to a few seconds to propagate to the
   * different clusters. Upon success, the response is 200 OK and contains the following userID
   * data.
   *
   * @param userID userID to assign. (required)
   * @return UserId
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UserId getUserId(String userID) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(getUserIdAsync(userID));
  }

  /**
   * (asynchronously) Returns the userID data stored in the mapping. The data returned will usually
   * be a few seconds behind real time, because userID usage may take up to a few seconds to
   * propagate to the different clusters. Upon success, the response is 200 OK and contains the
   * following userID data.
   *
   * @param userID userID to assign. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UserId> getUserIdAsync(String userID)
    throws AlgoliaRuntimeException {
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UserId>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Get the status of your clusters' migrations or user creations. Creating a large batch of users
   * or migrating your multi-cluster may take quite some time. This method lets you retrieve the
   * status of the migration, so you can know when it's done. Upon success, the response is 200 OK.
   * A successful response indicates that the operation has been taken into account, and the userIDs
   * are directly usable.
   *
   * @param getClusters Whether to get clusters or not. (optional)
   * @return CreatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public CreatedAtResponse hasPendingMappings(Boolean getClusters)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(hasPendingMappingsAsync(getClusters));
  }

  public CreatedAtResponse hasPendingMappings() throws AlgoliaRuntimeException {
    return this.hasPendingMappings(null);
  }

  /**
   * (asynchronously) Get the status of your clusters&#39; migrations or user creations. Creating a
   * large batch of users or migrating your multi-cluster may take quite some time. This method lets
   * you retrieve the status of the migration, so you can know when it&#39;s done. Upon success, the
   * response is 200 OK. A successful response indicates that the operation has been taken into
   * account, and the userIDs are directly usable.
   *
   * @param getClusters Whether to get clusters or not. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<CreatedAtResponse> hasPendingMappingsAsync(
    Boolean getClusters
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/clusters/mapping/pending";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (getClusters != null) {
      queryParams.put("getClusters", parameterToString(getClusters));
    }

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<CreatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * List API keys, along with their associated rights.
   *
   * @return ListApiKeysResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ListApiKeysResponse listApiKeys() throws AlgoliaRuntimeException {
    return LaunderThrowable.await(listApiKeysAsync());
  }

  /**
   * (asynchronously) List API keys, along with their associated rights.
   *
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ListApiKeysResponse> listApiKeysAsync()
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/keys";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<ListApiKeysResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * List the clusters available in a multi-clusters setup for a single appID. Upon success, the
   * response is 200 OK and contains the following clusters.
   *
   * @return ListClustersResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ListClustersResponse listClusters() throws AlgoliaRuntimeException {
    return LaunderThrowable.await(listClustersAsync());
  }

  /**
   * (asynchronously) List the clusters available in a multi-clusters setup for a single appID. Upon
   * success, the response is 200 OK and contains the following clusters.
   *
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ListClustersResponse> listClustersAsync()
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/clusters";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<ListClustersResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * List existing indexes from an application.
   *
   * @param page Requested page (zero-based). When specified, will retrieve a specific page; the
   *     page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   *     (optional)
   * @return ListIndicesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ListIndicesResponse listIndices(Integer page)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(listIndicesAsync(page));
  }

  public ListIndicesResponse listIndices() throws AlgoliaRuntimeException {
    return this.listIndices(null);
  }

  /**
   * (asynchronously) List existing indexes from an application.
   *
   * @param page Requested page (zero-based). When specified, will retrieve a specific page; the
   *     page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   *     (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ListIndicesResponse> listIndicesAsync(Integer page)
    throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/indexes";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (page != null) {
      queryParams.put("page", parameterToString(page));
    }

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<ListIndicesResponse>() {}.getType();
    return this.executeAsync(call, returnType);
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
   * @return ListUserIdsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ListUserIdsResponse listUserIds(Integer page, Integer hitsPerPage)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(listUserIdsAsync(page, hitsPerPage));
  }

  public ListUserIdsResponse listUserIds() throws AlgoliaRuntimeException {
    return this.listUserIds(null, null);
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ListUserIdsResponse> listUserIdsAsync(
    Integer page,
    Integer hitsPerPage
  ) throws AlgoliaRuntimeException {
    Object bodyObj = null;

    // create path and map variables
    String requestPath = "/1/clusters/mapping";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (page != null) {
      queryParams.put("page", parameterToString(page));
    }

    if (hitsPerPage != null) {
      queryParams.put("hitsPerPage", parameterToString(hitsPerPage));
    }

    Call call =
      this.buildCall(requestPath, "GET", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<ListUserIdsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Perform multiple write operations, potentially targeting multiple indices, in a single API
   * call.
   *
   * @param batchParams (required)
   * @return MultipleBatchResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public MultipleBatchResponse multipleBatch(BatchParams batchParams)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(multipleBatchAsync(batchParams));
  }

  /**
   * (asynchronously) Perform multiple write operations, potentially targeting multiple indices, in
   * a single API call.
   *
   * @param batchParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<MultipleBatchResponse> multipleBatchAsync(
    BatchParams batchParams
  ) throws AlgoliaRuntimeException {
    if (batchParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'batchParams' when calling multipleBatch(Async)"
      );
    }

    Object bodyObj = batchParams;

    // create path and map variables
    String requestPath = "/1/indexes/*/batch";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<MultipleBatchResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Perform a search operation targeting one or many indices.
   *
   * @param multipleQueriesParams (required)
   * @return MultipleQueriesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public MultipleQueriesResponse multipleQueries(
    MultipleQueriesParams multipleQueriesParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(multipleQueriesAsync(multipleQueriesParams));
  }

  /**
   * (asynchronously) Perform a search operation targeting one or many indices.
   *
   * @param multipleQueriesParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<MultipleQueriesResponse> multipleQueriesAsync(
    MultipleQueriesParams multipleQueriesParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<MultipleQueriesResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Peforms a copy or a move operation on a index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param operationIndexParams (required)
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse operationIndex(
    String indexName,
    OperationIndexParams operationIndexParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      operationIndexAsync(indexName, operationIndexParams)
    );
  }

  /**
   * (asynchronously) Peforms a copy or a move operation on a index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param operationIndexParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> operationIndexAsync(
    String indexName,
    OperationIndexParams operationIndexParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
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
   * @return UpdatedAtWithObjectIdResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtWithObjectIdResponse partialUpdateObject(
    String indexName,
    String objectID,
    List<Map<String, AttributeOrBuiltInOperation>> attributeOrBuiltInOperation,
    Boolean createIfNotExists
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      partialUpdateObjectAsync(
        indexName,
        objectID,
        attributeOrBuiltInOperation,
        createIfNotExists
      )
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtWithObjectIdResponse> partialUpdateObjectAsync(
    String indexName,
    String objectID,
    List<Map<String, AttributeOrBuiltInOperation>> attributeOrBuiltInOperation,
    Boolean createIfNotExists
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (createIfNotExists != null) {
      queryParams.put(
        "createIfNotExists",
        parameterToString(createIfNotExists)
      );
    }

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtWithObjectIdResponse>() {}
      .getType();
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
   * Remove a userID and its associated data from the multi-clusters. Upon success, the response is
   * 200 OK and a task is created to remove the userID data and mapping.
   *
   * @param userID userID to assign. (required)
   * @return RemoveUserIdResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public RemoveUserIdResponse removeUserId(String userID)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(removeUserIdAsync(userID));
  }

  /**
   * (asynchronously) Remove a userID and its associated data from the multi-clusters. Upon success,
   * the response is 200 OK and a task is created to remove the userID data and mapping.
   *
   * @param userID userID to assign. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<RemoveUserIdResponse> removeUserIdAsync(
    String userID
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "DELETE", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<RemoveUserIdResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Replace all allowed sources.
   *
   * @param source The sources to allow. (required)
   * @return ReplaceSourceResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public ReplaceSourceResponse replaceSources(List<Source> source)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(replaceSourcesAsync(source));
  }

  /**
   * (asynchronously) Replace all allowed sources.
   *
   * @param source The sources to allow. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<ReplaceSourceResponse> replaceSourcesAsync(
    List<Source> source
  ) throws AlgoliaRuntimeException {
    if (source == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'source' when calling replaceSources(Async)"
      );
    }

    Object bodyObj = source;

    // create path and map variables
    String requestPath = "/1/security/sources";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "PUT", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<ReplaceSourceResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Restore a deleted API key, along with its associated rights.
   *
   * @param key API Key string. (required)
   * @return AddApiKeyResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public AddApiKeyResponse restoreApiKey(String key)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(restoreApiKeyAsync(key));
  }

  /**
   * (asynchronously) Restore a deleted API key, along with its associated rights.
   *
   * @param key API Key string. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<AddApiKeyResponse> restoreApiKeyAsync(String key)
    throws AlgoliaRuntimeException {
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<AddApiKeyResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Add an object to the index, automatically assigning it an object ID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param body The Algolia record. (required)
   * @return SaveObjectResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SaveObjectResponse saveObject(String indexName, Object body)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(saveObjectAsync(indexName, body));
  }

  /**
   * (asynchronously) Add an object to the index, automatically assigning it an object ID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param body The Algolia record. (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SaveObjectResponse> saveObjectAsync(
    String indexName,
    Object body
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SaveObjectResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Create or update the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param rule (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return UpdatedRuleResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedRuleResponse saveRule(
    String indexName,
    String objectID,
    Rule rule,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      saveRuleAsync(indexName, objectID, rule, forwardToReplicas)
    );
  }

  public UpdatedRuleResponse saveRule(
    String indexName,
    String objectID,
    Rule rule
  ) throws AlgoliaRuntimeException {
    return this.saveRule(indexName, objectID, rule, null);
  }

  /**
   * (asynchronously) Create or update the Rule with the specified objectID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param rule (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedRuleResponse> saveRuleAsync(
    String indexName,
    String objectID,
    Rule rule,
    Boolean forwardToReplicas
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParams.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    Call call =
      this.buildCall(requestPath, "PUT", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedRuleResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Create a new synonym object or update the existing synonym object with the given object ID.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param objectID Unique identifier of an object. (required)
   * @param synonymHit (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return SaveSynonymResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SaveSynonymResponse saveSynonym(
    String indexName,
    String objectID,
    SynonymHit synonymHit,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      saveSynonymAsync(indexName, objectID, synonymHit, forwardToReplicas)
    );
  }

  public SaveSynonymResponse saveSynonym(
    String indexName,
    String objectID,
    SynonymHit synonymHit
  ) throws AlgoliaRuntimeException {
    return this.saveSynonym(indexName, objectID, synonymHit, null);
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SaveSynonymResponse> saveSynonymAsync(
    String indexName,
    String objectID,
    SynonymHit synonymHit,
    Boolean forwardToReplicas
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParams.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    Call call =
      this.buildCall(requestPath, "PUT", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SaveSynonymResponse>() {}.getType();
    return this.executeAsync(call, returnType);
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
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse saveSynonyms(
    String indexName,
    List<SynonymHit> synonymHit,
    Boolean forwardToReplicas,
    Boolean replaceExistingSynonyms
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      saveSynonymsAsync(
        indexName,
        synonymHit,
        forwardToReplicas,
        replaceExistingSynonyms
      )
    );
  }

  public UpdatedAtResponse saveSynonyms(
    String indexName,
    List<SynonymHit> synonymHit
  ) throws AlgoliaRuntimeException {
    return this.saveSynonyms(indexName, synonymHit, null, null);
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> saveSynonymsAsync(
    String indexName,
    List<SynonymHit> synonymHit,
    Boolean forwardToReplicas,
    Boolean replaceExistingSynonyms
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParams.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    if (replaceExistingSynonyms != null) {
      queryParams.put(
        "replaceExistingSynonyms",
        parameterToString(replaceExistingSynonyms)
      );
    }

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Perform a search operation targeting one specific index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchParams (required)
   * @return SearchResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SearchResponse search(String indexName, SearchParams searchParams)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(searchAsync(indexName, searchParams));
  }

  /**
   * (asynchronously) Perform a search operation targeting one specific index.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SearchResponse> searchAsync(
    String indexName,
    SearchParams searchParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SearchResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Search the dictionary entries.
   *
   * @param dictionaryName The dictionary to search in. (required)
   * @param searchDictionaryEntriesParams (required)
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse searchDictionaryEntries(
    DictionaryType dictionaryName,
    SearchDictionaryEntriesParams searchDictionaryEntriesParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      searchDictionaryEntriesAsync(
        dictionaryName,
        searchDictionaryEntriesParams
      )
    );
  }

  /**
   * (asynchronously) Search the dictionary entries.
   *
   * @param dictionaryName The dictionary to search in. (required)
   * @param searchDictionaryEntriesParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> searchDictionaryEntriesAsync(
    DictionaryType dictionaryName,
    SearchDictionaryEntriesParams searchDictionaryEntriesParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Search for values of a given facet, optionally restricting the returned values to those
   * contained in objects matching other search criteria.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param facetName The facet name. (required)
   * @param searchForFacetValuesRequest (optional)
   * @return SearchForFacetValuesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SearchForFacetValuesResponse searchForFacetValues(
    String indexName,
    String facetName,
    SearchForFacetValuesRequest searchForFacetValuesRequest
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      searchForFacetValuesAsync(
        indexName,
        facetName,
        searchForFacetValuesRequest
      )
    );
  }

  public SearchForFacetValuesResponse searchForFacetValues(
    String indexName,
    String facetName
  ) throws AlgoliaRuntimeException {
    return this.searchForFacetValues(indexName, facetName, null);
  }

  /**
   * (asynchronously) Search for values of a given facet, optionally restricting the returned values
   * to those contained in objects matching other search criteria.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param facetName The facet name. (required)
   * @param searchForFacetValuesRequest (optional)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SearchForFacetValuesResponse> searchForFacetValuesAsync(
    String indexName,
    String facetName,
    SearchForFacetValuesRequest searchForFacetValuesRequest
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SearchForFacetValuesResponse>() {}
      .getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Search for rules matching various criteria.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchRulesParams (required)
   * @return SearchRulesResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SearchRulesResponse searchRules(
    String indexName,
    SearchRulesParams searchRulesParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      searchRulesAsync(indexName, searchRulesParams)
    );
  }

  /**
   * (asynchronously) Search for rules matching various criteria.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param searchRulesParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SearchRulesResponse> searchRulesAsync(
    String indexName,
    SearchRulesParams searchRulesParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SearchRulesResponse>() {}.getType();
    return this.executeAsync(call, returnType);
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
   * @return SearchSynonymsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SearchSynonymsResponse searchSynonyms(
    String indexName,
    String query,
    SynonymType type,
    Integer page,
    Integer hitsPerPage
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      searchSynonymsAsync(indexName, query, type, page, hitsPerPage)
    );
  }

  public SearchSynonymsResponse searchSynonyms(String indexName)
    throws AlgoliaRuntimeException {
    return this.searchSynonyms(indexName, null, null, null, null);
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SearchSynonymsResponse> searchSynonymsAsync(
    String indexName,
    String query,
    SynonymType type,
    Integer page,
    Integer hitsPerPage
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (query != null) {
      queryParams.put("query", parameterToString(query));
    }

    if (type != null) {
      queryParams.put("type", parameterToString(type));
    }

    if (page != null) {
      queryParams.put("page", parameterToString(page));
    }

    if (hitsPerPage != null) {
      queryParams.put("hitsPerPage", parameterToString(hitsPerPage));
    }

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SearchSynonymsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
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
   * @return SearchUserIdsResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public SearchUserIdsResponse searchUserIds(
    SearchUserIdsParams searchUserIdsParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(searchUserIdsAsync(searchUserIdsParams));
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<SearchUserIdsResponse> searchUserIdsAsync(
    SearchUserIdsParams searchUserIdsParams
  ) throws AlgoliaRuntimeException {
    if (searchUserIdsParams == null) {
      throw new AlgoliaRuntimeException(
        "Missing the required parameter 'searchUserIdsParams' when calling searchUserIds(Async)"
      );
    }

    Object bodyObj = searchUserIdsParams;

    // create path and map variables
    String requestPath = "/1/clusters/mapping/search";

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "POST", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<SearchUserIdsResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Set dictionaries settings.
   *
   * @param dictionarySettingsParams (required)
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse setDictionarySettings(
    DictionarySettingsParams dictionarySettingsParams
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      setDictionarySettingsAsync(dictionarySettingsParams)
    );
  }

  /**
   * (asynchronously) Set dictionaries settings.
   *
   * @param dictionarySettingsParams (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> setDictionarySettingsAsync(
    DictionarySettingsParams dictionarySettingsParams
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "PUT", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Update settings of an index. Only specified settings are overridden; unspecified settings are
   * left unchanged. Specifying null for a setting resets it to its default value.
   *
   * @param indexName The index in which to perform the request. (required)
   * @param indexSettings (required)
   * @param forwardToReplicas When true, changes are also propagated to replicas of the given
   *     indexName. (optional)
   * @return UpdatedAtResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdatedAtResponse setSettings(
    String indexName,
    IndexSettings indexSettings,
    Boolean forwardToReplicas
  ) throws AlgoliaRuntimeException {
    return LaunderThrowable.await(
      setSettingsAsync(indexName, indexSettings, forwardToReplicas)
    );
  }

  public UpdatedAtResponse setSettings(
    String indexName,
    IndexSettings indexSettings
  ) throws AlgoliaRuntimeException {
    return this.setSettings(indexName, indexSettings, null);
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
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdatedAtResponse> setSettingsAsync(
    String indexName,
    IndexSettings indexSettings,
    Boolean forwardToReplicas
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    if (forwardToReplicas != null) {
      queryParams.put(
        "forwardToReplicas",
        parameterToString(forwardToReplicas)
      );
    }

    Call call =
      this.buildCall(requestPath, "PUT", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdatedAtResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }

  /**
   * Replace every permission of an existing API key.
   *
   * @param key API Key string. (required)
   * @param apiKey (required)
   * @return UpdateApiKeyResponse
   * @throws AlgoliaRuntimeException If fail to call the API, e.g. server error or cannot
   *     deserialize the response body
   */
  public UpdateApiKeyResponse updateApiKey(String key, ApiKey apiKey)
    throws AlgoliaRuntimeException {
    return LaunderThrowable.await(updateApiKeyAsync(key, apiKey));
  }

  /**
   * (asynchronously) Replace every permission of an existing API key.
   *
   * @param key API Key string. (required)
   * @param apiKey (required)
   * @return The awaitable future
   * @throws AlgoliaRuntimeException If fail to process the API call, e.g. serializing the request
   *     body object
   */
  public CompletableFuture<UpdateApiKeyResponse> updateApiKeyAsync(
    String key,
    ApiKey apiKey
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

    Map<String, String> queryParams = new HashMap<String, String>();
    Map<String, String> headers = new HashMap<String, String>();

    Call call =
      this.buildCall(requestPath, "PUT", queryParams, bodyObj, headers);
    Type returnType = new TypeToken<UpdateApiKeyResponse>() {}.getType();
    return this.executeAsync(call, returnType);
  }
}
