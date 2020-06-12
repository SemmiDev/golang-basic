package main

import (
	"fmt"
)

type val int

func (i val) jumlah() val {
	s := i * val(2)
	return s
}

func main() {
	numb := val(10)
	fmt.Println(numb.jumlah())
}
