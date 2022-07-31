package main

import (
	"encoding/json"
	"fmt"
)

//NOTE:  json package only stringfiy fields start with capital letter
type AddressBook struct{
	Name string `json:"Name"`
	Address string `json"Address"`
}

func main(){
	var name string
	var address string

	fmt.Scan(&name, &address)

	p:=AddressBook{Name:name, Address:address}
	json_obj, _ := json.Marshal(p)

	fmt.Println(string(json_obj))
}