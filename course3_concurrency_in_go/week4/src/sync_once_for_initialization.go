package main

import (
	"fmt"
	"sync"
)

var on sync.Once //call a function only once from inside goroutine/s

// this is the function we need to call only once
func initialize() {
	fmt.Println("Initialization step done!")

}

func worker(wg *sync.WaitGroup) {
	on.Do(initialize) //this is how sync.Once is used

	fmt.Println("Doing work!")
	wg.Done()
}
func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go worker(&wg)
	go worker(&wg)

	wg.Wait()

	fmt.Println("Finished working at main!")
}
