package parser

import (
	"bufio"
	"regexp"
)

type ParetheseLeftPattern struct {
	PatternBase
}

func (s * ParetheseLeftPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}


type ParetheseRightPattern struct {
	PatternBase
}

func (s * ParetheseRightPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}

func init() {

	var lr = regexp.MustCompile(`^\($`)

	RegistPattern(&ParetheseLeftPattern{PatternBase{Regexp:lr, Token:'('}})

	var rr = regexp.MustCompile(`^\)$`)

	RegistPattern(&ParetheseRightPattern{PatternBase{Regexp:rr, Token:')'}})
}

