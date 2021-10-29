package main

import (
	"fmt"
)

func main() {
	sli := make([]int, 0, 10)

	var input string
	var num int
	for i := 0; i < 10; i++ {
		fmt.Printf("Please Enter Integer(Up to 10, current: %d)to append in slice or 'X' to exit: ", len(sli))
		fmt.Scan(&input)
		if input == "X" {
			break
		} else {
			fmt.Sscan(input, &num)
		}
		sli = append(sli, num)

	}
	for i := 0; i < len(sli); i++ {
		BubbleSort(sli[:len(sli)-i])
	}
	fmt.Println(sli)
}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		if sli[i] > sli[i+1] {
			Swap(sli, i)
		}
	}
}

func Swap(sli []int, index int) {
	sli[index], sli[index+1] = sli[index+1], sli[index]
}
