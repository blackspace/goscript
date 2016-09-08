package parser

import (
	"bufio"
	"io"
	"regexp"
	"goscript/ast"
)

type VariablePattern struct {
	PatternBase
}

func (n * VariablePattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	s:=prefix
	for {
		r,_,err:=r.ReadRune()

		if err==nil {
			s=s+string(r)

			if n.Match(s) {
				continue
			} else {
				s=s[:len(s)-1]
				return &ast.VarExpr{Name:s},r,true
			}
		} else if err==io.EOF{
			return &ast.VarExpr{Name:s},0,false
		}

	}

	return nil,0,false

}

func init() {
	var variableRegexp=regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*$`)

	Patterns=append(Patterns,&VariablePattern{PatternBase{Regexp:variableRegexp,Token:VARIABLE}})
}

