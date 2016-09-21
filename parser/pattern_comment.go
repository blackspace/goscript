package parser

import (
	"bufio"
	"regexp"
	"io"
)

type LineCommentPattern struct {
	PatternBase
}

func (n * LineCommentPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	s:=prefix
	for {
		r,_,err:=r.ReadRune()

		s+=string(r)

		if err==nil {
			if n.Match(s) {
				continue
			} else {
				return nil,r,true
			}
		} else if err==io.EOF{
			return nil,0,false
		}
	}

	return nil,0,false

}
func init() {
	var r=regexp.MustCompile(`^#.*(\n|\r\n)$`)

	Patterns=append(Patterns,&LineCommentPattern{PatternBase{Regexp:r,Token:LINECOMMENT}})
}


