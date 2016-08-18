package command

import (
	"goserver/context"
)

const Prompt="\033[40;32mgochat-server>\033[0m"

func init() {
	_AddCommand(_NewCommand("show_prompt",func (c *context.ClientContext,args ...string) string {
		c.NeedPrompt=true
		c.Prompt=Prompt
		return ""
	},""))
	_AddCommand(_NewCommand("hide_prompt",func (c *context.ClientContext,args ...string) string {
		c.NeedPrompt=false
		c.Prompt=""
		return ""
	},""))
}
