package main

import (
	"github.com/santoshmano/redis/config"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {

	// Loading the configuration
	redisConfig := config.LoadConfig()

	ipAddr := redisConfig.Server.IPAddr
	port := strconv.Itoa(redisConfig.Server.Port)
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

	daytime := time.Now().String()

	if _, err := conn.Write([]byte(daytime + "\n")); err != nil {
		log.Fatalf("Write error: %s", err.Error())
	}
	if err := conn.Close(); err != nil {
		log.Printf("Close error: %s", err.Error())
	}
}
