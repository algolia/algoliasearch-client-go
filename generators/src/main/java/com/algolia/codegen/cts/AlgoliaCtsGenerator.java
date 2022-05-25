package com.algolia.codegen.cts;

import com.algolia.codegen.Utils;
import com.algolia.codegen.cts.manager.CtsManager;
import com.algolia.codegen.cts.manager.CtsManagerFactory;
import com.algolia.codegen.exceptions.*;
import com.fasterxml.jackson.core.JsonParseException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.google.common.collect.ImmutableMap.Builder;
import com.samskivert.mustache.Mustache.Lambda;
import io.swagger.v3.core.util.Json;
import java.io.File;
import java.io.IOException;
import java.util.*;
import java.util.Map.Entry;
import java.util.TreeMap;
import org.openapitools.codegen.*;

@SuppressWarnings("unchecked")
public class AlgoliaCtsGenerator extends DefaultCodegen {

  // cache the models
  private final Map<String, CodegenModel> models = new HashMap<>();
  private String language;
  private String client;
  private String packageName;
  private CtsManager ctsManager;

  @Override
  public CodegenType getTag() {
    return CodegenType.OTHER;
  }

  @Override
  public String getName() {
    return "algolia-cts";
  }

  @Override
  public void processOpts() {
    super.processOpts();

    language = (String) additionalProperties.get("language");
    client = (String) additionalProperties.get("client");
    packageName = (String) additionalProperties.get("packageName");
    ctsManager = CtsManagerFactory.getManager(language);

    String outputFolder = Utils.getClientConfigField(language, "tests", "outputFolder");
    String extension = Utils.getClientConfigField(language, "tests", "extension");

    setTemplateDir("tests/CTS/methods/requests/templates/" + language);
    setOutputDir("tests/output/" + language);
    String clientName = language.equals("php") ? Utils.createClientName(client, language) : client;
    supportingFiles.add(new SupportingFile("requests.mustache", outputFolder + "/methods/requests", clientName + extension));

    ctsManager.addSupportingFiles(supportingFiles);
  }

  @Override
  public Map<String, Object> postProcessAllModels(Map<String, Object> objs) {
    Map<String, Object> mod = super.postProcessAllModels(objs);
    for (Entry<String, Object> entry : mod.entrySet()) {
      List<Object> innerModel = ((Map<String, List<Object>>) entry.getValue()).get("models");
      if (!innerModel.isEmpty()) {
        models.put(entry.getKey(), (CodegenModel) ((Map<String, Object>) innerModel.get(0)).get("model"));
      }
    }
    return mod;
  }

  @Override
  protected Builder<String, Lambda> addMustacheLambdas() {
    Builder<String, Lambda> lambdas = super.addMustacheLambdas();

    lambdas.put("escapequotes", new EscapeQuotesLambda());
    return lambdas;
  }

  @Override
  public Map<String, Object> postProcessSupportingFileData(Map<String, Object> objs) {
    Map<String, Request[]> cts = null;
    try {
      cts = loadCTS();

      Map<String, CodegenOperation> operations = buildOperations(objs);

      // The return value of this function is not used, we need to modify the param itself.
      Object lambda = objs.get("lambda");
      List<CodegenServer> servers = (List<CodegenServer>) objs.get("servers");
      boolean hasRegionalHost = servers
        .stream()
        .anyMatch(server -> server.variables.stream().anyMatch(variable -> variable.name.equals("region")));

      Map<String, Object> bundle = objs;
      bundle.clear();

      // We can put whatever we want in the bundle, and it will be accessible in the template
      bundle.put("client", Utils.createClientName(client, language) + "Client");
      bundle.put("clientPrefix", Utils.createClientName(client, language));
      bundle.put("import", createImportName());
      bundle.put("hasRegionalHost", hasRegionalHost);
      bundle.put("defaultRegion", client.equals("predict") ? "ew" : "us");
      bundle.put("lambda", lambda);
      ctsManager.addDataToBundle(bundle);

      List<Object> blocks = new ArrayList<>();
      ParametersWithDataType paramsType = new ParametersWithDataType(models, language);

      for (Entry<String, CodegenOperation> entry : operations.entrySet()) {
        String operationId = entry.getKey();
        if (!cts.containsKey(operationId)) {
          throw new CTSException("operationId " + operationId + " does not exist in the spec");
        }
        Request[] op = cts.get(operationId);

        List<Object> tests = new ArrayList<>();
        for (int i = 0; i < op.length; i++) {
          Map<String, Object> test = paramsType.buildJSONForRequest(operationId, op[i], entry.getValue(), i);
          tests.add(test);
        }
        Map<String, Object> testObj = new HashMap<>();
        testObj.put("tests", tests);
        testObj.put("operationId", operationId);
        blocks.add(testObj);
      }
      bundle.put("blocks", blocks);

      return bundle;
    } catch (CTSException e) {
      if (e.isSkipable()) {
        System.out.println(e.getMessage());
        System.exit(0);
      }
      e.printStackTrace();
      System.exit(1);
    } catch (Exception e) {
      e.printStackTrace();
      System.exit(1);
    }
    return null;
  }

  private Map<String, Request[]> loadCTS() throws JsonParseException, JsonMappingException, IOException, CTSException {
    TreeMap<String, Request[]> cts = new TreeMap<>();
    String clientName = client;

    // This special case allow us to read the `search` CTS to generated the tests for the
    // `algoliasearch-lite` client, which is only available in JavaScript
    if (language.equals("javascript") && clientName.equals("algoliasearch-lite")) {
      clientName = "search";
    }

    File dir = new File("tests/CTS/methods/requests/" + clientName);
    File commonTestDir = new File("tests/CTS/methods/requests/common");
    if (!dir.exists()) {
      throw new CTSException("CTS not found at " + dir.getAbsolutePath(), true);
    }
    if (!commonTestDir.exists()) {
      throw new CTSException("CTS not found at " + commonTestDir.getAbsolutePath(), true);
    }
    for (File f : dir.listFiles()) {
      cts.put(f.getName().replace(".json", ""), Json.mapper().readValue(f, Request[].class));
    }
    for (File f : commonTestDir.listFiles()) {
      cts.put(f.getName().replace(".json", ""), Json.mapper().readValue(f, Request[].class));
    }
    return cts;
  }

  // operationId -> CodegenOperation
  private TreeMap<String, CodegenOperation> buildOperations(Map<String, Object> objs) {
    HashMap<String, CodegenOperation> result = new HashMap<>();
    List<Map<String, Object>> apis = ((Map<String, List<Map<String, Object>>>) objs.get("apiInfo")).get("apis");

    for (Map<String, Object> api : apis) {
      String apiName = ((String) api.get("baseName")).toLowerCase();
      if (!apiName.equals(client.replace("-", ""))) {
        continue;
      }

      List<CodegenOperation> operations = ((Map<String, List<CodegenOperation>>) api.get("operations")).get("operation");

      for (CodegenOperation ope : operations) {
        result.put(ope.operationId, ope);
      }
    }

    return new TreeMap<String, CodegenOperation>(result);
  }

  private String createImportName() {
    if (!language.equals("java")) {
      return this.packageName;
    }
    String[] clientParts = client.split("-");
    // do not capitalize the first part
    String name = clientParts[0];
    for (int i = 1; i < clientParts.length; i++) {
      name += Utils.capitalize(clientParts[i]);
    }

    return name;
  }

  @Override
  public String escapeUnsafeCharacters(String input) {
    return input;
  }

  public String escapeQuotationMark(String input) {
    return input.replace("\"", "\\\"");
  }
}
