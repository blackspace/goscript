package parser

import (
	"bufio"
	"regexp"
)

type AtPattern struct {
	PatternBase
}

func (n * AtPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return nil,0,false

}


func init() {

	var r = regexp.MustCompile(`^@$`)
	RegistPattern(&AtPattern{PatternBase{Regexp:r, Token:'@'}})
}
