package client

import (
	"goserver/context"
)


///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func CloseClientConnect(c *context.ClientContext){
	c.CloseConnect()
	c.IsClosed =true
}

func ClientConnectReadLine(c *context.ClientContext)(string , error) {
	l,_,err:=c.Reader.ReadLine()

	return string(l), err
}

func ClientConnectSendResult(c *context.ClientContext,r string) {
	c.Writer.WriteString(r)
	c.Writer.Flush()
}

func ClientConnectPrintPrompt(c *context.ClientContext) {
	c.Writer.WriteString(c.Prompt)
	c.Writer.Flush()
}

func ClientConnectWriteLine(c *context.ClientContext,s string) {
	c.Writer.WriteString(s+"\n")
	c.Writer.Flush()
}

func ClientConnectWrite(c *context.ClientContext,b []byte) (int,error){
	n,err := c.Writer.Write(b)
	c.Writer.Flush()
	return n,err
}

func ClientConnectRead(c *context.ClientContext,b []byte)(int,error){
	return c.Reader.Read(b)
}

func ClientReadByte(c *context.ClientContext) (byte,error) {
	return c.Reader.ReadByte()
}
