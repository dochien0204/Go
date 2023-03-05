package main

import (
	"fmt"
)

func main() {
	emptyPerson := NewPerson()
	fmt.Printf("Empty Person: %+v\n", emptyPerson)
	personWithFields := NewPersonWithFields("Do Chien", 20, 172)
	fmt.Printf("Person With Fields: %+v\n", personWithFields)

	//using Builders Pattern (easy to change)
	personWithBuilders := NewPerson().withHeight(180).withAge(20).withName("Doan xem").withHeight(190)
	fmt.Printf("Person using Builders Pattern: %+v\n", personWithBuilders)
}

type person struct {
	name   string
	age    int
	height float64
}

//Empty Constructor
func NewPerson() person {
	return person{}
}

//person with fields
func NewPersonWithFields(name string, age int, height float64) person {
	return person{
		name:   name,
		age:    age,
		height: height,
	}
}

//Builders start here
func (p person) withName(name string) person {
	p.name = name
	return p
}

func (p person) withAge(age int) person {
	p.age = age
	return p
}

func (p person) withHeight(height float64) person {
	p.height = height
	return p
}
