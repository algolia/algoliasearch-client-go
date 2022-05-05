package com.algolia.codegen;

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
  public CodegenOperation fromOperation(
    String path,
    String httpMethod,
    Operation operation,
    List<Server> servers
  ) {
    return Utils.specifyCustomRequest(
      super.fromOperation(path, httpMethod, operation, servers)
    );
  }

  /** Set default generator options */
  public void setDefaultGeneratorOptions(String client) {
    if (client.equals("search") || client.equals("recommend")) {
      additionalProperties.put("useCache", true);
    }

    additionalProperties.put(
      "configClassname",
      Utils.createClientName(client, "php") + "Config"
    );
  }

  /** Provides an opportunity to inspect and modify operation data before the code is generated. */
  @Override
  public Map<String, Object> postProcessOperationsWithModels(
    Map<String, Object> objs,
    List<Object> allModels
  ) {
    Map<String, Object> results = super.postProcessOperationsWithModels(
      objs,
      allModels
    );

    String client = Utils.getClientNameKebabCase(results);

    setDefaultGeneratorOptions(client);
    Utils.generateServer(client, additionalProperties);

    return results;
  }

  @Override
  public void processOpts() {
    // generator specific options
    setApiNameSuffix(Utils.API_SUFFIX);
    setParameterNamingConvention("camelCase");

    super.processOpts();

    // Remove base template as we want to change its path
    supportingFiles.removeIf(file ->
      file.getTemplateFile().equals("Configuration.mustache")
    );

    supportingFiles.add(
      new SupportingFile(
        "Configuration.mustache",
        "lib/Configuration",
        "Configuration.php"
      )
    );
  }
}
