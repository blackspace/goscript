package parser

import (
	"bufio"
	"regexp"
)

type GreaterThenPattern struct {
	PatternBase
}

func (n * GreaterThenPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}


type LessThanPattern struct {
	PatternBase
}

func (n * LessThanPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}


func init() {

	var r = regexp.MustCompile(`^>$`)

	RegistPattern(&GreaterThenPattern{PatternBase{Regexp:r, Token:'>'}})

	r = regexp.MustCompile(`^<$`)

	RegistPattern(&LessThanPattern{PatternBase{Regexp:r, Token:'<'}})
}

