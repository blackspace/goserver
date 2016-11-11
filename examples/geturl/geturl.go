package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"regexp"
	"time"
	"github.com/blackspace/goserver/client"
	"github.com/blackspace/goserver/action"
	"github.com/blackspace/goserver"
)

func DoGetUrl(clnt *client.Client , url string)  bool {
	res,err:=http.Get(string(url))

	if err !=nil {
		log.Fatal(err)
	}

	robots,err :=ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err!=nil {
		log.Fatal(err)
	}

	clnt.WriteLine(string(robots))
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
	s:=goserver.NewServer()

	s.Start("127.0.0.1","8058")
	defer s.Stop()

	time.Sleep(time.Hour)
}



