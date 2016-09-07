package pattern

import (
	"bufio"
	"regexp"
)

type Pattern interface {
	BuildFun(prefix string,r *bufio.Reader) (interface{},rune,bool)
	Match(s string) bool
}

type PatternBase struct {
	* regexp.Regexp
}

func (p PatternBase)Match(s string) bool {
	return p.MatchString(s)
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

