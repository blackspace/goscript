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
const DOUBLEAT = 57357
const FUNCTION = 57358
const CLASS = 57359
const DEF = 57360
const END = 57361
const DO = 57362
const POUNDCOMMENT = 57363
const DOUBLESLASHCOMMENT = 57364
const MULTILINECOMMENT = 57365
const GREATEREQUAL = 57366
const LESSEQUAL = 57367
const ELSE = 57368
const AND = 57369
const OR = 57370
const DOUBLEADD = 57371
const DOUBLESUB = 57372

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
	"DOUBLEAT",
	"FUNCTION",
	"CLASS",
	"DEF",
	"END",
	"DO",
	"POUNDCOMMENT",
	"DOUBLESLASHCOMMENT",
	"MULTILINECOMMENT",
	"'>'",
	"GREATEREQUAL",
	"'<'",
	"LESSEQUAL",
	"ELSE",
	"AND",
	"OR",
	"DOUBLEADD",
	"DOUBLESUB",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'.'",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"','",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:315

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 62,
	24, 0,
	25, 0,
	26, 0,
	27, 0,
	-2, 69,
	-1, 63,
	24, 0,
	25, 0,
	26, 0,
	27, 0,
	-2, 70,
	-1, 64,
	24, 0,
	25, 0,
	26, 0,
	27, 0,
	-2, 71,
	-1, 65,
	24, 0,
	25, 0,
	26, 0,
	27, 0,
	-2, 72,
}

const yyNprod = 78
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 404

var yyAct = [...]int{

	4, 40, 18, 82, 18, 130, 9, 10, 76, 83,
	116, 27, 127, 104, 128, 7, 39, 45, 46, 43,
	6, 5, 18, 87, 44, 72, 52, 42, 55, 136,
	56, 57, 58, 59, 60, 61, 62, 63, 64, 65,
	103, 104, 27, 69, 126, 18, 54, 53, 66, 44,
	2, 18, 42, 3, 41, 77, 28, 79, 80, 75,
	23, 22, 24, 20, 114, 27, 102, 74, 23, 22,
	24, 49, 20, 27, 43, 88, 18, 95, 27, 44,
	100, 101, 92, 33, 34, 99, 88, 18, 84, 84,
	91, 84, 108, 92, 98, 90, 89, 70, 100, 100,
	96, 91, 71, 28, 18, 112, 90, 89, 78, 106,
	51, 100, 110, 111, 115, 18, 18, 118, 119, 48,
	123, 122, 81, 18, 124, 117, 85, 131, 18, 131,
	18, 68, 1, 135, 138, 135, 107, 131, 18, 47,
	131, 18, 134, 135, 134, 138, 135, 133, 132, 133,
	132, 139, 134, 113, 93, 134, 12, 133, 132, 8,
	133, 132, 13, 17, 16, 121, 11, 28, 35, 36,
	37, 38, 0, 29, 30, 28, 45, 46, 31, 32,
	33, 34, 0, 44, 35, 36, 37, 38, 109, 29,
	30, 0, 0, 0, 31, 32, 33, 34, 23, 22,
	24, 14, 15, 27, 97, 0, 25, 26, 0, 0,
	19, 21, 23, 22, 24, 14, 15, 27, 0, 0,
	25, 26, 0, 0, 19, 21, 0, 23, 22, 24,
	14, 15, 27, 20, 125, 25, 26, 0, 0, 19,
	21, 23, 22, 24, 14, 15, 27, 20, 120, 25,
	26, 0, 0, 19, 21, 0, 23, 22, 24, 14,
	15, 27, 20, 73, 25, 26, 23, 22, 24, 21,
	94, 27, 0, 23, 22, 24, 20, 50, 27, 0,
	23, 22, 24, 14, 15, 27, 0, 0, 25, 26,
	0, 0, 105, 21, 94, 23, 22, 24, 14, 15,
	27, 0, 0, 25, 26, 23, 22, 24, 14, 15,
	27, 67, 0, 25, 26, 0, 86, 23, 22, 24,
	14, 15, 27, 0, 0, 25, 26, 0, 0, 0,
	20, 141, 0, 0, 23, 22, 24, 14, 15, 27,
	20, 140, 25, 26, 0, 0, 23, 22, 24, 14,
	15, 27, 20, 137, 25, 26, 0, 0, 19, 21,
	0, 35, 36, 37, 38, 0, 29, 30, 0, 20,
	129, 31, 32, 33, 34, 0, 20, 35, 36, 37,
	38, 20, 29, 30, 0, 0, 0, 31, 32, 33,
	34, 29, 30, 0, 0, 0, 31, 32, 33, 34,
	31, 32, 33, 34,
}
var yyPact = [...]int{

	342, -1000, 342, -1000, 353, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 262, 26, -1000, -14, 110,
	237, 101, -1000, -1000, -1000, 262, 2, -1000, -1000, 262,
	262, 262, 262, 262, 262, 262, 262, 262, 262, 353,
	11, 24, 269, 342, 93, -1000, -1000, -16, -1000, 223,
	-1000, 28, 337, -36, 64, 41, 366, 366, 47, 47,
	-1000, -1000, 362, 362, 362, 362, -1000, -1000, 15, 353,
	-1000, -1000, 80, -1000, 276, -1000, 56, 160, 33, -1000,
	262, 27, -2, -1000, -1000, 252, -1000, -1000, 353, -1000,
	-1000, -1000, -1000, -1000, 83, 144, 69, 69, 24, -1000,
	145, 353, 342, 25, 82, -1000, -1000, -31, -1000, 69,
	24, 24, -1000, 208, 342, -1000, 79, 24, -1000, -1000,
	-1000, 194, 5, -30, -1000, -1000, 330, -10, 313, -1000,
	-1000, 353, -1000, -1000, -1000, -1000, 301, -1000, -1000, 291,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 53, 166, 0, 164, 163, 21, 6, 15, 162,
	20, 159, 7, 23, 5, 156, 154, 139, 136, 9,
	3, 1, 50, 132, 131, 126, 14,
}
var yyR1 = [...]int{

	0, 23, 22, 22, 12, 12, 13, 13, 13, 13,
	13, 13, 25, 25, 16, 16, 16, 16, 18, 14,
	14, 14, 14, 14, 26, 26, 15, 15, 9, 10,
	21, 21, 11, 11, 17, 20, 20, 19, 24, 24,
	1, 1, 1, 1, 1, 1, 1, 6, 6, 6,
	6, 5, 5, 5, 5, 5, 4, 4, 8, 8,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 7, 7,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 5, 4, 1, 1, 1, 1,
	1, 1, 1, 2, 7, 6, 8, 7, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 4, 1, 3,
	1, 3, 7, 8, 1, 1, 3, 1, 1, 3,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 7, 6, 5, 4, 6, 3, 3, 2, 2,
	1, 1, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 1, 1, 1, 3, 2,
}
var yyChk = [...]int{

	-1000, -23, -22, -1, -3, -6, -10, -8, -11, -7,
	-12, -2, -15, -9, 7, 8, -4, -5, -21, 16,
	39, 17, 5, 4, 6, 12, 13, 9, -1, 29,
	30, 34, 35, 36, 37, 24, 25, 26, 27, -3,
	-21, 28, 41, 33, 38, 31, 32, -17, 9, -22,
	40, 9, -3, -10, 44, -21, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -7, 42, -24, -3,
	-1, 9, 41, 40, 39, -7, 44, -3, 44, 42,
	43, 42, -20, -19, 9, -25, 40, -13, -3, -6,
	-10, -8, -12, -16, 18, -3, 44, 44, -8, -7,
	-21, -3, 39, 42, 43, 40, -13, -18, 9, 44,
	-8, -8, -7, -22, 39, -19, 41, -8, -7, -7,
	40, -22, 42, -20, -7, 40, 39, 42, -26, 40,
	-14, -3, -6, -10, -8, -7, 39, 40, -14, -26,
	40, 40,
}
var yyDef = [...]int{

	0, -2, 1, 2, 40, 41, 42, 43, 44, 45,
	46, 60, 61, 62, 47, 0, 49, 50, 28, 0,
	0, 0, 73, 74, 75, 0, 0, 30, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 48,
	28, 0, 0, 0, 0, 58, 59, 0, 34, 0,
	77, 0, 0, 0, 0, 0, 63, 64, 65, 66,
	67, 68, -2, -2, -2, -2, 57, 26, 0, 38,
	29, 31, 0, 76, 0, 56, 0, 0, 0, 27,
	0, 0, 0, 35, 37, 0, 5, 12, 6, 7,
	8, 9, 10, 11, 0, 0, 0, 0, 0, 54,
	0, 39, 0, 0, 0, 4, 13, 0, 18, 0,
	0, 0, 53, 0, 0, 36, 0, 0, 55, 52,
	32, 0, 0, 0, 51, 33, 0, 0, 0, 15,
	24, 19, 20, 21, 22, 23, 0, 14, 25, 0,
	17, 16,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	41, 42, 36, 34, 43, 35, 38, 37, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 44,
	26, 33, 24, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 39, 3, 40,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 25, 27, 28, 29, 30, 31, 32,
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
		//line parser/parser.go.y:63
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:68
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:72
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:79
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, yyDollar[4].Exprs}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:83
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, nil}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:92
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:96
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 14:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:101
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, nil, yyDollar[6].Exprs}
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:105
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, nil, nil}
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:109
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[7].Exprs}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:113
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, yyDollar[4].Strings, nil}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:124
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:128
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:134
		{
			yyVAL.Expr = &ast.CalledExpr{yyDollar[1].Strings, nil}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:138
		{
			yyVAL.Expr = &ast.CalledExpr{yyDollar[1].Strings, yyDollar[3].Exprs}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:143
		{
			yyVAL.Expr = &ast.GetExpr{yyDollar[1].Strings}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:148
		{
			yyVAL.Expr = &ast.SetExpr{yyDollar[1].Strings, yyDollar[3].Expr}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:154
		{
			yyVAL.Strings = append(make([]string, 0, 16), yyDollar[1].String)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:158
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 32:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:165
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, nil, yyDollar[6].Exprs}
		}
	case 33:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:169
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[7].Exprs}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:176
		{
			yyVAL.Strings = append(make([]string, 0), yyDollar[1].String)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:180
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:187
		{

			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:192
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:208
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:212
		{
			yyVAL.Expr = &ast.ReturnExpr{yyDollar[2].Expr}
		}
	case 51:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:219
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 52:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:223
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:227
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:231
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 55:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:235
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:240
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:244
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:252
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Strings}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:256
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Strings}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:265
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:269
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:273
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:277
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:281
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:285
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:289
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:293
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:297
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:301
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:308
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:312
		{
			yyVAL.Expr = &ast.BlockExpr{}
		}
	}
	goto yystack /* stack new state and value */
}
