package main

import "fmt"

func main() {
	names := [5]string{
		"Sam",
		"dev",
		"adit",
		"dandi",
		"gusnur",
	}

	slice1 := names[1:4] // range 1 -> 3
	slice2 := names[:3]  // from 0 -> 2
	slice3 := names[1:]  // from 2 -> akhir
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// jika kita mengubah isi / value di slice dengan [index] maka array  sumber yg di slice juga akan ikut berubah

	/*
		nb => slice hanya sebagai pointer terhadap array,
			jadi isi array bukan di copy ke slice tapi cuma sebagai pointer/penunjuk.
			jika kita mengubah isi array makan yg di slice (include) yg dirubah juga akan otomatis berubah
	*/

}
