package main

import "fmt"

func double(n *int) {
	*n *= 2
}

func main() {
	i := 10
	p := &i

	fmt.Println("i = ", i)
	double(p)
	fmt.Println("i = ", i)
}