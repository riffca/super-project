package service

import (
	"fmt"
)

type User struct {
	Auth  bool
	Token string
	Name  string
}

func (u *User) Test() {
	fmt.Println("Test User")
}
