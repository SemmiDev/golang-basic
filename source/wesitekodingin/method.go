package main

import "fmt"

const UMR int64 = 20000000

type Employee struct {
	Name    string
	Sallary int64
}

// func (e Employee) showEmployee() {
// 	fmt.Println("Name\t : ", e.Name)
// 	fmt.Println("Sallary\t : ", e.Sallary)
// }

func (e Employee) showEmployee() bool {
	fmt.Println("Name\t : ", e.Name)
	fmt.Println("Sallary\t : ", e.Sallary)
	if e.Sallary < UMR {
		return false
	}
	return true
}

func main() {
	p1 := Employee{
		Name:    "Sammi Aldhi Yanto",
		Sallary: 1000000000,
	}
	fmt.Println(p1.showEmployee())
}
