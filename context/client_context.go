package context


import (
	"net"
	"bufio"
	"sync/atomic"
)

type ClientContext struct {
	_connect      net.Conn
	Reader        * bufio.Reader
	Writer        * bufio.Writer
	Id            int64
	IsClosed      bool
	NeedPrompt    bool
	Prompt        string
	UserName      string
	ServerContext *ServerContext
}

var LastId int64=0

func GetNewId() int64 {
	return 	atomic.AddInt64(&LastId,1)
}

func NewClientContext(s *ServerContext,conn net.Conn) (*ClientContext) {
	c:=&ClientContext{Id:GetNewId(),ServerContext:s, _connect:conn}
	c.Reader=bufio.NewReader(c._connect)
	c.Writer=bufio.NewWriter(c._connect)
	return c
}

func (c *ClientContext)CloseConnect() {
	c._connect.Close()
}







