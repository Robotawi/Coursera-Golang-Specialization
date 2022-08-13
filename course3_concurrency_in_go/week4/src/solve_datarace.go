package main

import (
	"fmt"
	"sync"
)

var i int = 0
var mtx sync.Mutex
var wg sync.WaitGroup

// the data race is now prevented with a mutex
func increment() {
	mtx.Lock()
	i = i + 1
	mtx.Unlock()
	wg.Done()
}

func main() {

	// A subtle data race rarely happens when the two threads read the 0 value of i.
	// It does show up if the number to repeat the two increments reach 100000 times
	for range [100000]int{} {
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
