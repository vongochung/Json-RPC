package main

import (
	"Json-RPC/Server/config"
	"Json-RPC/Server/user"
	"Json-RPC/Server/event"
	"Json-RPC/Server/chat"
	"code.google.com/p/go.net/websocket"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"log"

)

// Serve all methods to support via Json RPC
func jsonrpcHandler(ws *websocket.Conn) {
	jsonrpc.ServeConn(ws)
}

func main() {
	// Init config (db, app...)
	bin.InitProcs()
	bin.LoadConfiguration("./config.json")

	db := bin.DB()
	defer db.Close()

	// Regist methods json RPC
	rpc.Register(&user.UserMethod{db})
	rpc.Register(&event.EventMethod{db})

	// Handler via socket(url is channel)
	http.Handle("/jsonrpc", websocket.Handler(jsonrpcHandler))
	http.Handle("/chat", websocket.Handler(chat.ChatHandler))

	// Run app
	log.Printf("App listening at port: %d", bin.Config.Int("APP_PORT"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
