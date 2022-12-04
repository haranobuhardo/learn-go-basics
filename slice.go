package main

import "fmt"

func main() {
  //slice is about referring to an array
  // len -> length of slice and cap -> capacity of the referred array

  var testSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // in init slice, you just have to ignore the length unlike in array
  var testSlice_1 = testSlice[:2]
  var testSlice_2 = testSlice[:3:5] // the third index indicates the max cap
  var testSlice_copy = make([]int, 4)
  copy(testSlice_copy, testSlice) // this func return len of successfully copied slice. I just knew that we can call this function with a return wihout having to store the return value

  fmt.Println("testSlice:", testSlice)
  fmt.Println("testSlice_1:", testSlice_1)
  fmt.Println("len testSlice_1:", len(testSlice_1))
  fmt.Println("cap testSlice_1:", cap(testSlice_1))

  fmt.Println("testSlice_copy:", testSlice_copy)
  
  // try to change testSlice_2 and check whether it will affect all slices referring to the same source
  testSlice_2[0] = 11
  fmt.Println("Modified testSlice_2 data")
  fmt.Println("(2) testSlice:", testSlice)
  fmt.Println("(2) testSlice_1:", testSlice_1)
  fmt.Println("(2) testSlice_2:", testSlice_2)
  
  var testSlice_3 = append(testSlice, 999)
  fmt.Println("testSlice_3 (via append):", testSlice_3)
  fmt.Println("(3) testSlice:", testSlice)
}
