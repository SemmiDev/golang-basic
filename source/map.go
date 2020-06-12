package main

import "fmt"

func main() {

	mahasiswa := make(map[string]string)

	mahasiswa["101"] = "sam"
	mahasiswa["102"] = "dev"
	mahasiswa["103"] = "coder"

	fmt.Println(mahasiswa)
	// akses sama seperti array
	fmt.Println(mahasiswa["100"])
	fmt.Println(mahasiswa["101"])
	fmt.Println(mahasiswa["102"])
	// jika kita akses index 4, maka akan ditampilkan kosong
	// kalau di array akan error

	for nim, name := range mahasiswa {
		fmt.Println("nim : ", nim, " => ", name)
	}

	// cara lain membuat map
	siswa := map[string]string{
		"satu":  "sam",
		"dua":   "dev",
		"tiga":  "otong",
		"empat": "surotong",
		"lima":  "otong gede",
	}

	for index, value := range siswa {
		fmt.Println("nim : ", index, " => ", value)
	}
}
