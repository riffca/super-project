package service

import (
	"fmt"
)

type Auth struct {
	Data  interface{}
	Token string
}

func (a *Auth) CheckToken() {
	fmt.Println("ПРОВЕКА ТОКЕНА " + a.Token)
}
