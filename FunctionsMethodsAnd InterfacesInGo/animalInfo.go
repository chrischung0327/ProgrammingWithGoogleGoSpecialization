package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.sound
}

func main() {
	// var response string
	var animal string
	var action string
	cow := Animal{food: "grass", locomotion: "walk", sound: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", sound: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", sound: "hsss"}

	var animals = map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	fmt.Printf("Please Enter The Animal and One Request with one space in the middle or ctrl+c to exit: ")
	for {
		fmt.Printf("\n> ")
		fmt.Scanf("%s %s", &animal, &action)

		animalObj, animalExists := animals[animal]

		if !animalExists {
			fmt.Printf("Animal %s does not exist.\n", animal)
			continue
		}

		switch action {
		case "eat":
			fmt.Printf(animalObj.Eat())
		case "move":
			fmt.Printf(animalObj.Move())
		case "speak":
			fmt.Printf(animalObj.Speak())
		default:
			fmt.Printf("No such function!")
		}
	}
}
