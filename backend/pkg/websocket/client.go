package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct{
	ID string
	Conn *websocket.Conn		//gorilla websocket is the connection here
	Pool *Pool
	mu  sync.Mutex
}

type Message struct{
	Type int `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read(){
	defer func (){			// the function which don't have any name should call themselves
		c.Pool.UnRegister <- c		//unregister kardo is client ko
		c.Conn.Close()				// close the connection
	}()

	for {
		messageType,p,err := c.Conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}
		message := Message{Type: messageType,Body: string(p)}
		c.Pool.Broadcast <- message		// this is how you write on the channel
		fmt.Printf("message received :%+v\n",message)

	}
}