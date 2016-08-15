package server

import (
	"net"
	"log"
	"goserver/context"
)

func ServerBeginListen() {
	if context.ServerContext.Listener ==nil {
		l, err := net.Listen("tcp", `127.0.0.1:8058`)

		if err != nil {
			log.Fatalln(err)
		}
		context.ServerContext.Listener = l
	}
}

func ServerAcceptClientConnect() (net.Conn,error) {
	return context.ServerContext.Listener.Accept()
}


func ServerClose() {
	context.ServerContext.Listener.Close()
}

