package main

import "fmt"

type PersonAddress struct {
	line1   string
	pincode int
}

type PersonDetails struct {
	name   string
	job    string
	salary int
}

type PersonContact struct {
	phone int
	email string
}

type EmployeeDetails struct {
	person_details PersonDetails
	person_address PersonAddress
	person_contact PersonContact
}

func main() {

	fmt.Println("Examples Of Struct ")

	var person1 PersonDetails
	person1.name = "JOHN"
	person1.job = "SOFTWARE DEV"
	person1.salary = 210000000000000000

	fmt.Printf("person1 name is %s and job %s and salary %d", person1.name, person1.job, person1.salary)

	var employee1 EmployeeDetails

	employee1.person_details = PersonDetails{
		name:   "AVINASH",
		job:    "DEV",
		salary: 98000,
	}
	employee1.person_contact = PersonContact{
		phone: 1234123,
		email: "avi@gmail.com",
	}

}
