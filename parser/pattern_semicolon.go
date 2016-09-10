package parser

import (
	"bufio"
	"regexp"
)

type SemicolonPattern struct {
	PatternBase
}

func (n * SemicolonPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return ';',0,false

}


func init() {

	var r = regexp.MustCompile(`^;$`)

	RegistPattern(&SemicolonPattern{PatternBase{Regexp:r, Token:';'}})
}
