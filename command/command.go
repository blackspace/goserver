package command

import (
	"goserver/context"
	"strings"
	"log"
)

type _CommandFunc func (c * context.ClientContext,args ...string) string

type _Command struct {
	Name        string
	Func        _CommandFunc
	Description string
}

func _NewCommand(name string,f _CommandFunc,dst string) *_Command {
	return &_Command{name,f,dst}
}


var _commands =make([]*_Command,0,1<<8)

func _AddCommand(c *_Command) {
	_commands=append(_commands,c)
}

func _FindCommandByName(n string) *_Command {
	for _,c :=range _commands {
		if(c.Name==n) {
			return c
		}
	}
	return nil
}

func IsCommand(n string) bool {
	for _,c :=range _commands {
		if(c.Name==n) {
			return true
		}
	}
	return false
}

func _GetFragment(l string) []string {
	result := strings.Split(l, ` `)
	return result
}



func ExecString(c *context.ClientContext,l string)(string) {
	fs:=_GetFragment(l)

	log.Println("To Find a command:",fs[0])

	cmd := _FindCommandByName(fs[0])



	if(cmd!=nil&&cmd.Func!=nil) {
		return cmd.Func(c,fs[1:]...)
	} else {
		return  `The command '`+fs[0]+`' isn't exist.`
	}
}


