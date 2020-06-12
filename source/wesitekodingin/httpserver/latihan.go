package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port string = "9090"

func welcome(w http.ResponseWriter, r *http.Request) {
	dataHTML := `<h1>Kumpulan Bahasa Pemrograman Keren</h1>
        <ul>
            <li>Golang</li>
            <li>PHP</li>
            <li>JavaScript</li>
            <li>Rust</li>
            <li>Java</li>
        </ul>
        Tutorial by : <a href="localhost:9090">check</a>
        `
	if r.Method == "GET" {
		fmt.Fprintf(w, dataHTML)
	}
}

func files(w http.ResponseWriter, r *http.Request) {
	dataHTML := `
	<h1>Hei</h1>
        <ul>
           <li>sam</li>
	  </ul>
	  `
	if r.Method == "GET" {
		fmt.Fprintf(w, dataHTML)
	}
}

func main() {

	fmt.Println("Berjalan di Port :", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sedang Akses URL %v dengan method %v", r.URL, r.Method)
	})

	http.Handle("/handle", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sedang Akses URL %v dengan method %v", r.URL, r.Method)
	}))

	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/files", files)

	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
