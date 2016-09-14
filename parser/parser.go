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
	yys     int
	Expr    ast.Expr
	Params  ast.Params
	Exprs   []ast.Expr
	String  string
	Strings []string
}

const NUMBER = 57346
const BOOL = 57347
const STRING = 57348
const BREAK = 57349
const RETURN = 57350
const WORD = 57351
const BLANKSPACE = 57352
const LFCR = 57353
const IF = 57354
const FOR = 57355
const SWITCH = 57356
const FUNCTION = 57357
const CLASS = 57358
const DOUBLEADD = 57359
const DOUBLESUB = 57360
const GREATEREQUAL = 57361
const LESSEQUAL = 57362
const ELSE = 57363
const AND = 57364
const OR = 57365

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"BOOL",
	"STRING",
	"BREAK",
	"RETURN",
	"WORD",
	"BLANKSPACE",
	"LFCR",
	"IF",
	"FOR",
	"SWITCH",
	"FUNCTION",
	"CLASS",
	"DOUBLEADD",
	"DOUBLESUB",
	"'>'",
	"GREATEREQUAL",
	"'<'",
	"LESSEQUAL",
	"ELSE",
	"AND",
	"OR",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'{'",
	"'}'",
	"'.'",
	"'('",
	"')'",
	"','",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:254

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	34, 23,
	-2, 62,
	-1, 65,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 58,
	-1, 66,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 59,
	-1, 67,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 60,
	-1, 68,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 61,
	-1, 71,
	34, 17,
	-2, 18,
}

const yyNprod = 63
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 263

var yyAct = [...]int{

	10, 79, 91, 45, 15, 2, 15, 80, 114, 95,
	8, 58, 83, 36, 37, 38, 39, 3, 30, 31,
	29, 32, 33, 34, 35, 106, 107, 15, 73, 49,
	92, 110, 57, 94, 95, 13, 12, 14, 4, 56,
	25, 86, 75, 53, 23, 15, 77, 74, 87, 13,
	12, 14, 41, 15, 25, 82, 89, 58, 44, 72,
	32, 33, 34, 35, 34, 35, 54, 29, 97, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 68, 23,
	46, 15, 85, 93, 42, 43, 100, 58, 103, 101,
	105, 15, 81, 92, 6, 84, 99, 15, 88, 71,
	113, 101, 101, 109, 52, 104, 48, 115, 111, 112,
	116, 29, 118, 119, 101, 13, 12, 14, 120, 1,
	25, 117, 96, 55, 90, 81, 26, 36, 37, 38,
	39, 70, 30, 31, 81, 32, 33, 34, 35, 13,
	12, 14, 18, 19, 25, 98, 102, 27, 28, 47,
	22, 24, 13, 12, 14, 18, 19, 25, 69, 51,
	27, 28, 7, 22, 24, 17, 23, 108, 11, 16,
	9, 5, 21, 13, 12, 14, 18, 19, 25, 23,
	76, 27, 28, 20, 22, 24, 13, 12, 14, 18,
	19, 25, 0, 0, 27, 28, 0, 22, 24, 0,
	23, 50, 0, 0, 0, 0, 36, 37, 38, 39,
	0, 30, 31, 23, 32, 33, 34, 35, 23, 36,
	37, 38, 39, 0, 30, 31, 0, 32, 33, 34,
	35, 13, 12, 14, 42, 43, 25, 13, 12, 14,
	0, 0, 25, 41, 0, 0, 0, 0, 30, 31,
	40, 32, 33, 34, 35, 0, 0, 0, 0, 0,
	0, 0, 78,
}
var yyPact = [...]int{

	182, -1000, 182, -1000, 200, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 217, -1000, -1000, -1000, 233,
	57, -1000, 97, 169, 95, -1000, 9, 233, 2, -1000,
	233, 233, 233, 233, 233, 233, 233, 233, 233, 233,
	90, 182, -1000, -1000, 200, -5, 13, 8, -1000, 148,
	-1000, 15, -1000, 227, 187, -25, 45, 26, -1000, 33,
	33, 35, 35, -1000, -1000, 224, 224, 224, 224, 7,
	22, -1000, -1000, 89, -1000, 21, -1000, 182, -1000, -2,
	-1000, 200, -1000, 31, 108, 48, 111, 182, -1000, 13,
	-10, -1000, -1000, 135, -1000, 233, -6, 78, 78, 13,
	-1000, 67, -1000, -27, -1000, -1000, 13, 84, -1000, -1000,
	78, 13, 13, -1000, -1000, -1000, -1000, 13, -1000, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 17, 94, 38, 3, 183, 172, 171, 0, 10,
	170, 169, 7, 168, 165, 162, 126, 159, 158, 131,
	2, 124, 5, 119, 1,
}
var yyR1 = [...]int{

	0, 23, 22, 22, 1, 1, 1, 1, 1, 1,
	1, 1, 13, 14, 14, 15, 17, 18, 19, 10,
	10, 11, 11, 16, 21, 21, 20, 24, 24, 12,
	8, 8, 7, 7, 7, 7, 2, 9, 9, 6,
	6, 6, 6, 6, 5, 5, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 5, 5, 6, 5, 1, 1, 1, 5,
	6, 3, 4, 1, 1, 3, 1, 1, 3, 1,
	3, 2, 1, 2, 1, 1, 3, 2, 2, 7,
	6, 5, 4, 6, 3, 3, 1, 1, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 1,
}
var yyChk = [...]int{

	-1000, -23, -22, -1, -3, -7, -2, -15, -9, -10,
	-8, -13, 5, 4, 6, -4, -11, -14, 7, 8,
	-5, -6, 15, 31, 16, 9, -16, 12, 13, -1,
	24, 25, 27, 28, 29, 30, 19, 20, 21, 22,
	33, 26, 17, 18, -3, -4, 23, -16, 9, -22,
	32, -17, 9, 34, -3, -2, 37, -4, 9, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -18,
	-19, 9, -1, 33, -8, 34, 32, 31, 35, -24,
	-12, -3, -8, 37, -3, 37, 34, 26, 9, 35,
	-21, -20, 9, -22, 35, 36, -3, 37, 37, -9,
	-8, -4, 35, -24, -1, -8, 35, 36, 32, -12,
	37, -9, -9, -8, 35, -8, -20, -9, -8, -8,
	-8,
}
var yyDef = [...]int{

	0, -2, 1, 2, 4, 5, 6, 7, 8, 9,
	10, 11, 46, 47, 48, 49, 50, 51, 32, 0,
	34, 35, 0, 0, 0, -2, 0, 0, 0, 3,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 37, 38, 33, 49, 0, 0, 23, 0,
	31, 0, 16, 0, 0, 0, 0, 0, 62, 52,
	53, 54, 55, 56, 57, -2, -2, -2, -2, 0,
	0, -2, 36, 0, 45, 0, 30, 0, 21, 0,
	27, 29, 44, 0, 0, 0, 0, 0, 17, 0,
	0, 24, 26, 0, 22, 0, 0, 0, 0, 0,
	42, 0, 13, 0, 15, 19, 0, 0, 12, 28,
	0, 0, 0, 41, 14, 20, 25, 0, 43, 40,
	39,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	34, 35, 29, 27, 36, 28, 33, 30, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 37,
	21, 26, 19, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 31, 3, 32,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 20, 22, 23,
	24, 25,
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
		//line parser/parser.go.y:54
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:59
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:63
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:78
		{
			yyVAL.Expr = &ast.Class{}
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:93
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, nil, yyDollar[5].Expr}
		}
	case 20:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:97
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[6].Expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:102
		{
			yyVAL.Expr = &ast.FuncCallExpr{Name: yyDollar[1].String}
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:106
		{
			yyVAL.Expr = &ast.FuncCallExpr{yyDollar[1].String, yyDollar[3].Exprs}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:113
		{
			yyVAL.Strings = append(make([]string, 0), yyDollar[1].String)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:117
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:124
		{

			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:129
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:136
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:140
		{
			yyVAL.Expr = &ast.BlockExpr{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:145
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:149
		{
			yyVAL.Expr = &ast.ReturnExpr{yyDollar[2].Expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:157
		{
			yyVAL.Expr = &ast.AssignExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:162
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Expr}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:166
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Expr}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:172
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 40:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:176
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:180
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:184
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 43:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:188
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:195
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:199
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:208
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:212
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:216
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:220
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:224
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:228
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:232
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:236
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:240
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:244
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:249
		{
			yyVAL.Expr = &ast.Variable{yyDollar[1].String}
		}
	}
	goto yystack /* stack new state and value */
}
