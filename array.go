package main

import "fmt"

func main() {

  var numbers = [...]int{1, 2, 3, 4} // you can replace the (...) with your desired array length. using ... is good if you want to change the array length along development process

  for _, i := range numbers { // you can delete "_, if you want to
    fmt.Println("number: ", i)
  }

  var fruits = make([]string, 2) // this is a way to create a blank array. you can fill it along the way

  fmt.Printf("Fruits before right after created: %v\n", fruits) 
  fruits[0] = "orange"
  fruits[1] = "tomato"
  fmt.Printf("Fruits after filled: %v\n", fruits)
}
