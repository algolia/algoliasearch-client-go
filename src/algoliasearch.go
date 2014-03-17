package main

import "fmt"
import "log"

func main() {
    var client = NewClient("", "")
    var index = client.initIndex("go")
    obj := make(map[string]interface{})
    obj["name"] = "toto"
    obj["objectID"] = "id"
    var body, err = index.addObject(obj)
    if err != nil {
      log.Fatalf(err.Error())
    }
    //var body = client.getLogs(0, 1, false)
    fmt.Printf("%s", body)
 }
