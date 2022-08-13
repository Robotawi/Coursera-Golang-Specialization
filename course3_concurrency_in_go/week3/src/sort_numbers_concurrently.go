package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sort_slice(sli_int []int, wg *sync.WaitGroup) {
	sort.Slice(sli_int, func(i, j int) bool { return sli_int[i] < sli_int[j] })
	fmt.Println("Sorted sub slice", sli_int)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please give a string of unique integers: ")

	in_string, _ := reader.ReadString('\n')
	// fmt.Println("got input", in_string)

	fields := strings.Fields(in_string[0 : len(in_string)-1])

	sli_int := []int{}

	for _, val := range fields {
		val_int, _ := strconv.Atoi(val)
		sli_int = append(sli_int, val_int)
	}

	quarter_size := len(sli_int) / 4

	wg.Add(4)
	sub_slia := sli_int[:quarter_size]
	go sort_slice(sub_slia, &wg)

	sub_slib := sli_int[quarter_size : 2*quarter_size]
	go sort_slice(sub_slib, &wg)

	sub_slic := sli_int[2*quarter_size : 3*quarter_size]
	go sort_slice(sub_slic, &wg)

	sub_slid := sli_int[3*quarter_size:]
	go sort_slice(sub_slid, &wg)

	wg.Wait()
	final_sli := []int{}

	final_sli = append(final_sli, sub_slia...)
	final_sli = append(final_sli, sub_slib...)
	final_sli = append(final_sli, sub_slic...)
	final_sli = append(final_sli, sub_slid...)

	sort.Ints(final_sli) //implement merge sort instead
	fmt.Println("\nFinally sorted slice ", final_sli)
}
