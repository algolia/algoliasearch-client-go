package com.algolia.codegen.cts.manager;

import com.algolia.codegen.Utils;
import com.algolia.codegen.exceptions.GeneratorException;
import com.fasterxml.jackson.databind.JsonNode;
import java.util.*;
import org.openapitools.codegen.SupportingFile;

public class JavaScriptCTSManager implements CTSManager {

  private final String client;

  public JavaScriptCTSManager(String client) {
    this.client = client;
  }

  @Override
  public void addSupportingFiles(List<SupportingFile> supportingFiles) {}

  @Override
  public void addDataToBundle(Map<String, Object> bundle) throws GeneratorException {
    bundle.put("utilsPackageVersion", Utils.getClientConfigField("javascript", "utilsPackageVersion"));
    bundle.put("npmNamespace", Utils.getClientConfigField("javascript", "npmNamespace"));

    JsonNode openApiToolsConfig = Utils.readJsonFile("config/openapitools.json").get("generator-cli").get("generators");

    String output = openApiToolsConfig.get("javascript-" + client).get("output").asText();
    String clientName = output.substring(output.lastIndexOf('/') + 1);
    String npmNamespace = Utils.getClientConfigField("javascript", "npmNamespace");

    if (clientName.equals("algoliasearch")) {
      bundle.put("import", npmNamespace + "/" + "algoliasearch/lite");
    } else {
      bundle.put("import", npmNamespace + "/" + clientName);
    }
  }
}
