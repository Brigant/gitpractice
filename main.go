package main

import "fmt"

func main() {
	Intro()

	err := ShowMsg("To")
	if err != nil {
		fmt.Println(err)
	}

}
