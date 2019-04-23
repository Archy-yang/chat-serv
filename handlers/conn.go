package handlers

import (
	"net"
	"bufio"
	"strings"
	"crypto/md5"
)

type Client struct {
	ID string
	Name string
	Msg chan Message
	Conn net.Conn
	Done chan int
}
func Conn(conn net.Conn) {
	md := md5.New()
	md.Write([]byte(conn.RemoteAddr().String()))
	client := Client{
		Msg: make(chan Message),
		Conn: conn,
		ID: string(md.Sum(nil)),
	}

	defer client.Conn.Close()
	go client.sendMsg()

	input := bufio.NewScanner(client.Conn)
	conn.Write([]byte("what's your name:\n"))
	if input.Scan() {
		client.Name = input.Text()
		client.Msg <- Message{client.ID, "welcome to the wold! " + client.Name}
	}
	messages <- Message{client.ID, client.Name + " entering"}
	entering <- client

	for input.Scan() {
		text := input.Text()
		if strings.ToUpper(text) == "EXIT"  {
			client.Msg <- Message{client.ID, "bye-bye"}
			break
		}
		messages <- Message{client.ID, client.Name +": " + text}
	}

	leaving <- client
	messages <- Message{client.ID, client.Name + " leaving"}
}

func (c Client) sendMsg() {
	for msg := range c.Msg {
		if msg.ID != c.ID {
			c.Conn.Write([]byte(msg.Msg+"\n"))
		}
	}
}
