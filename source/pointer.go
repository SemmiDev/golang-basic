package main

import "fmt"

func main() {
	// concept
	var num = 100
	var name = "sammi"
	// * => alamat
	// & => referensi ke address memory
	var pointerAddress *int = &num
	var pointerName *string = &name

	fmt.Println("alamat : ", pointerAddress)
	fmt.Println("alamat : ", &num)
	fmt.Println("value  : ", *pointerAddress)

	fmt.Println("alamat : ", pointerName)
	fmt.Println("alamat : ", &name)
	fmt.Println("value  : ", *pointerName)

	*pointerAddress = 200
	fmt.Println("value  : ", *pointerAddress) // 200
	fmt.Println("value  : ", num)             // 200

	// practice
	var point = 100
	changePointer(&point)
	fmt.Println(point)

}

func changePointer(point *int) {
	*point = 200 + 5

}
