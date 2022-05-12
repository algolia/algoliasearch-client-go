package com.algolia.codegen;

public class GenerationException extends Exception {

  public GenerationException(String message) {
    super(message);
  }

  public GenerationException(String message, Throwable cause) {
    super(message, cause);
  }
}
