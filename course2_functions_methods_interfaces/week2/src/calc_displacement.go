package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var acc float64
	acc = 10
	var v_0 float64
	v_0 = 2.0
	var s_0 float64
	s_0 = 1.0
	var t float64
	t = 0

	fmt.Print("Enter the value of acceleration: ")
	reader := bufio.NewReader(os.Stdin)
	in_string, _ := reader.ReadString('\n')
	acc, _ = strconv.ParseFloat(in_string[0:len(in_string)-1], 64)
	// fmt.Println(acc)

	fmt.Print("Enter the value of initial velocity: ")
	in_string, _ = reader.ReadString('\n')
	v_0, _ = strconv.ParseFloat(in_string[0:len(in_string)-1], 64)
	// fmt.Println(v_0)

	fmt.Print("Enter the value of initial displacement: ")
	in_string, _ = reader.ReadString('\n')
	s_0, _ = strconv.ParseFloat(in_string[0:len(in_string)-1], 64)
	// fmt.Println(s_0)

	fmt.Print("Enter the time to calculate displacement at: ")
	in_string, _ = reader.ReadString('\n')
	t, _ = strconv.ParseFloat(in_string[0:len(in_string)-1], 64)
	// fmt.Println(t)

	fn := GenDisplaceFn(acc, v_0, s_0)

	fmt.Println("The displacement is : ", fn(t))

}

func GenDisplaceFn(acc, v_0, s_0 float64) func(time float64) float64 {
	return func(time float64) float64 { return 0.5*acc*math.Pow(time, 2) + time*v_0 + s_0 }
}
