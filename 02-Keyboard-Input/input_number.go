package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Give a number: ")
	n, _ := reader.ReadString('\n')
	n = strings.TrimSuffix(n, "\n")
	res, _ := strconv.Atoi(n)
	if res > 10 {
		fmt.Println(n, "is larger than 10")
	} else {
		fmt.Println(n, "is smaller than 10")
	}
}