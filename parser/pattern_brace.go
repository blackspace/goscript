package parser

import (
	"bufio"
	"regexp"
)

type LBracePattern struct {
	PatternBase
}

func (n * LBracePattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return '{',0,false

}

type RBracePattern struct {
	PatternBase
}

func (n * RBracePattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return '}',0,false

}


func init() {

	var lBraceRegexp = regexp.MustCompile(`^\{$`)

	RegistPattern(&LBracePattern{PatternBase{Regexp:lBraceRegexp, Token:'{'}})

	var rBraceRegexp = regexp.MustCompile(`^}$`)

	RegistPattern(&RBracePattern{PatternBase{Regexp:rBraceRegexp, Token:'}'}})
}

