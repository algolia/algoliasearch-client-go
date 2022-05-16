package com.algolia.codegen.cts.manager;

import java.util.*;

public class CtsManagerFactory {

  public static CtsManager getManager(String language) {
    switch (language) {
      case "javascript":
        return new JavaScriptCtsManager();
      case "java":
        return new JavaCtsManager();
      case "php":
        return new PhpCtsManager();
    }
    return null;
  }
}
