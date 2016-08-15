package command

import (
	"fmt"
	"goserver/context"
)

func init() {
	_AddCommand(_NewCommand("user_name",func (c * context.ClientContext,args ...string) string {
		if len(args)!=1 {
			return "ERROR:The Command need one argment."
		}
		c.UserName=args[0]
		return ""
	},""))
	_AddCommand(_NewCommand("whoami",func (c * context.ClientContext,args ...string) string {
		return c.UserName
	},""))
	_AddCommand(_NewCommand("id",func (c * context.ClientContext,args ...string) string {
		return fmt.Sprint(c.Id)
	},""))
}
