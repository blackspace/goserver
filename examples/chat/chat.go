package main

import (
	"fmt"
	"time"
	"strconv"
	"github.com/blackspace/goserver"
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/context"
	. "github.com/blackspace/goserver/command"
	_ "github.com/blackspace/goserver/command/client_info"
	_ "github.com/blackspace/goserver/command/help"
)


func init() {
	RegistCommand(NewCommand("broadcast",func (clnt *client.Client,args ...string) string {
		if len(args)!=0 {
			m :="MESSAGE:" + fmt.Sprint(clnt.Id)+" "+args[0]
			for _,cl:=range clnt.ServerContext.OnlineClient() {
				client.NewClient(cl).WriteLine(m)
			}
		}

		return ""
	},""))
	RegistCommand(NewCommand("say_to",func (clnt *client.Client,args ...string) string {
		if len(args)!=3 {
			return "ERROR:The command require 3 argments.\nLike as 'say_to id 1 ddddd'"
		} else if args[0]!="name"&&args[0]!="id" {
			return "ERROR:The command require a subcommand:name or id.\nLike as 'say_to id 1 ddddd'"
		}

		m := "MESSAGE:" + fmt.Sprint(clnt.Id) + " " + args[2]


		var rc *context.ClientContext =nil

		switch args[0] {
		case "name":
			rc= clnt.ServerContext.FindClientByName(args[1])
		case "id":
			id,err :=strconv.ParseInt(args[1],0,64)
			if err!=nil {
				return "ERROR:The id subcommand  need a integer64"
			} else {
				rc= clnt.ServerContext.FindClientById(id)
			}
		}

		if rc!=nil {
			client.NewClient(rc).WriteLine(m)
		} else {
			return "ERROR:Can't find the client"
		}

		return ""
	},""))
}

func main() {
	s:=goserver.NewServer()

	s.Start("127.0.0.1","8058")
	defer s.Stop()

	time.Sleep(time.Hour)
}

