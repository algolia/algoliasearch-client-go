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

  public Object error;
  public Match match;
  public String testSubject;
}

class Match {

  public String regexp;
  public Map<String, Object> objectContaining;
}
