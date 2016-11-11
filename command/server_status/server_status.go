package server_status

import (
	"time"
	"strconv"
	"fmt"
	"github.com/blackspace/goserver/client"
	. "github.com/blackspace/goserver/command"
)

func init() {
	Commands.RegistCommand("now",func (clt *client.Client,args ...string) string{
		return time.Now().String()
	},"")
	Commands.RegistCommand("client_count",func (clt *client.Client,args ...string) string {
		return strconv.Itoa(clt.ServerContext.ClientCount())
	},"")
	Commands.RegistCommand("who_online", func (clt *client.Client,args ...string) string {
		for _,cl:=range clt.ServerContext.OnlineClient() {
			clt.WriteLine(fmt.Sprint(cl.Id)+" "+cl.UserName)
		}
		return ""
	},"")
}
