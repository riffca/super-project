package vendor

type DataSchema struct {
	Action  string                 `json:"action"`
	Payload map[string]interface{} `json:"payload"`
	Token   string                 `json:"token"`
}

type MsgSchema DataSchema
