package service

import (
	"fmt"
)

type User struct {
	Auth  bool
	Token string
}

func (u *User) CheckAuth() {
	fmt.Println("ЭТО ЮЗЕР")
}
