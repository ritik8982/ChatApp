package main

import (
	"fmt"
	"net/http"
	"backend/pkg/websocket"
)

func serveWs(pool *websocket.Pool,w http.ResponseWriter,r *http.Request){

	fmt.Println("websocket end point reached");

	conn, err := websocket.Upgrade(w,r);

	if err != nil{
		fmt.Fprintln(w,"%+v\n",err);
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes(){
	pool := websocket.NewPool(); 
	go pool.Start()	

	http.HandleFunc("/ws",func(w http.ResponseWriter,r *http.Request){
		serveWs(pool,w,r);
	})
}
func main() {
	fmt.Println("ritik's full stack chat project")
	setupRoutes()
	http.ListenAndServe(":9000",nil);
}