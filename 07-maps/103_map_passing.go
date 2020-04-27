package main

import "fmt"

func add_item(m map[int]string) {
	m[4] = "D"
}

func main() {
	m := map[int]string{1: "A", 2: "B", 3: "C"}

	fmt.Println(m, len(m))
	add_item(m)
	fmt.Println(m, len(m))
}