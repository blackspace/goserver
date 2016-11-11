package terminal_mode

import (
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

const Prompt="\033[40;32mgochat-server>\033[0m"

func init() {
	RegistCommand(NewCommand("show_prompt",func (clnt *client.Client,args ...string) string {
		clnt.NeedPrompt=true
		clnt.Prompt=Prompt
		return ""
	},""))
	RegistCommand(NewCommand("hide_prompt",func (clnt *client.Client,args ...string) string {
		clnt.NeedPrompt=false
		clnt.Prompt=""
		return ""
	},""))
}
