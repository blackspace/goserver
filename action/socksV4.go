package action

import (
	"goserver/context"
	"net"
	"sync"
	"goserver/client"
)


func SocksV4(cc *context.ClientContext , buf []byte)  bool {
	port := buf[2:4]
	ip := buf[4:8]

	result := append(append([]byte{0x00, 0x5A}, port...), ip...)

	client.ClientConnectWrite(cc,result)

	rc, _ := net.DialTCP("tcp", nil, &net.TCPAddr{ip, int(port[1]) + int(port[0]) << 8, ""})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			rbuf := make([]byte, 1024)

			n, err := cc.Connect.Read(rbuf)

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

			cc.Connect.Write(rbuf[:n])

		}
		client.CloseClientConnect(cc)
		rc.Close()
		wg.Done()
	}()

	wg.Wait()
	return true
}