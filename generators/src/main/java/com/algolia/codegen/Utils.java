package com.algolia.codegen;

import com.google.common.collect.Sets;
import java.io.FileInputStream;
import java.net.URL;
import java.util.*;
import org.openapitools.codegen.CodegenOperation;
import org.yaml.snakeyaml.Yaml;

public class Utils {

  /** The suffix of our client names. */
  public static final String API_SUFFIX = "Client";

  public static final Set<String> CUSTOM_METHOD = Sets.newHashSet(
    "del",
    "get",
    "post",
    "put"
  );

  public static String capitalize(String str) {
    return str.substring(0, 1).toUpperCase() + str.substring(1);
  }

  /**
   * Will add the boolean `vendorExtensions.x-is-custom-request` to operations if they should not
   * escape '/' in the path variable
   */
  public static CodegenOperation specifyCustomRequest(CodegenOperation ope) {
    if (CUSTOM_METHOD.contains(ope.nickname)) {
      ope.vendorExtensions.put("x-is-custom-request", true);
    }
    return ope;
  }

  /** Returns the client name for the given language */
  public static String createClientName(String client, String language) {
    String[] clientParts = client.split("-");
    String clientName = "";
    if (language.equals("javascript")) {
      // do not capitalize the first part
      clientName = clientParts[0].toLowerCase();
      for (int i = 1; i < clientParts.length; i++) {
        clientName += capitalize(clientParts[i]);
      }
    } else {
      for (int i = 0; i < clientParts.length; i++) {
        clientName += capitalize(clientParts[i]);
      }
    }

    return clientName;
  }

  public static String getClientNameKebabCase(Map<String, Object> data) {
    String client = (String) ((Map<String, Object>) data.get("operations")).get(
        "pathPrefix"
      );
    return client
      .replaceAll("(.+?)([A-Z]|[0-9])", "$1-$2")
      .toLowerCase(Locale.ROOT);
  }

  public static String getClientNameCamelCase(Map<String, Object> data) {
    return (String) ((Map<String, Object>) data.get("operations")).get(
        "pathPrefix"
      );
  }

  /** Inject server info into the client to generate the right URL */
  public static void generateServer(
    String clientKebab,
    Map<String, Object> additionalProperties
  ) {
    Yaml yaml = new Yaml();
    try {
      Map<String, Object> spec = yaml.load(
        new FileInputStream("specs/bundled/" + clientKebab + ".yml")
      );
      List<Map<String, Object>> servers = (List<Map<String, Object>>) spec.get(
        "servers"
      );

      boolean hasRegionalHost = false;
      boolean fallbackToAliasHost = false;
      String host = "";
      String topLevelDomain = "";
      Set<String> allowedRegions = new HashSet<>();
      for (Map<String, Object> server : servers) {
        if (!server.containsKey("url")) {
          throw new GenerationException(
            "Invalid server, does not contains 'url'"
          );
        }

        // Determine if the current URL with `region` also have an alias without
        // variables.
        for (Map<String, Object> otherServer : servers) {
          if (server == otherServer) {
            continue;
          }
          String otherUrl = (String) otherServer.getOrDefault("url", "");
          if (otherUrl.replace(".{region}", "").equals(server.get("url"))) {
            fallbackToAliasHost = true;
            break;
          }
        }

        if (!server.containsKey("variables")) {
          continue;
        }

        Map<String, Map<String, Object>> variables = (Map<String, Map<String, Object>>) server.get(
          "variables"
        );

        if (
          !variables.containsKey("region") ||
          !variables.get("region").containsKey("enum")
        ) {
          continue;
        }
        ArrayList<String> regions = (ArrayList<String>) variables
          .get("region")
          .get("enum");
        hasRegionalHost = true;

        for (String region : regions) {
          allowedRegions.add(region);
        }

        // This is used for hosts like `insights` that uses `.io`
        URL url = new URL((String) server.get("url"));
        String[] hostParts = url.getHost().split("\\.");
        host = hostParts[0];
        topLevelDomain = hostParts[hostParts.length - 1];
      }
      additionalProperties.put("hasRegionalHost", hasRegionalHost);
      additionalProperties.put("fallbackToAliasHost", fallbackToAliasHost);
      additionalProperties.put("host", host);
      additionalProperties.put("topLevelDomain", topLevelDomain);
      additionalProperties.put(
        "allowedRegions",
        allowedRegions.toArray(new String[0])
      );

      if (clientKebab.equals("predict")) {
        additionalProperties.put("isExperimentalHost", true);
        additionalProperties.put(
          "host",
          new URL((String) servers.get(0).get("url")).getHost()
        );
      }
    } catch (Exception e) {
      e.printStackTrace();
      System.exit(1);
    }
  }
}
