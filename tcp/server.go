package tcp

import (
	"bobac/bobac/cache"
	"log"
	"net"
)

type Server struct {
	cache.Cache
}

func New(c cache.Cache)*Server  {
	return &Server{c}
}

func (s *Server)Listen()  {
	listener, lerr := net.Listen("tcp", ":6380")
	if lerr != nil {
		log.Println(lerr)
		panic(lerr)
	}
	for {
		conn, cerr := listener.Accept()
		if cerr != nil {
			log.Println(cerr)
			panic(cerr)
		}
		//start a goroutine to handle conect

	}
}

