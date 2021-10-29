package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	slicesOfInt := getSeries()
	sliceSize := int(len(slicesOfInt) / 4)
	slice1 := slicesOfInt[:sliceSize]
	slice2 := slicesOfInt[sliceSize : sliceSize*2]
	slice3 := slicesOfInt[sliceSize*2 : sliceSize*3]
	slice4 := slicesOfInt[sliceSize*3:]
	wg.Add(4)
	go sortSlice(slice1)
	go sortSlice(slice2)
	go sortSlice(slice3)
	go sortSlice(slice4)
	wg.Wait()
	sortedList := merge4Sort(slice1, slice2, slice3, slice4)
	fmt.Printf("Your sorted list is: .\n")
	fmt.Println(sortedList)

}

func getSeries() []int {
	slicesOfInt := []int{}
	fmt.Printf("Please enter a series of integers with space to saperate them.\n")
	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringOfInt := scanner.Text()
	listOfString := strings.Split(stringOfInt, " ")
	for _, i := range listOfString {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		slicesOfInt = append(slicesOfInt, j)
	}
	return slicesOfInt
}

func sortSlice(s []int) {
	defer wg.Done()
	sort.Ints(s)
}

func merge4Sort(s1, s2, s3, s4 []int) []int {
	return mergeSort(mergeSort(s1, s2), mergeSort(s3, s4))
}

func mergeSort(s1, s2 []int) []int {
	result := make([]int, len(s1)+len(s2))
	i := 0
	for len(s1) > 0 && len(s2) > 0 {
		if s1[0] < s2[0] {
			result[i] = s1[0]
			s1 = s1[1:]
		} else {
			result[i] = s2[0]
			s2 = s2[1:]
		}
		i++
	}

	for j := 0; j < len(s1); j++ {
		result[i] = s1[j]
		i++
	}
	for j := 0; j < len(s2); j++ {
		result[i] = s2[j]
		i++
	}
	return result
}
