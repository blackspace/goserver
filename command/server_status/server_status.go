package server_status

import (
	"time"
	"strconv"
	"fmt"
	"github.com/blackspace/goserver/client"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("now",func (clnt *client.Client,args ...string) string{
		return time.Now().String()
	},""))
	RegistCommand(NewCommand("client_count",func (clnt *client.Client,args ...string) string {
		return strconv.Itoa(clnt.ServerContext.ClientCount())
	},""))
	RegistCommand(NewCommand("who_online", func (clnt *client.Client,args ...string) string {
		for _,cl:=range clnt.ServerContext.OnlineClient() {
			clnt.WriteLine(fmt.Sprint(cl.Id)+" "+cl.UserName)
		}
		return ""
	},""))
}
