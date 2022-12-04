package main

import "fmt"


func main() {
  
  var test any = 1

  var test_json map[any]any
  test_json = map[any]any{}
  // var test_json = map[any]any{1: "test", "other": 0} // you can do through this way too
  test_json[1] = "test"
  test_json["other"] = 0
  // test_json = {1: "test", "other": 0} but sadly you can't do this

  fmt.Println("test:", test)
  fmt.Println("test_json:", test_json)

}

