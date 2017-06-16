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
	ucb.ID, _ = strconv.ParseUint(lead.Data["CreatedBy"].(string), 10, 64)
	uad := schema.User{}
	uad.ID, _ = strconv.ParseUint(lead.Data["Adress"].(string), 10, 64)
	code, _ := strconv.ParseUint(lead.Data["StatusCode"].(string), 10, 64)

	l := schema.Lead{
		StatusCode: code,
		CreatedBy:  ucb,
		Adress:     uad,
	}

	DB.Find(&ucb)
	DB.Find(&uad)
	DB.Create(&l)

	DB.Model(&ucb).Association("Leads").Append(&l)
	DB.Model(&l).Association("CreatedBy").Append(&ucb)
	DB.Model(&l).Association("Members").Append([]schema.User{ucb, uad})

	m := DB.Model(&l).Association("Members").Find(&ucb)
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

	lead.model = &schema.Lead{}
	lead.model.ID, _ = strconv.ParseUint(lead.Data["ID"].(string), 10, 64)
	lead.model.StatusCode, _ = strconv.ParseUint(lead.Data["StatusCode"].(string), 10, 64)

	fmt.Println(lead.model)

	DB.First(&lead.model, lead.model.ID)

	a := DB.Model(&lead.model).Association("Members").Find(&lead.model.Members)

	if a.Error != nil {
		panic(a.Error)
	}

	lead.Data["service_data"] = a

}
