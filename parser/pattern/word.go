package pattern

import (
	"bufio"
	"io"
	"regexp"
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
	var wordRegexp=regexp.MustCompile(`^[a-z]\w.$`)

	Patterns=append(Patterns,&NumberPattern{PatternBase{Regexp:wordRegexp}})
}

