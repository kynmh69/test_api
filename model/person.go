package model

import "fmt"

type Person struct {
	firstName string
	lastName  string
	password  string
	age       int
}

func New(firstName string, lastName string, age int) *Person {
	person := new(Person)
	person.firstName = firstName
	person.lastName = lastName
	person.age = age
	person.password = ""
	return person
}

func (p *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) Greetings() {
	fmt.Printf("Hello. %s\n", p.GetFullName())
}
