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

	id := p.Data["id"]
	if id != nil {
		fmt.Println("Find Page with id: ", id)
		d := DB.Where("id = ?", id.(string)).First(&schema.Page{})
		//d := DB.Find(&schema.Page{}, id.(float64))
		p.Data["service_data"] = d

	} else {
		s := []schema.Page{}
		a := DB.Find(&s)
		p.Data = a.Value
	}
}

func (p *Page) Create() {

	n, c := p.Data["Name"], p.Data["Content"]
	pa := schema.Page{Name: n.(string), Content: c.(string)}
	page := DB.Create(&pa)
	p.Data["service_data"] = page
	fmt.Println(page)

}
