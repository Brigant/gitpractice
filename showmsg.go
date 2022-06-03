package main

import "fmt"

func Intro() {
	fmt.Println("Hello in INTRO. Let's GO")
}

func ShowMsg(msg string) error {
	if len(msg) < 3 {
		return fmt.Errorf("too short message: \"%s\"", msg)
	}
	fmt.Println("Hello World!")
	fmt.Println("Here dev1 add new func and make some fix \"ShowMsg(msg string) error\"")
	return nil
}
