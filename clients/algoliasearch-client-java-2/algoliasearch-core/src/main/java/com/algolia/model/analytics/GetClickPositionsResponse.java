// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** GetClickPositionsResponse */
public class GetClickPositionsResponse {

  @JsonProperty("positions")
  private List<ClickPosition> positions = new ArrayList<>();

  public GetClickPositionsResponse setPositions(List<ClickPosition> positions) {
    this.positions = positions;
    return this;
  }

  public GetClickPositionsResponse addPositions(ClickPosition positionsItem) {
    this.positions.add(positionsItem);
    return this;
  }

  /**
   * A list of the click positions with their click count.
   *
   * @return positions
   */
  @javax.annotation.Nonnull
  public List<ClickPosition> getPositions() {
    return positions;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GetClickPositionsResponse getClickPositionsResponse = (GetClickPositionsResponse) o;
    return Objects.equals(this.positions, getClickPositionsResponse.positions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(positions);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GetClickPositionsResponse {\n");
    sb.append("    positions: ").append(toIndentedString(positions)).append("\n");
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
