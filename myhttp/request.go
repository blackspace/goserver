package myhttp

import (
	"strings"
	"bufio"
	"log"
	"strconv"
	"net/url"
)


type MultiBodyPart struct {
	Headers Header
	Body    interface{}
}

type Request struct {
	Method,Url,Proto string
	Header           Header
	Body             interface{}
}



func NewRequest() *Request {
	return new(Request)
}

func (r *  Request)ExtractStartLine(s string) {
	r.Method =strings.Split(string(s)," ")[0]
	r.Url =strings.Split(string(s)," ")[1]
	r.Proto =strings.Split(string(s)," ")[2]
}


func (r * Request)ExtractHeaderFromStream(reader * bufio.Reader) {
	header :=make([]byte,0,1024)

	for {
		c,err := reader.ReadByte()

		if err!=nil {
			log.Println(err)
			break
		}

		header=append(header,c)


		if  len(header)>=4 && string(header[len(header)-4:])==string([]byte{'\r','\n','\r','\n'}) {
			break
		}
	}

	r.Header = _ParseHeaderFromString(string(header[:len(header)-2]))
}

func (r *  Request)ExtractBodyFromStream(reader * bufio.Reader)  {
	if cl,f :=r.Header.GetField("Content-Length");f==true {
		ct,_ :=r.Header.GetField("Content-Type")
		mt :=ContentType(ct).GetMediaType()

		switch mt {
		case "multipart/form-data":
			r._ParseMultiPartBodyFromStream(reader)
		case "application/x-www-form-urlencoded":
			l,_ := strconv.Atoi(cl.Value)
			buf :=make([]byte,l)
			n,_ :=   reader.Read(buf)
			r._ParseQueryBody(string(buf[:n]))
		}
	}
}


func (r *  Request)_ParseMultiPartBodyFromStream(reader * bufio.Reader)  {

	ct,_:=r.Header.GetField("Content-Type")

	b := ContentType(ct).GetBoundary()

	boundary := "--"+b
	last_boundary := boundary+"--"

	r.Body=make([]MultiBodyPart,0,16)

	for {
		l,_,_ :=reader.ReadLine()

		if string(l)==boundary {
			r._ParseBodyPartFromStream(reader)
		}

		if string(l)==last_boundary {
			break
		}
	}
}

func (r * Request)_ParseBodyPartFromStream(reader * bufio.Reader) {
	mbp := MultiBodyPart{}

	for {
		l,_,_ :=reader.ReadLine()

		if len(l)==0 {
			break
		}

		mbp.Headers = append(mbp.Headers, FIELD.Parse(string(l)).(Field))
	}

	if ct,_:=mbp.Headers.GetField("Content-Type");ct.Value=="application/octet-stream" {
		buf := make([]byte,0,1024)
		for {
			c,_:=reader.ReadByte()

			buf=append(buf,c)

			if  len(buf)>=3 && string(buf[len(buf)-3:])==string([]byte{10,13,10}) {
				break
			}
		}

		mbp.Body =buf[:len(buf)-3]

	} else {
		l,_,_ :=reader.ReadLine()
		mbp.Body = string(l)

	}


	if v,ok:= r.Body.([]MultiBodyPart);ok {
		r.Body= append(v,mbp)
	}

}


func (r *  Request)_ParseQueryBody(s string)  {
	 m,_ := url.ParseQuery(s)

	var result []Field

	for  k := range m {
		result = append(result,Field{k,m.Get(k)})
	}

	r.Body = result
}





