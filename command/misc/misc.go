package misc

import (
	"strings"
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

func init() {
	Commands.RegistCommand(NewCommand("echo", func(clt *client.Client, args ...string) (result string) {
		return strings.Join(args," ")
	},""))
}

