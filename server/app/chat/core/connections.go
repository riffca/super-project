package core

var Connections []string

func AddConnection(session string) {
	Connections = append(Connections, session)
}
