package main

import (
	"flag"
	"log"
	"v1/config"
	"v1/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the gkredis server")
	flag.IntVar(&config.Port, "port", 1609, "port for the gkredis server")
	flag.Parse()

}

func main() {
	setupFlags()

	log.Println("gk redis online")
	server.RunSyncTcpServer()
}
