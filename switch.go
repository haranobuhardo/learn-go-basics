package main

import ("fmt"
        "strconv"
        "os"
        "bufio"
        "strings"
      )

func main() {
  var nilai string
  var rating string

  if len(os.Args) == 1 {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Enter your score: ")
    nilai, _ = reader.ReadString('\n')
  } else {
    nilai = os.Args[1]
  }

  floatNilai, _ := strconv.ParseFloat(strings.TrimSpace(nilai), 32)
  
  fmt.Printf("Your score: %.2f\n", floatNilai)

  switch {
    case floatNilai >= 80 && floatNilai <= 100:
      rating = "Funtastic"
    case floatNilai >= 60 && floatNilai < 80:
      rating = "Good"
    case floatNilai < 60:
      rating = "Please retake your test!"
    default:
      panic("Score can't be recognize!")
  }

  fmt.Printf("Your final rating: %s\n", rating)

}
