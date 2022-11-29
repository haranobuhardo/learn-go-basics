package main

import ("fmt"
        "os"
        "strconv"
      )

func main(){
  if len(os.Args) == 1 {panic("You forgot to put nilai as argument!")}
  nilai := os.Args[1]
  intNilai, err := strconv.Atoi(nilai)

  if err != nil {
    panic(err)
  }

  var lulus bool = false

  if intNilai > 10 {
    panic("Tidak mungkin nilai lebih dari 10!")
  } else if intNilai <= 10 && intNilai >= 7 {
    lulus = true
  } else if intNilai < 7 {
    lulus = false
  }

  fmt.Printf("status kelulusan: %t\n", lulus)
}
