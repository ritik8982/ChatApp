package main

import (
	"fmt"
	"net/http"
	"backend/pkg/websocket"
)
// pool jo hai wo ek object hai Pool struct ka 
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
	client.Read()  // Read ek struct ki method hai 
}

func setupRoutes(){
	pool := websocket.NewPool(); //NewPool is the function in ur websocket package, that will acts as constructor and return a new Pool's object
	go pool.Start()	// starting as the parallel process(go routine)(thread)

	//this is websocket endPoint, jab is end point ko hit karoge to ye serveWs call hoga
	http.HandleFunc("/ws",func(w http.ResponseWriter,r *http.Request){
		serveWs(pool,w,r);
	})
}
func main() {
	fmt.Println("ritik's full stack chat project")
	setupRoutes()
	http.ListenAndServe(":9000",nil);
}