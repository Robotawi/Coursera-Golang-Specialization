package main

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

submit by May 5
*/

import (
	"fmt"
)

func main() {
	//declaring a map and initalizing it
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["C++"] = 1
	map1["Golang"] = 2

	//declare and intitalize in one step 
	map2 := make(map[string]int)
	map2["Python"] = 1
	map2["Golang"] = 2


	//using map lateral
	map3 := map[string]int{
		"Robotics": 1,
		"AI":       2,
	}

	//printing maps
	fmt.Println(map1)

	for key, value := range map1 {
		fmt.Println(key, " ", value)
	}

	fmt.Println(map2)
	for key, value := range map2 {
		fmt.Println(key, " ", value)
	}

	fmt.Println(map3)
	for key, value := range map3 {
		fmt.Println(key, " ", value)
	}
}
