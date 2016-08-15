package action

import (
	"goserver/context"
	"goserver/client"
	"log"
	"goserver/command"
)



func Command(cc *context.ClientContext , buf []byte)  bool {
	buf = buf[0:len(buf) - 2]
	line := string(buf)

	if len(line) == 0 {
		client.ClientConnectPrintPrompt(cc)
		return false
	}

	log.Println(`Get a line from client:`, line)

	result := command.ExecString(cc, line)

	log.Println("The result of the command is:", result)

	//When the client execute the quit command,the client has been closed,
	//Can't write the connect and must break the for loop
	if !cc.IsClosed {
		if len(result) > 0 {
			client.ClientConnectSendResult(cc, result + "\n")
		}

		if cc.NeedPrompt {
			client.ClientConnectPrintPrompt(cc)
		}

	}

	return false
}
