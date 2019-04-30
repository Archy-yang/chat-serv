package main

import (
    "chat-serv/http-handlers"
    "chat-serv/handlers"
)

func main(){
    go handlers.Server()
    http_handlers.Start()
}
