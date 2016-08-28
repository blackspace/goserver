package myhttp

import (
	"bufio"
	"strings"
	"log"
)

type Header []Field

func (h Header)GetField(n string) (Field,bool) {
	for _,h := range h {
		if h.Name==n {
			return h,true
		}
	}

	return  Field{},false
}

func (h Header)ToString() (result string){
	for _,f := range h {
		result+=f.Name+": "+f.Value+"\r\n"
	}

	return
}

func _ParseHeaderFromString(s string) (result Header) {
	scanner:=bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan()  {
		l := scanner.Text()
		m :=FIELD.FindStringSubmatch(l)

		result = append(result,Field{m[1],m[2]})

	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return
}
