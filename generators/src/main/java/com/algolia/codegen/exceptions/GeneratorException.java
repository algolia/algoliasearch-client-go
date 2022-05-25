package com.algolia.codegen.exceptions;

public class GeneratorException extends RuntimeException {

  public GeneratorException(String message) {
    super(message);
  }

  public GeneratorException(String message, Throwable cause) {
    super(message, cause);
  }
}
