package main

import "fmt"

type (
	Category struct {
		Name string
	}
	Post struct {
		Title string
	}
)

func (c Category) showData() {
	fmt.Println(c)
}

func (p Post) showData() {
	fmt.Println(p)
}

func main() {
	cats := Category{
		Name: "News",
	}
	post := Post{
		Title: "learn golang",
	}

	cats.showData()
	post.showData()
}
