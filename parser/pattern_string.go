package parser

import (
	"bufio"
	"regexp"
	"goscript/ast"
)

type StringPattern struct {
	PatternBase
}

func (n * StringPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	s:=prefix

	s=s[1:len(s)-1]

	return &ast.String{S:s},0,false

}

func init() {
	var r=regexp.MustCompile(`".*"`)

	Patterns=append(Patterns,&StringPattern{PatternBase{Regexp:r,Token:STRING}})
}

