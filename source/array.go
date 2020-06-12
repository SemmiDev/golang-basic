package main

import (
	"fmt"
)

func main() {
	var num [3]int
	var names [5]string

	for i := 0; i < len(num); i++ {
		num[i] = i
		fmt.Println(num[i])
	}

	names[0] = "sammi"
	names[1] = "dev"
	names[2] = "eko"
	names[3] = "adet"
	names[4] = "null"

	// bahasa golang ini sedikit rewel, jika suatu variable yg telah dibuat dan tidak digunakan
	// maka akan terjadi error :)
	// untuk meng ignore kan nya gunakan tanda underscrore _
	for _, value := range names {
		// fmt.Println("index : ", index, " => ", value)
		fmt.Println(value)
	}

	hp := [5]string{"iphone", "samsung", "oppo", "vivo", "lenovo"}
	for _, value := range hp {
		fmt.Println(value)
	}
}
