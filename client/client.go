package client

import (
	"github.com/blackspace/goserver/context"
)


type Client struct {
	*context.ClientContext
}

func NewClient(cc *context.ClientContext) *Client{
	return &Client{cc}
}

func (c *Client)CloseClientConnect(){
	c.CloseConnect()
	c.IsClosed =true
}

func (c *Client)ClientConnectReadLine()(string , error) {
	l,_,err:=c.Reader.ReadLine()

	return string(l), err
}

func (c *Client)ClientConnectSendResult(r string) {
	c.Writer.WriteString(r)
	c.Writer.Flush()
}

func (c *Client)ClientConnectPrintPrompt() {
	c.Writer.WriteString(c.Prompt)
	c.Writer.Flush()
}

func (c *Client)ClientConnectWriteLine(s string) {
	c.Writer.WriteString(s+"\n")
	c.Writer.Flush()
}

func (c *Client)ClientConnectWrite(b []byte) (int,error){
	n,err := c.Writer.Write(b)
	c.Writer.Flush()
	return n,err
}

func (c *Client)ClientConnectRead(b []byte)(int,error){
	return c.Reader.Read(b)
}

func (c *Client)ClientReadByte() (byte,error) {
	return c.Reader.ReadByte()
}
