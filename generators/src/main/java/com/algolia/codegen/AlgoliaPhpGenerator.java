package com.algolia.codegen;

import com.algolia.codegen.exceptions.*;
import io.swagger.v3.oas.models.Operation;
import io.swagger.v3.oas.models.servers.Server;
import java.util.List;
import java.util.Map;
import org.openapitools.codegen.CodegenOperation;
import org.openapitools.codegen.SupportingFile;
import org.openapitools.codegen.languages.PhpClientCodegen;

public class AlgoliaPhpGenerator extends PhpClientCodegen {

  @Override
  public String getName() {
    return "algolia-php";
  }

  @Override
  public void processOpts() {
    // generator specific options
    String client = (String) additionalProperties.get("client");
    setApiNameSuffix(Utils.API_SUFFIX);
    setParameterNamingConvention("camelCase");
    additionalProperties.put("modelPackage", "Model\\" + Utils.createClientName(client, "php"));
    additionalProperties.put("invokerPackage", "Algolia\\AlgoliaSearch");

    super.processOpts();

    // Remove base template as we want to change its path
    supportingFiles.removeIf(file -> file.getTemplateFile().equals("Configuration.mustache"));

    supportingFiles.add(new SupportingFile("Configuration.mustache", "lib/Configuration", "Configuration.php"));
  }

  @Override
  public CodegenOperation fromOperation(String path, String httpMethod, Operation operation, List<Server> servers) {
    return Utils.specifyCustomRequest(super.fromOperation(path, httpMethod, operation, servers));
  }

  /** Set default generator options */
  public void setDefaultGeneratorOptions(String client) {
    if (client.equals("search") || client.equals("recommend")) {
      additionalProperties.put("useCache", true);
    }
    additionalProperties.put("isSearchClient", client.equals("search"));
    additionalProperties.put("configClassname", Utils.createClientName(client, "php") + "Config");
  }

  /** Provides an opportunity to inspect and modify operation data before the code is generated. */
  @Override
  public Map<String, Object> postProcessOperationsWithModels(Map<String, Object> objs, List<Object> allModels) {
    Map<String, Object> results = super.postProcessOperationsWithModels(objs, allModels);

    String client = (String) additionalProperties.get("client");

    setDefaultGeneratorOptions(client);
    try {
      Utils.generateServer(client, additionalProperties);
      additionalProperties.put("packageVersion", Utils.getClientConfigField("php", "packageVersion"));
    } catch (GeneratorException e) {
      e.printStackTrace();
      System.exit(1);
    }

    return results;
  }

  public String getComposerPackageName() {
    return "algolia/algoliasearch-client-php";
  }
}
