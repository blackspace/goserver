package action

import (
	"github.com/blackspace/goserver/context"
)

type LineActionFun func(cc * context.ClientContext,line string) (need_keep_link bool)
type LinePredicateFun func(string)  bool

type LineAction struct {
	MatchPatternFun LinePredicateFun
	DoActionFun     LineActionFun
}

type _LineActions struct {
	_data []LineAction
}

func NewLineActions() *_LineActions {
	return &_LineActions{_data:make([]LineAction,0,1<<8)}
}

func (as *_LineActions)AddAction(mf LinePredicateFun,df LineActionFun) {
	as._data=append(as._data,LineAction{mf,df})
}

var LineActions = NewLineActions()

func FindActionForLine(line string) LineActionFun {
	for  _,r:=range LineActions._data {
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


