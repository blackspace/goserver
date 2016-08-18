package action

import (
	"regexp"
	"goserver/command"
	"strings"
	"goserver/context"
)

type LineActionFun func(cc * context.ClientContext,line string) (need_keep_link bool)
type LinePredicateFun func(string)  bool
var LineActions = []struct {
	MatchPatternFun LinePredicateFun
	DoActionFun     LineActionFun
}{
	{func(line string) bool { return IsHttpRequest(line) },DoHttpMethod},
	{func(line string) bool { return IsCommand(line) }, DoCommand},
	{func(line string) bool { return IsUrl(line)}, DoGetUrl},
	{func(line string) bool { return IsEmptyLine(line)},DoEmptyLine},
}
func FindActionForLine(line string) LineActionFun {
	for  _,r:=range LineActions {
		if r.MatchPatternFun(line) {
			return r.DoActionFun
		}
	}

	return nil
}

func IsLine(buf []byte) bool {
	return len(buf) >= 2 &&
		(string(buf[len(buf) - 2:]) == string([]byte{0x0D, 0x0A}) ||
			string(buf[len(buf) - 2:]) == string([]byte{0x0D, 0x00}))
}

func IsHttpRequest(line string) bool {
	b,_:=regexp.MatchString(`((GET)|(POST))`,line)
	return b
}


func IsEmptyLine(line string) bool {
	return  len(line)==0
}

func IsUrl(line string) bool {
	b,_ :=regexp.MatchString(`https?://.*`,line)
	return b
}

func IsCommand(line string) bool {
	return command.IsCommand(strings.Split(line," ")[0])
}
