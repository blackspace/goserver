package action

import (
	"goserver/context"
	"goserver/client"
	"log"
	"strings"
	"strconv"
)


type pattern struct {
	
}

func DoHttpMethod(cc * context.ClientContext,buf []byte) bool {
	start := buf[:len(buf)-2]

	method :=strings.Split(string(start)," ")[0]
	url :=strings.Split(string(start)," ")[1]
	ver :=strings.Split(string(start)," ")[2]



	headers := map[string]string{}

	for {
		l,err:=client.ClientConnectReadLine(cc)

		if err!=nil {
			log.Println(err)
			break
		}

		kv := strings.Split(l,": ")

		log.Println(l)

		if len(l)==0 {
			break
		} else {
			headers[kv[0]]=kv[1]
		}

	}

	log.Println(start)
	log.Println(method,url,ver)
	log.Println(headers)

	if cl :=headers["Content-Length"];cl!="" {
		body:=make([]byte,0,1024)
		l,_ := strconv.Atoi(cl)
		for i :=0;i<l;i++{
			c,err:=client.ClientReadByte(cc)
			if err!=nil {
				break
			}
			body=append(body,c)
		}

		log.Println(string(body))


		ct:=strings.Split(headers["Content-Type"],";")[0];
		ct_kv:=map[string]string{}

		if len(strings.Split(headers["Content-Type"],";"))>1 {
			for _,t:=range strings.Split(headers["Content-Type"],";")[1:] {
				ct_kv[strings.TrimSpace(strings.Split(t,"=")[0])]=strings.Split(t,"=")[1]
			}
		}

		log.Println(ct_kv)

		switch ct {
		case "multipart/form-data":
			log.Println(ct_kv["boundary"])
			log.Println(ct)
		case "application/x-www-form-urlencoded":
			log.Println(ct)
		}
	}

	return true
}
