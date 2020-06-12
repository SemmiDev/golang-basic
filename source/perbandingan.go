package main

import "fmt"

func main() {
	a := 2
	b := 5

	c := a + 3

	result := a == b  // false
	result2 := b == c // true
	result3 := a != b // true
	result4 := a < b  // true
	result5 := a <= b // true
	result6 := a > b  // false
	result7 := a >= b // false

	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)

}
