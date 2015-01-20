package manager

import (
)

var Manage ManageSocket

type ManageSocket struct{
	Channels []Channel
}


type Channel struct{
	ID int
	Name string
	Broadcast chan string
	Emit chan Emit
	Connections map[*Connection]bool
}

type Emit struct{
	From string
	To	[]string
	Message string
}



func Init(){
	//Manage = bin.Redis().HMSet("channel1", "Ok")
	Manage =  ManageSocket{}
	channel := Channel{
        ID : len(Manage.Channels) + 1,
        Name: "chat",
        Broadcast: make(chan string),
        Emit: make(chan Emit),
        Connections: make(map[*Connection]bool),
    }

    Manage.Channels= append(Manage.Channels,channel)

}
