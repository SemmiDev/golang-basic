package main

import "fmt"

var scopeName string = "Sammi"

func main() {
	// deklarasi variable
	var name string = "belajar golang"
	var kelas string

	fmt.Println(name)
	name = "basic"
	fmt.Println(name)

	kelas = "xii"
	fmt.Println(kelas)

	var angka1 int = 5
	var angka2 int = 5
	var result int = angka1 + angka2
	fmt.Println(result)

	fmt.Println(name == kelas)    // false
	fmt.Println(angka1 == angka2) // true

	// cara yang lebih singkat, tapi harus langsung di inisialisasi
	namaLengkap := "Sammi Aldhi Yanto"
	fmt.Println(namaLengkap)

	// int = alias untuk int32
	nilai := 10
	fmt.Println(nilai)

	// constant
	const company string = "blibli"
	fmt.Println(company)

	// scope
	// misal var scopeName kita letakkan di luar method main biar bisa diakses di luar method main
	aksesScope() // akses method aksesScope
}

func aksesScope() {
	fmt.Println(scopeName)
}
