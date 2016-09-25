package parser

import (
	"bufio"
	"regexp"
)

type ColonPattern struct {
	PatternBase
}

func (s * ColonPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}

func init() {

	var lr = regexp.MustCompile(`^:$`)

	RegistPattern(&ColonPattern{PatternBase{Regexp:lr, Token:':'}})
}
