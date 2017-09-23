package http

import (
	"net"
)

type Session struct {
	Data map[string]string
}

func NewSession(conn net.Conn) *Session {
	session := new(Session)
	return session
}