package main

type server struct {
	rooms map[string]*room
	commands chan commandID
}

func newServer() *server{
	return &server{
		rooms: make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("New client has connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn :
	}
}

//10:24