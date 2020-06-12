package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port string = "9090"

func main() {
	fmt.Println("Berjalan di Port :", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sedang Akses URL %v dengan method %v", r.URL, r.Method)
	})
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
