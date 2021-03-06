package client_info

import (
	"fmt"
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

func init() {
	Commands.RegistCommand("user_name",func (clt *client.Client,args ...string) string {
		if len(args)!=1 {
			return "ERROR:The Command need one argment."
		}
		clt.UserName=args[0]
		return ""
	},"")
	Commands.RegistCommand("whoami",func (clt *client.Client,args ...string) string {
		return clt.UserName
	},"")
	Commands.RegistCommand("id",func (clt *client.Client,args ...string) string {
		return fmt.Sprint(clt.Id)
	},"")
}
