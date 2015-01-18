package main

import (
	"NeightBourSocial/Server/config"
	"NeightBourSocial/Server/user"
	"code.google.com/p/go.net/websocket"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func jsonrpcHandler(ws *websocket.Conn) {
	jsonrpc.ServeConn(ws)
}

func main() {
	bin.InitProcs()
	bin.LoadConfiguration("./config.json")

	db := bin.DB()
	defer db.Close()
	rpc.Register(&user.UserMethod{db})
	http.Handle("/jsonrpc", websocket.Handler(jsonrpcHandler))
	//http.Handle("/notify", websocket.Handler(notifyHandler))
	//http.Handle("/push", websocket.Handler(pushHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
