package main

import "fmt"

func main() {
	m := map[int]rune {1: 'A', 2: 'B', 3: 'C'}

	// 1. Check key exists
	v, exists := m[4]
	fmt.Println(v, exists) // exists is boolean

	if v, exists := m[3]; exists == true {
		fmt.Println(string(v))
	}

	// 2. Iterating
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 3. Remove item
	fmt.Println(m, len(m))
	delete(m, 1)
	fmt.Println(m, len(m))
}