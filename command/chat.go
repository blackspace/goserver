package command

import (
	"fmt"
	"goserver/context"
	"strconv"
	"goserver/client"
)

func init() {
	_AddCommand(_NewCommand("broadcast",func (c *context.ClientContext,args ...string) string {
		m :="MESSAGE:" + fmt.Sprint(c.Id)+" "+args[0]
		for _,cl:=range context.OnlineClient() {
			client.ClientConnectWriteLine(cl,m)
		}
		return ""
	},""))
	_AddCommand(_NewCommand("say_to",func (c *context.ClientContext,args ...string) string {
		if len(args)!=3 {
			return "ERROR:The command require 3 argments.\nLike as 'say_to id 1 ddddd'"
		} else if args[0]!="name"&&args[0]!="id" {
			return "ERROR:The command require a subcommand:name or id.\nLike as 'say_to id 1 ddddd'"
		}

		m := "MESSAGE:" + fmt.Sprint(c.Id) + " " + args[2]


		var rc *context.ClientContext =nil

		switch args[0] {
		case "name":
			rc=context.FindClientByName(args[1])
		case "id":
			id,err :=strconv.ParseInt(args[1],0,64)
			if err!=nil {
				return "ERROR:The id subcommand  need a integer64"
			} else {
				rc=context.FindClientById(id)
			}
		}

		if rc!=nil {
			client.ClientConnectWriteLine(rc,m)
		} else {
			return "ERROR:Can't find the client"
		}

		return ""
	},""))
}
