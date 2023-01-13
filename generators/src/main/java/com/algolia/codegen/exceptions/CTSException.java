package com.algolia.codegen.exceptions;

public class CTSException extends GeneratorException {

  private boolean skipable;

  public CTSException(String message) {
    super(message);
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
}