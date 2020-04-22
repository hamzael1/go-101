package main

import "fmt"

func reverse_array(a [10]int) [10]int {
	var new_arr [10]int
	for i := 10; i > 0; i-- {
		new_arr[10-i] = a[i-1]
	}
	return new_arr
}

func main() {
	const n = 10
	var arr [n]int

	for i := 0 ; i < 10 ; i++ {
		arr[i] = i+1
	}

	fmt.Println(arr)

	reversed := reverse_array(arr)
	fmt.Println(reversed)
}