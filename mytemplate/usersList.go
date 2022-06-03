package mytemplate

import "fmt"

type User struct {
	Name string
	Age  int
}

func ShowUsers() {
	fmt.Println("user1")
	fmt.Println("user2")
	fmt.Println("user3")
}

func CreateUser(name string, age int) *User {
	return &User{name, age}
}
