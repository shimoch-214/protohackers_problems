package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"protohackers_problems/problems/means-to-end/data"
	"protohackers_problems/server"
)

type Request struct {
	Method string
	Filed1 int32
	Field2 int32
}

func NewRequest(message []byte) (*Request, error) {
	method := string(message[:1])
	field1 := int32(binary.BigEndian.Uint32(message[1:5]))
	field2 := int32(binary.BigEndian.Uint32(message[5:]))
	req := Request{method, int32(field1), int32(field2)}
	return &req, nil
}

type MeanToEnd struct{ *server.Config }

func (MeanToEnd) Handle(conn net.Conn) {
	defer conn.Close()
	assets := data.NewAssets()

	for {
		rawReq := make([]byte, 9)
		if _, err := io.ReadAtLeast(conn, rawReq, 9); err != nil {
			log.Fatal(err)
		}
		req, err := NewRequest(rawReq)
		if err != nil {
			log.Fatal(err)
		}
		res := []byte{}
		switch req.Method {
		case "I":
			assets.Insert(data.NewRow(req.Filed1, req.Field2))
		case "Q":
			binary.PutVarint(res, int64(assets.Query(req.Filed1, req.Field2)))
		default:
			log.Printf("invalid method: %v", req.Method)
		}
		if _, err := conn.Write(res); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	config := server.NewConfig(10000)
	server.RunTCP(MeanToEnd{config})
}
