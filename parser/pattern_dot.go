package parser

import (
"bufio"
"regexp"
)

type DotPattern struct {
	PatternBase
}

func (s * DotPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}

func init() {

	var r = regexp.MustCompile(`^\.$`)

	RegistPattern(&DotPattern{PatternBase{Regexp:r, Token:'.'}})
}
