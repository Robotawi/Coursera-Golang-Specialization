package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//make a struct
type Name struct {
	fname string
	lname string
}

func main() {
	//make a slice of struct
	var input_file string

	names_sli := []Name{}

	fmt.Print("Enter the input file name: ")
	fmt.Scan(&input_file)
	input_text, _ := os.Open(input_file)
	input_text_reader := bufio.NewReader(input_text)

	for {
		input_line, err := input_text_reader.ReadString('\n')

		// names are space-separated strings
		names := strings.Fields(input_line)
		extracted_name := Name{names[0], names[1]}
		names_sli = append(names_sli, extracted_name)

		// keep reading till the error is not nil
		if err != nil {
			break
		}
	}

	for n, names_elem := range names_sli {
		fmt.Println("Entry number", n+1, "First name is:", names_elem.fname, "and last name is:", names_elem.lname)
	}

}
