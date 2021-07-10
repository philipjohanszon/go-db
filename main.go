package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	Server "github.com/philipjohanszon/go-db/server"
)

func main() {
	environmentError := godotenv.Load()

	if environmentError != nil {
		panic("FATAL: " + environmentError.Error())
	}

	port, portError := strconv.Atoi(os.Getenv("CONN_PORT"))

	if portError != nil {
		panic("FATAL: " + portError.Error())
	}

	server := Server.Server{
		Port: port,
		Host: os.Getenv("CONN_HOST"),
	}

	//Server starts
	server.Start()
}
