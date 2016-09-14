package parser

import (
	"bufio"
	"regexp"
)

type CommaPattern struct {
	PatternBase
}

func (s * CommaPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}

func init() {

	var lr = regexp.MustCompile(`^,$`)

	RegistPattern(&SquareBracketLeftPattern{PatternBase{Regexp:lr, Token:','}})
}