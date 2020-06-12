package main

import "fmt"

func main() {
	fmt.Print("masukkan nama : ")
	var input string
	fmt.Scanf("%s", &input)
	output := input
	fmt.Println("your name is : " + output)
}
