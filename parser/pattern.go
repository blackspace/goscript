package parser

import (
	"bufio"
	"regexp"
)

type Pattern interface {
	BuildFun(prefix string,r *bufio.Reader) (interface{},rune,bool)
	Match(s string) bool
	GetToken() int
}

type PatternBase struct {
	* regexp.Regexp
	Token int
}

func (p PatternBase)Match(s string) bool {
	return p.MatchString(s)
}

func (p PatternBase)GetToken() int {
	return p.Token
}


var Patterns = make([]Pattern,0,256)


func FindPattern(s string) Pattern {
	for _,e:=range Patterns {

		if e.Match(s) {
			return e
		}
	}

	return nil
}

