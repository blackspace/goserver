package command

import (
	"goserver/context"
	"goserver/client"
)

func init() {
	_AddCommand(_NewCommand("quit", func(c *context.ClientContext, args ...string) string {
		client.CloseClientConnect(c)
		return ""
	},""))
}
