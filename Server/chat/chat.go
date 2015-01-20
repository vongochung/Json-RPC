package chat

import (
    "code.google.com/p/go.net/websocket"
    "fmt"
    "Json-RPC/Server/helper"
    "Json-RPC/Server/manage"
)

func ChatHandler(ws *websocket.Conn) {
    // Create connection
    cnn := &manager.Connection{Send: make(chan string, 1024), Ws: ws}
    id, err := manager.GenConnID(ws.Request())
    helper.PanicIf(err)
    cnn.ID = id

    // Regist connection
    manager.DeleteExistConn(cnn.ID)
    manager.Manage.Channels[0].RegistConn(cnn)

    fmt.Println(len(manager.Manage.Channels[0].Connections))
    for {
        go cnn.ReceiveMsg()
        go cnn.SendMsg()

        select{
        case m := <- manager.Manage.Channels[0].Broadcast:
            for c := range manager.Manage.Channels[0].Connections {
                c.Send <- m                
            }
        }
    }
    
}