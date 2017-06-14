package service

/*
/*Modify data as
/*t.Data = inteface{}
/*with every service
*/
import "github.com/jinzhu/gorm"
import "encoding/json"
import "../schema"

var MethodMap map[string][]string

var DB *gorm.DB

func InitDB(db *gorm.DB) {
	DB = db
}

func init() {

	type None struct {
		Schema string
	}

	none, _ := json.Marshal(&None{Schema: "None"})
	u, _ := json.Marshal(&schema.User{})
	p, _ := json.Marshal(&schema.Page{Content: "{&quot;json&quot;:&quot;default&quot;}"})
	l, _ := json.Marshal(&schema.Lead{})
	m, _ := json.Marshal(&schema.Message{})

	//Превратить в interface
	MethodMap = map[string][]string{

		//schema services
		"User": {
			"Get",
			"Create",
			"Update",
			"Delete",
			string(u),
		},
		"Page": {
			"Get",
			"Create",
			"Update",
			"Delete",
			string(p),
		},
		"Lead": {
			"Get",
			"Create",
			"Update",
			"Delete",
			string(l),
		},
		"Message": {
			"Create",
			string(m),
		},
		//app services
		"Auth": {
			"Register",
			"Login",
			"Logout",
			string(none),
		},
		"Data": {
			"DumpTables",
			string(none),
		},
	}
}

func CheckMethod(service string, name string) bool {
	val := false
	for _, s := range MethodMap[service] {
		if s == name {
			val = true
		}
	}
	return val
}
