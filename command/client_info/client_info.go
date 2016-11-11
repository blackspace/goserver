package client_info

import (
	"fmt"
	"github.com/blackspace/goserver/context"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("user_name",func (c * context.ClientContext,args ...string) string {
		if len(args)!=1 {
			return "ERROR:The Command need one argment."
		}
		c.UserName=args[0]
		return ""
	},""))
	RegistCommand(NewCommand("whoami",func (c * context.ClientContext,args ...string) string {
		return c.UserName
	},""))
	RegistCommand(NewCommand("id",func (c * context.ClientContext,args ...string) string {
		return fmt.Sprint(c.Id)
	},""))
}
