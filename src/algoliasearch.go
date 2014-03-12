package main

import "fmt"

func main() {
    var client = NewClient("", "")
    var index = client.initIndex("go")
    obj := make(map[string]interface{})
    obj["name"] = "toto"
    obj["objectID"] = "id"
    var body = index.addObject(obj)
    //var body = client.getLogs(0, 1, false)
    fmt.Printf("%s", body)
 }
