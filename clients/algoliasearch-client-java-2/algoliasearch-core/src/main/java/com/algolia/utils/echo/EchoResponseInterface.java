package com.algolia.utils.echo;

import java.util.Map;

public interface EchoResponseInterface {
  public String getPath();

  public String getMethod();

  public String getBody();

  public Map<String, String> getQueryParams();
}
