package help

import (
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

func init() {
	Commands.RegistCommand("help", func(clt *client.Client, args ...string) (result string) {
		if len(args)>0 {
		 	cmd := Commands.FindCommandByName(args[0])
			if cmd==nil {
				return "There isn't the command"
			} else {
				return cmd.Description
			}
		} else {
			for _,n:=range Commands.AllCommandName() {
				result=result+"\r\n"+n
			}
		}

		return
	},"Print all commands.")
}
