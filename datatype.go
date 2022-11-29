package main

import "fmt"

// 1.Tipe data numerik non-desimal
var positiveNum uint8 = 89
var negativeNum = -12345555

// 2. Tipe data string
var message string = "Halo"

// 3. Tipe data bool
var exist bool = true

func main(){
	fmt.Println("1. Numerical non-decimal")
	fmt.Printf("positive num: %d (%T) \n", positiveNum, positiveNum)
	fmt.Printf("negative num: %d (%T) \n", negativeNum, negativeNum)

	fmt.Println()
	fmt.Println("2. String")
	fmt.Printf("message: %[1]s (%[1]T) \n", message)

	fmt.Println()
	fmt.Println("3. Bool")
	fmt.Printf("exist? %[1]t (%[1]T) \n", exist)
}
