package com.algolia.codegen.cts.tests;

import java.util.List;
import java.util.Map;

public class ClientTestData {

  public String testName;
  public boolean autoCreateClient = true;
  public List<Step> steps;
}

class Step {

  public String type;
  public String object;
  public String path;
  public Map<String, Object> parameters;
  public Expected expected;
}

class Expected {

  public String type;
  public Object error;
  public Object match;
}
