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

func (a *Auth) GetToken() {
	fmt.Println("ПОЛУЧЕН ТОКЕН")
}

func (a *Auth) SignUp() {

}

func (a *Auth) SignIn() {

}
