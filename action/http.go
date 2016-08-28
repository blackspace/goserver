package action

import (
	"goserver/context"
	"goserver/myhttp"
	"bufio"
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
	if h,ok :=myhttp.FindHander(r.Url);ok {
		h.HandlerFun(w,r)
	} else {
		response := myhttp.Response{StatusLine:myhttp.StatusLine{"HTTP/1.1","404","Not Found"}}
		w.WriteString(response.ToString())
		w.Flush()
	}


}

func init() {
	myhttp.AddHandler("/",func(w * bufio.Writer,r *myhttp.Request)  {
		w.WriteString(`hello world`)
		w.Flush()
	})

}
