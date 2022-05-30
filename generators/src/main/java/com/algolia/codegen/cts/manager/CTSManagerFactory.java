package com.algolia.codegen.cts.manager;

import java.util.*;

public class CTSManagerFactory {

  private CTSManagerFactory() {}

  public static CTSManager getManager(String language, String client) {
    switch (language) {
      case "javascript":
        return new JavaScriptCTSManager(client);
      case "java":
        return new JavaCTSManager(client);
      case "php":
        return new PhpCTSManager();
    }
    return null;
  }
}
