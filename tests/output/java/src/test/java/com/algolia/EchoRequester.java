package com.algolia;

import com.algolia.exceptions.*;
import com.algolia.utils.JSON;
import com.algolia.utils.Requester;
import java.io.IOException;
import java.lang.reflect.Type;
import okhttp3.Request;
import okhttp3.Response;

public class EchoRequester implements Requester {

  private int connectionTimeout, readTimeout, writeTimeout;

  private Response lastResponse;

  public EchoRequester() {
    this.connectionTimeout = 100;
    this.readTimeout = 100;
    this.writeTimeout = 100;
  }

  public CallEcho newCall(Request request) {
    return new CallEcho(request);
  }

  public <T> T handleResponse(Response response, Type returnType) throws AlgoliaRuntimeException {
    lastResponse = response;
    return null;
  }

  public EchoResponse getLastEchoResponse() {
    try {
      return JSON.deserialize(lastResponse.body().string(), EchoResponse.class);
    } catch (IOException e) {
      e.printStackTrace();
      return null;
    }
  }

  // NO-OP for now
  public void setDebugging(boolean debugging) {}

  public int getConnectTimeout() {
    return this.connectionTimeout;
  }

  public void setConnectTimeout(int connectionTimeout) {
    this.connectionTimeout = connectionTimeout;
  }

  public int getReadTimeout() {
    return this.readTimeout;
  }

  public void setReadTimeout(int readTimeout) {
    this.readTimeout = readTimeout;
  }

  public int getWriteTimeout() {
    return this.writeTimeout;
  }

  public void setWriteTimeout(int writeTimeout) {
    this.writeTimeout = writeTimeout;
  }
}
