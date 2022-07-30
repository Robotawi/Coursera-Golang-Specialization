package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter a string: ")

	input_reader := bufio.NewReader(os.Stdin)
	input_line, err := input_reader.ReadString('\n')
	if err != nil {
		fmt.Print("The provided input is wrong!\n")
		os.Exit(1) //exit with an error state 1
	}
	lower_input := strings.ToLower(input_line)

	if lower_input[0] == 'i' && lower_input[len(lower_input)-2] == 'n' && strings.Contains(lower_input, "a") {
		fmt.Print("Found!\n")
	} else {
		fmt.Print("Not Found!\n")
	}

}
