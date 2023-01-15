package main

import (
	"log"
	"net"
	"os"
	"time"
)

func main() {
	ipAddr, port := os.Args[1], os.Args[2]
	service := ipAddr + ":" + port

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		log.Fatalf("Fatal error: %s, check correctness of ipaddr or port\n", err.Error())
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("Fatal error: %s\n", err.Error())
	}

	acceptErrorCount := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			acceptErrorCount++
			if acceptErrorCount > 1000 {
				log.Fatalf("Fatal Error: %s\n", err.Error())
			} else {
				log.Printf("Accept Error #%d: %s\n", acceptErrorCount, err.Error())
				continue
			}
		}
		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	daytime := time.Now().String()
	_, err := conn.Write([]byte(daytime + "\n"))
	if err != nil {
		log.Fatalf("Write error: %s", err.Error())
	}
}
