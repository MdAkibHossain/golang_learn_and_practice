package main

import "fmt"

type contactinfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

func main() {
	//name := person{"Akib", "hossain"}
	// var name person
	// name.firstName = "Akib"
	// name.lastName = "Hossain"
	akib := person{firstName: "Akib", lastName: "Hossain", contact: contactinfo{email: "akib@gmail.com", zipCode: 4282}}

	//............***
	// ak:= &akib
	// ak.updateName("Akibi")
	// or
	akib.updateName("Akibi")
	//...........****

	akib.print()

}
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName

}
func (p person) print() {
	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
	fmt.Printf("%+v", p)
}
