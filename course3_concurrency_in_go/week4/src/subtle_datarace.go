package main

import (
	"fmt"
	"sync"
)

var i int = 0
var wg sync.WaitGroup

func increment() {
	i = i + 1
	wg.Done()
}

func main() {

	// A subtle data race rarely happens when the two threads read the 0 value of i.
	// It does show up if the number to repeat the two increments reach 100000 times
	for range [10000]int{} {
		i = 0
		wg.Add(2)
		go increment()
		go increment()
		wg.Wait()
		if i != 2 {
			fmt.Println(i)
		}
	}
}
