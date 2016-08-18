package action

import (
	"goserver/context"
	"goserver/client"
	"log"
	"goserver/command"
)



func DoCommand(cc *context.ClientContext , line string)  bool {
	log.Println(`Get a line from client:`, line)

	result := command.ExecString(cc, line)

	log.Println("The result of the command is:", result)

	//When the client execute the quit command,the client has been closed,
	//Can't write the connect and must break the for loop
	if !cc.IsClosed {
		if len(result) > 0 {
			client.ClientConnectSendResult(cc, result + "\r\n")
		}

		if cc.NeedPrompt {
			client.ClientConnectPrintPrompt(cc)
		}

	}

	return true
}
