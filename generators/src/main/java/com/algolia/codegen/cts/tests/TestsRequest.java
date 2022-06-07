package com.algolia.codegen.cts.tests;

import com.algolia.codegen.Utils;
import com.algolia.codegen.exceptions.CTSException;
import java.io.File;
import java.util.*;
import org.openapitools.codegen.CodegenModel;
import org.openapitools.codegen.CodegenOperation;
import org.openapitools.codegen.SupportingFile;

public class TestsRequest extends TestsGenerator {

  public TestsRequest(String language, String client) {
    super(language, client);
  }

  protected Map<String, Request[]> loadRequestCTS() throws Exception {
    String clientName = client;
    // This special case allow us to read the `search` CTS to generated the tests for the
    // `algoliasearch-lite` client, which is only available in JavaScript
    if (language.equals("javascript") && client.equals("algoliasearch-lite")) {
      clientName = "search";
    }
    return super.loadCTS("methods/requests", clientName, Request[].class);
  }

  @Override
  public boolean available() {
    File templates = new File("templates/" + language + "/tests/requests/requests.mustache");
    return templates.exists();
  }

  @Override
  public void addSupportingFiles(List<SupportingFile> supportingFiles, String outputFolder, String extension) {
    if (!available()) {
      return;
    }
    String clientName = language.equals("php") ? Utils.createClientName(client, language) : client;
    supportingFiles.add(new SupportingFile("requests/requests.mustache", outputFolder + "/methods/requests", clientName + extension));
  }

  @Override
  public void run(Map<String, CodegenModel> models, Map<String, CodegenOperation> operations, Map<String, Object> bundle) throws Exception {
    Map<String, Request[]> cts = loadRequestCTS();

    List<Object> blocks = new ArrayList<>();
    ParametersWithDataType paramsType = new ParametersWithDataType(models, language);

    for (Map.Entry<String, CodegenOperation> entry : operations.entrySet()) {
      String operationId = entry.getKey();
      if (!cts.containsKey(operationId)) {
        throw new CTSException("operationId " + operationId + " does not exist in the spec");
      }
      Request[] op = cts.get(operationId);

      List<Object> tests = new ArrayList<>();
      for (int i = 0; i < op.length; i++) {
        Map<String, Object> test = new HashMap<>();
        Request req = op[i];
        test.put("method", operationId);
        test.put("testName", req.testName == null ? operationId : req.testName);
        test.put("testIndex", i);
        test.put("request", req.request);
        test.put("hasParameters", req.parameters.size() != 0);

        if (req.requestOptions != null) {
          test.put("hasRequestOptions", true);
          Map<String, Object> requestOptions = new HashMap<>();
          if (req.requestOptions.queryParameters != null) {
            Map<String, Object> queryParameters = new HashMap<>();
            paramsType.enhanceParameters(req.requestOptions.queryParameters, queryParameters);
            requestOptions.put("queryParameters", queryParameters);
          }
          if (req.requestOptions.headers != null) {
            Map<String, Object> headers = new HashMap<>();
            // convert the headers to an acceptable type
            paramsType.enhanceParameters(new HashMap<String, Object>(req.requestOptions.headers), headers);
            requestOptions.put("headers", headers);
          }
          test.put("requestOptions", requestOptions);
        }

        CodegenOperation ope = entry.getValue();
        paramsType.enhanceParameters(req.parameters, test, ope);
        tests.add(test);
      }
      Map<String, Object> testObj = new HashMap<>();
      testObj.put("tests", tests);
      testObj.put("operationId", operationId);
      blocks.add(testObj);
    }
    bundle.put("blocksRequests", blocks);
  }
}
