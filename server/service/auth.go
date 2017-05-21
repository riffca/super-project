package service

import (
	"fmt"
)

type Auth struct {
	Token string
}

func (a *Auth) CheckToken() {
	fmt.Println("ПРОВЕКА ТОКЕНА " + a.Token)
}
