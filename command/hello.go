package command

import (
	"goserver/context"
)

func init() {
	_AddCommand(_NewCommand("hello", func(c *context.ClientContext, args ...string) string {
		return "hello,I am gochat-server"
	},"say hello"))
}
