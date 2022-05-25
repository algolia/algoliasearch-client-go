package com.algolia.codegen.cts.manager;

import com.algolia.codegen.Utils;
import com.algolia.codegen.exceptions.GeneratorException;
import com.fasterxml.jackson.databind.JsonNode;
import java.util.*;
import java.util.Map.Entry;
import org.openapitools.codegen.SupportingFile;

public class JavaScriptCtsManager implements CtsManager {

  @Override
  public void addSupportingFiles(List<SupportingFile> supportingFiles) {
    supportingFiles.add(new SupportingFile("package.mustache", ".", "package.json"));
  }

  @Override
  public void addDataToBundle(Map<String, Object> bundle) throws GeneratorException {
    bundle.put("packageDependencies", this.getPackageDependencies());
    bundle.put("utilsPackageVersion", Utils.getClientConfigField("javascript", "utilsPackageVersion"));
  }

  private List<Map<String, String>> getPackageDependencies() {
    List<Map<String, String>> result = new ArrayList<>();

    JsonNode openApiToolsConfig = Utils.readJsonFile("config/openapitools.json");
    for (Entry<String, JsonNode> field : (Iterable<Entry<String, JsonNode>>) () ->
      openApiToolsConfig.get("generator-cli").get("generators").fields()) {
      if (!field.getKey().startsWith("javascript-")) {
        continue;
      }
      JsonNode generator = field.getValue();
      JsonNode additionalProperties = generator.get("additionalProperties");
      String packageName = additionalProperties.get("packageName").asText();
      String packageVersion = additionalProperties.get("packageVersion").asText();

      Map<String, String> newEntry = new HashMap<>();
      newEntry.put("packageName", packageName);
      newEntry.put("packageVersion", packageVersion);
      result.add(newEntry);
    }
    return result;
  }
}
