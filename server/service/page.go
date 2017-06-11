package service

import (
	"../schema"
	//"fmt"
	//"reflect"
	"strings"
	//"time"
	//"encoding/json"
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

	page := schema.Page{}
	DB.Where(&schema.Page{Name: "sta"}).First(&page)
	DB.Model(&page).Updates(schema.Page{
		Name:    c.(string),
		Content: n.(string),
	})

	page.Content = c.(string)
	page.Name = n.(string)
	d := DB.Save(&page)
	p.Data["service_data"] = d.Value

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
