package com.algolia.codegen.cts;

import com.algolia.codegen.Utils;
import com.algolia.codegen.cts.manager.CTSManager;
import com.algolia.codegen.cts.manager.CTSManagerFactory;
import com.algolia.codegen.cts.tests.*;
import com.algolia.codegen.exceptions.*;
import com.google.common.collect.ImmutableMap.Builder;
import com.samskivert.mustache.Mustache.Lambda;
import java.util.*;
import java.util.Map.Entry;
import java.util.TreeMap;
import org.openapitools.codegen.*;

@SuppressWarnings("unchecked")
public class AlgoliaCTSGenerator extends DefaultCodegen {

  // cache the models
  private final Map<String, CodegenModel> models = new HashMap<>();
  private String language;
  private String client;
  private CTSManager ctsManager;
  private List<TestsGenerator> testsGenerators = new ArrayList<>();

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
    ctsManager = CTSManagerFactory.getManager(language, client);

    String outputFolder = Utils.getClientConfigField(language, "tests", "outputFolder");
    String extension = Utils.getClientConfigField(language, "tests", "extension");

    setTemplateDir("templates/" + language + "/tests");
    setOutputDir("tests/output/" + language);
    ctsManager.addSupportingFiles(supportingFiles);

    testsGenerators.add(new TestsRequest(language, client));
    testsGenerators.add(new TestsClient(language, client));

    for (TestsGenerator testGen : testsGenerators) {
      testGen.addSupportingFiles(supportingFiles, outputFolder, extension);
    }
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
    try {
      Map<String, CodegenOperation> operations = buildOperations(objs);

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
      bundle.put("hasRegionalHost", hasRegionalHost);
      bundle.put("defaultRegion", client.equals("predict") ? "ew" : "us");
      bundle.put("lambda", lambda);
      ctsManager.addDataToBundle(bundle);

      for (TestsGenerator testGen : testsGenerators) {
        try {
          testGen.run(models, operations, bundle);
        } catch (CTSException e) {
          if (e.isSkipable()) {
            System.out.println(e.getMessage());
            continue;
          }
          e.printStackTrace();
          System.exit(1);
        } catch (Exception e) {
          e.printStackTrace();
          System.exit(1);
        }
      }

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

  @Override
  public String escapeUnsafeCharacters(String input) {
    return input;
  }

  public String escapeQuotationMark(String input) {
    return input.replace("\"", "\\\"");
  }
}
