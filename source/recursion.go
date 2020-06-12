package main

import "fmt"

func main() {
	fmt.Println(faktorial(10))
}

func faktorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * faktorial(x-1)
}
