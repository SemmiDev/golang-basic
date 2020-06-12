package main

import "fmt"

type Kendaraan interface {
	getWarna()
}

type Mobil struct {
	Depan    string
	Belakang string
	Atas     string
	Dalam    string
}

func cetakWarna(w Kendaraan) {
	fmt.Println(w)
}

type Motor struct {
	Depan    string
	Belakang string
}

func (m Mobil) getWarna() {
	fmt.Println(m)
}

func (m Motor) getWarna() {
	fmt.Println(m)
}

func main() {
	var k Kendaraan
	mobil := Mobil{
		Depan:    "Biru",
		Belakang: "Biru",
		Atas:     "Biru",
		Dalam:    "Biru",
	}
	motor := Motor{
		Depan:    "Biru",
		Belakang: "Biru",
	}

	k = mobil
	k.getWarna()
	k = motor
	k.getWarna()

	cetakWarna(k)
}
