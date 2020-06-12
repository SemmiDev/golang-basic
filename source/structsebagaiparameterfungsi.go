package main

import "fmt"

func main() {
	a := Person{
		Id:     1,
		Name:   "Sammi",t
		Gender: "male",
	}
	b := Person{
		Id:     2,
		Name:   "tak tau",
		Gender: "female",
	}

	printPerson(a)
	printPerson(b)
}

type Person struct {
	Id     int
	Name   string
	Gender string
}

func printPerson(person1 Person) {
	fmt.Println("ID      = ", person1.Id)
	fmt.Println("Name    = ", person1.Name)
	fmt.Println("Gender  = ", person1.Gender)
}
