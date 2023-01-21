package server

import (
	"fmt"
	"log"
	"net"
)

type Config struct {
	port int
}

func NewConfig(port int) *Config {
	return &Config{port: port}
}

type Configure interface {
	Port() int
}

func (config *Config) Port() int {
	return config.port
}

type Server interface {
	Configure
	Handle(conn net.Conn)
}

func RunTCP(server Server) {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%v", server.Port()))
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening for TCP connections on port %v", server.Port())

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go server.Handle(conn)
	}
}
