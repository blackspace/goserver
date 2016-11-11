package misc

import (
	"strings"
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

func init() {
	RegistCommand(NewCommand("echo", func(cln *client.Client, args ...string) (result string) {
		return strings.Join(args," ")
	},""))
}

