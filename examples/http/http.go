package main

import (
	"github.com/blackspace/goserver/context"
	"github.com/blackspace/goserver/action"
	"./myhttp"
	"bufio"
	"regexp"
	"time"
	"github.com/blackspace/goserver"
)


func DoHttpMethod(cc * context.ClientContext, start_line string) bool {
	request := myhttp.NewRequest()

	request.ExtractStartLine(start_line)
	request.ExtractHeaderFromStream(cc.Reader)

	if request.Method=="POST" {
		request.ExtractBodyFromStream(cc.Reader)
	}

	DoRequest(cc.Writer,request)

	return false
}

func DoRequest(w * bufio.Writer,r * myhttp.Request) {
	if h,ok := myhttp.FindHander(r.Url);ok {
		h.HandlerFun(w,r)
	} else {
		response := myhttp.Response{StatusLine:myhttp.StatusLine{"HTTP/1.1","404","Not Found"}}
		w.WriteString(response.ToString())
		w.Flush()
	}
}

func IsHttpRequest(line string) bool {
	b,_:=regexp.MatchString(`((GET)|(POST))`,line)
	return b
}


func init() {
	action.LineActions.AddAction(action.LineAction{func(line string) bool { return IsHttpRequest(line) },DoHttpMethod})


	myhttp.AddHandler("/",func(w *bufio.Writer,r *myhttp.Request)  {
		response := myhttp.NewResponse()
		response.Body = `hello world!`
		w.WriteString(response.ToString())
		w.Flush()
	})
}

func main() {
	goserver.Start()
	defer goserver.Stop()
	time.Sleep(time.Hour)
}
