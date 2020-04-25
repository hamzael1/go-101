package main

import "fmt"

func main() {
	a1 := [3]int {1,2,3}
	fmt.Println(a1)

	// We let go calculate size based on provided elements
	a2 := [...]int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(a2)

	// We initialize some elements
	a3 := [5]int {0: 1, 4: 5}
	fmt.Println(a3)
}