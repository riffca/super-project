package service

import (
	"../schema"
	"fmt"
	//"time"
	//"encoding/json"
)

type Page struct {
	Data map[string]interface{}
}

func (p *Page) Get() {
	d := schema.DB.Find(&schema.User{})
	p.Data["service_data"] = d
}

func (p *Page) Create() {

	n, c := p.Data["Name"], p.Data["Content"]
	page := schema.DB.Create(&schema.Page{Name: n.(string), Content: c.(string)})

	p.Data["service_data"] = page
	fmt.Println(page)
}
