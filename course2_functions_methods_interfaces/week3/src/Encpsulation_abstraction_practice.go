package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name   string
	food   string
	sound  string
	motion string
}

func (animal *Animal) Eat() {
	println(animal.food)
}

func (animal *Animal) Speak() {
	println(animal.sound)
}

func (animal *Animal) Walk() {
	println(animal.motion)
}

func main() {
	snake := Animal{name: "snake", food: "mice", sound: "hsss", motion: "slither"}
	cow := Animal{name: "cow", food: "grass", sound: "moo", motion: "walk"}
	bird := Animal{name: "bird", food: "worms", sound: "peep", motion: "fly"}

	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("> ")
		in_string, _ := reader.ReadString('\n')

		fields := strings.Fields(in_string)

		if len(fields) != 2 {
			fmt.Println("The provided input is invalid!")
			fmt.Println("Please use the format: animal_name animal_needed_information")
			fmt.Println("Example: cow sound")
			// os.Exit(1)
			continue
		}
		var temp_animal *Animal
		if fields[0] == "snake" {
			temp_animal = &snake
		} else if fields[0] == "cow" {
			temp_animal = &cow
		} else if fields[0] == "bird" {
			temp_animal = &bird
		} else {
			fmt.Println("The animal type should be bird, cow, or snake!")
			// os.Exit(1)
			continue
		}

		if fields[1] == "eat" {
			temp_animal.Eat()
		} else if fields[1] == "move" {
			temp_animal.Walk()
		} else if fields[1] == "speak" {
			temp_animal.Speak()
		} else {
			fmt.Println("The requested animal's information type should be eat, move, or speak!")
			// os.Exit(1)
			continue
		}
	}

}
