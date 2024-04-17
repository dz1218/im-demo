package server

import (
	"net"
	"sync"
)

type Server struct {
	once    sync.Once
	id      string
	address string
	sync.Mutex

	// 会话列表
	users map[string]net.Conn
}

func newServer(id, address string) *Server {
	return &Server{
		id:      id,
		address: address,
		users:   make(map[string]net.Conn, 100),
	}
}

func NewServer(id, address string) *Server {
	return newServer(id, address)
}
