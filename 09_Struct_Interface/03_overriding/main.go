package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

// Add function to "object" person
func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

// Add function for "object" doublZero - Note. It is not overriding
func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()

}
