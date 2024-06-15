package main

import "fmt"

type person struct {
	fistName string
	lastName string
}

func main() {
	// go does the assignment based on the order of declaration
	// this is bad because another property may be added which could
	// change the order and break everything
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)

	// this is better :)
	mary := person{fistName: "Mary", lastName: "Anderson"}
	fmt.Println(mary)

	// as this has nothing assigned, it will print an empty struct
	// go assigns empty string ("") automatically when nothing is
	// assigned by the dev
	var john person
	fmt.Println(john)
	fmt.Printf("%+v\n", john) // %+v prints the field names too

	john.fistName = "John"
	john.lastName = "Anderson"
	fmt.Println(john)
}
