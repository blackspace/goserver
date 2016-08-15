package command

import (
	"strings"
	"goserver/context"
)

func init() {
	_AddCommand(_NewCommand("echo", func(c *context.ClientContext, args ...string) (result string) {
		return strings.Join(args," ")
	},""))
}

