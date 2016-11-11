package help

import (
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

func init() {
	RegistCommand(NewCommand("help", func(cln *client.Client, args ...string) (result string) {
		if len(args)>0 {
		 	cmd := FindCommandByName(args[0])
			if cmd==nil {
				return "There isn't the command"
			} else {
				return cmd.Description
			}
		} else {
			for _,cmd := range Commands {
				result += cmd.Name+"\n"
			}
		}

		return
	},"Print all commands."))
}
