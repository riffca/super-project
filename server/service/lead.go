package service

import (
	"../schema"
	"fmt"
	"strconv"
)

type Lead struct {
	Data    map[string]interface{}
	active  interface{}
	current string
	model   *schema.Lead
}

// func (sc *Lead) ConnectUser() {

// }

// func (sc *Lead) GetMessages() {

// }

func (lead *Lead) Create() {
	ucb := schema.User{}
	ucb.ID, _ = strconv.ParseUint(lead.Data["created_by"].(string), 10, 64)
	uad := schema.User{}
	uad.ID, _ = strconv.ParseUint(lead.Data["adress"].(string), 10, 64)
	code, _ := strconv.ParseUint(lead.Data["status_code"].(string), 10, 64)

	l := schema.Lead{
		StatusCode: code,
		CreatedBy:  ucb,
		Adress:     uad,
	}
	fmt.Println("LEAD CREATE------------------>")
	fmt.Println(&l)

	DB.Find(&ucb)
	DB.Find(&uad)
	DB.Create(&l)

	DB.Model(&ucb).Association("Leads").Append(&l)
	DB.Model(&l).Association("CreatedBy").Append(&ucb)
	m := DB.Model(&l).Association("Members").Append([]schema.User{ucb, uad})

	lead.Data["service_data"] = m

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

func (lead *Lead) Get() {

	fmt.Println("GET LEAD------------------>")
	model := schema.Lead{}
	model.ID, _ = strconv.ParseUint(lead.Data["id"].(string), 10, 64)
	// lead.model = &schema.Lead{}
	// lead.model.StatusCode, _ = strconv.ParseUint(lead.Data["status_code"].(string), 10, 64)

	DB.Where(&model).First(&model)
	model.Members = []schema.User{}
	DB.Model(&model).Association("Members").Find(&model.Members)

	lead.Data["service_data"] = model

}
