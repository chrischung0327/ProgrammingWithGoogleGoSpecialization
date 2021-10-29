package main

import (
	"fmt"
)

var x int

func funcA() {
	fmt.Printf("This is A: %d\n", x)
	x = x + 1
}

func funcB() {
	fmt.Printf("This is B: %d\n", x)
	x = x + 1
}

func main() {
	x = 1
	go funcA()
	go funcB()
	fmt.Printf("Final result: %d\n", x)
}

/* Output "go run -race raceCondition.go"
Final result: 1
This is A: 1
This is B: 1
==================
WARNING: DATA RACE
Read at 0x0001005395c0 by goroutine 8:
  main.funcB()
      /Users/chungchris/project/learning/golang/courses/concurrency_in_Go/raceCondition.go:16 +0xc8

Previous write at 0x0001005395c0 by goroutine 7:
  main.funcA()
      /Users/chungchris/project/learning/golang/courses/concurrency_in_Go/raceCondition.go:11 +0xe0

Goroutine 8 (running) created at:
  main.main()
      /Users/chungchris/project/learning/golang/courses/concurrency_in_Go/raceCondition.go:22 +0x6c

Goroutine 7 (finished) created at:
  main.main()
      /Users/chungchris/project/learning/golang/courses/concurrency_in_Go/raceCondition.go:21 +0x58
==================
Found 1 data race(s)
exit status 66
*/
