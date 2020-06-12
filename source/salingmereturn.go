package main

import "fmt"

func main() {
	fmt.Println(f1())
}

func f1() int {
	return f2() + 5
}
func f2() int {
	return f3() + f4()
}
func f3() int {
	return 5
}
func f4() int {
	return 10
}
