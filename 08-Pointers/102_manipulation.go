package main

import "fmt"

func main() {
	i := 10

	p := &i
	
	fmt.Println("Address of i", &i, "Content of p", p, "are the same")

	// Increment i using pointer p
	*p++
	fmt.Printf("i = %d\n", i)

}