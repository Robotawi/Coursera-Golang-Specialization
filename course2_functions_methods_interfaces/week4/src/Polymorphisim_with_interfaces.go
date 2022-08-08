package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	Name() string
}

type Cow struct {
	name  string
	food  string
	sound string
}

type Snake struct {
	name  string
	food  string
	sound string
}

type Bird struct {
	name  string
	food  string
	sound string
}

func (c Cow) Eat()         { fmt.Println(" eats grass") }
func (c Cow) Move()        { fmt.Println(" walks") }
func (c Cow) Speak()       { fmt.Println(" makes moo sound") }
func (c Cow) Name() string { return c.name }

func (b Bird) Eat()         { fmt.Println(" eats worms") }
func (b Bird) Move()        { fmt.Println(" flies") }
func (b Bird) Speak()       { fmt.Println(" makes peep sound") }
func (b Bird) Name() string { return b.name }

func (s Snake) Eat()         { fmt.Println(" eats mice") }
func (s Snake) Move()        { fmt.Println(" slithers") }
func (s Snake) Speak()       { fmt.Println(" makes hsss sound") }
func (s Snake) Name() string { return s.name }

func Eat(animal Animal) {
	cow, ok := animal.(Cow)
	if ok {
		Eat(cow)
	}

	snake, ok := animal.(Snake)
	if ok {
		Eat(snake)
	}

	bird, ok := animal.(Bird)
	if ok {
		Eat(bird)
	}
}

func Name(animal Animal) {
	cow, ok := animal.(Cow)
	if ok {
		Name(cow)
	}

	bird, ok := animal.(Bird)
	if ok {
		Name(bird)
	}

	snake, ok := animal.(Snake)
	if ok {
		Name(snake)
	}
}

func Speak(animal Animal) {
	cow, ok := animal.(Cow)
	if ok {
		Speak(cow)
	}

	bird, ok := animal.(Bird)
	if ok {
		Speak(bird)
	}

	snake, ok := animal.(Snake)
	if ok {
		Speak(snake)
	}
}

func Move(animal Animal) {
	cow, ok := animal.(Cow)
	if ok {
		Move(cow)
	}

	bird, ok := animal.(Bird)
	if ok {
		Move(bird)
	}

	snake, ok := animal.(Snake)
	if ok {
		Move(snake)
	}
}

func main() {
	sli_interface := []Animal{}

	for true {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		in_string, _ := reader.ReadString('\n')

		in_fields := strings.Fields(in_string[0 : len(in_string)-1])

		fmt.Println("length of input is ", len(in_fields))
		fmt.Println(in_fields[0], in_fields[1], in_fields[2])

		if len(in_fields) > 3 {
			fmt.Println("Invalid input format! ")
			os.Exit(1)
		}

		if in_fields[0] == "newanimal" {
			switch in_fields[2] {
			case "bird":
				var animal Animal
				var bird Bird
				bird = Bird{name: in_fields[1], food: "worms", sound: "peep"}
				animal = bird

				sli_interface = append(sli_interface, animal)
				println("Created!")

			case "cow":
				var animal Animal
				var cow Cow
				cow = Cow{name: in_fields[1], food: "worms", sound: "peep"}
				animal = cow

				sli_interface = append(sli_interface, animal)
				println("Created!")

			case "snake":
				var animal Animal
				var snake Snake
				snake = Snake{in_fields[1], "mice", "hsss"}
				animal = snake

				sli_interface = append(sli_interface, animal)
				println("Created!")

			}
		} else if in_fields[0] == "query" {
			animal_found := false
			// println("Got the following animals so far: ")
			for _, search_animal := range sli_interface {
				if search_animal.Name() == in_fields[1] {
					fmt.Print("Animal with name ", search_animal.Name())

					switch in_fields[2] {
					case "move":
						search_animal.Move()
					case "speak":
						search_animal.Speak()
					case "eat":
						search_animal.Eat()
					}

					animal_found = true
				}
			}

			if !animal_found {
				fmt.Println("The requested animal information have not been provided.")
			}
		} else {
			fmt.Println("Invalid input provided!\nThe first string must start with newanimal or query. Please try again.")
		}
	}
}
