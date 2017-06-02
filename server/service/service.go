package service

func init() {

}

var MethodMap map[string]interface{}

func CheckMethod(service string, name string) bool {

	MethodMap := map[string][]string{
		"User": {
			"Test",
		},
		"Auth": {
			"checkToken",
		},
	}

	val := false
	for _, s := range MethodMap[service] {
		if s == name {
			val = true
		}
	}
	return val
}
