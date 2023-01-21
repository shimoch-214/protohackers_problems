package main

import (
	"io"
	"log"
	"net"

	server "protohackers_problems/server"
)

type SmokeTest struct{ *server.Config }

func (SmokeTest) Handle(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Println(err)
	}
}

func main() {
	config := server.NewConfig(8080)

	server.RunTCP(SmokeTest{config})
}
