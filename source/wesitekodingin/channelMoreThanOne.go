package main

import "fmt"

/**
author : Sam
*/

func penjumlahan(a, b int, ch chan int) {
	j := a + b
	ch <- j
}
func pengurangan(a, b int, ch chan int) {
	k := a - b
	ch <- k
}

func main() {
	j := make(chan int)
	k := make(chan int)

	go penjumlahan(10, 8, j)
	go pengurangan(10, 9, k)
	hasilJumlah, hasilKurang := <-j, <-k

	fmt.Printf("Hasil dari penjumlahan %d dengan %d adalah %d\n",
		10, 8, hasilJumlah)

	fmt.Printf("Hasil dari pengurangan %d dengan %d adalah %d\n",
		10, 8, hasilKurang)
}
