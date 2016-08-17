package action

import (
	"goserver/context"
	"goserver/client"
	"net/http"
	"log"
	"io/ioutil"
)

func DoGetUrl(cc *context.ClientContext , buf []byte)  bool {
	buf = buf[:len(buf)-2]
	res,err:=http.Get(string(buf))

	if err !=nil {
		log.Fatal(err)
	}

	robots,err :=ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err!=nil {
		log.Fatal(err)
	}

	client.ClientConnectWriteLine(cc,string(robots))
	return true
}
