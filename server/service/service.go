package service

func init() {

}

func CheckMethod(service string, name string) bool {
	schema := map[string][]string{
		"User": {
			"Test",
		},
		"Auth": {
			"checkToken",
		},
	}
	val := false
	for _, s := range schema[service] {
		if s == name {
			val = true
		}
	}
	return val
}
