package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** ClickPosition */
public class ClickPosition {

  @JsonProperty("position")
  private List<Integer> position = new ArrayList<>();

  @JsonProperty("clickCount")
  private Integer clickCount;

  public ClickPosition setPosition(List<Integer> position) {
    this.position = position;
    return this;
  }

  public ClickPosition addPosition(Integer positionItem) {
    this.position.add(positionItem);
    return this;
  }

  /**
   * Range of positions with the following pattern: - Positions from 1 to 10 included are displayed
   * in separated groups. - Positions from 11 to 20 included are grouped together. - Positions from
   * 21 and up are grouped together.
   *
   * @return position
   */
  @javax.annotation.Nonnull
  public List<Integer> getPosition() {
    return position;
  }

  public ClickPosition setClickCount(Integer clickCount) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClickPosition clickPosition = (ClickPosition) o;
    return Objects.equals(this.position, clickPosition.position) && Objects.equals(this.clickCount, clickPosition.clickCount);
  }

  @Override
  public int hashCode() {
    return Objects.hash(position, clickCount);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClickPosition {\n");
    sb.append("    position: ").append(toIndentedString(position)).append("\n");
    sb.append("    clickCount: ").append(toIndentedString(clickCount)).append("\n");
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
