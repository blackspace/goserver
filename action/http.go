package action

import (
	"goserver/context"
	"goserver/client"
	"log"
)

func DoHttpMethod(cc * context.ClientContext,buf []byte) bool {
	client.ClientConnectWrite(cc,buf)
	for {
		l,err:=client.ClientConnectReadLine(cc)

		if err!=nil {
			log.Println(err)
			break
		}

		if len(l)==0 {
			break
		}

		client.ClientConnectWriteLine(cc,l)
	}

	client.ClientConnectWriteLine(cc,"")
	client.ClientConnectWriteLine(cc,"")

	client.ClientConnectWriteLine(cc,"hello http")


	return true
}
