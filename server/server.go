package server

import (
	"net"
	"strconv"
	"io"
	"bytes"
	"strings"
	"Airttp/logger"
	"Airttp/http"
	"Airttp/modules"
)

type Server struct {
	port int
}

func New(port int) *Server {
	server := new(Server)
	server.port = port
	return server
}

func (server *Server) Run() {
	ln, err := net.Listen("tcp", ":" + strconv.Itoa(server.port))
	if err != nil {
		logger.GetInstance().Error(err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.GetInstance().Error(err.Error())
		}
		logger.GetInstance().NoticeF("New connection from %s", conn.RemoteAddr().String())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	httpSession := http.NewSession(conn)
	buff := bytes.NewBufferString("")
	for {
		tmp := make([]byte, 256)
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				logger.GetInstance().ErrorF("read error: %s", err.Error())
			}
			break
		}
		appendBuf(buff, tmp, n)
		if (bytes.Contains(buff.Bytes(), []byte("\r\n\r\n"))) {
			if (checkForContent(buff)) {
				go modules.GetManagerInstance().ExecRequest(conn, httpSession, http.NewRequest(buff.Bytes()), http.NewResponce())
				buff.Reset()
			}
		}
	}
}

func checkForContent(buff *bytes.Buffer) bool {
	if (bytes.Contains(buff.Bytes(), []byte("Content-Length:"))) {
		contentLength := bytes.Split(buff.Bytes(), []byte("Content-Length:"))[1]
		contentLength = bytes.Split(contentLength, []byte("\r\n"))[0]
		contentLengthNb, err := strconv.Atoi(strings.TrimSpace(string(contentLength)))
		if err != nil {
			logger.GetInstance().Error(err.Error())
			return true
		}
		content := bytes.Split(buff.Bytes(), []byte("\r\n\r\n"))[1]
		if (len(content) >= contentLengthNb) {
			return true
		} else {
			return false
		}
	}
	return true
}

func appendBuf(buff *bytes.Buffer, data []byte, n int) {
	i := 0
	for (i < n) {
		buff.WriteByte(data[i])
		i++
	}
}