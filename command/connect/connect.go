package connect

import (
	"github.com/blackspace/goserver/client"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("quit", func(cln *client.Client, args ...string) string {
		cln.CloseConnect()
		return ""
	},""))
}
