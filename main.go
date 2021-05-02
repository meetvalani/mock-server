package main

import (
	"log"
	server "mockserver/server"
	"os"
)

func getHostFromCli(args []string) (address string) {
	if len(args) > 2 && args[1] == "--host" {
		return args[2]
	}
	log.Println("Invalid args, example: ./main --host 127.0.0.1:8080")
	panic("No valid args found.")
}

func main() {
	server.StartServer(getHostFromCli(os.Args))
}
