package main

import "fmt"

func main () {
	// Using make
	m := make(map[int]string)
	m[1] = "A"
	m[2] = "B"
	m[3] = "C"
	fmt.Println(m)

	// Using literals
	mm := map[int]string{}
	mm[1] = "A"
	mm[2] = "B"
	mm[3] = "C"
	fmt.Println(mm)
	
	// This does not work !
	var mmm map[int]string = {1: "A", 2: "B", 3: "C"}
	fmt.Println(mmm)

}