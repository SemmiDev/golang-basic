package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(1.1)
	fmt.Println("tipe data string")
	fmt.Println(true && false)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	// multiple variable
	var (
		a = 1
		b = 2
		c = 3
	)

	var (
		resultAdd    = a + b + c
		resultKurang = (b - a) + c
	)

	fmt.Println(resultAdd)
	fmt.Println(resultKurang)
}
