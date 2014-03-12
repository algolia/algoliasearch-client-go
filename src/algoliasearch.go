package main

import "fmt"

type Index struct {
  name string
  nameEncoded string
  client *Client
}

func main() {
    var client = NewClient("", "")
    var body = client.listIndexes()
    fmt.Printf("%s", body)
 }
