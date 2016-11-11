package client_info

import (
	"fmt"
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

func init() {
	RegistCommand(NewCommand("user_name",func (clnt *client.Client,args ...string) string {
		if len(args)!=1 {
			return "ERROR:The Command need one argment."
		}
		clnt.UserName=args[0]
		return ""
	},""))
	RegistCommand(NewCommand("whoami",func (clnt *client.Client,args ...string) string {
		return clnt.UserName
	},""))
	RegistCommand(NewCommand("id",func (clnt *client.Client,args ...string) string {
		return fmt.Sprint(clnt.Id)
	},""))
}
