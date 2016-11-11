package myhttp

import "regexp"


type Pattern struct {
	* regexp.Regexp
	_parse func(s string,r * regexp.Regexp) interface{}
}


func (p Pattern)Parse(s string) interface{} {
	return p._parse(s,p.Regexp)

}

var FIELD = Pattern{Regexp:regexp.MustCompile(`(.+): (.+)`), _parse:func(s string,r * regexp.Regexp) interface{} {
	m:=r.FindStringSubmatch(s)
	return Field{m[1],m[2]}
	},}



