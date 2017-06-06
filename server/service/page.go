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

func (p *Page) GetScheme() schema.Page {
	return schema.Page{}
}

func (p *Page) Create() {
	n, c := p.Data["Name"].(string), p.Data["Content"].(string)
	fmt.Println(n, c)
	//page := schema.DB.Where("name = ?", p.Data["name"]).First(&schema.Page{})

	page := schema.DB.Create(&schema.Page{Name: n, Content: c})
	fmt.Println(page)

	// if page != nil {
	//  page.Update("content", p.Data["content"])
	// }

	// // Create
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// // Read
	// var product Product
	// db.First(&product, 1) // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// // Delete - delete product
	// db.Delete(&product)

}
