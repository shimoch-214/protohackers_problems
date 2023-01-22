package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"protohackers_problems/server"
)

func isPrimeNumber(number float64) bool {
	if number < 2 {
		return false
	}
	sqrt := math.Floor(math.Sqrt(number))
	for i := sqrt; i > 1; i-- {
		if math.Mod(number, i) == 0 {
			return false
		}
	}
	return true
}

type PrimeTime struct{ *server.Config }

type Request struct {
	Method string      `json:"method"`
	Number json.Number `json:"number"`
}

func (req *Request) isValidRequest() bool {
	if req.Method != "isPrime" {
		return false
	}
	_, err := req.Number.Float64()
	return err == nil
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func (PrimeTime) Handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	var resMessage []byte
	for scanner.Scan() {
		reqMessage := scanner.Bytes()

		var req Request
		err := json.Unmarshal(reqMessage, &req)
		if err != nil || !req.isValidRequest() {
			fmt.Println(err)
			resMessage = []byte("invalid request")
			if _, err := conn.Write(append(resMessage, []byte("\n")...)); err != nil {
				log.Fatal(err)
			}
			conn.Close()
		} else {
			number, _ := req.Number.Float64()
			isPrime := isPrimeNumber(number)
			res := Response{
				Method: "isPrime",
				Prime:  isPrime,
			}
			resMessage, _ = json.Marshal(res)
			if _, err := conn.Write(append(resMessage, []byte("\n")...)); err != nil {
				log.Fatal(err)
			}
		}

	}
}

func main() {
	config := server.NewConfig(10000)
	server.RunTCP(PrimeTime{config})
}
