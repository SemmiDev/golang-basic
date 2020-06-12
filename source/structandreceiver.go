package main

import "fmt"

func main() {
	ref := &Person{
		ID:        1,
		Name:      "Sam",
		Gender:    "male",
		isMarried: false,
	}
	fmt.Println(ref.getID())
	fmt.Println(ref.getName())
	fmt.Println(ref.getGender())
	fmt.Println(ref.getIsMarried())
	ref.ChangeName("sammi")
	fmt.Println(ref.getName())
}

type Person struct {
	ID           int
	Name, Gender string
	isMarried    bool
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

func (p *Person) getID() int {
	return p.ID
}
func (p *Person) getName() string {
	return p.Name
}
func (p *Person) getGender() string {
	return p.Gender
}
func (p *Person) getIsMarried() bool {
	return p.isMarried
}
