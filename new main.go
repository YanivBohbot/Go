package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {

	jim := person{
		firstname: "jim",
		lastname:  "party",
		contactInfo: contactInfo{
			email:   "jgmail.com",
			zipcode: 9848499,
		},
	}
	var alex person
	alex.firstname = "alex"
	alex.lastname = "anderson"
	fmt.Println(alex)

	fmt.Print("%+v", jim)

	jimPointer := &jim
	jimPointer.updateName("jimmmmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {

	(*pointerToPerson).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
