package main

import (
	"fmt"
	"sort"
)

func main() {
	sli := make([]int, 0, 3)

	var input string
	var num int
	for {
		fmt.Printf("Current length in list: %d, capacity: %d, item: %v\n", len(sli), cap(sli), sli)
		fmt.Printf("Please Enter Integer to append or 'X' to exit: ")
		fmt.Scan(&input)
		if input == "X" {
			break
		} else {
			fmt.Sscan(input, &num)
		}
		sli = append(sli, num)
		sort.Ints(sli)

	}

}
