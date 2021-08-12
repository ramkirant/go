package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) toString() string {
	return p.firstName + " " + p.lastName
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	/*deadpool := person{"Ryan", "Renolds"}
	fmt.Println(deadpool)*/

	var captainAmerica person
	captainAmerica.firstName = "Steve"
	captainAmerica.lastName = "Rodgers"

	ironMan := person{
		firstName: "Tony",
		lastName:  "Stark",
		contactInfo: contactInfo{
			email:   "tonystark@starkindustries.com",
			zipCode: 10010,
		},
	}
	fmt.Println(ironMan.toString())

	/*
	   A variable can be assigned to a receiver which accepts a pointer.
	   Go implicitely sends the pointer to the variable to the receiver.
	*/
	ironMan.updateName("Toni")
	fmt.Printf("%+v", ironMan)
}
