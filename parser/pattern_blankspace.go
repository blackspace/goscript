package parser

import (
	"bufio"
	"regexp"
	"io"
)

type BlankSpacePattern struct {
	PatternBase
}

func (n * BlankSpacePattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
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

	var blankSpaceRegexp = regexp.MustCompile(`^ +$`)

	Patterns=append(Patterns,&BlankSpacePattern{PatternBase{Regexp:blankSpaceRegexp,Token:BLANKSPACE}})

}
