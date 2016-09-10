package parser

import (
	"bufio"
	"io"
	"regexp"
	"goscript/ast"
)

type WordPattern struct {
	PatternBase
}

func (n * WordPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	s:=prefix
	for {
		r,_,err:=r.ReadRune()

		if err==nil {
			s=s+string(r)

			if n.Match(s) {
				continue
			} else {
				s=s[:len(s)-1]
				return s,r,true
			}
		} else if err==io.EOF{
			return s,0,false
		}

	}

	return nil,0,false

}

func init() {
	var r=regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*$`)

	Patterns=append(Patterns,&WordPattern{PatternBase{Regexp:r,Token:WORD}})
}

func ParseWord(w string) (Token int,v interface{}) {
	switch w {
	case "true":
		return BOOL,&ast.Bool{Bool:true}
	case "false":
		return BOOL,&ast.Bool{Bool:false}
	case "if":
		return IF,nil
	case "else":
		return ELSE,nil
	case "for":
		return FOR,nil
	}
	return VARIABLE,&ast.VarExpr{Name:w}
}

