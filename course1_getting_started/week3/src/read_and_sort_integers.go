package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//init slice of integers
	//loop to: ask user for inputs.
	//if X, then quit the program, else append in the int slice and sort it

	var input string
	int_sli := []int{}

	for {
		fmt.Print("Please enter an integer: ")
		fmt.Scan(&input)

		if input == "X" {
			break
		} else {
			int_val, _ := strconv.Atoi(input)
			int_sli = append(int_sli, int_val)
			sort.Slice(int_sli, func(i, j int) bool {
				return int_sli[i] < int_sli[j]
			})

			for _, v := range int_sli {
				fmt.Print(v, " ")
			}
			fmt.Println()
		}
	}
}
