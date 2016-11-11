package terminal_mode

import (
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

const Prompt="\033[40;32mgochat-server>\033[0m"

func init() {
	Commands.RegistCommand("show_prompt",func (clt *client.Client,args ...string) string {
		clt.NeedPrompt=true
		clt.Prompt=Prompt
		return ""
	},"")
	Commands.RegistCommand("hide_prompt",func (clt *client.Client,args ...string) string {
		clt.NeedPrompt=false
		clt.Prompt=""
		return ""
	},"")
}
