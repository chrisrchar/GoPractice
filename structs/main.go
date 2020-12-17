package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	// Go can interpret type from field name
	contactInfo
}

func main() {
	// Ways to declare structs

	// Go Interpreted Assignments
	// ---
	// alex := person{"Alex", "Anderson"}

	// Field Defined Assignments
	// ---
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }

	// Declare type only
	// ---
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

// Structs are pass by value so need to use pointers in functions
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
