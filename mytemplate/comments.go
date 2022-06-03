package mytemplate

import "fmt"

type Comment struct {
	Title string
	Text  string
}

func ShowComments() {
	fmt.Println("Comment 1")
	fmt.Println("Comment 2")
	fmt.Println("Comment 3")
	fmt.Println("Comment 4")
	fmt.Println("Comment 5")
	fmt.Println("Comment 6")
}

func CreateComment(title, text string) *Comment {
	return &Comment{Title: title, Text: text}
}
