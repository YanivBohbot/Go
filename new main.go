package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {

	jim := person{
		firstname: "jim",
		lastname:  "party",
		contact: contactInfo{
			email:   "jgmail.com",
			zipcode: 9848499,
		},
	}
	var alex person
	alex.firstname = "alex"
	alex.lastname = "anderson"
	fmt.Println(alex)

	fmt.Print("%+v", jim)

}
