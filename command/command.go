package command

import (
	"strings"
	"log"
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/context"
	"github.com/blackspace/goserver/action"
)

type _CommandFunc func (c * context.ClientContext,args ...string) string

type _Command struct {
	Name        string
	Func        _CommandFunc
	Description string
}

func NewCommand(name string,f _CommandFunc,dst string) *_Command {
	return &_Command{name,f,dst}
}


var Commands =make([]*_Command,0,1<<8)

func RegistCommand(c *_Command) {
	Commands =append(Commands,c)
}

func FindCommandByName(n string) *_Command {
	for _,c :=range Commands {
		if(c.Name==n) {
			return c
		}
	}
	return nil
}

func IsCommand(n string) bool {
	for _,c :=range Commands {
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

	cmd := FindCommandByName(fs[0])



	if(cmd!=nil&&cmd.Func!=nil) {
		return cmd.Func(c,fs[1:]...)
	} else {
		return  `The command '`+fs[0]+`' isn't exist.`
	}
}


func IsEmptyLine(line string) bool {
	return  len(line)==0
}

func DoCommand(cc *context.ClientContext , line string)  bool {
	log.Println(`Get a line from client:`, line)

	result := ExecString(cc, line)

	log.Println("The result of the command is:", result)

	//When the client execute the quit command,the client has been closed,
	//Can't write the connect and must break the for loop
	if !cc.IsClosed {
		if len(result) > 0 {
			client.ClientConnectSendResult(cc, result + "\r\n")
		}

		if cc.NeedPrompt {
			client.ClientConnectPrintPrompt(cc)
		}

	}

	return true
}

func DoEmptyLine(cc *context.ClientContext , line string)  bool {
	return true
}



func init() {
	action.LineActions.AddAction(func(line string) bool { return IsCommand(line) }, DoCommand)
	action.LineActions.AddAction(func(line string) bool { return IsEmptyLine(line)},DoEmptyLine)
}