package com.algolia.codegen.cts.tests;

import com.algolia.codegen.Utils;
import com.algolia.codegen.exceptions.CTSException;
import io.swagger.v3.core.util.Json;
import java.io.File;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import org.openapitools.codegen.CodegenModel;
import org.openapitools.codegen.CodegenOperation;
import org.openapitools.codegen.SupportingFile;

public class TestsClient implements TestsGenerator {

  private final String language, client;

  public TestsClient(String language, String client) {
    this.language = language;
    this.client = client;
  }

  private Map<String, ClientTestData[]> loadCTS() throws Exception {
    if (!available()) {
      throw new CTSException("Templates not found for client test", true);
    }

    File dir = new File("tests/CTS/client/" + client);
    if (!dir.exists()) {
      throw new CTSException("CTS not found at " + dir.getAbsolutePath(), true);
    }
    Map<String, ClientTestData[]> cts = new HashMap<>();
    for (File f : dir.listFiles()) {
      cts.put(f.getName().replace(".json", ""), Json.mapper().readValue(f, ClientTestData[].class));
    }
    return cts;
  }

  @Override
  public boolean available() {
    // no algoliasearch-lite client test for now
    if (language.equals("javascript") && client.equals("algoliasearch-lite")) {
      return false;
    }

    File templates = new File("templates/" + language + "/tests/client/suite.mustache");
    return templates.exists();
  }

  @Override
  public void addSupportingFiles(List<SupportingFile> supportingFiles, String outputFolder, String extension) {
    if (!available()) {
      return;
    }
    String clientName = language.equals("php") ? Utils.createClientName(client, language) : client;
    supportingFiles.add(new SupportingFile("client/suite.mustache", outputFolder + "/client", clientName + extension));
  }

  public void run(Map<String, CodegenModel> models, Map<String, CodegenOperation> operations, Map<String, Object> bundle) throws Exception {
    Map<String, ClientTestData[]> cts = loadCTS();
    ParametersWithDataType paramsType = new ParametersWithDataType(models, language);

    List<Object> blocks = new ArrayList<>();
    for (Map.Entry<String, ClientTestData[]> blockEntry : cts.entrySet()) {
      Map<String, Object> testObj = new HashMap<>();
      List<Object> tests = new ArrayList<>();
      int testIndex = 0;
      for (ClientTestData test : blockEntry.getValue()) {
        Map<String, Object> testOut = new HashMap<>();
        List<Object> steps = new ArrayList<>();
        testOut.put("testName", test.testName);
        testOut.put("testIndex", testIndex++);
        testOut.put("autoCreateClient", test.autoCreateClient);
        for (Step step : test.steps) {
          Map<String, Object> stepOut = new HashMap<>();
          CodegenOperation ope = null;
          if (step.type.equals("createClient")) {
            stepOut.put("isCreateClient", true);
          } else if (step.type.equals("variable")) {
            stepOut.put("isVariable", true);
          } else if (step.type.equals("method")) {
            ope = operations.get(step.path);
            if (ope == null) {
              throw new CTSException("Cannot find operation for method: " + step.path);
            }
            stepOut.put("returnType", ope.returnType);
            stepOut.put("isMethod", true);
          }

          stepOut.put("object", step.object);
          stepOut.put("path", step.path);
          paramsType.enhanceParameters(step.parameters, stepOut, ope);

          if (step.expected.testSubject == null) {
            stepOut.put("testSubject", "result");
          } else {
            switch (step.expected.testSubject) {
              case "userAgent":
                stepOut.put("testUserAgent", true);
                break;
              default:
                stepOut.put("testSubject", step.expected.testSubject);
            }
          }
          if (step.expected.error != null) {
            stepOut.put("isError", true);
            stepOut.put("expectedError", step.expected.error);
          } else if (step.expected.match != null) {
            Map<String, Object> match = new HashMap<>();
            match.put("regexp", step.expected.match.regexp);
            if (step.expected.match.objectContaining != null) {
              Map<String, Object> objectContaining = new HashMap<>();
              paramsType.enhanceParameters(step.expected.match.objectContaining, objectContaining);
              match.put("objectContaining", objectContaining);
            }
            stepOut.put("match", match);
          }
          steps.add(stepOut);
        }
        testOut.put("steps", steps);
        tests.add(testOut);
      }
      testObj.put("tests", tests);
      testObj.put("testType", blockEntry.getKey());
      blocks.add(testObj);
    }
    bundle.put("blocksClient", blocks);
  }
}
