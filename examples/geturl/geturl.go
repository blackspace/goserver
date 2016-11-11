package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"regexp"
	"time"
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/context"
	"github.com/blackspace/goserver/action"
	"github.com/blackspace/goserver"
)

func DoGetUrl(cc *context.ClientContext , url string)  bool {
	res,err:=http.Get(string(url))

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

func IsUrl(line string) bool {
	b,_ :=regexp.MatchString(`https?://.*`,line)
	return b
}


func init() {
	action.LineActions.AddAction(func(line string) bool { return IsUrl(line)}, DoGetUrl)
}

func main() {
	goserver.Start()
	defer goserver.Stop()

	time.Sleep(time.Hour)
}



