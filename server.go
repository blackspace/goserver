package goserver

import (
	"net"
	"log"
	"github.com/blackspace/goserver/context"
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/action"
)

type Server struct {
	 *context.ServerContext
}

func NewServer() *Server{
	return &Server{ServerContext:context.NewServerContext()}
}


func (s *Server)_DoWork(conn net.Conn)() {
	cc := context.NewClientContext()

	cc.SetConn(conn)

	s.AddClientContext(cc)
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


func (s *Server)_BeginListen() {
	if s.Listener ==nil {
		l, err := net.Listen("tcp", `127.0.0.1:8058`)

		if err != nil {
			panic(err)
		}
		s.Listener = l
	}
	log.Println("Listenning on 127.0.0.1:8058")
}

func (s *Server)_AcceptClientConnect() (net.Conn,error) {
	return s.Listener.Accept()
}


func (s *Server)_ServerClose() {
	s.Listener.Close()
}

func (s *Server)Start() {
	s._BeginListen()
	go s.ClearClosedClient()

	go func() {
		for {
			conn,err:= s._AcceptClientConnect()
			log.Println("A connection to client has been established")
			if err != nil {
				panic(err)
			}
			go s._DoWork(conn)
		}

	}()
}

func (s *Server)Stop() {
	s._ServerClose()
}