package com.algolia.codegen.cts.manager;

import com.algolia.codegen.Utils;
import com.algolia.codegen.exceptions.GeneratorException;
import java.util.*;
import org.openapitools.codegen.SupportingFile;

public class JavaCtsManager implements CtsManager {

  @Override
  public void addSupportingFiles(List<SupportingFile> supportingFiles) {
    supportingFiles.add(new SupportingFile("build.mustache", ".", "build.gradle"));
  }

  @Override
  public void addDataToBundle(Map<String, Object> bundle) throws GeneratorException {
    bundle.put("packageVersion", Utils.getClientConfigField("java", "packageVersion"));
  }
}
