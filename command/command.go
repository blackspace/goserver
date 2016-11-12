package command

import (
	"strings"
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/action"
)

type CommandFunc func (clt *client.Client,args ...string) string

type Command struct {
	Name        string
	Func        CommandFunc
	Description string
}

func NewCommand(name string,f CommandFunc,dst string) *Command {
	return &Command{name,f,dst}
}

var Commands =NewCommands()

type _Commands struct {
	_data []*Command
}

func NewCommands() *_Commands {
	return &_Commands{_data:make([]*Command,0,1<<8)}
}

func (cs *_Commands)RegistCommand(name string,f CommandFunc,dst string) {
	c:=NewCommand(name,f,dst)
	cs._data =append(cs._data,c)
}

func (cs *_Commands)FindCommandByName(n string) *Command {
	for _,c :=range cs._data {
		if(c.Name==n) {
			return c
		}
	}
	return nil
}

func (cs *_Commands)AllCommandName() (result []string) {
	for _,cmd := range Commands._data {
		result=append(result ,cmd.Name)
	}
	return
}

func IsCommand(l string) bool {
	ss:=strings.Split(l," ")

	for _,c :=range Commands._data {
		if(c.Name==ss[0]) {
			return true
		}
	}
	return false
}

func _ExecString(clt *client.Client,l string)(string) {
	fs:=strings.Split(l, ` `)

	cmd := Commands.FindCommandByName(fs[0])

	if(cmd!=nil&&cmd.Func!=nil) {
		return cmd.Func(clt,fs[1:]...)
	} else {
		return  `The command '`+fs[0]+`' isn't exist.`
	}
}

func DoCommand(clt *client.Client , line string)  bool {
	result := _ExecString(clt, line)

	//When the client execute the quit command,the client has been closed,
	//Can't write the connect and must break the for loop
	if !clt.IsClosed {
		if len(result) > 0 {
			clt.SendResult(result + "\r\n")
		}

		if clt.NeedPrompt {
			clt.PrintPrompt()
		}

	}

	return true
}

func IsEmptyLine(line string) bool {
	return  len(line)==0
}

func DoEmptyLine(clt *client.Client , line string)  bool {
	return true
}

func init() {
	action.LineActions.AddAction(func(line string) bool { return IsCommand(line) }, DoCommand)
	action.LineActions.AddAction(func(line string) bool { return IsEmptyLine(line)},DoEmptyLine)
}