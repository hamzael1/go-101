package main

import "fmt"

func main() {
	myslice := []int {0,1,2,3,4,5,6,7,8,9,10}
	fmt.Println(myslice)
	fmt.Println(len(myslice))

	myslice = append(myslice, 11)
	fmt.Println(myslice)
	fmt.Println(len(myslice))

	// Another way to create a slice is to use make

	newslice := make( []int, 1, 3)
	newslice[0] = 1
	newslice[0] = 2
	fmt.Println(newslice)
}