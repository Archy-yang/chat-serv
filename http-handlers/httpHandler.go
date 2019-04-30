package http_handlers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func Start() {
	http.HandleFunc("/push", pushHandler)
	http.HandleFunc("/open", wsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	err := http.ListenAndServe(":9005", nil)
	//err := http.ListenAndServeTLS(":9005", "/Users/didi/https/server.crt", "/Users/didi/https/server.key", nil)
	fmt.Println(err)
}

func pushHandler (w http.ResponseWriter, r *http.Request) {
	if push, ok := w.(http.Pusher); ok {
		fmt.Println("in")
		err := push.Push("/Users/didi/project/go/mine/chat-serv/src/chat-serv/http-handlers/a.js", nil);
		fmt.Println(err)
	}
	w.Write([]byte("hello"))
}

var upgrade = websocket.Upgrader{}
func wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

