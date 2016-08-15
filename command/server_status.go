package command

import (
	"time"
	"strconv"
	"goserver/context"
	"fmt"
	"goserver/client"
)

func init() {
	_AddCommand(_NewCommand("now",func (c * context.ClientContext,args ...string) string{
		return time.Now().String()
	},""))
	_AddCommand(_NewCommand("client_count",func (c *context.ClientContext,args ...string) string {
		return strconv.Itoa(context.ClientCount())
	},""))
	_AddCommand(_NewCommand("who_online", func (c *context.ClientContext,args ...string) string {
		for _,cl:=range context.OnlineClient() {
			client.ClientConnectWriteLine(c,fmt.Sprint(cl.Id)+" "+cl.UserName)
		}
		return ""
	},""))
}
