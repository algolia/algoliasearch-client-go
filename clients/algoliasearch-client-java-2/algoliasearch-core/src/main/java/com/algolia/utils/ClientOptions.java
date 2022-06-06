package com.algolia.utils;

import com.algolia.utils.retry.StatefulHost;
import java.util.List;

public class ClientOptions {

  private Requester requester;
  private AlgoliaAgent.Segment[] algoliaAgentSegments;
  private List<StatefulHost> hosts;

  private ClientOptions() {}

  public static ClientOptions build() {
    return new ClientOptions();
  }

  public Requester getRequester() {
    return this.requester;
  }

  public ClientOptions setRequester(Requester requester) {
    this.requester = requester;
    return this;
  }

  public AlgoliaAgent.Segment[] getAlgoliaAgentSegments() {
    return this.algoliaAgentSegments;
  }

  public ClientOptions setAlgoliaAgentSegments(AlgoliaAgent.Segment[] algoliaAgentSegments) {
    this.algoliaAgentSegments = algoliaAgentSegments;
    return this;
  }

  public List<StatefulHost> getHosts() {
    return this.hosts;
  }

  public ClientOptions setHosts(List<StatefulHost> hosts) {
    this.hosts = hosts;
    return this;
  }
}
