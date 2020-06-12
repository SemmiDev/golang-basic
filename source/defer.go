package main

import (
	"fmt"
)

func main() {
	defer second()
	first()
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}
