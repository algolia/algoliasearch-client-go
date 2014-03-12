package main

import "testing"

func TestInd(t *testing.T) {
  var name = "test "
  var client = NewClient("", "")
  var index = client.initIndex(name)
  if index.name != "test " {
    t.Errorf("Ind: \"%s\" expected \"%s\"", index.name, name)
  }
  if index.nameEncoded != "test+" {
    t.Errorf("Ind: \"%s\" expected \"%s\"", index.name, name)
  }
}
