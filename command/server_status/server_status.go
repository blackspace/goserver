package server_status

import (
	"time"
	"strconv"
	"github.com/blackspace/goserver/context"
	"fmt"
	"github.com/blackspace/goserver/client"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("now",func (c * context.ClientContext,args ...string) string{
		return time.Now().String()
	},""))
	RegistCommand(NewCommand("client_count",func (c *context.ClientContext,args ...string) string {
		return strconv.Itoa(context.ClientCount())
	},""))
	RegistCommand(NewCommand("who_online", func (c *context.ClientContext,args ...string) string {
		for _,cl:=range context.OnlineClient() {
			client.ClientConnectWriteLine(c,fmt.Sprint(cl.Id)+" "+cl.UserName)
		}
		return ""
	},""))
}
