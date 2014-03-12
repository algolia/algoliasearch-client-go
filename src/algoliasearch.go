package main

import "fmt"

func main() {
    var client = NewClient("", "")
    var body = client.getLogs(0, 1, false)
    fmt.Printf("%s", body)
 }
