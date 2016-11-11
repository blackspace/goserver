package connect

import (
	"github.com/blackspace/goserver/context"
	"github.com/blackspace/goserver/client"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("quit", func(c *context.ClientContext, args ...string) string {
		client.CloseClientConnect(c)
		return ""
	},""))
}
