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

func New(db *gorm.DB) {
	// Configure any package-level settings
	DB = db
}

func init() {

	type Default struct {
		Schema string
	}

	u, _ := json.Marshal(&schema.User{})
	p, _ := json.Marshal(&schema.Page{Content: "{&quot;json&quot;:&quot;default&quot;}"})
	d, _ := json.Marshal(&Default{Schema: "none"})

	//Превратить в interface
	MethodMap = map[string][]string{
		"User": {
			"Test",
			string(u),
		},
		"Auth": {
			"SignUp",
			"SignIn",
			string(d),
		},
		"Page": {
			"Get",
			"Create",
			"Update",
			"Delete",
			string(p),
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
