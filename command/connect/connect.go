package connect

import (
	"github.com/blackspace/goserver/client"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("quit", func(client *client.Client, args ...string) string {
		client.CloseClientConnect()
		return ""
	},""))
}
