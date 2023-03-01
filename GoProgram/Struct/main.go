package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type personWithContactInfo struct {
	firstName string
	lastName  string

	//embedded struct
	contact contactInfo
}

type person struct {
	firstName string
	lastName  string
}

func main() {

	//declaration
	chien := person{"Chien", "Do"}
	fmt.Println(chien)

	//another way to declare
	anotherChien := person{firstName: "Chien", lastName: "Do Another Way"} // recommend used
	fmt.Println(anotherChien)

	//if not pass any value for newPerson, its properties are the zero value
	var newPerson person
	fmt.Println(newPerson)

	//update Struct value
	chien.firstName = "Chien update"
	chien.lastName = "Do update"

	//declaration struct have embedded struct
	personWithContact := personWithContactInfo{
		firstName: "Chien",
		lastName:  "Do",
		contact: contactInfo{
			email:   "doxuanchienh02042002@gmail.com",
			zipcode: 10000,
		},
	}

	//after update
	fmt.Println(chien)

	fmt.Println("\n Person with contact info")
	fmt.Printf("%+v", personWithContact)

}
