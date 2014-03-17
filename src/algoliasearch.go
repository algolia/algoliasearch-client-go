package main

import "fmt"
import "log"

func main() {
    var client = NewClient("", "")
    var index = client.InitIndex("go")
    obj := make(map[string]interface{})
    obj["name"] = "toto"
    obj["objectID"] = "id"
    var body, err = index.AddObject(obj)
    if err != nil {
      log.Fatalf(err.Error())
    }
    //var body = client.GetLogs(0, 1, false)
    fmt.Printf("%s", body)
 }
