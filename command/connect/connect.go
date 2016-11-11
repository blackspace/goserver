package connect

import (
	"github.com/blackspace/goserver/client"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("quit", func(clnt *client.Client, args ...string) string {
		clnt.CloseConnect()
		return ""
	},""))
}
