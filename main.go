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

	S:
	for {
		buf = buf[0:0]

		T:
		for {
			c, err := client.ClientReadByte(cc)

			if err != nil {
				log.Println(err)
				break S
			}

			buf = append(buf, c)

			if action.IsFlag(buf) {
				fa :=action.FindActionForFlag(buf)
				if fa !=nil {
					if fa(cc,buf) {
						break T
					} else {
						break S
					}
				}
			}

			if action.IsLine(buf) {
				line := string(buf[:len(buf) - 2])
				la := action.FindActionForLine(line)
				if la != nil {
					if la(cc, line) {
						break T
					} else {
						break S
					}
				} else {
					client.ClientConnectSendResult(cc, "The line is invalid line\r\n")
					break T
				}
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
