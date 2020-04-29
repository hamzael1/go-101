package main


import "fmt"

// Any type that implements bark is considered a dog
type dog interface {
	bark()
}

type labrador struct {
	name string
}

type bulldog struct {
	name string
}

func (_ labrador) bark() {
	fmt.Println("Haw Haw I'm a Labrador barking")
}

func (_ bulldog) bark() {
	fmt.Println("Haw Haw I'm a Bulldog barking")
}

func func_that_accepts_any_dog(d dog) {
	d.bark()
}

func main() {
	l := labrador {"Laby"}
	b := bulldog {"Bully"}
	func_that_accepts_any_dog(l)
	func_that_accepts_any_dog(b)
}