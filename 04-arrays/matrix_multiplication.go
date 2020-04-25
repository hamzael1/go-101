package main
/*
	Very basic unsophisticated (2x2) matrix multiplication
*/

import "fmt"

func print_matrix(M *[2][2]int) {
	pref := "\t"
	for _, r := range M {
		fmt.Printf("%s|", pref)
		for _, e := range r {
			fmt.Printf("  %d", e)
		}
		fmt.Printf("  |\n")
	}
}

func dot_product(a [2]int, b [2]int) int {
	return a[0]*b[0] + a[1]*b[1]
}

func multiply_matrices(M *[2][2]int, N*[2][2]int) [2][2]int {
	var res [2][2]int
	for i, r := range M {
		c := [...]int {N[0][0], N[1][0]}
		res[i][0] = dot_product(r, c)
		c = [...]int {N[0][1], N[1][1]}
		res[i][1] = dot_product(r, c)
	}
	return res
}

func main() {

	A := [2][2]int { {1,2}, {3,4} }
	B := [2][2]int { {5,6}, {7,8} }

	fmt.Printf("A = \n")
	print_matrix(&A)
	fmt.Printf("B = \n")
	print_matrix(&B)

	res := multiply_matrices(&A, &B)
	fmt.Printf("A x B = \n")
	print_matrix(&res)
}