package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** Parameters for the search. */
public class SearchRulesParams {

  @SerializedName("query")
  private String query = "";

  @SerializedName("anchoring")
  private Anchoring anchoring;

  @SerializedName("context")
  private String context;

  @SerializedName("page")
  private Integer page = 0;

  @SerializedName("hitsPerPage")
  private Integer hitsPerPage = 20;

  @SerializedName("enabled")
  private Boolean enabled;

  @SerializedName("requestOptions")
  private List<Object> requestOptions = null;

  public SearchRulesParams setQuery(String query) {
    this.query = query;
    return this;
  }

  /**
   * Full text query.
   *
   * @return query
   */
  @javax.annotation.Nullable
  public String getQuery() {
    return query;
  }

  public SearchRulesParams setAnchoring(Anchoring anchoring) {
    this.anchoring = anchoring;
    return this;
  }

  /**
   * Get anchoring
   *
   * @return anchoring
   */
  @javax.annotation.Nullable
  public Anchoring getAnchoring() {
    return anchoring;
  }

  public SearchRulesParams setContext(String context) {
    this.context = context;
    return this;
  }

  /**
   * Restricts matches to contextual rules with a specific context (exact match).
   *
   * @return context
   */
  @javax.annotation.Nullable
  public String getContext() {
    return context;
  }

  public SearchRulesParams setPage(Integer page) {
    this.page = page;
    return this;
  }

  /**
   * Requested page (zero-based).
   *
   * @return page
   */
  @javax.annotation.Nullable
  public Integer getPage() {
    return page;
  }

  public SearchRulesParams setHitsPerPage(Integer hitsPerPage) {
    this.hitsPerPage = hitsPerPage;
    return this;
  }

  /**
   * Maximum number of hits in a page. Minimum is 1, maximum is 1000.
   *
   * @return hitsPerPage
   */
  @javax.annotation.Nullable
  public Integer getHitsPerPage() {
    return hitsPerPage;
  }

  public SearchRulesParams setEnabled(Boolean enabled) {
    this.enabled = enabled;
    return this;
  }

  /**
   * When specified, restricts matches to rules with a specific enabled status. When absent
   * (default), all rules are retrieved, regardless of their enabled status.
   *
   * @return enabled
   */
  @javax.annotation.Nullable
  public Boolean getEnabled() {
    return enabled;
  }

  public SearchRulesParams setRequestOptions(List<Object> requestOptions) {
    this.requestOptions = requestOptions;
    return this;
  }

  public SearchRulesParams addRequestOptionsItem(Object requestOptionsItem) {
    if (this.requestOptions == null) {
      this.requestOptions = new ArrayList<>();
    }
    this.requestOptions.add(requestOptionsItem);
    return this;
  }

  /**
   * A mapping of requestOptions to send along with the request.
   *
   * @return requestOptions
   */
  @javax.annotation.Nullable
  public List<Object> getRequestOptions() {
    return requestOptions;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SearchRulesParams searchRulesParams = (SearchRulesParams) o;
    return (
      Objects.equals(this.query, searchRulesParams.query) &&
      Objects.equals(this.anchoring, searchRulesParams.anchoring) &&
      Objects.equals(this.context, searchRulesParams.context) &&
      Objects.equals(this.page, searchRulesParams.page) &&
      Objects.equals(this.hitsPerPage, searchRulesParams.hitsPerPage) &&
      Objects.equals(this.enabled, searchRulesParams.enabled) &&
      Objects.equals(this.requestOptions, searchRulesParams.requestOptions)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(
      query,
      anchoring,
      context,
      page,
      hitsPerPage,
      enabled,
      requestOptions
    );
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SearchRulesParams {\n");
    sb.append("    query: ").append(toIndentedString(query)).append("\n");
    sb
      .append("    anchoring: ")
      .append(toIndentedString(anchoring))
      .append("\n");
    sb.append("    context: ").append(toIndentedString(context)).append("\n");
    sb.append("    page: ").append(toIndentedString(page)).append("\n");
    sb
      .append("    hitsPerPage: ")
      .append(toIndentedString(hitsPerPage))
      .append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb
      .append("    requestOptions: ")
      .append(toIndentedString(requestOptions))
      .append("\n");
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
