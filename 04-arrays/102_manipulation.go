package main

import "fmt"

func main() {

	a1 := [...]int {1,2,3}
	fmt.Println(a1[1])

	// Assignment is done by copying elements ( by value )
	a2 := a1
	fmt.Println(a2)
	a2[0] = 99
	fmt.Println(a1)
	fmt.Println(a2)

}