package service

import (
	schema "../schema"
	//"fmt"
	"strconv"
	"strings"
)

type User struct {
	Data     map[string]interface{}
	active   interface{}
	current  string
	searchID string
	model    schema.User
}

func (u *User) Update() {

	un, em := u.Data["user_name"], u.Data["email"]
	pw := u.Data["password"]

	id, _ := strconv.ParseUint(u.searchID, 10, 64)

	user := schema.User{}
	DB.First(&user, id)
	d := DB.Model(&user).Updates(schema.User{
		UserName: un.(string),
		Email:    em.(string),
		Password: pw.(string),
	})

	u.Data["service_data"] = d

}

func (p *User) Get() {

	p.searchID = p.Data["id"].(string)
	username := p.Data["user_name"].(string)
	email := p.Data["email"].(string)

	if len(username) > 0 {
		p.active, p.current = username, "user_name"
	}

	if len(email) > 0 {
		p.active, p.current = email, "email"
	}

	if len(p.searchID) > 0 {
		p.active, p.current = p.searchID, "id"
	}

	if len(p.current) > 0 {
		m := []string{p.current, " = ?"}
		d := DB.Where(strings.Join(m, ""), p.active).First(&p.model)
		p.Data["service_data"] = d

		return

	}

	s := []schema.User{}
	a := DB.Find(&s)
	p.Data["service_data"] = a

}

func (u *User) Create() {
	n, e := u.Data["user_name"], u.Data["email"]
	pw := u.Data["password"]

	us := schema.User{
		UserName: n.(string),
		Email:    e.(string),
		Password: pw.(string),
	}
	user := DB.Create(&us)
	u.Data["service_data"] = user
}

func (u *User) GetLeads() {

	u.model.ID, _ = strconv.ParseUint(u.Data["id"].(string), 10, 64)
	DB.First(&u.model)
	//ls := []schema.Lead{}

	DB.Where(&u.model).First(&u.model)
	d := DB.Model(&u.model).Association("Leads")
	u.Data["service_data"] = d
	u.Data["service_message"] = "NOT WORKING METHOD"

}
