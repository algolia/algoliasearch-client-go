package com.algolia.codegen.exceptions;

public class ConfigException extends GeneratorException {

  public ConfigException(String message) {
    super(message);
  }

  public ConfigException(String message, Throwable cause) {
    super(message, cause);
  }
}
