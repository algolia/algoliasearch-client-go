package main

import "fmt"

func main() {
    var client = NewClient("", "")
    var body = client.listIndexes()
    fmt.Printf("%s", body)
 }
