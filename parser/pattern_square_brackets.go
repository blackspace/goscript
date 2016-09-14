package parser

import (
	"bufio"
	"regexp"
)

type SquareBracketLeftPattern struct {
	PatternBase
}

func (s * SquareBracketLeftPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}


type SquareBracketRightPattern struct {
	PatternBase
}

func (s * SquareBracketRightPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}

func init() {

	var lr = regexp.MustCompile(`^\[$`)

	RegistPattern(&SquareBracketLeftPattern{PatternBase{Regexp:lr, Token:'['}})

	var rr = regexp.MustCompile(`^\]$`)

	RegistPattern(&SquareBracketRightPattern{PatternBase{Regexp:rr, Token:']'}})
}

