package action

import (
	"goserver/context"
	"goserver/client"
	"goserver/action/myhttp"
	"log"
)


func DoHttpMethod(cc * context.ClientContext, start_line string) bool {
	request := myhttp.NewRequest()

	request.ExtractStartLine(start_line)
	request.ExtractHeaderFromStream(cc.Reader)

	if request.Method=="POST" {
		request.ExtractBodyFromStream(cc.Reader)
	}

	log.Println(request)

	client.ClientConnectWriteLine(cc,"hello http")

	return false
}
