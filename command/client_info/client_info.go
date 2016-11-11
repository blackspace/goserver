package client_info

import (
	"fmt"
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

func init() {
	RegistCommand(NewCommand("user_name",func (cln *client.Client,args ...string) string {
		if len(args)!=1 {
			return "ERROR:The Command need one argment."
		}
		cln.UserName=args[0]
		return ""
	},""))
	RegistCommand(NewCommand("whoami",func (cln *client.Client,args ...string) string {
		return cln.UserName
	},""))
	RegistCommand(NewCommand("id",func (cln *client.Client,args ...string) string {
		return fmt.Sprint(cln.Id)
	},""))
}
