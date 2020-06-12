package main

import "fmt"

type vector struct {
	x, y int
}

type structure struct {
	name  string
	kelas int
}

func main() {
	// struct mirip seperti list atau collection di java
	var x vector
	x.x = 10
	x.y = 15

	fmt.Println(x.x + x.y)

	var satu structure
	satu.name = "satu"
	satu.kelas = 12

	fmt.Println(satu)
	fmt.Println(satu.name)
}
