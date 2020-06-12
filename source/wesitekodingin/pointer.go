package main

import "fmt"

func main() {
	// * : alamat
	// & : value dari alamat

	var nomor int = 10
	var view *int = &nomor

	fmt.Println("alamat memori var nomor : ", view)
	fmt.Println("nilai dari variable view : ", *view)

	// passing pointer ke sebuah fungsi

	var a int = 10
	var b *int = &a
	ubah(b)
	fmt.Println("alamat memory b : ", b)
	fmt.Println("nilai dari var b :", *b)

	// mengenbalikan nilai pointer dengan fungsi
	var a1 int = 10
	var a2 *int = &a1
	var a3 int = 15
	datac := ubah2(a3)

	fmt.Println("alamat a2 ", a2)
	fmt.Println("nilI a2 ", *a2)

	fmt.Println("alamat dari a3 ", datac)
	fmt.Println("value a3 ", &datac)

}

func ubah(data *int) {
	*data = 100
}

func ubah2(data int) *int {
	data = 20
	return &data
}
