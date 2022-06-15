package com.algolia.model.analytics;

import com.google.gson.annotations.SerializedName;
import java.util.Objects;

/** AverageClickEvent */
public class AverageClickEvent {

  @SerializedName("average")
  private Double average;

  @SerializedName("clickCount")
  private Integer clickCount;

  @SerializedName("date")
  private String date;

  public AverageClickEvent setAverage(Double average) {
    this.average = average;
    return this;
  }

  /**
   * The average of all the click count event.
   *
   * @return average
   */
  @javax.annotation.Nonnull
  public Double getAverage() {
    return average;
  }

  public AverageClickEvent setClickCount(Integer clickCount) {
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

  public AverageClickEvent setDate(String date) {
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
    AverageClickEvent averageClickEvent = (AverageClickEvent) o;
    return (
      Objects.equals(this.average, averageClickEvent.average) &&
      Objects.equals(this.clickCount, averageClickEvent.clickCount) &&
      Objects.equals(this.date, averageClickEvent.date)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(average, clickCount, date);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AverageClickEvent {\n");
    sb.append("    average: ").append(toIndentedString(average)).append("\n");
    sb.append("    clickCount: ").append(toIndentedString(clickCount)).append("\n");
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
