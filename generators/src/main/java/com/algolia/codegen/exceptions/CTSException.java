package com.algolia.codegen.exceptions;

public class CTSException extends GeneratorException {

  private boolean skipable;
  private String testName;

  public CTSException(String message) {
    super(message);
  }

  public CTSException(String message, String testName) {
    super(message);
    this.testName = testName;
  }

  public CTSException(String message, Throwable cause) {
    super(message, cause);
  }

  public CTSException(String message, boolean skipable) {
    this(message);
    this.skipable = skipable;
  }

  public boolean isSkipable() {
    return skipable;
  }

  public void setTestName(String testName) {
    this.testName = testName;
  }

  @Override
  public String getMessage() {
    if (testName != null) {
      return "Error in " + testName + ": " + super.getMessage();
    }

    return super.getMessage();
  }
}
