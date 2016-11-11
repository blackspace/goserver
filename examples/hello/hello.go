package main

import (
	"time"
	"github.com/blackspace/goserver"
	"github.com/blackspace/goserver/context"
	"github.com/blackspace/goserver/command"
)

func main() {
	command.RegistCommand(command.NewCommand("hello", func(c *context.ClientContext, args ...string) string {
		return "hello,I am a robot"
	},"say hello"))


	s:=goserver.NewServer()

	s.Start("127.0.0.1","5050")
	defer s.Stop()

	time.Sleep(time.Hour)
}

