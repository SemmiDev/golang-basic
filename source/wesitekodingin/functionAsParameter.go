package main

import "fmt"

func main() {
	// contoh pertama
	v := func(nilai int) int {
		return nilai
	}

	result := checknilai(v)
	fmt.Println(result)

	// contoh kedua
	v2 := func(nilai int) int {
		return nilai
	}

	name, nilai := checknilai2("Sam", v2)
	fmt.Println(name, nilai)
}

func checknilai(nilai func(int) int) int {
	return nilai(10)
}
func checknilai2(name string, nilai func(int) int) (string, int) {
	return name, nilai(100)
}
