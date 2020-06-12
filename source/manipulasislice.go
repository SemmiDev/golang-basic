package main

import "fmt"

func main() {
	// cara pertama
	slice1 := []string{
		"sam",
		"dev",
		"coder",
	}

	fmt.Println(slice1)

	// cara kedua
	slice2 := make([]string, 3)
	slice2[0] = "sam"
	slice2[1] = "dev"
	slice2[2] = "coder"

	fmt.Println(slice2)

	// men slicekan slice
	// : => slice semua
	slice3 := slice2[:]
	fmt.Println(slice3)
	slice3[0] = "dirubah"

	fmt.Println(slice2)
	fmt.Println(slice3)

	// operasi di slice

	// nambah data kedalam slice(append)
	// ingat ketika kita append / menambah data di slice, itu artniya kita make slice baru bukan update slice nya
	slice4 := append(slice2, "newnew", "newnew2") // slice 4 kita copy dari slice2 kemudian tambah data dengan newnew dan newnew2
	fmt.Println(slice4)

	// copy
	// dia akan meng copy sesuai kebutuhan, jika kita pasang length = 2, maka cuma akan 2 saja yg di include
	slice5 := make([]string, 5)
	// destinations,source
	copy(slice5, slice4)
	fmt.Println(slice5)

}
