package main
/*
	Demonstrates that when using append on sliced slice,
	the original slice sees the changes, because the underlying array
	is the same !
*/
import "fmt"

func main() {
	s := []int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(s)

	// To avoid the pitfall we specify capacity and length to be the
	new_s := s[3:7:7]
	fmt.Println(new_s)

	new_s = append(new_s, 99)
	fmt.Println("s: ", s)
	fmt.Println("new_s: ", new_s)
}