package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PrintSlice(sli []int) {
	for i := range sli {
		fmt.Print(sli[i], " ")
	}
	fmt.Print("\n")
}

func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]

}

func BubbleSort(sli []int) {

	for i := 0; i < len(sli); i++ {
		for j := 0; j < len(sli)-i-1; j++ {

			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
			// PrintSlice(sli)
		}
	}
}

func main() {
	int_sli := []int{}
	fmt.Print("* I do bubble sort! Bubble sort is called that way because at the end of each iteration the largest value bubbles to the end.\n* This program sorts a sequence of up to 10 integers. \n\n")
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Print("Write a number to add to sequence or press enter if all inputs are given: ")
		in_string, _ := reader.ReadString('\n')
		if in_string == "\n" {
			break
		} else {
			in_int, _ := strconv.Atoi(in_string[0 : len(in_string)-1])
			int_sli = append(int_sli, in_int)
		}
	}

	BubbleSort(int_sli)
	PrintSlice(int_sli)
}
