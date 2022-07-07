package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** TopHitWithAnalytics */
public class TopHitWithAnalytics {

  @JsonProperty("hit")
  private String hit;

  @JsonProperty("count")
  private Integer count;

  @JsonProperty("clickThroughRate")
  private Double clickThroughRate;

  @JsonProperty("conversionRate")
  private Double conversionRate;

  @JsonProperty("trackedSearchCount")
  private Integer trackedSearchCount;

  @JsonProperty("clickCount")
  private Integer clickCount;

  @JsonProperty("conversionCount")
  private Integer conversionCount;

  public TopHitWithAnalytics setHit(String hit) {
    this.hit = hit;
    return this;
  }

  /**
   * The hit.
   *
   * @return hit
   */
  @javax.annotation.Nonnull
  public String getHit() {
    return hit;
  }

  public TopHitWithAnalytics setCount(Integer count) {
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

  public TopHitWithAnalytics setClickThroughRate(Double clickThroughRate) {
    this.clickThroughRate = clickThroughRate;
    return this;
  }

  /**
   * The click-through rate.
   *
   * @return clickThroughRate
   */
  @javax.annotation.Nonnull
  public Double getClickThroughRate() {
    return clickThroughRate;
  }

  public TopHitWithAnalytics setConversionRate(Double conversionRate) {
    this.conversionRate = conversionRate;
    return this;
  }

  /**
   * The conversion rate.
   *
   * @return conversionRate
   */
  @javax.annotation.Nonnull
  public Double getConversionRate() {
    return conversionRate;
  }

  public TopHitWithAnalytics setTrackedSearchCount(Integer trackedSearchCount) {
    this.trackedSearchCount = trackedSearchCount;
    return this;
  }

  /**
   * The number of tracked search click.
   *
   * @return trackedSearchCount
   */
  @javax.annotation.Nonnull
  public Integer getTrackedSearchCount() {
    return trackedSearchCount;
  }

  public TopHitWithAnalytics setClickCount(Integer clickCount) {
    this.clickCount = clickCount;
    return this;
  }

  /**
   * The number of click event.
   *
   * @return clickCount
   */
  @javax.annotation.Nonnull
  public Integer getClickCount() {
    return clickCount;
  }

  public TopHitWithAnalytics setConversionCount(Integer conversionCount) {
    this.conversionCount = conversionCount;
    return this;
  }

  /**
   * The number of converted clicks.
   *
   * @return conversionCount
   */
  @javax.annotation.Nonnull
  public Integer getConversionCount() {
    return conversionCount;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TopHitWithAnalytics topHitWithAnalytics = (TopHitWithAnalytics) o;
    return (
      Objects.equals(this.hit, topHitWithAnalytics.hit) &&
      Objects.equals(this.count, topHitWithAnalytics.count) &&
      Objects.equals(this.clickThroughRate, topHitWithAnalytics.clickThroughRate) &&
      Objects.equals(this.conversionRate, topHitWithAnalytics.conversionRate) &&
      Objects.equals(this.trackedSearchCount, topHitWithAnalytics.trackedSearchCount) &&
      Objects.equals(this.clickCount, topHitWithAnalytics.clickCount) &&
      Objects.equals(this.conversionCount, topHitWithAnalytics.conversionCount)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(hit, count, clickThroughRate, conversionRate, trackedSearchCount, clickCount, conversionCount);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TopHitWithAnalytics {\n");
    sb.append("    hit: ").append(toIndentedString(hit)).append("\n");
    sb.append("    count: ").append(toIndentedString(count)).append("\n");
    sb.append("    clickThroughRate: ").append(toIndentedString(clickThroughRate)).append("\n");
    sb.append("    conversionRate: ").append(toIndentedString(conversionRate)).append("\n");
    sb.append("    trackedSearchCount: ").append(toIndentedString(trackedSearchCount)).append("\n");
    sb.append("    clickCount: ").append(toIndentedString(clickCount)).append("\n");
    sb.append("    conversionCount: ").append(toIndentedString(conversionCount)).append("\n");
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
