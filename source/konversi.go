package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int32 = 10
	var y int64 = int64(x) // konversi ke y
	var z float64 = float64(y)
	var a int32 = int32(z)

	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)

	var b float32 = 5.4
	var c int32 = int32(b) // tidak akan di bulatkan tapi diambil angka awal => 5
	fmt.Println(c)

	name := "sammi"
	fmt.Println(string(name[0])) // byte -> string

	var nilai string = "100"
	// sebenarnya setelah nilaiInt, underscore itu unutk nampung error nya
	nilaiInt, _ := strconv.Atoi(nilai)
	fmt.Println(nilaiInt)

	// kalau conv ke string cuma butuh satu paramerter
	nilaiString := strconv.Itoa(nilaiInt)
	fmt.Println(nilaiString)

	// parse (Cara yg lebih mudah)
	d, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64) // 64 => bit
	var gtoInt int32 = int32(f)
	fmt.Println(gtoInt)
	e, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println(d, f, e)

	// format menjadi string
	s := strconv.FormatBool(true)
	fmt.Println(s)

}
