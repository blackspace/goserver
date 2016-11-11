package main

import (
	"net"
	"sync"
	"time"
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/context"
	"github.com/blackspace/goserver/action"
	"github.com/blackspace/goserver"
)


func IsSocksV4Instruction(buf []byte) bool {
	return len(buf) >= 8 && buf[0] == 0x04 && buf[1] == 0x01 &&  buf[len(buf) - 1] == 0x00
}


func DoSocksV4(clnt client.Client , buf []byte)  bool {
	port := buf[2:4]
	ip := buf[4:8]

	result := append(append([]byte{0x00, 0x5A}, port...), ip...)

	clnt.Write(result)

	rc, _ := net.DialTCP("tcp", nil, &net.TCPAddr{ip, int(port[1]) + int(port[0]) << 8, ""})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			rbuf := make([]byte, 1024)

			n, err := clnt.Read(rbuf)

			if err != nil {
				break
			}
			rc.Write(rbuf[:n])

		}
		wg.Done()
	}()

	go func() {
		for {
			rbuf := make([]byte, 1024)

			n, err := rc.Read(rbuf)
			if err != nil {
				break
			}

			clnt.Write(rbuf[:n])

		}
		clnt.CloseConnect()
		rc.Close()
		wg.Done()
	}()

	wg.Wait()
	return false
}

func init() {
	action.BinaryActions.AddAction(func(buf []byte) bool { return IsSocksV4Instruction(buf)}, DoSocksV4)
}


func main() {
	s:=goserver.NewServer()

	s.Start("127.0.0.1","8058")
	defer s.Stop()

	time.Sleep(time.Hour)
}
