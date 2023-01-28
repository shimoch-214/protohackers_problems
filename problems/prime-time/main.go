package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math"
	"net"
	"protohackers_problems/server"
)

func isPrimeNumber(number float64) bool {
	if number < 2 || math.Floor(number) != number {
		return false
	}
	sqrt := math.Floor(math.Sqrt(number))
	for i := float64(2); i <= sqrt; i++ {
		if math.Mod(number, i) == 0 {
			return false
		}
	}
	return true
}

type PrimeTime struct{ *server.Config }

type Request struct {
	Method string   `json:"method"`
	Number *float64 `json:"number"`
}

func (req *Request) isValidRequest() bool {
	return req.Method == "isPrime" && req.Number != nil
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
		log.Println(string(reqMessage))
		log.Println(req)
		if err != nil || !req.isValidRequest() {
			log.Println(err)
			resMessage = []byte("invalid request")
			if _, err := conn.Write(append(resMessage, []byte("\n")...)); err != nil {
				log.Fatal(err)
			}
		} else {
			isPrime := isPrimeNumber(*req.Number)
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
