package server_status

import (
	"time"
	"strconv"
	"fmt"
	"github.com/blackspace/goserver/client"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("now",func (cln *client.Client,args ...string) string{
		return time.Now().String()
	},""))
	RegistCommand(NewCommand("client_count",func (cln *client.Client,args ...string) string {
		return strconv.Itoa(cln.ServerContext.ClientCount())
	},""))
	RegistCommand(NewCommand("who_online", func (cln *client.Client,args ...string) string {
		for _,cl:=range cln.ServerContext.OnlineClient() {
			cln.WriteLine(fmt.Sprint(cl.Id)+" "+cl.UserName)
		}
		return ""
	},""))
}
