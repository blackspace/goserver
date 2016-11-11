package misc

import (
	"strings"
	"github.com/blackspace/goserver/context"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("echo", func(c *context.ClientContext, args ...string) (result string) {
		return strings.Join(args," ")
	},""))
}

