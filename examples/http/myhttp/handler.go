package myhttp

import "bufio"

type HandlerFun func(w * bufio.Writer,r * Request)

type Handler struct {
	url string
	HandlerFun HandlerFun
}

var handlers=make([]Handler,0,256)


func AddHandler(url string,f HandlerFun) {
	handlers=append(handlers,Handler{url,f})
}

func FindHander(u string) (Handler,bool) {
	for _,h:=range handlers {
		if h.url==u {
			return h,true
		}
	}

	return Handler{},false
}
