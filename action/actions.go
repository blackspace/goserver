package action

import (
	"goserver/context"
	"regexp"
	"goserver/command"
	"strings"
)

type ActionFun func(cc * context.ClientContext,buf []byte) (need_keep_link bool)
type PredicateFun func([]byte)  bool

type  Action struct {
	MatchPatternFun PredicateFun
	DoActionFun     ActionFun
}

var Actions = []Action{
	{func(buf []byte) bool { return IsLine(buf) && IsHttpRequest(buf) },DoHttpMethod},


	{func(buf []byte) bool { return IsLine(buf) && IsCommand(buf) }, DoCommand},
	{func(buf []byte) bool { return IsLine(buf) && IsUrl(buf)}, DoGetUrl},
	{func(buf []byte) bool { return IsLine(buf) && IsEmptyLine(buf)},DoEmptyLine},
	{func(buf []byte) bool { return IsLine(buf) },DoInvalidLine},

	{func(buf []byte) bool { return IsSocksV4Instruction(buf)}, DoSocksV4},
}

func FindAction(buf []byte) ActionFun {
	for  _,r:=range Actions {
		if r.MatchPatternFun(buf) {
			return r.DoActionFun
		}
	}

	return nil
}

func IsHttpRequest(buf []byte) bool {
	b,_:=regexp.MatchString(`((GET)|(POST))`,string(buf))
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