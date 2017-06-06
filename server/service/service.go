package service

var MethodMap map[string][]string

func init() {
	MethodMap = map[string][]string{
		"User": {
			"Test",
			"Go",
		},
		"Auth": {
			"checkToken",
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
