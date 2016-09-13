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
const BREAK = 57350
const BLANKSPACE = 57351
const LFCR = 57352
const WORD = 57353
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
	"VARIABLE",
	"BOOL",
	"STRING",
	"BREAK",
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
	"';'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:174

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 50,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 36,
	-1, 51,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 37,
	-1, 52,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 38,
	-1, 53,
	19, 0,
	20, 0,
	21, 0,
	22, 0,
	-2, 39,
}

const yyNprod = 40
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 179

var yyAct = [...]int{

	10, 5, 56, 8, 28, 29, 30, 31, 36, 22,
	23, 43, 24, 25, 26, 27, 59, 68, 69, 18,
	32, 39, 26, 27, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 35, 68, 55, 7, 1, 42,
	58, 33, 34, 18, 60, 28, 29, 30, 31, 9,
	22, 23, 6, 24, 25, 26, 27, 62, 41, 65,
	3, 63, 67, 21, 16, 66, 15, 72, 70, 71,
	0, 74, 75, 73, 76, 28, 29, 30, 31, 0,
	22, 23, 0, 24, 25, 26, 27, 18, 12, 14,
	11, 13, 4, 54, 2, 0, 19, 20, 21, 17,
	12, 14, 11, 13, 4, 33, 34, 0, 19, 20,
	0, 17, 0, 37, 32, 18, 57, 12, 40, 11,
	13, 12, 14, 11, 13, 4, 0, 18, 38, 19,
	20, 0, 17, 28, 29, 30, 31, 0, 22, 23,
	0, 24, 25, 26, 27, 0, 64, 0, 18, 12,
	40, 11, 13, 22, 23, 0, 24, 25, 26, 27,
	24, 25, 26, 27, 12, 40, 11, 13, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 61,
}
var yyPact = [...]int{

	117, -1000, 117, -1000, -1000, 114, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 88, 11, -1000, -26, 96, 160,
	6, -1000, 160, 160, 160, 160, 160, 160, 160, 160,
	160, 160, 117, -1000, -1000, -12, -33, 84, -1000, 56,
	-1000, -17, 145, -6, 133, 133, -7, -7, -1000, -1000,
	129, 129, 129, 129, -1000, -1000, -12, -1000, -1000, 113,
	26, 12, -1000, -15, 30, 30, -12, -1000, 24, 30,
	-12, -12, -1000, -12, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 60, 37, 1, 66, 64, 52, 0, 3, 49,
	94, 38,
}
var yyR1 = [...]int{

	0, 11, 10, 10, 1, 1, 1, 1, 1, 1,
	1, 7, 7, 6, 6, 2, 8, 8, 5, 5,
	5, 5, 5, 9, 4, 4, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 3, 2, 1, 1, 3, 2, 2, 7, 6,
	5, 4, 6, 4, 3, 3, 1, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
}
var yyChk = [...]int{

	-1000, -11, -10, -1, 8, -3, -6, -2, -8, -9,
	-7, 6, 4, 7, 5, -4, -5, 15, 31, 12,
	13, -1, 24, 25, 27, 28, 29, 30, 19, 20,
	21, 22, 26, 17, 18, 23, 34, -10, 32, -3,
	5, -2, 33, 5, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -1, -7, 35, 32, -7, 33,
	-3, 33, -7, -3, 33, 33, -8, -7, 5, 33,
	-8, -8, -7, -8, -7, -7, -7,
}
var yyDef = [...]int{

	0, -2, 1, 2, 4, 5, 6, 7, 8, 9,
	10, 26, 27, 28, 29, 13, 14, 0, 0, 0,
	0, 3, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 16, 17, 0, 0, 0, 12, 0,
	29, 0, 0, 0, 30, 31, 32, 33, 34, 35,
	-2, -2, -2, -2, 15, 25, 0, 11, 24, 0,
	0, 0, 23, 0, 0, 0, 0, 21, 0, 0,
	0, 0, 20, 0, 22, 19, 18,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	34, 35, 29, 27, 3, 28, 3, 30, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 33,
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
		//line parser/parser.go.y:39
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:44
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:48
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:54
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:66
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:70
		{
			yyVAL.Expr = &ast.BlockExpr{}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:79
		{
			yyVAL.Expr = &ast.AssignExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:84
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Expr}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:88
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Expr}
		}
	case 18:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:94
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 19:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:98
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:102
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:106
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:110
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:117
		{
			yyVAL.Expr = &ast.FuncExpr{}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:122
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:126
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:135
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:139
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:143
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:147
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:151
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:155
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:159
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:163
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:167
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:171
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	}
	goto yystack /* stack new state and value */
}
