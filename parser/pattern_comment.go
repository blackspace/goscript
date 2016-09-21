package parser

import (
	"bufio"
	"regexp"
	"io"
)

type PoundCommentPattern struct {
	PatternBase
}

func (n *PoundCommentPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	s:=prefix
	for {
		r,_,err:=r.ReadRune()

		s+=string(r)

		if err==nil {
			if n.Match(s) {
				continue
			} else {
				return s,r,true
			}
		} else if err==io.EOF{
			return s,0,false
		}
	}

	return s,0,false

}


type DoubleSlashCommentPattern struct {
	PatternBase
}

func (n *DoubleSlashCommentPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	s:=prefix
	for {
		r,_,err:=r.ReadRune()

		s+=string(r)

		if err==nil {
			if n.Match(s) {
				continue
			} else {
				return s,r,true
			}
		} else if err==io.EOF{
			return s,0,false
		}
	}

	return s,0,false

}


type MultiLineCommentPattern struct {
	PatternBase
}

func (n *MultiLineCommentPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	s:=prefix
	for {
		r,_,err:=r.ReadRune()

		s+=string(r)

		if err==nil {
			if n.Match(s) {
				continue
			} else {
				return s,r,true
			}
		} else if err==io.EOF{
			return s,0,false
		}
	}

	return s,0,false

}

func init() {
	var pr=regexp.MustCompile(`^#.*(\n|\r\n)?$`)

	Patterns=append(Patterns,&PoundCommentPattern{PatternBase{Regexp:pr,Token:POUNDCOMMENT}})


	var dr=regexp.MustCompile(`^//.*(\n|\r\n)?$`)

	Patterns=append(Patterns,&DoubleSlashCommentPattern{PatternBase{Regexp:dr,Token:DOUBLESLASHCOMMENT}})

	var mr=regexp.MustCompile(`^/\*.*\*/$`)

	Patterns=append(Patterns,&MultiLineCommentPattern{PatternBase{Regexp:mr,Token:MULTILINECOMMENT}})
}


