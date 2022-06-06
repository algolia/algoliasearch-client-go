package com.algolia;

import com.algolia.exceptions.*;
import com.algolia.utils.JSON;
import com.algolia.utils.Requester;
import com.algolia.utils.retry.StatefulHost;
import java.io.IOException;
import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.List;
import okhttp3.Request;
import okhttp3.Response;

public class EchoRequester implements Requester {

  private int connectionTimeout, readTimeout, writeTimeout;
  private List<StatefulHost> hosts;
  private Response lastResponse;

  public EchoRequester() {
    this.connectionTimeout = 100;
    this.readTimeout = 100;
    this.writeTimeout = 100;
    this.hosts = new ArrayList<>();
  }

  public EchoResponse getLastEchoResponse() {
    try {
      return JSON.deserialize(lastResponse.body().string(), EchoResponse.class);
    } catch (IOException e) {
      e.printStackTrace();
      return null;
    }
  }

  @Override
  public CallEcho newCall(Request request) {
    return new CallEcho(request);
  }

  @Override
  public <T> T handleResponse(Response response, Type returnType) throws AlgoliaRuntimeException {
    lastResponse = response;
    return null;
  }

  // NO-OP for now
  @Override
  public void setDebugging(boolean debugging) {}

  @Override
  public int getConnectTimeout() {
    return this.connectionTimeout;
  }

  @Override
  public void setConnectTimeout(int connectionTimeout) {
    this.connectionTimeout = connectionTimeout;
  }

  @Override
  public int getReadTimeout() {
    return this.readTimeout;
  }

  @Override
  public void setReadTimeout(int readTimeout) {
    this.readTimeout = readTimeout;
  }

  @Override
  public int getWriteTimeout() {
    return this.writeTimeout;
  }

  @Override
  public void setWriteTimeout(int writeTimeout) {
    this.writeTimeout = writeTimeout;
  }

  @Override
  public void setHosts(List<StatefulHost> hosts) {
    this.hosts = hosts;
  }
}
