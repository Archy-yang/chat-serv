package handlers

func broadcaster() {
	var clients = map[Client]bool{}

	for {
		select {
		case message := <-messages:
			for client := range clients {
				go func(c Client, msg Message) {
					c.Msg <- message
				}(client, message)
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Msg)
		}
	}
}
