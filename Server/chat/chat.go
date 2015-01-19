package chat

import (
    "code.google.com/p/go.net/websocket"
    "fmt"
)

type Connection struct {
    // The websocket connection.
    ws *websocket.Conn

    // Buffered channel of outbound messages.
    send chan string
}

type ManageSocket struct{
    Connections map[*Connection]bool
    Broadcast chan string
}

var manager = ManageSocket{
    Broadcast: make(chan string),
    Connections: make(map[*Connection]bool),
}

func (c *Connection) Receive() {
    for {
        var message string
        err := websocket.Message.Receive(c.ws, &message)
        if err != nil {
            break
        }
        //c.send <- message
        manager.Broadcast <- message
    }
    c.ws.Close()
}

func (c *Connection) Send() {
    for message := range c.send {
        err := websocket.Message.Send(c.ws, message)
        if err != nil {
            break
        }
    }
    c.ws.Close()
}

func ChatHandler(ws *websocket.Conn) {
    c := &Connection{send: make(chan string, 1024), ws: ws}
    manager.Connections[c] = true
    fmt.Println(len( manager.Connections))
    for {
        go c.Receive()
        go c.Send()

        select{
        case m := <- manager.Broadcast:
            for c := range manager.Connections {
                c.send <- m                
            }
        }
    }
    
}