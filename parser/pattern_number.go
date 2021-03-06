package parser

import (
	"bufio"
	"strconv"
	"io"
	"goscript/ast"
	"regexp"
)

type IntPattern struct {
	PatternBase
}

func (n * IntPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre_rune rune,has_prerune bool) {
	s:=prefix
	for {
		r,_,err:=r.ReadRune()

		if err==nil {
			s=s+string(r)

			if n.Match(s) {
				continue
			} else {
				s=s[:len(s)-1]
				i,_:=strconv.Atoi(s)
				return &ast.Int{int64(i)},r,true
			}
		} else if err==io.EOF{
			i,_:=strconv.Atoi(s)
			return &ast.Int{int64(i)},0,false
		}

	}

	return nil,0,false

}

func init() {
	var numberRegexp=regexp.MustCompile(`^[0-9]+$`)

	Patterns=append(Patterns,&IntPattern{PatternBase{Regexp:numberRegexp,Token:INT}})
}

