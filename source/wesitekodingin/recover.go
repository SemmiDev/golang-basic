package main

import (
	"fmt"
)

func recoverCetak() {
	if r := recover(); r != nil {
		fmt.Println("error", r)
	}
}

func cetak(nama *string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error", r)
		}
	}() // kalau diawal pake defer, kalo diakhir fungsi ga usah

	if nama == nil {
		panic("Tidak boleh nil om")
	}
}

func main() {
	fmt.Println("Belajar Golang ke 33.\n")
	cetak(nil)
}
