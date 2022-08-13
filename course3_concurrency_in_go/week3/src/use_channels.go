package main

import (
	"fmt"
)

func prod(x int, y int, c chan int) {
	c <- x * y
}

func main() {
	ch := make(chan int)

	go prod(1, 2, ch)
	go prod(3, 4, ch)

	a := <-ch
	b := <-ch

	fmt.Println("Product is ", a*b)

}
