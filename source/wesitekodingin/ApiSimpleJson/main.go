// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {

// }

// // mahasiswa
// type Mahasiswa struct {
// 	ID     int    `json:"id"`
// 	NIM    int    `json:"nim"`
// 	Name   string `json:"name"`
// 	Gender string `json:"gender"`
// }

// // nama mahasiswa
// func NewMahasiswa() []Mahasiswa {
// 	mhs := []Mahasiswa{
// 		Mahasiswa{
// 			ID:     1,
// 			NIM:    12321,
// 			Name:   "_SammiiDev",
// 			Gender: "male",
// 		},
// 		Mahasiswa{
// 			ID:     2,
// 			NIM:    32221,
// 			Name:   "_Eji",
// 			Gender: "female",
// 		},
// 	}
// 	return mhs
// }

// // GetMahasiswa
// func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		mhs := NewMahasiswa()
// 		datamahasiswa, err := json.Marshal(mhs)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(datamahasiswa)
// 		return
// 	}

// 	http.Error(w, "hayo mau ngapain", http.StatusNotFound)
// }

// func main() {
// 	http.HandleFunc("/mahasiswa", GetMahasiswa)
// 	fmt.Println("server running...")
// 	if err := http.ListenAndServe(":7000", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
