package service

type Page struct {
	Data map[string]interface{}
}

func (p *Page) EditPage() {

	// page := schema.DB.Where("name = ?", p.Data["name"]).First(&schema.Page{})

	// if page != nil {
	//  page.Update("content", p.Data["content"])
	// }

}
