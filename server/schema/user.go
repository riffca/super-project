package schema

type User struct {
	Model
	Name  string
	Leads []Lead
}