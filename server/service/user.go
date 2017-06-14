package service

import (
	schema "../schema"
	"strconv"
	"strings"
)

type User struct {
	Data     map[string]interface{}
	active   interface{}
	current  string
	searchID string
	model    *schema.User
}

func (u *User) Update() {

	un, em := u.Data["UserName"], u.Data["Email"]
	pw := u.Data["Password"]

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
	p.model = &schema.User{}
	p.searchID = p.Data["ID"].(string)
	username := p.Data["UserName"].(string)
	email := p.Data["Email"].(string)

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
		d := DB.Where(strings.Join(m, ""), p.active).First(p.model)
		p.Data["service_data"] = d

		return

	}

	s := []schema.User{}
	a := DB.Find(&s)
	p.Data["service_data"] = a

}

func (u *User) Create() {
	n, e := u.Data["UserName"], u.Data["Email"]
	pw := u.Data["Password"]

	pa := schema.User{
		UserName: n.(string),
		Email:    e.(string),
		Password: pw.(string),
	}
	page := DB.Create(&pa)
	u.Data["service_data"] = page
}
