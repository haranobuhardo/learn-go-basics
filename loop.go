package main

import "fmt"

func main() {
  var intNumber = 123
  var strTest = "Hardo"

  for i:=0; i < 5; i++ {
    fmt.Printf("intNumber: %d\n", intNumber)
  }

  for pos, char := range strTest {
    fmt.Printf("char: %c at pos: %d\n", char, pos)
  }

  fmt.Println("### Looping without condition")
  
  var i1 int32 = 0
  for {
    fmt.Println("Printed from loop with inside condition")
    i1++
    if i1 == 5 {
      break
    } else {
      continue
    }
  }

}
