package com.algolia.codegen;

import com.algolia.codegen.exceptions.*;
import io.swagger.v3.oas.models.Operation;
import io.swagger.v3.oas.models.media.Schema;
import io.swagger.v3.oas.models.servers.Server;
import java.util.*;
import org.openapitools.codegen.*;
import org.openapitools.codegen.languages.JavaClientCodegen;
import org.openapitools.codegen.utils.ModelUtils;

@SuppressWarnings("unchecked")
public class AlgoliaJavaGenerator extends JavaClientCodegen {

  @Override
  public String getName() {
    return "algolia-java";
  }

  @Override
  public void processOpts() {
    // generator specific options
    String client = (String) additionalProperties.get("client");
    setDateLibrary("java8");
    setLibrary("okhttp-gson");
    setSourceFolder("algoliasearch-core/src/main/java");
    setGroupId("com.algolia");
    setModelPackage("com.algolia.model." + Utils.camelize(client));
    additionalProperties.put("invokerPackage", "com.algolia");
    setApiPackage("com.algolia.api");
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
  public CodegenOperation fromOperation(String path, String httpMethod, Operation operation, List<Server> servers) {
    return Utils.specifyCustomRequest(super.fromOperation(path, httpMethod, operation, servers));
  }

  @Override
  public Map<String, Object> postProcessOperationsWithModels(Map<String, Object> objs, List<Object> allModels) {
    Map<String, Object> results = super.postProcessOperationsWithModels(objs, allModels);

    String client = (String) additionalProperties.get("client");
    additionalProperties.put("isSearchClient", client.equals("search"));

    try {
      Utils.generateServer(client, additionalProperties);

      additionalProperties.put("packageVersion", Utils.getClientConfigField("java", "packageVersion"));
    } catch (GeneratorException e) {
      e.printStackTrace();
      System.exit(1);
    }

    return results;
  }

  @Override
  public Map<String, Object> postProcessAllModels(Map<String, Object> objs) {
    Map<String, Object> models = super.postProcessAllModels(objs);

    for (Object modelContainer : models.values()) {
      CodegenModel model = ((Map<String, List<Map<String, CodegenModel>>>) modelContainer).get("models").get(0).get("model");
      if (!model.oneOf.isEmpty()) {
        List<HashMap<String, String>> oneOfList = new ArrayList();

        for (String iterateModel : model.oneOf) {
          HashMap<String, String> oneOfModel = new HashMap();

          oneOfModel.put("type", iterateModel);
          oneOfModel.put("name", iterateModel.replace("<", "").replace(">", ""));

          oneOfList.add(oneOfModel);
        }

        model.vendorExtensions.put("x-is-one-of-interface", true);
        model.vendorExtensions.put("x-is-one-of-list", oneOfList);
      }
    }

    return models;
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
      return value.replaceAll("-", "_").replaceAll("(.+?)([A-Z]|[0-9])", "$1_$2").toUpperCase(Locale.ROOT);
    }
    return super.toEnumVarName(value, datatype);
  }
}
