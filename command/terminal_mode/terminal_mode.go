package terminal_mode

import (
	. "github.com/blackspace/goserver/command"
	"github.com/blackspace/goserver/client"
)

const Prompt="\033[40;32mgochat-server>\033[0m"

func init() {
	RegistCommand(NewCommand("show_prompt",func (client *client.Client,args ...string) string {
		client.NeedPrompt=true
		client.Prompt=Prompt
		return ""
	},""))
	RegistCommand(NewCommand("hide_prompt",func (client *client.Client,args ...string) string {
		client.NeedPrompt=false
		client.Prompt=""
		return ""
	},""))
}
