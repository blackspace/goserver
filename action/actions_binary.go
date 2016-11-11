package action

import (
	"github.com/blackspace/goserver/client"
)

type BinaryActionFun func(clt * client.Client,buf []byte) (need_keep_link bool)
type BinaryPredicateFun func([]byte)  bool

type BinaryAction struct {
	MatchFun BinaryPredicateFun
	DoFun    BinaryActionFun
}


var BinaryActions = NewBinaryActions()

type _BinaryActions struct {
	_data []BinaryAction
}

func NewBinaryActions() *_BinaryActions {
	return &_BinaryActions{_data:make([]BinaryAction,0,1<<8)}
}


func (as *_BinaryActions)AddAction(mf BinaryPredicateFun,df BinaryActionFun) {
	as._data=append(as._data, BinaryAction{mf,df})
}

func FindActionForBinary(buf []byte) BinaryActionFun {
	for  _,r:=range BinaryActions._data {
		if r.MatchFun(buf) {
			return r.DoFun
		}
	}

	return nil
}

func IsBinary(buf []byte) bool {
	if len(buf)<=1 {
		return false
	} else {
		for  _,r:=range BinaryActions._data {
			if r.MatchFun(buf) {
				return true
			}
		}
	}

	return false
}

