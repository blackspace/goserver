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

func (c *Client)CloseConnect(){
	c.CloseConnect()
	c.IsClosed =true
}

func (c *Client)ReadLine()(string , error) {
	l,_,err:=c.Reader.ReadLine()

	return string(l), err
}

func (c *Client)SendResult(r string) {
	c.Writer.WriteString(r)
	c.Writer.Flush()
}

func (c *Client)PrintPrompt() {
	c.Writer.WriteString(c.Prompt)
	c.Writer.Flush()
}

func (c *Client)WriteLine(s string) {
	c.Writer.WriteString(s+"\n")
	c.Writer.Flush()
}

func (c *Client)Write(b []byte) (int,error){
	n,err := c.Writer.Write(b)
	c.Writer.Flush()
	return n,err
}

func (c *Client)Read(b []byte)(int,error){
	return c.Reader.Read(b)
}

func (c *Client)ReadByte() (byte,error) {
	return c.Reader.ReadByte()
}
