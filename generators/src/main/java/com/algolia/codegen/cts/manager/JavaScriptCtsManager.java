package com.algolia.codegen.cts.manager;

import com.algolia.codegen.Utils;
import com.fasterxml.jackson.databind.JsonNode;
import java.util.*;
import org.openapitools.codegen.SupportingFile;

public class JavaScriptCtsManager extends CtsManager {

  public void addSupportingFiles(List<SupportingFile> supportingFiles) {
    supportingFiles.add(
      new SupportingFile("package.mustache", ".", "package.json")
    );
  }

  private List<Object> getPackageDependencies() {
    List<Object> result = new ArrayList<Object>();

    JsonNode openApiToolsConfig = Utils.readJsonFile(
      "config/openapitools.json"
    );
    Iterator<Map.Entry<String, JsonNode>> fieldIterator = openApiToolsConfig
      .get("generator-cli")
      .get("generators")
      .fields();

    while (fieldIterator.hasNext()) {
      Map.Entry<String, JsonNode> field = fieldIterator.next();
      if (!field.getKey().startsWith("javascript-")) {
        continue;
      }
      JsonNode generator = field.getValue();
      JsonNode additionalProperties = generator.get("additionalProperties");
      String packageName = additionalProperties.get("packageName").asText();
      String packageVersion = additionalProperties
        .get("packageVersion")
        .asText();

      Map<String, String> newEntry = new HashMap<>();
      newEntry.put("packageName", packageName);
      newEntry.put("packageVersion", packageVersion);
      result.add(newEntry);
    }
    return result;
  }

  protected void addExtraToBundle(Map<String, Object> bundle) {
    bundle.put("packageDependencies", this.getPackageDependencies());
    bundle.put("utilsPackageVersion", this.getUtilsPackageVersion());
  }

  private String getUtilsPackageVersion() {
    JsonNode openApiToolsConfig = Utils.readJsonFile(
      "config/openapitools.json"
    );

    String utilsPackageVersion = openApiToolsConfig
      .get("generator-cli")
      .get("generators")
      .get("javascript-search")
      .get("additionalProperties")
      .get("utilsPackageVersion")
      .asText();

    return utilsPackageVersion;
  }
}
