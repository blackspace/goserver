package server_status

import (
	"time"
	"strconv"
	"fmt"
	"github.com/blackspace/goserver/client"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("now",func (client *client.Client,args ...string) string{
		return time.Now().String()
	},""))
	RegistCommand(NewCommand("client_count",func (client *client.Client,args ...string) string {
		return strconv.Itoa(client.ServerContext.ClientCount())
	},""))
	RegistCommand(NewCommand("who_online", func (client *client.Client,args ...string) string {
		for _,cl:=range client.ServerContext.OnlineClient() {
			client.ClientConnectWriteLine(fmt.Sprint(cl.Id)+" "+cl.UserName)
		}
		return ""
	},""))
}
