package main

import "fmt"

func main() {

	jumlah := func(x, y int) int {
		return x + y
	}

	kurang := func(x, y int) int {
		return x - y
	}

	bagi := func(x, y int) int {
		return x / y
	}

	kali := func(x, y int) int {
		return x * y
	}

	fmt.Println(jumlah(7, 4))
	fmt.Println(kurang(7, 4))
	fmt.Println(kali(7, 4))
	fmt.Println(bagi(7, 4))

	fmt.Print("masukkan nama anda : ")
	var input string
	fmt.Scanf("%s", &input)
	output := input
	nameCall(output)

}

func nameCall(name string) {
	fmt.Println("your name is " + name)
}
