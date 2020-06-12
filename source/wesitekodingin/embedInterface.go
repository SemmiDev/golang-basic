package main

import (
	"fmt"
)

type WarnaModif interface {
	GetWarnaModif()
}

type WarnaOriginal interface {
	GetWarnaOriginal()
}

type Warna interface {
	WarnaModif
	WarnaOriginal
}

type Mobil struct {
	Depan    string
	Belakang string
	Atas     string
	Dalam    string
}

func (m Mobil) GetWarnaOriginal() {
	fmt.Println(m)
}

func (m Mobil) GetWarnaModif() {
	fmt.Println(m)
}

func main() {
	var k Warna
	mobil := Mobil{
		Depan:    "Biru",
		Belakang: "Biru",
		Atas:     "Pink",
		Dalam:    "Putih",
	}

	k = mobil
	k.GetWarnaModif()

	mobil2 := Mobil{
		Depan:    "Green",
		Belakang: "Biru",
		Atas:     "Pink",
		Dalam:    "Putih",
	}

	k = mobil2
	k.GetWarnaOriginal()
}
