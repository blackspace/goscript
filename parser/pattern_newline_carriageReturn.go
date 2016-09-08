package parser

import (
	"bufio"
	"regexp"
	"io"
)

type LFCRPattern struct {
	PatternBase
}

func (n * LFCRPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	s:=prefix
	for {
		r,_,err:=r.ReadRune()

		if err==nil {
			s=s+string(r)

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

	var r= regexp.MustCompile(`^(\r|(\r\n)|\n)+$`)

	Patterns=append(Patterns,&LFCRPattern{PatternBase{Regexp:r,Token:LFCR}})

}
