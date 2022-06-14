package com.algolia.codegen.cts.manager;

import com.algolia.codegen.Utils;
import com.algolia.codegen.exceptions.GeneratorException;
import com.fasterxml.jackson.databind.JsonNode;
import java.util.*;
import java.util.Map.Entry;
import org.openapitools.codegen.SupportingFile;

public class JavaScriptCTSManager implements CTSManager {

  private final String client;
  private final JsonNode openApiToolsConfig;

  public JavaScriptCTSManager(String client) {
    this.client = client;
    this.openApiToolsConfig = Utils.readJsonFile("config/openapitools.json").get("generator-cli").get("generators");
  }

  @Override
  public void addSupportingFiles(List<SupportingFile> supportingFiles) {
    supportingFiles.add(new SupportingFile("package.mustache", ".", "package.json"));
  }

  @Override
  public void addDataToBundle(Map<String, Object> bundle) throws GeneratorException {
    bundle.put("packageDependencies", this.getPackageDependencies());
    bundle.put("utilsPackageVersion", Utils.getClientConfigField("javascript", "utilsPackageVersion"));

    String output = this.openApiToolsConfig.get("javascript-" + client).get("output").asText();
    String packageName = Utils.getClientConfigField("javascript", "npmNamespace") + "/" + output.substring(output.lastIndexOf('/') + 1);
    bundle.put("import", packageName);
  }

  private List<Map<String, String>> getPackageDependencies() {
    List<Map<String, String>> result = new ArrayList<>();

    for (Entry<String, JsonNode> field : (Iterable<Entry<String, JsonNode>>) () -> openApiToolsConfig.fields()) {
      if (!field.getKey().startsWith("javascript-")) {
        continue;
      }
      JsonNode generator = field.getValue();
      JsonNode additionalProperties = generator.get("additionalProperties");
      String output = generator.get("output").asText();
      String packageVersion = additionalProperties.get("packageVersion").asText();
      String packageName = output.substring(output.lastIndexOf('/') + 1);

      Map<String, String> newEntry = new HashMap<>();
      newEntry.put("packageName", packageName);
      newEntry.put("packageVersion", packageVersion);
      result.add(newEntry);
    }
    return result;
  }
}
