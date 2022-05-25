package com.algolia.codegen;

import com.algolia.codegen.exceptions.*;
import io.swagger.v3.oas.models.Operation;
import io.swagger.v3.oas.models.media.Schema;
import io.swagger.v3.oas.models.parameters.Parameter;
import io.swagger.v3.oas.models.servers.Server;
import java.util.List;
import java.util.Map;
import org.openapitools.codegen.CodegenOperation;
import org.openapitools.codegen.SupportingFile;
import org.openapitools.codegen.languages.TypeScriptNodeClientCodegen;

public class AlgoliaJavaScriptGenerator extends TypeScriptNodeClientCodegen {

  private String CLIENT;

  @Override
  public String getName() {
    return "algolia-javascript";
  }

  @Override
  public void processOpts() {
    super.processOpts();

    // generator specific options
    setSupportsES6(true);
    setModelPropertyNaming("original");
    setApiPackage("src");
    // clear all supported files to avoid unwanted ones
    supportingFiles.clear();

    supportingFiles.add(new SupportingFile("modelBarrel.mustache", "model", "index.ts"));
    supportingFiles.add(new SupportingFile("browser.mustache", "builds", "browser.ts"));
    supportingFiles.add(new SupportingFile("node.mustache", "builds", "node.ts"));

    // root
    supportingFiles.add(new SupportingFile("index.mustache", "", "index.js"));
    supportingFiles.add(new SupportingFile("index.d.mustache", "", "index.d.ts"));

    supportingFiles.add(new SupportingFile("package.mustache", "", "package.json"));
    supportingFiles.add(new SupportingFile("tsconfig.mustache", "", "tsconfig.json"));
  }

  @Override
  public CodegenOperation fromOperation(String path, String httpMethod, Operation operation, List<Server> servers) {
    return Utils.specifyCustomRequest(super.fromOperation(path, httpMethod, operation, servers));
  }

  /** Set default generator options */
  private void setDefaultGeneratorOptions() {
    String apiName = CLIENT + Utils.API_SUFFIX;

    additionalProperties.put("apiName", apiName);
    additionalProperties.put("capitalizedApiName", Utils.capitalize(apiName));
    additionalProperties.put("algoliaAgent", Utils.capitalize(CLIENT));
    additionalProperties.put("gitRepoId", "algoliasearch-client-javascript");
    additionalProperties.put("isSearchClient", CLIENT.equals("search"));
  }

  /** Provides an opportunity to inspect and modify operation data before the code is generated. */
  @Override
  public Map<String, Object> postProcessOperationsWithModels(Map<String, Object> objs, List<Object> allModels) {
    Map<String, Object> results = super.postProcessOperationsWithModels(objs, allModels);

    CLIENT = Utils.getClientNameCamelCase(results);

    setDefaultGeneratorOptions();
    try {
      Utils.generateServer(Utils.getClientNameKebabCase(results), additionalProperties);
      additionalProperties.put("utilsPackageVersion", Utils.getClientConfigField("javascript", "utilsPackageVersion"));
    } catch (GeneratorException e) {
      e.printStackTrace();
      System.exit(1);
    }

    List<CodegenOperation> operations = ((Map<String, List<CodegenOperation>>) results.get("operations")).get("operation");

    // We read operations and detect if we should wrap parameters under an object.
    // We only wrap if there is a mix between body parameters and other parameters.
    for (CodegenOperation ope : operations) {
      // Nothing to wrap as there is no parameters
      if (!ope.hasParams) {
        continue;
      }

      boolean hasBodyParams = !ope.bodyParams.isEmpty();
      boolean hasHeaderParams = !ope.headerParams.isEmpty();
      boolean hasQueryParams = !ope.queryParams.isEmpty();
      boolean hasPathParams = !ope.pathParams.isEmpty();

      // If there is nothing but body params, we just check if it's a single param
      if (hasBodyParams && !hasHeaderParams && !hasQueryParams && !hasPathParams) {
        // At this point the single parameter is already an object, to avoid double wrapping
        // we skip it
        if (ope.bodyParams.size() == 1 && !ope.bodyParams.get(0).isArray) {
          ope.vendorExtensions.put("x-is-single-body-param", true);
          continue;
        }
      }

      // Any other cases here are wrapped
      ope.vendorExtensions.put("x-create-wrapping-object", true);
    }

    return results;
  }

  /**
   * The `apiSuffix` option is not supported on the TypeScript client, so we override the names
   * method to use it with our suffix.
   */

  /** The `apiName` is capitalized. */
  @Override
  public String toApiName(String name) {
    if (name.length() == 0) {
      return "Default" + Utils.API_SUFFIX;
    }

    return Utils.capitalize(CLIENT + Utils.API_SUFFIX);
  }

  /** The `apiFileName` is in camelCase. */
  @Override
  public String toApiFilename(String name) {
    if (name.length() == 0) {
      return "default" + Utils.API_SUFFIX;
    }

    return CLIENT + Utils.API_SUFFIX;
  }

  /** The `apiFileName` is in camelCase. */
  @Override
  public String apiFilename(String templateName, String tag) {
    return super.apiFilename(templateName, toApiFilename(CLIENT));
  }

  @Override
  protected String getParameterDataType(Parameter parameter, Schema p) {
    String type = super.getParameterDataType(parameter, p);
    // openapi generator is wrong, 'object' is not a fit all object, we need 'any'
    // we use replace because there might be more to this type, like '| undefined'
    return type.replace("{ [key: string]: object; }", "Record<string, any>");
  }

  @Override
  public String toInstantiationType(Schema schema) {
    String type = super.toInstantiationType(schema);

    // Same as the `getParameterDataType` but for the models, the generator
    // consider `object` as a fit all which is wrong in TypeScript
    return type.replace("null<String, object>", "Record<string, any>");
  }
}
