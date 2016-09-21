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
const DOUBLEADD = 57357
const DOUBLESUB = 57358
const DOUBLEAT = 57359
const FUNCTION = 57360
const CLASS = 57361
const DEF = 57362
const END = 57363
const DO = 57364
const POUNDCOMMENT = 57365
const DOUBLESLASHCOMMENT = 57366
const MULTILINECOMMENT = 57367
const GREATEREQUAL = 57368
const LESSEQUAL = 57369
const ELSE = 57370
const AND = 57371
const OR = 57372

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
	"DOUBLEADD",
	"DOUBLESUB",
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
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"'.'",
	"','",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:335

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	40, 53,
	42, 40,
	-2, 93,
	-1, 68,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	-2, 87,
	-1, 69,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	-2, 88,
	-1, 70,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	-2, 89,
	-1, 71,
	26, 0,
	27, 0,
	28, 0,
	29, 0,
	-2, 90,
	-1, 77,
	40, 38,
	-2, 39,
	-1, 127,
	42, 40,
	-2, 38,
}

const yyNprod = 94
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 536

var yyAct = [...]int{

	46, 15, 93, 15, 94, 83, 47, 23, 75, 23,
	11, 98, 151, 157, 10, 9, 10, 3, 86, 76,
	31, 149, 61, 13, 12, 14, 15, 82, 27, 95,
	137, 60, 23, 13, 12, 14, 19, 20, 27, 10,
	95, 29, 30, 15, 166, 148, 122, 24, 26, 23,
	13, 12, 14, 15, 73, 27, 10, 59, 49, 23,
	72, 159, 74, 112, 136, 90, 10, 25, 79, 91,
	31, 85, 141, 78, 38, 39, 40, 41, 56, 32,
	33, 15, 34, 35, 36, 37, 25, 23, 95, 116,
	88, 106, 15, 89, 105, 61, 104, 15, 23, 13,
	12, 14, 115, 23, 27, 10, 120, 106, 124, 119,
	105, 170, 104, 116, 116, 126, 128, 125, 118, 165,
	92, 13, 12, 14, 25, 158, 27, 135, 132, 146,
	116, 156, 8, 48, 8, 134, 117, 25, 147, 142,
	122, 133, 80, 110, 139, 140, 143, 15, 36, 37,
	15, 160, 145, 23, 42, 144, 23, 8, 81, 15,
	154, 6, 162, 6, 15, 23, 15, 61, 2, 15,
	23, 15, 23, 15, 8, 23, 162, 23, 95, 23,
	163, 162, 43, 44, 8, 162, 6, 168, 121, 28,
	122, 58, 172, 109, 52, 110, 152, 4, 127, 4,
	42, 32, 33, 6, 34, 35, 36, 37, 34, 35,
	36, 37, 103, 6, 50, 43, 44, 45, 77, 55,
	114, 51, 4, 8, 96, 1, 57, 54, 103, 62,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 4,
	18, 101, 107, 17, 130, 131, 16, 155, 7, 4,
	7, 22, 6, 84, 21, 0, 87, 101, 0, 153,
	5, 138, 5, 13, 12, 14, 19, 20, 27, 0,
	0, 29, 30, 7, 13, 12, 14, 99, 0, 27,
	0, 0, 0, 111, 0, 5, 0, 84, 4, 0,
	7, 0, 0, 99, 0, 0, 0, 25, 174, 0,
	7, 0, 5, 0, 0, 0, 0, 84, 38, 39,
	40, 41, 5, 32, 33, 0, 34, 35, 36, 37,
	0, 0, 0, 0, 0, 0, 129, 0, 102, 0,
	0, 0, 0, 13, 12, 14, 19, 20, 27, 7,
	100, 29, 30, 0, 102, 0, 0, 24, 26, 108,
	0, 5, 38, 39, 40, 41, 100, 32, 33, 0,
	34, 35, 36, 37, 0, 0, 0, 25, 123, 0,
	113, 13, 12, 14, 19, 20, 27, 0, 0, 29,
	30, 0, 0, 0, 0, 24, 26, 108, 13, 12,
	14, 19, 20, 27, 0, 0, 29, 30, 0, 0,
	0, 0, 24, 26, 0, 25, 97, 13, 12, 14,
	19, 20, 27, 0, 0, 29, 30, 0, 0, 0,
	0, 0, 25, 53, 13, 12, 14, 19, 20, 27,
	0, 0, 29, 30, 0, 13, 12, 14, 19, 20,
	27, 25, 173, 29, 30, 13, 12, 14, 19, 20,
	27, 0, 0, 29, 30, 0, 0, 0, 25, 171,
	0, 0, 0, 13, 12, 14, 19, 20, 27, 25,
	169, 29, 30, 13, 12, 14, 19, 20, 27, 25,
	167, 29, 30, 13, 12, 14, 19, 20, 27, 0,
	0, 29, 30, 0, 0, 0, 0, 25, 164, 0,
	0, 13, 12, 14, 19, 20, 27, 25, 161, 29,
	30, 0, 0, 0, 0, 24, 26, 25, 150, 38,
	39, 40, 41, 0, 32, 33, 0, 34, 35, 36,
	37, 0, 0, 0, 0, 25,
}
var yyPact = [...]int{

	497, -1000, 497, -1000, 493, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 167, -1000, -1000, -1000, -1000,
	270, 103, -1000, 16, 212, 384, 210, -1000, 38, 270,
	13, -1000, 270, 270, 270, 270, 270, 270, 270, 270,
	270, 270, 497, -1000, -1000, 493, -1000, 12, 99, 209,
	33, -1000, 29, -1000, 104, -1000, 117, 48, -26, 46,
	121, -1000, 174, 174, 112, 112, -1000, -1000, 170, 170,
	170, 170, -1000, 209, -1000, 25, 36, -1000, 79, -1000,
	367, -1000, 152, -1000, 493, -1000, 19, 326, 86, -1000,
	95, 497, 99, 147, -1000, -1000, 329, -1000, -1000, 493,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 189, -1000,
	270, 282, 158, 158, 99, -1000, 200, -1000, 100, -1000,
	-1000, 99, 169, -1000, -1000, 24, -12, -1000, -1000, 158,
	99, 99, -1000, -1000, -1000, -1000, 31, 146, 99, -1000,
	-1000, 91, 97, 5, -1000, -1000, 479, 87, 20, 469,
	-1000, -1000, 493, -1000, -1000, -1000, -1000, -1000, 459, 81,
	3, -1000, -1000, 441, -1000, 431, 73, -1000, 420, -1000,
	403, -1000, 259, -1000, -1000,
}
var yyPgo = [...]int{

	0, 17, 160, 196, 0, 254, 251, 259, 13, 131,
	15, 246, 5, 10, 11, 12, 243, 242, 247, 240,
	189, 227, 8, 19, 6, 4, 2, 168, 225, 27,
	224, 21,
}
var yyR1 = [...]int{

	0, 28, 27, 27, 13, 13, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 30, 30, 17, 17, 17,
	17, 17, 17, 17, 17, 15, 15, 15, 15, 15,
	15, 31, 31, 16, 16, 18, 19, 21, 22, 23,
	24, 1, 1, 1, 1, 1, 1, 1, 1, 10,
	10, 11, 11, 20, 26, 26, 25, 29, 29, 12,
	7, 7, 7, 7, 2, 9, 9, 6, 6, 6,
	6, 6, 5, 5, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 8, 8, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 5, 4, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 7, 6, 8,
	7, 9, 8, 10, 9, 1, 1, 1, 1, 1,
	1, 1, 2, 5, 6, 5, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 5,
	6, 3, 4, 1, 1, 3, 1, 1, 3, 1,
	1, 2, 1, 1, 3, 2, 2, 7, 6, 5,
	4, 6, 3, 3, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 1,
}
var yyChk = [...]int{

	-1000, -28, -27, -1, -3, -7, -2, -18, -9, -10,
	-8, -13, 5, 4, 6, -4, -11, -16, -19, 7,
	8, -5, -6, -24, 18, 38, 19, 9, -20, 12,
	13, -1, 31, 32, 34, 35, 36, 37, 26, 27,
	28, 29, 33, 15, 16, -3, -4, -24, 30, 42,
	-20, 9, -27, 39, -21, 9, 40, -3, -2, 44,
	-4, 9, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -1, 42, -8, -22, -23, 9, 40, 39,
	38, 41, -29, -12, -3, -8, 44, -3, 44, -23,
	40, 33, 41, -26, -25, 9, -30, 39, -14, -3,
	-7, -2, -18, -9, -10, -8, -13, -17, 20, 41,
	43, -3, 44, 44, -9, -8, -4, 41, -29, -1,
	-8, 41, 43, 39, -14, -22, -24, 9, -12, 44,
	-9, -9, -8, 41, -8, -25, 40, 42, -9, -8,
	-8, 41, -26, -22, 9, -8, 38, 41, 40, -31,
	39, -15, -3, -7, -2, -18, -9, -8, 38, 41,
	-26, 39, -15, -31, 39, 38, 41, 39, -31, 39,
	38, 39, -31, 39, 39,
}
var yyDef = [...]int{

	0, -2, 1, 2, 41, 42, 43, 44, 45, 46,
	47, 48, 74, 75, 76, 77, 78, 79, 80, 60,
	0, 62, 63, 0, 0, 0, 0, -2, 0, 0,
	0, 3, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 65, 66, 61, 77, 0, 0, 0,
	0, 53, 0, 92, 0, 37, 0, 0, 0, 0,
	0, 93, 81, 82, 83, 84, 85, 86, -2, -2,
	-2, -2, 64, 0, 73, 0, 36, -2, 0, 91,
	0, 51, 0, 57, 59, 72, 0, 0, 0, 36,
	0, 0, 0, 0, 54, 56, 0, 5, 15, 6,
	7, 8, 9, 10, 11, 12, 13, 14, 0, 52,
	0, 0, 0, 0, 0, 70, 0, 33, 0, 35,
	49, 0, 0, 4, 16, 0, 0, -2, 58, 0,
	0, 0, 69, 34, 50, 55, 0, 0, 0, 71,
	68, 0, 0, 0, 38, 67, 0, 0, 0, 0,
	18, 31, 25, 26, 27, 28, 29, 30, 0, 0,
	0, 17, 32, 0, 20, 0, 0, 19, 0, 22,
	0, 21, 0, 24, 23,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	40, 41, 36, 34, 43, 35, 42, 37, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 44,
	28, 33, 26, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 38, 3, 39,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 27, 29, 30, 31, 32,
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
		//line parser/parser.go.y:57
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:62
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:66
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:73
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, yyDollar[4].Exprs}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:77
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, nil}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:86
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:90
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:95
		{
			yyVAL.Expr = &ast.MethodDefineExpr{"", yyDollar[2].String, nil, yyDollar[6].Exprs}
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:99
		{
			yyVAL.Expr = &ast.MethodDefineExpr{"", yyDollar[2].String, nil, nil}
		}
	case 19:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:103
		{
			yyVAL.Expr = &ast.MethodDefineExpr{"", yyDollar[2].String, yyDollar[4].Strings, yyDollar[7].Exprs}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:107
		{
			yyVAL.Expr = &ast.MethodDefineExpr{"", yyDollar[2].String, yyDollar[4].Strings, nil}
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser/parser.go.y:111
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, yyDollar[4].String, nil, yyDollar[8].Exprs}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:115
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, yyDollar[4].String, nil, nil}
		}
	case 23:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line parser/parser.go.y:119
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, yyDollar[4].String, yyDollar[6].Strings, yyDollar[9].Exprs}
		}
	case 24:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser/parser.go.y:123
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, yyDollar[4].String, yyDollar[6].Strings, nil}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:131
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:135
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:141
		{
			yyVAL.Expr = &ast.MethodCalledExpr{yyDollar[1].String, yyDollar[3].String, nil}
		}
	case 34:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:145
		{
			yyVAL.Expr = &ast.MethodCalledExpr{yyDollar[1].String, yyDollar[3].String, yyDollar[5].Exprs}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:149
		{
			yyVAL.Expr = &ast.AttributeSetExpr{yyDollar[1].String, yyDollar[3].String, yyDollar[5].Expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:154
		{
			yyVAL.Expr = &ast.AttributeExpr{yyDollar[1].String, yyDollar[3].String}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:176
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, nil, yyDollar[5].Expr}
		}
	case 50:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:180
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[6].Expr}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:185
		{
			yyVAL.Expr = &ast.FuncCallExpr{yyDollar[1].String, nil}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:189
		{
			yyVAL.Expr = &ast.FuncCallExpr{yyDollar[1].String, yyDollar[3].Exprs}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:196
		{
			yyVAL.Strings = append(make([]string, 0), yyDollar[1].String)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:200
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:207
		{

			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:212
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:220
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:224
		{
			yyVAL.Expr = &ast.ReturnExpr{yyDollar[2].Expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:232
		{
			yyVAL.Expr = &ast.AssignExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:237
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Expr}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:241
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Expr}
		}
	case 67:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:247
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 68:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:251
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:255
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:259
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:263
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:268
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:272
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:282
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:286
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:290
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:294
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:298
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:302
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:306
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:310
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:314
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:318
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:323
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:327
		{
			yyVAL.Expr = &ast.BlockExpr{}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:332
		{
			yyVAL.Expr = &ast.Variable{yyDollar[1].String}
		}
	}
	goto yystack /* stack new state and value */
}
