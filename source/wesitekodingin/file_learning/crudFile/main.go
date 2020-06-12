package main

import (
	"fmt"
	"log"
	"os"
)

const path = "note.txt"

func membuatFile() {
	c, err := os.Create("note.txt")
	if err != nil {
		log.Fatal(err)
	}
	c.WriteString("1. Belajar PHP\n")
	c.WriteString("2. Belajar Javascript\n")
	c.WriteString("3. Belajar Golang\n")
	c.WriteString("4. Belajar C++\n")
	defer c.Close()
	fmt.Println("selesai membuat file")
}
func menambahIsiFile() {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	file.WriteString("\nnew line\n")
	file.WriteString("3. Belajar C#\n")
	file.WriteString("4. Belajar Python\n")
	file.Sync()
}
func hapusFile() {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println(err)
	}

	err = os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	membuatFile()
	menambahIsiFile()
	hapusFile()
}
