package main

import "fmt"

func main() {
	var input float32

	fmt.Print("Enter a float number: ")
	fmt.Scan(&input)

	// fmt.Printf("After conversion to integer: %d\n", int(input))
	fmt.Print("After conversion to integer: ", int(input), "\n")
}
