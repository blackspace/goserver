package command

import (
	"goserver/context"
)

func init() {
	_AddCommand(_NewCommand("help", func(c *context.ClientContext, args ...string) (result string) {
		if len(args)>0 {
		 	cmd := _FindCommandByName(args[0])
			if cmd==nil {
				return "There isn't the command"
			} else {
				return cmd.Description
			}
		} else {
			for _,cmd := range _commands {
				result += cmd.Name+"\n"
			}
		}

		return
	},"Print all commands."))
}
