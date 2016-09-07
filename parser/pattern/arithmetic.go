package pattern

import (
	"bufio"
	"regexp"
)

type AddPattern struct {
	PatternBase
}

func (n * AddPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return '+',0,false

}


type SubPattern struct {
	PatternBase
}

func (n * SubPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return '-',0,false

}

type MultiPattern struct {
	PatternBase
}

func (n * MultiPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return '*',0,false

}

type DivPattern struct {
	PatternBase
}

func (n * DivPattern)BuildFun(prefix string,r *bufio.Reader) (v interface{},pre rune,has_prerune bool) {
	return '/',0,false

}

func init() {

	var addRegexp=regexp.MustCompile(`\+`)

	Patterns=append(Patterns,&AddPattern{PatternBase{Regexp:addRegexp}})


	var subRegexp=regexp.MustCompile(`-`)

	Patterns=append(Patterns,&SubPattern{PatternBase{Regexp:subRegexp}})

	var multiRegexp=regexp.MustCompile(`\*`)

	Patterns=append(Patterns,&MultiPattern{PatternBase{Regexp:multiRegexp}})

	var divRegexp=regexp.MustCompile(`/`)

	Patterns=append(Patterns,&DivPattern{PatternBase{Regexp:divRegexp}})
}


