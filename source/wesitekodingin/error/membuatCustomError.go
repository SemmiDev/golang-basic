package main

/*
   author : Sam
*/

import (
	"fmt"
)

type (
	customError struct {
		err string
		val int
	}
)

func (ce *customError) Error() string {
	return fmt.Sprintf("Nilainya %d: %s", ce.val, ce.err)
}
func checkAge(nilai int) (int, error) {
	if nilai <= 10 {
		return nilai, &customError{err: "umur tidak di ijinkan", val: nilai}
	}

	return nilai, nil
}

func main() {
	checkAge, err := checkAge(11)
	if err != nil {
		if err, ok := err.(*customError); ok {
			fmt.Printf("Umur anda adalah %d\n", err.val)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Umur anda adalah %d \n", checkAge)
}
