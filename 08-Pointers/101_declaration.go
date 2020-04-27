package main

import "fmt"

func main() {
	var i int = 110

	// 1. Using var
	var p *int = &i
	
	// 2. Using new
	pp := new(int)
	pp = &i

	// 3. direct assignment
	ppp := &i

	fmt.Println(*p, *pp, *ppp)
}