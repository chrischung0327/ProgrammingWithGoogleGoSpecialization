package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var command, animal, action string
	var reqAnimal Animal

	animalMap := make(map[string]Animal)
	animalMap["cow"] = Cow{}
	animalMap["bird"] = Bird{}
	animalMap["snake"] = Snake{}

	fmt.Printf("Please Enter The Command and Animal and One Request with one space in the middle or ctrl+c to exit: ")
	fmt.Printf("Example: `query cow eat`")
	for {
		fmt.Printf("\n> ")
		fmt.Scanf("%s %s %s", &command, &animal, &action)

		switch command {
		case "query":
			reqAnimal = animalMap[animal]
			switch action {
			case "eat":
				reqAnimal.Eat()
			case "move":
				reqAnimal.Move()
			case "speak":
				reqAnimal.Speak()
			default:
				fmt.Printf("No such function!")
			}
		case "newanimal":
			animalMap[animal] = animalMap[action]
			fmt.Printf("Created it!\n")
		default:
			fmt.Printf("No such command!")
		}
	}
}
