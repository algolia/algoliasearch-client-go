package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** ConversionRateEvent */
public class ConversionRateEvent {

  @JsonProperty("rate")
  private Double rate;

  @JsonProperty("trackedSearchCount")
  private Integer trackedSearchCount;

  @JsonProperty("conversionCount")
  private Integer conversionCount;

  @JsonProperty("date")
  private String date;

  public ConversionRateEvent setRate(Double rate) {
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

  public ConversionRateEvent setTrackedSearchCount(Integer trackedSearchCount) {
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

  public ConversionRateEvent setConversionCount(Integer conversionCount) {
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

  public ConversionRateEvent setDate(String date) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ConversionRateEvent conversionRateEvent = (ConversionRateEvent) o;
    return (
      Objects.equals(this.rate, conversionRateEvent.rate) &&
      Objects.equals(this.trackedSearchCount, conversionRateEvent.trackedSearchCount) &&
      Objects.equals(this.conversionCount, conversionRateEvent.conversionCount) &&
      Objects.equals(this.date, conversionRateEvent.date)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(rate, trackedSearchCount, conversionCount, date);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ConversionRateEvent {\n");
    sb.append("    rate: ").append(toIndentedString(rate)).append("\n");
    sb.append("    trackedSearchCount: ").append(toIndentedString(trackedSearchCount)).append("\n");
    sb.append("    conversionCount: ").append(toIndentedString(conversionCount)).append("\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
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
