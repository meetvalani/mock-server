package main

import (
	"log"
	database "mockserver/database"
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
	// TODO: read database schema from file.
	db := database.GetDatabase(
		"tester",
		"test.db",
		`CREATE TABLE IF NOT EXISTS mock (
			id INTEGER PRIMARY KEY,
			method TEXT,
			endpoint TEXT,
			responseCode INTEGER,
			httpResponseContentType TEXT,
			httpHeaders TEXT,
			httpResponseBody TEXT
			)`,
	)
	server.StartServer(getHostFromCli(os.Args), db)
}
