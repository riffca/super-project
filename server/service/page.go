package service

import (
	"../schema"
	//"fmt"
	//"reflect"
	"strings"
	//"time"
	//"encoding/json"
	"strconv"
)

type Page struct {
	Data     map[string]interface{}
	active   interface{}
	current  string
	searchID string
	model    *schema.Page
}

func (p *Page) Update() {

	n, c := p.Data["Name"], p.Data["Content"]

	id, _ := strconv.ParseUint(p.Data["ID"].(string), 10, 64)

	page := schema.Page{}
	DB.First(&page, id)
	d := DB.Model(&page).Updates(schema.Page{
		Name:    n.(string),
		Content: c.(string),
	})
	p.Data["service_data"] = d

}

func (p *Page) Get() {
	p.model = &schema.Page{}
	p.searchID = p.Data["ID"].(string)

	name := p.Data["Name"].(string)

	if len(p.searchID) > 0 {
		p.active, p.current = p.searchID, "id"
	}

	if len(name) > 0 {
		p.active, p.current = name, "name"
	}

	if len(p.current) > 0 {
		m := []string{p.current, " = ?"}
		d := DB.Where(strings.Join(m, ""), p.active).First(p.model)
		p.Data["service_data"] = d

		return

	}

	s := []schema.Page{}
	a := DB.Find(&s)
	p.Data["service_data"] = a

}

func (p *Page) Create() {
	n, c := p.Data["Name"], p.Data["Content"]
	pa := schema.Page{Name: n.(string), Content: c.(string)}
	page := DB.Create(&pa)
	p.Data["service_data"] = page
}
