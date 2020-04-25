package main
/*
	Counts duplicate lines from stdio
*/

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	m := make(map[string]int)

	var t string
	for reader.Scan() {
		t = reader.Text()
		if len(t) == 0 {
			break
		}
		m[t]++
	}
	fmt.Println(m)

}