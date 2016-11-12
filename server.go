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
	cc := context.NewClientContext(s.ServerContext,conn)
	s.AddClientContext(cc)

	clt := client.NewClient(cc)
	defer clt.CloseConnect()

	buf := make([]byte, 0, 1<<10)

	S:
	for {
		buf = buf[0:0]

		T:
		for {
			c, err := clt.ReadByte()

			if err != nil {
				break S
			}

			buf = append(buf, c)

			if action.IsBinary(buf) {
				fa := action.FindActionForBinary(buf)
				if fa != nil {
					if fa(clt, buf) {
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
					if la(clt, line) {
						break T
					} else {
						break S
					}
				} else {
					clt.SendResult("The line is invalid line\r\n")
					break T
				}
			}
		}
	}
}


func (s *Server)Start(addr string,port string) {
	if s.GetListener() ==nil {
		l, err := net.Listen("tcp", addr+":"+port)

		if err != nil {
			panic(err)
		}
		s.SetListener(l)

		log.Println("Listenning on "+addr+":"+port)

		go s.ClearClosedClient()

		go func() {
			for {
				conn,err:= s.GetListener().Accept()
				if err != nil {
					panic(err)
				}
				go s._DoWork(conn)
			}

		}()
	} else {
		panic("This server has already been start")
	}
}

func (s *Server)Stop() {
	s.GetListener().Close()
}