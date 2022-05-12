package com.algolia.codegen;

import io.swagger.v3.oas.models.Operation;
import io.swagger.v3.oas.models.media.Schema;
import io.swagger.v3.oas.models.servers.Server;
import java.util.*;
import org.openapitools.codegen.*;
import org.openapitools.codegen.languages.JavaClientCodegen;
import org.openapitools.codegen.utils.ModelUtils;

@SuppressWarnings("unchecked")
public class AlgoliaJavaGenerator extends JavaClientCodegen {

  /**
   * Configures a friendly name for the generator. This will be used by the generator to select the
   * library with the -g flag.
   *
   * @return the friendly name for the generator
   */
  @Override
  public String getName() {
    return "algolia-java";
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

    try {
      Utils.generateServer(
        Utils.getClientNameKebabCase(results),
        additionalProperties
      );
      additionalProperties.put(
        "packageVersion",
        Utils.getPackageVersion("java")
      );
    } catch (GenerationException e) {
      e.printStackTrace();
      System.exit(1);
    }

    return results;
  }

  @Override
  public Map<String, Object> postProcessAllModels(Map<String, Object> objs) {
    Map<String, Object> models = super.postProcessAllModels(objs);

    for (Object modelContainer : models.values()) {
      CodegenModel model =
        ((Map<String, List<Map<String, CodegenModel>>>) modelContainer).get(
            "models"
          )
          .get(0)
          .get("model");
      if (!model.oneOf.isEmpty()) {
        List<HashMap<String, String>> oneOfList = new ArrayList();

        for (String iterateModel : model.oneOf) {
          HashMap<String, String> oneOfModel = new HashMap();

          oneOfModel.put("type", iterateModel);
          oneOfModel.put(
            "name",
            iterateModel.replace("<", "").replace(">", "")
          );

          oneOfList.add(oneOfModel);
        }

        model.vendorExtensions.put("x-is-one-of-interface", true);
        model.vendorExtensions.put("x-is-one-of-list", oneOfList);
      }
    }

    return models;
  }

  /**
   * Returns human-friendly help for the generator. Provide the consumer with help tips, parameters
   * here
   *
   * @return A string value for the help message
   */
  @Override
  public String getHelp() {
    return "Generates an algolia-java client library.";
  }

  @Override
  public void processOpts() {
    // generator specific options
    setDateLibrary("java8");
    setSourceFolder("algoliasearch-core/src/main/java");
    setInvokerPackage("com.algolia");
    setApiNameSuffix(Utils.API_SUFFIX);

    super.processOpts();

    // Prevent all useless file to generate
    apiTestTemplateFiles.clear();
    modelTestTemplateFiles.clear();
    apiDocTemplateFiles.clear();
    modelDocTemplateFiles.clear();

    supportingFiles.removeIf(file ->
      file.getTemplateFile().equals("build.gradle.mustache") ||
      file.getTemplateFile().equals("settings.gradle.mustache") ||
      file.getTemplateFile().equals("gitignore.mustache") ||
      file.getTemplateFile().equals("ApiCallback.mustache") ||
      file.getTemplateFile().equals("ApiResponse.mustache") ||
      file.getTemplateFile().equals("JSON.mustache") ||
      file.getTemplateFile().equals("ProgressRequestBody.mustache") ||
      file.getTemplateFile().equals("ProgressResponseBody.mustache") ||
      file.getTemplateFile().equals("Pair.mustache")
    );
  }

  @Override
  public String toDefaultValue(Schema schema) {
    // Replace the {} from openapi with new Object()
    if (ModelUtils.isObjectSchema(schema) && schema.getDefault() != null) {
      return "new Object()";
    }
    return super.toDefaultValue(schema);
  }

  @Override
  public String toEnumVarName(String value, String datatype) {
    if ("String".equals(datatype)) {
      // convert camelCase77String to CAMEL_CASE_77_STRING
      return value
        .replaceAll("-", "_")
        .replaceAll("(.+?)([A-Z]|[0-9])", "$1_$2")
        .toUpperCase(Locale.ROOT);
    }
    return super.toEnumVarName(value, datatype);
  }
}
