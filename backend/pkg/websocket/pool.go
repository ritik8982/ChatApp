package websocket

import "fmt"

type Pool struct {
	Register   chan *Client //chan is channel
	UnRegister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message //broadcast will broadcast the message acroos the channel

}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client), // making the channel
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	// we will have multiple concurrent collection, this is how connection pool works , because you can have multiple peoples in ur chat room  to do this we need to create connection pool
	// yaha per client kon hoga browser jitne browser khologe utne clients honge, yehi chat room create karega
	for {
		select {
		case client := <-pool.Register:		// as soon as the new user(browser) joins we tell other user that a new user has joined
			pool.Clients[client] = true
			fmt.Println("size of the connection pool = ",len(pool.Clients))
			for client,_ := range pool.Clients{
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1,Body: "New User joined..."})
			}
			break
		case client := <-pool.UnRegister:
			delete(pool.Clients,client)
			fmt.Println("size of the connection pool = ",len(pool.Clients))
			for client,_ := range pool.Clients{
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1,Body: "User Diconnected..."})
			}
			break;
		case message := <-pool.Broadcast:
			fmt.Println("sending message to all clients in the pool") 
			for client,_ := range pool.Clients{
				if err := client.Conn.WriteJSON(message); err != nil{
					fmt.Println(err);
					return
				}
			}
		}
	}
}