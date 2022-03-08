package main

import "fmt"

/**
 * Contact Information
 * @type {struct} contactInformation
 */
type contactInfo struct {
	email   string
	zipCode int
}

/**
 * Person structure.
 * @type {struct} person
 */
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	/** shortcut - person{"Sarwesh","Maharjan"} */
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	/**
	 * Another way to assign struct
	 * If no value is assigned it goes to default which is empty string
	 */
	var john person
	fmt.Println(john)

	john.firstName = "Sarwesh"
	john.lastName = "Maharjan"

	/** To format print */
	john.print()

	jim := person{
		firstName: "Jim",
		lastName:  "Shrestha",
		contact: contactInfo{
			email:   "jim@example.go",
			zipCode: 94100,
		},
	}
	jim.update()
	jim.print()

	/** shortcut is that you can directly do jim.updatef(), *person pointer can call person type */
	jimPointer := &jim
	jimPointer.updatef()
	jim.print()
}

func (p person) print() {
	fmt.Printf("\n %+v", p)
}

func (pointerToPerson *person) updatef() {
	(*pointerToPerson).firstName = "Sneha"
}

func (p person) update() {
	p.firstName = "Shristi"
}
