package server

import (
	"fmt"
	"net"
)

type Server struct {
	Port int
	Host string
}

func (server *Server) Start() {
	//Creates listener
	listener, listenerError := net.Listen("tcp", server.Host+":"+fmt.Sprint(server.Port))

	if listenerError != nil {
		panic("FATAL ERROR: " + listenerError.Error())
	}

	//Closes listener at the end of the function (at the end of the application)
	defer listener.Close()

	fmt.Println("Listening on: " + server.Host + ":" + fmt.Sprint(server.Port))

	for {
		connection, connectionError := listener.Accept()

		if connectionError != nil {
			fmt.Println("Error accepting connection: " + connectionError.Error())
		} else {
			go server.handleRequest(connection)
		}
	}
}

func (server *Server) handleRequest(connection net.Conn) {
	buffer := make([]byte, 1024)

	_, readError := connection.Read(buffer)

	if readError != nil {
		//if there is a error print the error and write the error to the connection
		fmt.Println("Error reading: ", readError.Error())
		connection.Write([]byte(readError.Error()))
	} else {

		connection.Write([]byte("Message Received\n"))
	}

	connection.Close()
}
