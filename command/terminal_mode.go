package command

import (
	"goserver/context"
	"goserver/client"
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
	_AddCommand(_NewCommand("character_mode",func (c *context.ClientContext,args ...string) string {
		c.IsCharacterMode=true
		a :=  []byte{0xff,0xfd,0x18}
		client.ClientConnectWrite(c,a)
		return ""

	},""))
	_AddCommand(_NewCommand("line_mode",func (c *context.ClientContext,args ...string) string {
		c.IsCharacterMode=true
		return "\xff\xfd\x18"
	},""))
}
