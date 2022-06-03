package main

import "fmt"

func ShowMsg(msg string) error {
	if len(msg) < 3 {
		return fmt.Errorf("Too short message: \"%s\"\n", msg)
	}
	fmt.Println("Hello World!")
	fmt.Println("Here dev1 add new func and make some fix \"ShowMsg(msg string) error\"")
	return nil
}
