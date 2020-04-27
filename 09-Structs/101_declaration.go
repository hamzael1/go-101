package main

import "fmt"

func main() {

	// Define a new type point
	type point struct {
		x int
		y int
	}

	// declare a point using var
	var p1 point
	p1.x = 1
	p1.y = 2
	fmt.Println("p1:", p1)

	// declare a point using literals
	p2 := point {
		x: 10,
		y: 50,
	}
	fmt.Println("p2", p2)

	// declare a pointer
	p3 := &point {
		x: 100,
		y: 400,
	}
	p3.x += 1
	fmt.Println("p3", p3)
}