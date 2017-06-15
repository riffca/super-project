package service

import (
	"../schema"
	"strconv"
	// "strings"
)

type Lead struct {
	Data     map[string]interface{}
	active   interface{}
	current  string
	searchID string
	model    *schema.Lead
}

// func (sc *Lead) ConnectUser() {

// }

// func (sc *Lead) GetMessages() {

// }

func (lead *Lead) Create() {

	ucb := schema.User{}
	ucb.ID, _ = strconv.ParseUint(lead.Data["CreatedBy"].(string), 10, 64)
	uad := schema.User{}
	uad.ID, _ = strconv.ParseUint(lead.Data["Adress"].(string), 10, 64)
	code, _ := strconv.ParseUint(lead.Data["StatusCode"].(string), 10, 64)

	l := schema.Lead{
		StatusCode: code,
		CreatedBy:  ucb,
		Adress:     uad,
	}
	//DB.First(&u, cb.(string))

	//mod := DB.Create(&l)

	mod := DB.Model(&ucb).Association("Leads").Append(&l)

	lead.Data["service_data"] = mod

}

// func (sc *Lead) Update() {

//  id, _ := strconv.ParseUint(sc.Data["ID"].(string), 10, 64)

//  mod := schema.Lead{}
//  DB.First(&mod, id)
//  d := DB.Model(&mod).Updates(schema.Lead{
//    Name:    n.(string),
//    Content: c.(string),
//  })
//  sc.Data["service_data"] = d

// }

// func (sc *Lead) Get() {
//  sc.model = &schema.Lead{}
//  sc.searchID = sc.Data["ID"].(string)

//  name := sc.Data["Name"].(string)

//  if len(sc.searchID) > 0 {
//    sc.active, sc.current = sc.searchID, "id"
//  }

//  if len(name) > 0 {
//    sc.active, sc.current = name, "name"
//  }

//  if len(sc.current) > 0 {
//    m := []string{sc.current, " = ?"}
//    d := DB.Where(strings.Join(m, ""), p.active).First(p.model)
//    sc.Data["service_data"] = d

//    return

//  }

//  s := []schema.Lead{}
//  a := DB.Find(&s)
//  sc.Data["service_data"] = a

// }
