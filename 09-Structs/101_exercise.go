package main
/*
	Mini Interactive Program to create structs and append them to slice
*/
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)


type Address struct {
	city string
	country string
}

type Person struct {
	name string
	age int
	addresses []Address
}


var reader = bufio.NewReader(os.Stdin)
func read_input(question string) string {
	fmt.Printf(question)
	in, _ := reader.ReadString('\n')
	return strings.TrimSuffix(in, "\n")
}

func main() {

	people := []Person{}
	var answer string
	
	for {
		answer = read_input("Do you want to add a person ? [y/N]: ")
		if answer == "N" {
			fmt.Println("Thanks for playing! Bye")
			break
		} else if answer == "y" {
			var p Person
			fmt.Println(cap(p.addresses))
			var nbr_addresses int
			p.name = read_input("Name of the person: ")
			for {
				v, err := strconv.Atoi(read_input("Age of the person: "))
				if err == nil {
					p.age = v
					break
				}
			}
			for {
				v, err := strconv.Atoi(read_input("How many addresses : "))
				if err == nil {
					nbr_addresses = v
					break
				}
			}
			for i:= 0; i < nbr_addresses; i++ {
				var a Address
				a.country = read_input("\tCountry: ")
				a.city = read_input("\tCity: ")
				p.addresses = append(p.addresses, a)
			fmt.Println(cap(p.addresses))
			}
			people = append(people, p)
		}
	}

	fmt.Println(people)

}