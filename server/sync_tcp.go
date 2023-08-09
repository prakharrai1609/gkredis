package server

import (
	"io"
	"log"
	"net"
	"strconv"
	"v1/config"
)

func readCommand(conn net.Conn) (string, error) {
	var buffer []byte = make([]byte, 1024)
	n, err := conn.Read(buffer[:])

	if err != nil {
		return "", err
	}

	return string(buffer[:n]), err
}

func RunSyncTcpServer() {
	log.Printf("Starting a synchronous tcp server at %s::%d!", config.Host, config.Port)

	listener, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))

	if err != nil {
		panic(err)
	}

	var concurrent_clients int = 0

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		concurrent_clients += 1

		log.Println("Client with address -> ", conn.RemoteAddr(), " disconnected, concurrent client count -> ", concurrent_clients)
		go func() {
			for {
				command, err := readCommand(conn)

				if err != nil {
					conn.Close()
					concurrent_clients -= 1
					log.Println("Client with address -> ", conn.RemoteAddr(), ", concurrent client count -> ", concurrent_clients)

					if err == io.EOF {
						break
					}

					log.Println("err", err)
				}

				log.Println("command > ", command)

			}
		}()
	}
}
