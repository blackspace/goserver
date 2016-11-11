package action

import (
	"github.com/blackspace/goserver/context"
)

type FlagActionFun func(cc * context.ClientContext,buf []byte) (need_keep_link bool)
type FlagPredicateFun func([]byte)  bool

type FlagAction struct {
	MatchPatternFun FlagPredicateFun
	DoActionFun     FlagActionFun
}


type _FlagActions struct {
	_data []FlagAction
}

func (as *_FlagActions)AddAction(mf FlagPredicateFun,df FlagActionFun) {
	as._data=append(as._data,FlagAction{mf,df})
}

func NewFlagActions() *_FlagActions {
	return &_FlagActions{_data:make([]FlagAction,0,1<<8)}
}

var FlagActions=NewFlagActions()

func FindActionForFlag(buf []byte) FlagActionFun {
	for  _,r:=range FlagActions._data {
		if r.MatchPatternFun(buf) {
			return r.DoActionFun
		}
	}

	return nil
}

func IsFlag(buf []byte) bool {
	if len(buf)<=1 {
		return false
	} else {
		for  _,r:=range FlagActions._data {
			if r.MatchPatternFun(buf) {
				return true
			}
		}
	}

	return false
}

