package main

import "fmt"
import "bufio"
import ("os"
        "strconv"
        "strings"
      )

func main(){
  var left = false
  var right = true
  fmt.Printf("left is %t, right is %t\n", left, right)
  

  var leftAndRight = left && right
  fmt.Printf("left AND right \t(%t) \n", leftAndRight)
  
  var leftOrRight = left || right
  fmt.Printf("left OR right \t(%t) \n", leftOrRight)

  var leftReverse = !left
  fmt.Printf("NOT left \t(%t) \n", leftReverse)


  // Test input and aritmathic calc
  
  reader := bufio.NewReader(os.Stdin)
  
  fmt.Print("Enter the targeted x value (ditanya): ")
  x1, _ := reader.ReadString('\n')
  x1_int, _ := strconv.Atoi(strings.TrimSpace(x1))
 
  fmt.Print("Enter the known x value (diketahui): ")
  x2, _ := reader.ReadString('\n')
  x2_int, _ := strconv.Atoi(strings.TrimSpace(x2))

  fmt.Print("Enter the known y value (diketahui): ")
  y2, _ := reader.ReadString('\n')
  y2_int, _ := strconv.Atoi(strings.TrimSpace(y2)) 


  // fmt.Printf("%s %s %s", x1, x2, y2)
  result := x1_int/x2_int*y2_int

  fmt.Printf("The targetedvalue (ditanya) is %d\n", result)
}
