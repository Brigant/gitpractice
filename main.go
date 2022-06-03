package main

import "fmt"

func main() {
	err := ShowMsg("To")
	if err != nil {
		fmt.Println(err)
	}

}
