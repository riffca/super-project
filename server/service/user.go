package service

import (
	//schema "../schema"
	"fmt"
)

type User struct {
	Data interface{}
}

func (u *User) Test() {
	fmt.Println("Test User")
}

func (u *User) Create() {

}
