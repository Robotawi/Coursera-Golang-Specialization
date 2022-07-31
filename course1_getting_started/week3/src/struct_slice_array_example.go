package main

import (
	"fmt"
)

type Person struct {
	name    string
	address string
	phone   int
}

func main() {
	//declare a struct Person 
	//init two persons 
	p1 := Person{"X", "sky", 123}

	var p2 Person
	p2 = Person{"Y", "earth", 456} 

	person_sli := []Person{} //slice of Person struct
	person_sli = append(person_sli, p1)
	person_sli = append(person_sli, p2)

	person_arr := [...]Person{p1, p2} //array of Person struct

	for _, elem := range person_sli {
		fmt.Println("Name is ", elem.name, " has a phone number ", elem.phone)
	}

	for _, elem := range person_arr {
		fmt.Println("Name is ", elem.name, " has an address in ", elem.address)
	}	

}
