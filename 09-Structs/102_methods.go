package main

import "fmt"

type duck struct {
	name string
}

func (a duck) speak() {
	fmt.Println(a.name, "Quack Quack !")
}

func main() {
	d := duck {
		name: "Boby",
	}
	d.speak()
}