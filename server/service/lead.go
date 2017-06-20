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
	model   schema.Lead
}

func (lead *Lead) Create() {

	ucr := schema.User{}
	ucr.ID, _ = strconv.ParseUint(lead.Data["creator_id"].(string), 10, 64)

	uad := schema.User{}
	uad.ID, _ = strconv.ParseUint(lead.Data["adress_id"].(string), 10, 64)

	code, _ := strconv.ParseUint(lead.Data["status_code"].(string), 10, 64)

	// fmt.Println("LEAD CREATE------------------>")
	DB.Find(&ucr)
	DB.Find(&uad)

	l := schema.Lead{
		CreatedBy:  ucr,
		Adress:     uad,
		CreatorID:  ucr.ID,
		AdressID:   uad.ID,
		StatusCode: code,
	}

	d := DB.Create(&l)
	DB.Model(&l).Association("Members").Append([]schema.User{ucr, uad})
	DB.Model(&ucr).Association("Leads").Append(&d)
	if d.Error == nil {
		lead.Data["service_data"] = d
		lead.Data["service_message"] = "LEAD CREATED"
	}
}

func (l *Lead) Get() {
	if id, err := strconv.ParseUint(l.Data["id"].(string), 10, 64); err == nil {
		l.model.ID = id
	}
	fmt.Println("ID-------------------------->", l.model.ID)
	if sc, err := strconv.ParseUint(l.Data["status_code"].(string), 10, 64); err == nil {
		l.model.StatusCode = sc
	}
	fmt.Println("STATUS CODE-------------------------->", l.model.StatusCode)
	if l.model.ID == 0 && l.model.StatusCode == 0 {
		all := []schema.Lead{}
		d := DB.Find(&all)
		l.Data["service_data"] = d
		return
	}
	DB.Where(&l.model).First(&l.model)
	DB.Model(&l.model).Association("Members").Find(&l.model.Members)
	l.Data["service_data"] = l.model
}

func (l *Lead) Delete() {
	if id, err := strconv.ParseUint(l.Data["id"].(string), 10, 64); err == nil {
		l.model.ID = id
	}
	DB.Delete(&l.model)
}

func (l *Lead) GetUserLeads() {

}
