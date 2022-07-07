package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** NoResultsRateEvent */
public class NoResultsRateEvent {

  @JsonProperty("date")
  private String date;

  @JsonProperty("noResultCount")
  private Integer noResultCount;

  @JsonProperty("count")
  private Integer count;

  @JsonProperty("rate")
  private Double rate;

  public NoResultsRateEvent setDate(String date) {
    this.date = date;
    return this;
  }

  /**
   * Date of the event.
   *
   * @return date
   */
  @javax.annotation.Nonnull
  public String getDate() {
    return date;
  }

  public NoResultsRateEvent setNoResultCount(Integer noResultCount) {
    this.noResultCount = noResultCount;
    return this;
  }

  /**
   * The number of occurrences.
   *
   * @return noResultCount
   */
  @javax.annotation.Nonnull
  public Integer getNoResultCount() {
    return noResultCount;
  }

  public NoResultsRateEvent setCount(Integer count) {
    this.count = count;
    return this;
  }

  /**
   * The number of occurrences.
   *
   * @return count
   */
  @javax.annotation.Nonnull
  public Integer getCount() {
    return count;
  }

  public NoResultsRateEvent setRate(Double rate) {
    this.rate = rate;
    return this;
  }

  /**
   * The click-through rate.
   *
   * @return rate
   */
  @javax.annotation.Nonnull
  public Double getRate() {
    return rate;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NoResultsRateEvent noResultsRateEvent = (NoResultsRateEvent) o;
    return (
      Objects.equals(this.date, noResultsRateEvent.date) &&
      Objects.equals(this.noResultCount, noResultsRateEvent.noResultCount) &&
      Objects.equals(this.count, noResultsRateEvent.count) &&
      Objects.equals(this.rate, noResultsRateEvent.rate)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(date, noResultCount, count, rate);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NoResultsRateEvent {\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
    sb.append("    noResultCount: ").append(toIndentedString(noResultCount)).append("\n");
    sb.append("    count: ").append(toIndentedString(count)).append("\n");
    sb.append("    rate: ").append(toIndentedString(rate)).append("\n");
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
