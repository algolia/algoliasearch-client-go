package com.algolia.codegen.cts.tests;

import java.util.List;
import java.util.Map;
import org.openapitools.codegen.CodegenModel;
import org.openapitools.codegen.CodegenOperation;
import org.openapitools.codegen.SupportingFile;

public interface TestsGenerator {
  public void addSupportingFiles(List<SupportingFile> supportingFiles, String outputFolder, String extension);

  public void run(Map<String, CodegenModel> models, Map<String, CodegenOperation> operations, Map<String, Object> bundle) throws Exception;
}
