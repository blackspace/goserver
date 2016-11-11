package command

import (
	"strings"
	"log"
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/action"
)

type _CommandFunc func (clnt *client.Client,args ...string) string

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

func IsCommand(l string) bool {
	ss:=strings.Split(l," ")

	for _,c :=range Commands {
		if(c.Name==ss[0]) {
			return true
		}
	}
	return false
}

func _GetFragment(l string) []string {
	result := strings.Split(l, ` `)
	return result
}



func ExecString(clnt *client.Client,l string)(string) {
	fs:=_GetFragment(l)

	log.Println("To Find a command:",fs[0])

	cmd := FindCommandByName(fs[0])

	if(cmd!=nil&&cmd.Func!=nil) {
		return cmd.Func(clnt,fs[1:]...)
	} else {
		return  `The command '`+fs[0]+`' isn't exist.`
	}
}


func IsEmptyLine(line string) bool {
	return  len(line)==0
}

func DoCommand(clnt *client.Client , line string)  bool {
	log.Println(`Get a line from client:`, line)

	result := ExecString(clnt, line)

	log.Println("The result of the command is:", result)

	//When the client execute the quit command,the client has been closed,
	//Can't write the connect and must break the for loop
	if !clnt.IsClosed {
		if len(result) > 0 {
			clnt.SendResult(result + "\r\n")
		}

		if clnt.NeedPrompt {
			clnt.PrintPrompt()
		}

	}

	return true
}

func DoEmptyLine(clnt *client.Client , line string)  bool {
	return true
}



func init() {
	action.LineActions.AddAction(func(line string) bool { return IsCommand(line) }, DoCommand)
	action.LineActions.AddAction(func(line string) bool { return IsEmptyLine(line)},DoEmptyLine)
}