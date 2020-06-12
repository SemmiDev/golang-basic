package main

import "fmt"

func main() {
	// number
	fmt.Println("1 + 1  : ", 10+2)
	fmt.Println("10 - 2 : ", 10-2)
	fmt.Println("10 * 3 : ", 10*3)
	fmt.Println("5 / 5  : ", 5/5)
	// biar hasilnya float
	fmt.Println("10.0 / 3.4", 10.0/3.4)
	fmt.Println("10 % 2 : ", 10%2)

	// boolean
	fmt.Println(true && true)
	fmt.Println(false && false)
	fmt.Println(true && false)
	fmt.Println(false && true)

	fmt.Println(true || true)
	fmt.Println(false || false)
	fmt.Println(true || false)
	fmt.Println(false || true)

	fmt.Println(!true)
	fmt.Println(!false)

	fmt.Println((true && false) || (false && true) || !(false && false))

}
