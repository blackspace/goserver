package action

import (
	"goserver/context"
	"regexp"
	"goserver/command"
	"strings"
)

type  Route struct {
	PatternFunc func([]byte)  bool
	ActionFunc func(cc * context.ClientContext,buf []byte) bool
}

var Routes = []Route{
	{func(buf []byte) bool { return IsLine(buf) && IsHttpRequest(buf) },DoHttpMethod},


	{func(buf []byte) bool { return IsLine(buf) && IsCommand(buf) },Command},
	{func(buf []byte) bool { return IsLine(buf) && IsUrl(buf)},GetUrl},
	{func(buf []byte) bool { return IsLine(buf) && IsEmptyLine(buf)},DoEmptyLine},
	{func(buf []byte) bool { return IsLine(buf) },DoInvalidLine},

	{func(buf []byte) bool { return IsSocksV4Instruction(buf)},SocksV4},
}

func FindRoute(buf []byte) *Route {
	for  _,r:=range Routes {
		if r.PatternFunc(buf) {
			return &r
		}
	}

	return nil
}

func IsHttpRequest(buf []byte) bool {
	b,_:=regexp.MatchString(`GET /.*`,string(buf))
	return b
}

func IsSocksV4Instruction(buf []byte) bool {
	return len(buf) >= 8 && buf[0] == 0x04 && buf[1] == 0x01 &&  buf[len(buf) - 1] == 0x00
}


func IsLine(buf []byte)  bool {
	return len(buf) >= 2 &&
		(string(buf[len(buf) - 2:]) == string([]byte{0x0D, 0x0A}) ||
			string(buf[len(buf) - 2:]) == string([]byte{0x0D, 0x00}))

}

func IsEmptyLine(buf []byte) bool {
	if len(buf)<2 {
		return false
	}
	return  len(string(buf[:len(buf)-2]))==0
}

func IsUrl(buf []byte) bool {
	b,_ :=regexp.MatchString(`https?://.*`,string(buf))
	return b
}

func IsCommand(buf []byte) bool {
	n :=string(buf[:len(buf)-2])
	return command.IsCommand(strings.Split(n," ")[0])
}