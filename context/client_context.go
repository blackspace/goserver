package context


import (
	"net"
	"bufio"
	"sync/atomic"
)

type ClientContext struct {
	Connect         net.Conn
	Reader          * bufio.Reader
	Writer          * bufio.Writer
	Id              int64
	IsClosed        bool
	NeedPrompt      bool
	Prompt          string
	UserName        string
	IsCharacterMode bool
}



var LastId int64=0

func GetNewId() int64 {
	return 	atomic.AddInt64(&LastId,1)
}

func NewClientContext() (*ClientContext) {
	return &ClientContext{Id:GetNewId() }
}

func (c *ClientContext)SetConn(conn net.Conn) {
	c.Connect =conn
	c.Reader=bufio.NewReader(c.Connect)
	c.Writer=bufio.NewWriter(c.Connect)
}








