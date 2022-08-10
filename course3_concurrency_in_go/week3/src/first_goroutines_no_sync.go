package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Print from goroutine")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Print from main")
}
