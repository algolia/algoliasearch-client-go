package com.algolia;

import com.algolia.utils.JSON;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;
import okhttp3.Call;
import okhttp3.Callback;
import okhttp3.Headers;
import okhttp3.HttpUrl;
import okhttp3.MediaType;
import okhttp3.Protocol;
import okhttp3.Request;
import okhttp3.Response;
import okhttp3.ResponseBody;
import okio.Buffer;
import okio.Timeout;

public class CallEcho implements Call {

  private final Request request;

  public CallEcho(Request request) {
    this.request = request;
  }

  @Override
  public Request request() {
    return request;
  }

  @Override
  public void cancel() {}

  @Override
  public Call clone() {
    return null;
  }

  private String processResponseBody() {
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

  private Map<String, String> buildQueryParameters() {
    Map<String, String> params = new HashMap<>();
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

  @Override
  public void enqueue(Callback callback) {
    Response.Builder builder = new Response.Builder();
    builder.code(200);
    builder.request(request);
    builder.protocol(Protocol.HTTP_2);
    builder.message("EchoRequest");
    try {
      EchoResponse body = new EchoResponse();
      body.path = request.url().encodedPath();
      body.method = request.method();
      body.body = processResponseBody();
      body.queryParameters = buildQueryParameters();
      body.headers = buildHeaders(request.headers());

      builder.body(
        ResponseBody.create(
          JSON.serialize(body),
          MediaType.parse("application/json")
        )
      );
      callback.onResponse(this, builder.build());
    } catch (Exception e) {
      e.printStackTrace();
    }
  }

  @Override
  public Response execute() throws IOException {
    return null;
  }

  @Override
  public boolean isExecuted() {
    return false;
  }

  @Override
  public boolean isCanceled() {
    return false;
  }

  @Override
  public Timeout timeout() {
    return null;
  }
}
