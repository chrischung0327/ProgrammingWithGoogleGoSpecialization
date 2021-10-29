package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Printf("Please Enter a Valid String: ")
	fmt.Scan(&input)
	lower_str := strings.ToLower(input)
	if strings.Contains(lower_str, "i") && strings.Contains(lower_str, "a") && strings.Contains(lower_str, "n") {
		if strings.Index(lower_str, "i") == 0 && strings.Index(lower_str, "n") == len(lower_str)-1 {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("Not Found!")
	}
}
