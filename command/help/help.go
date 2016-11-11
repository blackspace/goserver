package help

import (
	"github.com/blackspace/goserver/context"
	. "github.com/blackspace/goserver/command"
)

func init() {
	RegistCommand(NewCommand("help", func(c *context.ClientContext, args ...string) (result string) {
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
