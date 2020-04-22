package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What's your birth year : ")
	birth_year_str, _ := reader.ReadString('\n')
	fmt.Print("What's the current year: ")
	current_year_str, _ := reader.ReadString('\n')

	birth_year_str = strings.TrimSuffix(birth_year_str, "\n")
	birth_year, _ := strconv.Atoi(birth_year_str)
	current_year_str = strings.TrimSuffix(current_year_str, "\n")
	current_year, _ := strconv.Atoi(current_year_str)

	age := current_year - birth_year

	fmt.Printf("Your are %d years old.\n", age)
}