package terminal_mode

import (
	"github.com/blackspace/goserver/context"
	. "github.com/blackspace/goserver/command"
)

const Prompt="\033[40;32mgochat-server>\033[0m"

func init() {
	RegistCommand(NewCommand("show_prompt",func (c *context.ClientContext,args ...string) string {
		c.NeedPrompt=true
		c.Prompt=Prompt
		return ""
	},""))
	RegistCommand(NewCommand("hide_prompt",func (c *context.ClientContext,args ...string) string {
		c.NeedPrompt=false
		c.Prompt=""
		return ""
	},""))
}
