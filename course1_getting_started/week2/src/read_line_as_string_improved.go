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
	input_line, _ := input_reader.ReadString('\n')

	lower_input := strings.ToLower(input_line)

	if lower_input[0] == 'i' && lower_input[len(lower_input)-2] == 'n' && strings.Contains(lower_input, "a") {
		fmt.Print("Found!\n")
	} else {
		fmt.Print("Not Found!\n")
	}

}
