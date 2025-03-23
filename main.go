package main

import (
	"giva-url-shortner/database"
	"giva-url-shortner/server"
)

func main() {
	database.ConnectDB()
	server.RunServer()
}
