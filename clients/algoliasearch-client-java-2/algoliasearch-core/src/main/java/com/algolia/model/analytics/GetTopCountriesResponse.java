// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** GetTopCountriesResponse */
public class GetTopCountriesResponse {

  @JsonProperty("countries")
  private List<TopCountry> countries = new ArrayList<>();

  public GetTopCountriesResponse setCountries(List<TopCountry> countries) {
    this.countries = countries;
    return this;
  }

  public GetTopCountriesResponse addCountries(TopCountry countriesItem) {
    this.countries.add(countriesItem);
    return this;
  }

  /**
   * A list of countries with their count.
   *
   * @return countries
   */
  @javax.annotation.Nonnull
  public List<TopCountry> getCountries() {
    return countries;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GetTopCountriesResponse getTopCountriesResponse = (GetTopCountriesResponse) o;
    return Objects.equals(this.countries, getTopCountriesResponse.countries);
  }

  @Override
  public int hashCode() {
    return Objects.hash(countries);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GetTopCountriesResponse {\n");
    sb.append("    countries: ").append(toIndentedString(countries)).append("\n");
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
