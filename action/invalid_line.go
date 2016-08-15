package action

import (
	"goserver/context"
	"goserver/client"
	"log"
)

func DoInvalidLine(cc * context.ClientContext,buf []byte) bool {
	log.Print(string(buf))
	client.ClientConnectWriteLine(cc,"The line is invalid")
return false
}
