package com.algolia;

import com.algolia.utils.UseReadTransporter;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;
import okhttp3.Headers;
import okhttp3.HttpUrl;
import okhttp3.Interceptor;
import okhttp3.MediaType;
import okhttp3.Protocol;
import okhttp3.Request;
import okhttp3.Response;
import okhttp3.ResponseBody;
import okio.Buffer;

public class EchoInterceptor {

  private EchoResponse lastResponse;
  private int httpCode;

  public EchoInterceptor(int httpCode) {
    this.httpCode = httpCode;
  }

  public EchoInterceptor() {
    this(200);
  }

  public Interceptor getEchoInterceptor() {
    return new Interceptor() {
      @Override
      public Response intercept(Chain chain) throws IOException {
        Request request = chain.request();
        try {
          UseReadTransporter useReadTransporter = (UseReadTransporter) request.tag();
          EchoResponse echo = new EchoResponse();
          echo.path = request.url().encodedPath();
          echo.host = request.url().host();
          echo.method = request.method();
          echo.body = processResponseBody(request);
          echo.queryParameters = buildQueryParameters(request);
          echo.headers = buildHeaders(request.headers());
          echo.connectTimeout = chain.connectTimeoutMillis();
          echo.responseTimeout =
            (useReadTransporter != null || request.method().equals("GET")) ? chain.readTimeoutMillis() : chain.writeTimeoutMillis();

          lastResponse = echo;
        } catch (Exception e) {
          e.printStackTrace();
          lastResponse = null;
        }
        Response.Builder builder = new Response.Builder();
        builder.code(httpCode);
        builder.request(request);
        builder.protocol(Protocol.HTTP_2);
        builder.message("EchoMessage");
        builder.body(ResponseBody.create("", MediaType.parse("application/json")));
        return builder.build();
      }
    };
  }

  private String processResponseBody(Request request) {
    try {
      final Request copy = request.newBuilder().build();
      final Buffer buffer = new Buffer();
      if (copy.body() == null) {
        return "";
      }
      copy.body().writeTo(buffer);
      return buffer.readUtf8();
    } catch (final IOException e) {
      return "error";
    }
  }

  private Map<String, Object> buildQueryParameters(Request request) {
    Map<String, Object> params = new HashMap<>();
    HttpUrl url = request.url();
    for (String name : url.queryParameterNames()) {
      for (String value : url.queryParameterValues(name)) {
        params.put(name, value);
      }
    }
    return params;
  }

  private Map<String, String> buildHeaders(Headers headers) {
    Map<String, String> mapHeaders = new HashMap<>();
    for (String headerName : headers.names()) {
      mapHeaders.put(headerName, headers.get(headerName));
    }
    return mapHeaders;
  }

  public EchoResponse getLastResponse() {
    return this.lastResponse;
  }

  public void setHttpCode(int httpCode) {
    this.httpCode = httpCode;
  }
}
