package core

import "fmt"

type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Session string `json:"sock_session"`
}

var Users []User

func CurrentUser(session string) *User {
	var u *User
	for i, _ := range Users {
		fmt.Println(i)
		// if i.Session == Session {
		// 	u = i
		// }
	}
	return u
}

func DeleteUser(session string) {
	for u, i := range Users {
		// if u.Session == session {
		// 	Users = append(Users[:i], Users[i+1:]...)
		// }
		fmt.Println(u, i)
	}
}

func AddUser(u User) {
	Users = append(Users, u)
}
