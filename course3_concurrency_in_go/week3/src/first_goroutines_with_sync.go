package main

import (
	"fmt"
	"sync"
)

// sync.WaitGroup variable should be passed by pointer to the worker/goroutine
func foo(wg *sync.WaitGroup) {
	fmt.Println("Print from goroutine")
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	fmt.Println("Print from main")
}
