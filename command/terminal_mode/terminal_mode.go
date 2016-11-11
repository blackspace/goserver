package terminal_mode

import (
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

const Prompt="\033[40;32mgochat-server>\033[0m"

func init() {
	RegistCommand(NewCommand("show_prompt",func (cln *client.Client,args ...string) string {
		cln.NeedPrompt=true
		cln.Prompt=Prompt
		return ""
	},""))
	RegistCommand(NewCommand("hide_prompt",func (cln *client.Client,args ...string) string {
		cln.NeedPrompt=false
		cln.Prompt=""
		return ""
	},""))
}
