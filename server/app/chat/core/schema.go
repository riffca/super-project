package core

type DataSchema struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
	Token   string      `json:"token"`
}

type MsgSchema DataSchema

func (d *DataSchema) SetUser(session string) {

}
