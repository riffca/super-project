package service

func init() {

}

func CheckMethod(service string, name string) bool {
	all := map[string][]string{
		"User": {
			"Test",
		},
		"Auth": {
			"checkToken",
		},
	}
	val := false
	for _, s := range all[service] {
		if s == name {
			val = true
		}
	}
	return val
}
