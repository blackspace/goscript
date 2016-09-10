//line parser/parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser/parser.go.y:3
import (
	"goscript/ast"
)

var ParseResult []ast.Expr

//line parser/parser.go.y:14
type yySymType struct {
	yys   int
	Expr  ast.Expr
	Exprs []ast.Expr
}

const NUMBER = 57346
const VARIABLE = 57347
const BOOL = 57348
const STRING = 57349
const BLANKSPACE = 57350
const LFCR = 57351
const WORD = 57352
const IF = 57353
const FOR = 57354
const SWITCH = 57355
const FUNCTION = 57356
const CLASS = 57357
const DOUBLEADD = 57358
const DOUBLESUB = 57359
const ELSE = 57360
const AND = 57361
const OR = 57362

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"VARIABLE",
	"BOOL",
	"STRING",
	"BLANKSPACE",
	"LFCR",
	"WORD",
	"IF",
	"FOR",
	"SWITCH",
	"FUNCTION",
	"CLASS",
	"DOUBLEADD",
	"DOUBLESUB",
	"ELSE",
	"AND",
	"OR",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"';'",
	"'{'",
	"'}'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:116

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 25
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 91

var yyAct = [...]int{

	3, 46, 36, 15, 44, 2, 8, 10, 7, 9,
	20, 21, 25, 13, 14, 28, 16, 17, 38, 18,
	19, 20, 21, 35, 37, 8, 10, 7, 9, 1,
	48, 5, 13, 14, 18, 19, 20, 21, 12, 41,
	15, 15, 39, 40, 11, 45, 23, 24, 15, 43,
	6, 22, 47, 8, 10, 7, 9, 0, 0, 0,
	13, 14, 16, 17, 4, 18, 19, 20, 21, 0,
	0, 0, 8, 10, 7, 9, 0, 42, 26, 13,
	14, 29, 30, 31, 32, 33, 34, 8, 27, 7,
	9,
}
var yyPact = [...]int{

	68, -1000, 68, -1000, 43, -1000, -1000, -1000, -1000, -1000,
	30, -6, -1000, 83, 68, -1000, 83, 83, 83, 83,
	83, 83, 68, -1000, -1000, -25, -3, -1000, -8, 12,
	12, -14, -14, -1000, -1000, -1000, 68, 68, 68, 49,
	21, -22, -1000, -1000, 68, -26, 68, 2, -1000,
}
var yyPgo = [...]int{

	0, 0, 50, 64, 44, 38, 31, 5, 29,
}
var yyR1 = [...]int{

	0, 8, 7, 7, 1, 1, 1, 6, 6, 5,
	4, 4, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 1, 1, 1, 1, 1, 9,
	5, 5, 3, 2, 2, 1, 1, 1, 1, 3,
	3, 3, 3, 3, 3,
}
var yyChk = [...]int{

	-1000, -8, -7, -1, -3, -6, -2, 6, 4, 7,
	5, -4, -5, 11, 12, -1, 19, 20, 22, 23,
	24, 25, 21, 16, 17, 18, -3, 5, -1, -3,
	-3, -3, -3, -3, -3, -1, 27, 27, 26, -7,
	-7, -1, 28, 28, 26, -1, 27, -7, 28,
}
var yyDef = [...]int{

	0, -2, 1, 2, 4, 5, 6, 15, 16, 17,
	18, 7, 8, 0, 0, 3, 0, 0, 0, 0,
	0, 0, 0, 13, 14, 0, 0, 18, 0, 19,
	20, 21, 22, 23, 24, 12, 0, 0, 0, 0,
	0, 0, 11, 10, 0, 0, 0, 0, 9,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 24, 22, 3, 23, 3, 25, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 26,
	3, 21, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 27, 3, 28,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:37
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:42
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:46
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)

		}
	case 9:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser/parser.go.y:61
		{
			yyVAL.Expr = &ast.ForExpr{}
		}
	case 10:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:67
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[4].Exprs, nil}
		}
	case 11:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:71
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[4].Exprs

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:78
		{
			yyVAL.Expr = &ast.AssignExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:82
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Expr}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:86
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:93
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:97
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:101
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:105
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:109
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:113
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	}
	goto yystack /* stack new state and value */
}
