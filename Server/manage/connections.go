package manager

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
	"Json-RPC/Server/helper"
)

type Connection struct {
	// ID of connection
    ID string

    // The websocket connection.
    Ws *websocket.Conn

    // Buffered channel of outbound messages.
    Send chan string

    // Channel
    ChannelID int
}

// RegistConn regist conn into channel
func (c *Channel) RegistConn(conn *Connection){
    c.Connections[conn] = true
} 

// RegistConn regist conn into channel
func (c *Channel) UnRegistConn(conn *Connection){
    delete(c.Connections, conn)
} 

// GenConnID return ID of websocket conn(
func GenConnID(req *http.Request)(string, error){
    ip, err := helper.ParseIP(req.RemoteAddr)
    if err != nil {
        return "", err
    }
    id := helper.NewTokenHash(req.UserAgent()+ip)
    return id, nil
}


func (c *Connection) ReceiveMsg() {
    for {
        var message string
        err := websocket.Message.Receive(c.Ws, &message)
        if err != nil {
            break
        }
        //c.Send <- message
        Manage.Channels[0].Broadcast <- message
    }
    c.Ws.Close()
}

func (c *Connection) SendMsg() {
    for message := range c.Send {
        err := websocket.Message.Send(c.Ws, message)
        if err != nil {
            break
        }
    }
    c.Ws.Close()
}