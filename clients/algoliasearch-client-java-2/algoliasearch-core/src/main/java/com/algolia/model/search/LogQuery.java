package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** LogQuery */
public class LogQuery {

  @JsonProperty("index_name")
  private String indexName;

  @JsonProperty("user_token")
  private String userToken;

  @JsonProperty("query_id")
  private String queryId;

  public LogQuery setIndexName(String indexName) {
    this.indexName = indexName;
    return this;
  }

  /**
   * Index targeted by the query.
   *
   * @return indexName
   */
  @javax.annotation.Nullable
  public String getIndexName() {
    return indexName;
  }

  public LogQuery setUserToken(String userToken) {
    this.userToken = userToken;
    return this;
  }

  /**
   * User identifier.
   *
   * @return userToken
   */
  @javax.annotation.Nullable
  public String getUserToken() {
    return userToken;
  }

  public LogQuery setQueryId(String queryId) {
    this.queryId = queryId;
    return this;
  }

  /**
   * QueryID for the given query.
   *
   * @return queryId
   */
  @javax.annotation.Nullable
  public String getQueryId() {
    return queryId;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    LogQuery logQuery = (LogQuery) o;
    return (
      Objects.equals(this.indexName, logQuery.indexName) &&
      Objects.equals(this.userToken, logQuery.userToken) &&
      Objects.equals(this.queryId, logQuery.queryId)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(indexName, userToken, queryId);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class LogQuery {\n");
    sb.append("    indexName: ").append(toIndentedString(indexName)).append("\n");
    sb.append("    userToken: ").append(toIndentedString(userToken)).append("\n");
    sb.append("    queryId: ").append(toIndentedString(queryId)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}
