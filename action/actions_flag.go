package action

import (
	"goserver/context"
)

type FlagActionFun func(cc * context.ClientContext,buf []byte) (need_keep_link bool)
type FlagPredicateFun func([]byte)  bool
var FlagActions = []struct {
	MatchPatternFun FlagPredicateFun
	DoActionFun     FlagActionFun
}{
	{func(buf []byte) bool { return IsSocksV4Instruction(buf)}, DoSocksV4},
}


func FindActionForFlag(buf []byte) FlagActionFun {
	for  _,r:=range FlagActions {
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
		for  _,r:=range FlagActions {
			if r.MatchPatternFun(buf) {
				return true
			}
		}
	}

	return false
}

func IsSocksV4Instruction(buf []byte) bool {
	return len(buf) >= 8 && buf[0] == 0x04 && buf[1] == 0x01 &&  buf[len(buf) - 1] == 0x00
}
