package main

import "net/http"
import "net/url"
import "io/ioutil"
import "bytes"
import "log"
import "encoding/json"
import "strconv"

type Transport struct {
  httpClient *http.Client
  appID string
  apiKey string
  host [3]string
}

func NewTransport(appID, apiKey string) *Transport {
  transport := new(Transport)
  transport.appID = appID
  transport.apiKey = apiKey
  tr := &http.Transport{DisableKeepAlives: false, MaxIdleConnsPerHost: 2}
  transport.httpClient = &http.Client{Transport: tr}
  transport.host = [3]string{"https://" + appID + "-1.algolia.io", "https://" + appID + "-2.algolia.io", "https://" + appID + "-3.algolia.io", }
    //TODO Suffle
  return transport
}

func (t *Transport) urlEncode(value string) string {
  return url.QueryEscape(value)
}

func (t *Transport) request(method, path string, body interface{}) interface{}{
  var req *http.Request
  var err error
  var bodyBytes []byte
  if body != nil {
    bodyBytes, err := json.Marshal(body)
    if err != nil {
      log.Fatal(err)
    }
    reader := bytes.NewReader(bodyBytes)
    req, err = http.NewRequest(method, t.host[0] + path, reader)
    req.Header.Add("Content-Length", strconv.Itoa(len(string(bodyBytes))))
    req.Header.Add("Content-Type", "application/json; charset=utf-8")
  } else {
    req, err = http.NewRequest(method, t.host[0] + path, nil)
  }
  if err != nil {
    log.Fatal(err)
  }
  req.Header.Add("X-Algolia-API-Key", t.apiKey)
  req.Header.Add("X-Algolia-Application-Id", t.appID)
  req.Header.Add("Connection", "keep-alive") 
  req.Header.Add("User-Agent", "Algolia for go")
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
    var str string
    if body == nil {
      str = "nil"
    } else {
      str = "notnil"
    }
    log.Fatal(resp.Status + " on "+ method + ": " + req.URL.Host + req.URL.Path + "\n" + string(res) + string(bodyBytes) + "\n" + str)
  }
  var jsonResp interface{}
  err = json.Unmarshal(res, &jsonResp)
  if err != nil {
    log.Fatal(err)
  }
  return jsonResp
}


