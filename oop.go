package main

import "fmt"

type base struct {
  a string
  b int
}

//method xyz
func (this base) xyz() {
  fmt.Println("a called from xyz is:", this.a)
}

//method display
func (this base) display() {
  fmt.Println("a called from display is:", this.a)
}

type derived struct {
  base // embedding
  d     int
  a     float32 //SHADOWED
}

//method display - SHADOWED
func (this derived) display() {
  fmt.Println("a called from new display is:", this.a)
}

func main() {
  var a derived = derived{base{"base-a", 10}, 20, 2.5}

  a.display() //calls Derived.display(a); 2.5

  a.base.display() //calls Derived.Base.Display(a.base); base-a

  a.xyz() //xyz is still pure from base
}
