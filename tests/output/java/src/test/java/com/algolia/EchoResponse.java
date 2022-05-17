package com.algolia;

import java.util.Map;

public class EchoResponse {

  public String path;
  public String method;
  public String body;
  public Map<String, Object> queryParameters;
  public Map<String, String> headers;
}
