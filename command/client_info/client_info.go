package client_info

import (
	"fmt"
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

func init() {
	RegistCommand(NewCommand("user_name",func (client *client.Client,args ...string) string {
		if len(args)!=1 {
			return "ERROR:The Command need one argment."
		}
		client.UserName=args[0]
		return ""
	},""))
	RegistCommand(NewCommand("whoami",func (client *client.Client,args ...string) string {
		return client.UserName
	},""))
	RegistCommand(NewCommand("id",func (client *client.Client,args ...string) string {
		return fmt.Sprint(client.Id)
	},""))
}
