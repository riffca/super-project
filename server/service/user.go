package service

import (
	schema "../schema"
	"fmt"
)

type User struct {
	Data map[string]interface{}
}

func (u *User) Test() {
	u.Data["server_message"] = "modify data test ok 200"
	fmt.Println("Test Ok")
}

func (u *User) Create() {

}

func (p *User) GetScheme() schema.User {
	return schema.User{}
}
