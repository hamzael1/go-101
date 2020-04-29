package main

import "fmt"

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	c1 := closure() // this is now function with its own i
	fmt.Println("c1:", c1())
	fmt.Println("c1:",c1())
	fmt.Println("c1:",c1())

	c2 := closure() // this is another function with its own i independent of the last one
	fmt.Println("c2:",c2())
}