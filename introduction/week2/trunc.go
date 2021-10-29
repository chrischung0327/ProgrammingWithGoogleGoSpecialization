package main

import (
	"fmt"
)

var number float32

func main() {
	fmt.Printf("Please Enter a valid float number: ")

	num, _ := fmt.Scan(&number)

	fmt.Printf("After truncating is: %v\n", num)
}
