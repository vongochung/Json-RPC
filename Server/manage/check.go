package manager

import (
)

func DeleteExistConn(ID string) {
    for c := range Manage.Channels[0].Connections{
        if c.ID == ID {
            Manage.Channels[0].UnRegistConn(c)
            return
        }
    }
}