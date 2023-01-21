package main

import (
	"fmt"
	"net"
)

type Client struct {
	ip   string
	port int
}

func NewClient(ip string, port int) *Client {
	return &Client{
		ip:   ip,
		port: port,
	}
}

func (c *Client) Addr() string {
	return fmt.Sprintf("%s:%v", c.ip, c.port)
}

func (c *Client) Send(message []byte) (string, error) {
	conn, err := net.Dial("tcp", c.Addr())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func main() {
	c := NewClient("0.0.0.0", 10000)

	response, err := c.Send([]byte("send message with tcp connection"))
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
