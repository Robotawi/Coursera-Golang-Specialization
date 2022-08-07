package main

import (
	"fmt"
)

type Speaker interface{ Speak() }

// this Dog struct satisfies the Speaker interface
type Dog struct{ name string }

func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("<noise>")
	} else {
		fmt.Println(d.name)
	}
}

func main() {
	var sp Speaker
	var dg *Dog
	// dynamic binding
	sp = dg
	sp.Speak()
}
