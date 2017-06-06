package service

import "encoding/json"
import "../schema"

var MethodMap map[string][]string

func init() {

	u, _ := json.Marshal(&schema.User{})
	p, _ := json.Marshal(&schema.Page{})

	//Превратить в interface
	MethodMap = map[string][]string{
		"User": {
			"Test",
			"Go",
			string(u),
		},
		"Auth": {
			"checkToken",
		},
		"Page": {
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
