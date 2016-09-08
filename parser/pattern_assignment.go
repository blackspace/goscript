package parser

import (
	"bufio"
	"regexp"
)

type AssignPattern struct {
	PatternBase
}

func (n * AssignPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return '=',0,false

}

func init() {

	var assignRegexp = regexp.MustCompile(`^=$`)

	Patterns = append(Patterns, &AssignPattern{PatternBase{Regexp:assignRegexp,Token:'='}})
}
