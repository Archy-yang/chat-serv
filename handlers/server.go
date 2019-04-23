package handlers

import (
	"net"
	"fmt"
)

type Message struct {
	ID string
	Msg string
}

var (
	messages = make(chan Message)
	leaving = make(chan Client)
	entering = make(chan Client)
)

func Server() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":9001")
	if err != nil {
		panic(err.Error())
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err.Error())
	}
	defer listener.Close()

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go func(conn net.Conn) {
			Conn(conn)
		} (conn)
	}
}
