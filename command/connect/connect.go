package connect

import (
	"github.com/blackspace/goserver/client"
	. "github.com/blackspace/goserver/command"
)

func init() {
	Commands.RegistCommand(NewCommand("quit", func(clt *client.Client, args ...string) string {
		clt.CloseConnect()
		return ""
	},""))
}
