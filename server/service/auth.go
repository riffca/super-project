package service

import (
	"fmt"
)

type Auth struct {
	Data  map[string]interface{}
	Token string
}

func (a *Auth) CheckToken() {

	fmt.Println("ПРОВЕКА ТОКЕНА " + a.Token)

}

func (a *Auth) Register() {

}

func (a *Auth) Login() {

}

func (a *Auth) Logout() {

}
