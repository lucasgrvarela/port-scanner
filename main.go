package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <protocol> <hostname> <port>")
		fmt.Println("If port == 0 then it will scan all ports from 1 to 1024, if any other number is specified it will scan only this specific port")
		return
	}

	protocol := os.Args[1]
	if protocol != "udp" && protocol != "tcp" {
		log.Fatal("Invalid protocol, you can use 'udp' or 'tcp'")
	}

	hostname := os.Args[2]
	port, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal("Invalid port number")
	}

	if port == 0 {
		results := ScanAll(hostname)
		for _, v := range results {
			fmt.Println(v)
		}
		return
	}
	fmt.Printf("%v\n", ScanOnePort(protocol, hostname, port))
}

type ScanResult struct {
	Port  string
	State string
}

func ScanOnePort(protocol string, hostname string, port int) ScanResult {
	result := ScanResult{Port: fmt.Sprintf("%s/%d", protocol, port)}
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.Dial(protocol, address)
	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()

	result.State = "Open"
	return result
}

func ScanAll(hostname string) []ScanResult {
	var results []ScanResult

	for port := 1; port <= 1024; port++ {
		result := ScanOnePort("tcp", hostname, port)
		results = append(results, result)
	}

	return results
}
