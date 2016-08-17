package main

import (
	"log"
	"net"
	"goserver/context"
	"goserver/server"
	"goserver/client"
	"goserver/action"
)

func doWork(conn net.Conn)() {
	cc := context.NewClientContext()

	cc.SetConn(conn)

	context.ServerContextAddClientContext(cc)

	defer client.CloseClientConnect(cc)

	buf := make([]byte, 0, 10240)

	L:
	for {
		buf = buf[0:0]

		M:
		for {
			c, err := client.ClientReadByte(cc)


			if err != nil {
				log.Println(err)
				break L
			}

			buf = append(buf, c)

			a :=action.FindAction(buf)

			if a !=nil {
				if a(cc,buf) {
					break M
				} else {
					break L
				}
			} else {
				continue
			}

		}
	}
	log.Println("A goroutine has finished")
}


func init() {
	server.ServerBeginListen()
	go context.ServerContextClearClosedClient()
}

func main() {
	defer server.ServerClose()

	for {
		conn,err:= server.ServerAcceptClientConnect()
		if err != nil {
			log.Fatalln(err)
		}
		go doWork(conn)
	}

	log.Println("goserver exited")
}
