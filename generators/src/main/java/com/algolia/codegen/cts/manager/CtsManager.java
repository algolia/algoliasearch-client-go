package com.algolia.codegen.cts.manager;

import com.algolia.codegen.Utils;
import com.fasterxml.jackson.databind.JsonNode;
import java.util.*;
import org.openapitools.codegen.SupportingFile;

public abstract class CtsManager {

  public abstract void addSupportingFiles(List<SupportingFile> supportingFiles);

  public List<Object> getPackageDependencies() {
    return null;
  }

  protected void addExtraToBundle(Map<String, Object> bundle) {}

  public void addDataToBundle(Map<String, Object> bundle) {
    bundle.put("packageDependencies", this.getPackageDependencies());
    this.addExtraToBundle(bundle);
  }

  protected Object[] getFilteredPackageVersions(List<String> packages) {
    HashMap<String, String> result = new HashMap<>();

    // Read config/openapitools.js for JavaScript
    JsonNode openApiToolsConfig = Utils.readJsonFile(
      "config/openapitools.json"
    );
    Iterator<JsonNode> generatorIterator = openApiToolsConfig
      .get("generator-cli")
      .get("generators")
      .elements();
    while (generatorIterator.hasNext()) {
      JsonNode generator = generatorIterator.next();
      JsonNode additionalProperties = generator.get("additionalProperties");
      if (!additionalProperties.has("packageVersion")) {
        continue;
      }
      String packageName = additionalProperties.get("packageName").asText();
      String packageVersion = additionalProperties
        .get("packageVersion")
        .asText();
      if (packages.contains(packageName)) {
        result.put(packageName, packageVersion);
      }
    }

    JsonNode clientsConfig = Utils.readJsonFile("config/clients.config.json");
    Iterator<JsonNode> clientsIterator = clientsConfig.elements();
    while (clientsIterator.hasNext()) {
      JsonNode client = clientsIterator.next();

      if (!client.has("packageVersion")) {
        continue;
      }
      String packageName = client.get("packageName").asText();
      String packageVersion = client.get("packageVersion").asText();
      if (packages.contains(packageName)) {
        result.put(packageName, packageVersion);
      }
    }

    return result
      .entrySet()
      .stream()
      .map(entry -> {
        Map<String, String> newEntry = new HashMap<>();
        newEntry.put("packageName", entry.getKey());
        newEntry.put("packageVersion", entry.getValue());
        return newEntry;
      })
      .toArray(Object[]::new);
  }
}
