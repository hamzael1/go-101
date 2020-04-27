package main

import "fmt"


type Address struct {
	city string
	country string
}

type Person struct {
	name string
	age int
	addresses []Address
}

func main() {

	p := Person {
		name: "Hamza",
		age: 99,
		addresses: []Address{
			{"NY", "USA",},
			{"Helsinki", "Finland"},
		},
	}
	fmt.Println(p)

}