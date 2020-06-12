package main

import "fmt"

// input
func input(name string) {
	fmt.Print("input your name : ")
	fmt.Scanln(&name)
	if len(name) <= 3 {
		panic("text length mush more than 3 characters")
	}
}

func cetak(name *string) {
	if name == nil {
		panic("error,ga boleh nil")
	}
}

func main() {
	var name string
	input(name)
	fmt.Println("learn golang panic series")
	cetak(&name)
}
