package main

import "net/http"
import "net/url"
import "io/ioutil"
import "bytes"
import "log"
import "encoding/json"

type Transport struct {
  httpClient *http.Client
  appID string
  apiKey string
  host [3]string
}

func NewTransport(appID, apiKey string) *Transport {
  transport := new(Transport)
  var tr = &http.Transport{}
  transport.appID = appID
  transport.apiKey = apiKey
  transport.httpClient = &http.Client{Transport: tr}
  transport.host = [3]string{"https://" + appID + "-1.algolia.io", "https://" + appID + "-2.algolia.io", "https://" + appID + "-3.algolia.io", }
    //TODO Suffle
  return transport
}

func (t *Transport) urlEncode(value string) string {
  v := url.Values{}
  v.Add("", value)
  encoded := v.Encode()
  return encoded[1:len(encoded)] 
}

func (t *Transport) request(method, path string, body interface{}) interface{}{
  bodyBytes, err := json.Marshal(body)
  if err != nil {
    log.Fatal(err)
  }
  reader := bytes.NewReader(bodyBytes)
  req, err := http.NewRequest(method, t.host[0] + path, reader)
  if err != nil {
    log.Fatal(err)
  }
  req.Header.Add("X-Algolia-API-Key", t.apiKey)
  req.Header.Add("X-Algolia-Application-Id", t.appID)
  resp, err := t.httpClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  res, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  if resp.StatusCode >= 300 {
    log.Fatal(req.URL)
    log.Fatal(resp.Status)
  }
  var jsonResp interface{}
  err = json.Unmarshal(res, &jsonResp)
  if err != nil {
    log.Fatal(err)
  }
  return jsonResp
}


