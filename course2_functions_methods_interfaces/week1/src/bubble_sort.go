package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func print_slice(sli []int) {
	for i := range sli {
		fmt.Print(sli[i], " ")
	}
	fmt.Print("\n")
}

func bubble_sort(sli []int) {
	keep_looping := true
	//loop over the sli, compare and swap every two nums if they are
	for keep_looping == true {
		keep_looping = false
		for i := 0; i <= len(sli)-2; i++ {
			if sli[i] > sli[i+1] {
				//swap
				sli[i], sli[i+1] = sli[i+1], sli[i] 
				//mark if a next loop is needed
				if i >= 2 && sli[i-2] > sli[i-1] || i >= 3 && sli[i-3] > sli[i-2]{
					keep_looping = true
					fmt.Print("i is ", i, ", sli[i-2] is ", sli[i-2], " and sli[i-1] is ", sli[i-1], ". keep_looping is ",keep_looping)
				}
				
			}
			print_slice(sli)
		}
	}
}

func main() {
	int_sli := []int{}
	fmt.Print("Enter a sequence of integers: ")
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		in_string, _ := reader.ReadString('\n')
		if in_string == "\n" {
			break
		} else {
			in_int, _ := strconv.Atoi(in_string[0 : len(in_string)-1])
			int_sli = append(int_sli, in_int)
		}
	}

	print_slice(int_sli)
	// fmt.Print("Slice length is ", len(int_sli), "\n")
	bubble_sort(int_sli)
	// print_slice(int_sli)
}
