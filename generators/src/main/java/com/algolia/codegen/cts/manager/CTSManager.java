package com.algolia.codegen.cts.manager;

import com.algolia.codegen.exceptions.GeneratorException;
import java.util.*;
import org.openapitools.codegen.SupportingFile;

public interface CTSManager {
  public void addSupportingFiles(List<SupportingFile> supportingFiles);

  public void addDataToBundle(Map<String, Object> bundle) throws GeneratorException;
}
