package main

import "fmt"

//The program initializes a slice of integers and concurrently pushes 1000 values into it and prints its length.
//The proof of the data race is that the pushed numbers count and the length of the slice are not consistent.
//Also, after pushing 1000 elements,the length of the slice is less than 1000. The two coroutines are push_element and print_length. Both operate on the slice of integers

func push_element(sli_int *[]int, val int) {
	*sli_int = append(*sli_int, val)
	fmt.Println("Pushed element number ", val+1)
}

func print_length(sli_int *[]int) {
	fmt.Println("The slice length  is ", len(*sli_int))
}

func main() {

	sli_int := []int{}

	for i := 0; i < 1000; i++ {
		go push_element(&sli_int, i)
		go print_length(&sli_int)
	}

	fmt.Println("At the end, size of the slice of integers is ", len(sli_int))

}
