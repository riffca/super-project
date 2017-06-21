package service

import (
	"../schema"
	//"fmt"
	"strconv"
	"strings"
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

	if sc, err := strconv.ParseUint(l.Data["status_code"].(string), 10, 64); err == nil {
		l.model.StatusCode = sc
	}

	l.model.CreatorID, _ = strconv.ParseUint(l.Data["creator_id"].(string), 10, 64)

	if l.model.CreatorID > 0 {
		l.current = "creator_id"
		l.active = l.model.CreatorID
		m := []string{l.current, " = ?"}
		all := []schema.Lead{}
		d := DB.Where(strings.Join(m, ""), l.active).Preload("Members").Find(&all)

		l.Data["service_data"] = d
		return
	}

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

func (l *Lead) WriteMessage() {

	m := l.Data["message"].(string)

	uid, _ := strconv.ParseUint(l.Data["sender_id"].(string), 10, 64)
	user := schema.User{}
	user.ID = uid

	l.model.ID, _ = strconv.ParseUint(l.Data["lead_id"].(string), 10, 64)

	DB.First(&l.model)
	DB.First(&user)
	message := DB.Create(schema.Message{
		Text:     m,
		SenderID: user.ID,
		Sender:   user,
		Lead:     l.model,
		LeadID:   l.model.ID,
	})

	l.Data["service_data"] = message
	l.Data["service_message"] = "MESSAGE SEND"

}
