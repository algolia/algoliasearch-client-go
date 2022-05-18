package com.algolia.codegen.cts.manager;

import com.algolia.codegen.GenerationException;
import com.algolia.codegen.Utils;
import java.util.*;
import org.openapitools.codegen.SupportingFile;

public class JavaCtsManager extends CtsManager {

  public void addSupportingFiles(List<SupportingFile> supportingFiles) {
    supportingFiles.add(
      new SupportingFile("build.mustache", ".", "build.gradle")
    );
  }

  protected void addExtraToBundle(Map<String, Object> bundle)
    throws GenerationException {
    bundle.put(
      "packageVersion",
      Utils.getClientConfigField("java", "packageVersion")
    );
  }
}
