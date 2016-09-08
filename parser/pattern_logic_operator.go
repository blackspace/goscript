package parser

import (
	"bufio"
	"regexp"
)

type ORPattern struct {
	PatternBase
}

func (n * ORPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}

type ANDPattern struct {
	PatternBase
}

func (n * ANDPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}

func init() {
	var orRegexp = regexp.MustCompile(`^\|\|$`)

	Patterns = append(Patterns, &ORPattern{PatternBase{Regexp:orRegexp, Token:OR}})

	var andRegexp = regexp.MustCompile(`^&&$`)

	Patterns = append(Patterns, &ANDPattern{PatternBase{Regexp:andRegexp, Token:AND}})
}

