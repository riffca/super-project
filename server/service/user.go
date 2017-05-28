package service

import (
	//schema "../schema"
	"fmt"
)

type User struct {
	Data map[string]interface{}
}

func (u *User) Test() {

	fmt.Println(u.Data["test"])
	fmt.Println("Test User")
}

func (u *User) Create() {

}
