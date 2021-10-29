package main

import (
	"fmt"
)

func main() {
	var acc float64
	var vel float64
	var s float64
	var t float64

	fmt.Printf("Please Enter Initial Acceleration: ")
	fmt.Scan(&acc)
	fmt.Printf("Please Enter Initial velocity: ")
	fmt.Scan(&vel)
	fmt.Printf("Please Enter Initial displacement: ")
	fmt.Scan(&s)
	fmt.Printf("Please Enter time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(acc, vel, s)
	fmt.Printf("The displacement after %f seconds is: %f\n", t, fn(t))

}

func GenDisplaceFn(acc, vel, s float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 1/2*acc*t*t + vel*t + s
	}
}
