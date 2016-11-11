package goserver

import (
	"net"
	"log"
	"github.com/blackspace/goserver/context"
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/action"
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
				fa := action.FindActionForFlag(buf)
				if fa != nil {
					if fa(cc, buf) {
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


func ServerBeginListen() {
	if context.ServerContext.Listener ==nil {
		l, err := net.Listen("tcp", `127.0.0.1:8058`)

		if err != nil {
			panic(err)
		}
		context.ServerContext.Listener = l
	}
	log.Println("Listenning on 127.0.0.1:8058")
}

func ServerAcceptClientConnect() (net.Conn,error) {
	return context.ServerContext.Listener.Accept()
}


func ServerClose() {
	context.ServerContext.Listener.Close()
}

func Start() {
	ServerBeginListen()
	go context.ServerContextClearClosedClient()

	go func() {
		for {
			conn,err:= ServerAcceptClientConnect()
			log.Println("A connection to client has been established")
			if err != nil {
				panic(err)
			}
			go doWork(conn)
		}

	}()
}

func Stop() {
	ServerClose()
}


