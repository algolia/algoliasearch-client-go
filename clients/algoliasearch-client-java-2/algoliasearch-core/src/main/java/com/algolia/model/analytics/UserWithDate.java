package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** UserWithDate */
public class UserWithDate {

  @JsonProperty("date")
  private String date;

  @JsonProperty("count")
  private Integer count;

  public UserWithDate setDate(String date) {
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

  public UserWithDate setCount(Integer count) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UserWithDate userWithDate = (UserWithDate) o;
    return Objects.equals(this.date, userWithDate.date) && Objects.equals(this.count, userWithDate.count);
  }

  @Override
  public int hashCode() {
    return Objects.hash(date, count);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UserWithDate {\n");
    sb.append("    date: ").append(toIndentedString(date)).append("\n");
    sb.append("    count: ").append(toIndentedString(count)).append("\n");
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
