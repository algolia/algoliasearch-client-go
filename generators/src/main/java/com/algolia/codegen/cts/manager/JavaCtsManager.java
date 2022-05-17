package com.algolia.codegen.cts.manager;

import com.algolia.codegen.Utils;
import java.util.*;
import org.openapitools.codegen.SupportingFile;

public class JavaCtsManager extends CtsManager {

  public void addSupportingFiles(List<SupportingFile> supportingFiles) {
    supportingFiles.add(
      new SupportingFile("build.mustache", ".", "build.gradle")
    );
  }

  protected void addExtraToBundle(Map<String, Object> bundle) {
    bundle.put(
      "packageVersion",
      Utils
        .readJsonFile("config/clients.config.json")
        .get("java")
        .get("packageVersion")
        .asText()
    );
  }
}
