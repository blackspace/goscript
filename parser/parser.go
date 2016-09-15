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
const DOUBLEADD = 57357
const DOUBLESUB = 57358
const DOUBLEAT = 57359
const FUNCTION = 57360
const CLASS = 57361
const DEF = 57362
const END = 57363
const DO = 57364
const GREATEREQUAL = 57365
const LESSEQUAL = 57366
const ELSE = 57367
const AND = 57368
const OR = 57369

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
	"'@'",
	"'.'",
	"','",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:301

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	37, 47,
	-2, 86,
	-1, 65,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 82,
	-1, 66,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 83,
	-1, 67,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 84,
	-1, 68,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 85,
	-1, 71,
	37, 33,
	-2, 34,
}

const yyNprod = 87
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 431

var yyAct = [...]int{

	4, 90, 3, 91, 80, 29, 79, 11, 7, 69,
	83, 139, 13, 12, 14, 140, 94, 25, 45, 15,
	44, 15, 58, 73, 70, 9, 42, 43, 54, 92,
	130, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 41, 15, 8, 72, 138, 92, 57, 118, 86,
	108, 40, 29, 10, 81, 56, 127, 84, 134, 106,
	15, 36, 37, 38, 39, 75, 30, 31, 15, 32,
	33, 34, 35, 58, 117, 89, 6, 118, 95, 105,
	123, 53, 106, 150, 107, 102, 98, 81, 5, 137,
	115, 23, 77, 114, 95, 158, 15, 157, 87, 23,
	74, 102, 98, 100, 112, 55, 15, 81, 82, 41,
	120, 122, 15, 46, 121, 13, 12, 14, 26, 100,
	25, 99, 129, 32, 33, 34, 35, 112, 112, 110,
	154, 101, 135, 34, 35, 42, 43, 99, 141, 111,
	141, 47, 112, 116, 2, 58, 144, 101, 144, 113,
	92, 141, 124, 125, 97, 152, 15, 141, 15, 144,
	160, 161, 156, 88, 126, 144, 96, 131, 49, 15,
	97, 128, 152, 153, 155, 15, 15, 15, 132, 133,
	71, 145, 96, 145, 52, 136, 48, 13, 12, 14,
	93, 146, 25, 146, 145, 1, 13, 12, 14, 51,
	145, 25, 30, 31, 146, 32, 33, 34, 35, 147,
	146, 103, 17, 16, 143, 21, 143, 13, 12, 14,
	20, 0, 25, 0, 0, 85, 142, 143, 142, 0,
	78, 0, 0, 143, 0, 0, 0, 0, 0, 142,
	0, 36, 37, 38, 39, 142, 30, 31, 0, 32,
	33, 34, 35, 13, 12, 14, 18, 19, 25, 0,
	109, 27, 28, 0, 0, 0, 149, 13, 12, 14,
	18, 19, 25, 0, 0, 27, 28, 0, 0, 0,
	149, 0, 0, 0, 23, 159, 0, 0, 148, 0,
	0, 0, 13, 12, 14, 18, 19, 25, 23, 151,
	27, 28, 148, 0, 0, 149, 13, 12, 14, 18,
	19, 25, 0, 0, 27, 28, 0, 0, 0, 0,
	22, 24, 104, 23, 0, 0, 0, 148, 0, 0,
	0, 13, 12, 14, 18, 19, 25, 23, 119, 27,
	28, 0, 0, 0, 0, 22, 24, 0, 0, 0,
	13, 12, 14, 18, 19, 25, 0, 0, 27, 28,
	0, 0, 23, 76, 22, 24, 0, 0, 0, 13,
	12, 14, 18, 19, 25, 0, 0, 27, 28, 0,
	0, 23, 50, 22, 24, 104, 0, 13, 12, 14,
	18, 19, 25, 0, 0, 27, 28, 0, 0, 0,
	23, 22, 24, 36, 37, 38, 39, 0, 30, 31,
	0, 32, 33, 34, 35, 23, 0, 0, 23, 36,
	37, 38, 39, 0, 30, 31, 0, 32, 33, 34,
	35,
}
var yyPact = [...]int{

	383, -1000, 383, -1000, 396, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 11, -1000, -1000, -1000, 213,
	86, -1000, 177, 346, 175, -1000, 44, 213, 13, -1000,
	213, 213, 213, 213, 213, 213, 213, 213, 213, 213,
	171, 383, -1000, -1000, 396, -17, 56, 28, -1000, 327,
	-1000, 57, -1000, 192, 380, -32, 183, 79, -1000, 92,
	92, 100, 100, -1000, -1000, 174, 174, 174, 174, 12,
	68, -1000, -1000, 154, -1000, 37, -1000, 365, -1000, 41,
	-1000, 396, -1000, 8, 218, 64, 111, 383, -1000, 56,
	36, -1000, -1000, 302, -1000, 396, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 154, -1000, 213, 38, 136, 136,
	56, -1000, 120, -1000, 18, -1000, -1000, 56, 141, -1000,
	-1000, -7, -1000, 136, 56, 56, -1000, -1000, -1000, -1000,
	20, 56, -1000, -1000, 54, 7, -1000, 288, 48, 263,
	-1000, 396, -1000, -1000, -1000, -1000, -1000, -1000, 121, 121,
	288, -1000, -1000, 67, -1000, 65, 249, 383, 383, -1000,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 2, 76, 0, 18, 220, 215, 88, 53, 43,
	25, 213, 4, 7, 16, 15, 212, 8, 211, 209,
	118, 199, 9, 24, 3, 1, 144, 195, 6, 190,
	11,
}
var yyR1 = [...]int{

	0, 27, 26, 26, 13, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 29, 29, 18, 18, 15, 15,
	15, 15, 15, 15, 15, 30, 30, 19, 19, 16,
	16, 17, 21, 22, 23, 1, 1, 1, 1, 1,
	1, 1, 1, 10, 10, 11, 11, 20, 25, 25,
	24, 28, 28, 12, 8, 8, 7, 7, 7, 7,
	2, 9, 9, 6, 6, 6, 6, 6, 5, 5,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 5, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 7, 8, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 4, 4, 5,
	6, 5, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 5, 6, 3, 4, 1, 1, 3,
	1, 1, 3, 1, 3, 2, 1, 2, 1, 1,
	3, 2, 2, 7, 6, 5, 4, 6, 3, 3,
	1, 1, 1, 1, 1, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 1,
}
var yyChk = [...]int{

	-1000, -27, -26, -1, -3, -7, -2, -17, -9, -10,
	-8, -13, 5, 4, 6, -4, -11, -16, 7, 8,
	-5, -6, 18, 35, 19, 9, -20, 12, 13, -1,
	28, 29, 31, 32, 33, 34, 23, 24, 25, 26,
	40, 30, 15, 16, -3, -4, 27, -20, 9, -26,
	36, -21, 9, 37, -3, -2, 42, -4, 9, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -22,
	-23, 9, -1, 40, -8, 37, 36, 35, 38, -28,
	-12, -3, -8, 42, -3, 42, 37, 30, 9, 38,
	-25, -24, 9, -29, -14, -3, -7, -2, -17, -9,
	-10, -8, -13, -18, 20, 38, 41, -3, 42, 42,
	-9, -8, -4, 38, -28, -1, -8, 38, 41, 36,
	-14, -22, -12, 42, -9, -9, -8, 38, -8, -24,
	37, -9, -8, -8, 38, -25, -8, 35, 38, -30,
	-15, -3, -7, -2, -17, -9, -8, -19, 39, 17,
	35, 36, -15, -23, 9, -23, -30, 30, 30, 36,
	-1, -1,
}
var yyDef = [...]int{

	0, -2, 1, 2, 35, 36, 37, 38, 39, 40,
	41, 42, 70, 71, 72, 73, 74, 75, 56, 0,
	58, 59, 0, 0, 0, -2, 0, 0, 0, 3,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 61, 62, 57, 73, 0, 0, 47, 0,
	55, 0, 32, 0, 0, 0, 0, 0, 86, 76,
	77, 78, 79, 80, 81, -2, -2, -2, -2, 0,
	0, -2, 60, 0, 69, 0, 54, 0, 45, 0,
	51, 53, 68, 0, 0, 0, 0, 0, 33, 0,
	0, 48, 50, 0, 14, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 0, 46, 0, 0, 0, 0,
	0, 66, 0, 29, 0, 31, 43, 0, 0, 4,
	15, 0, 52, 0, 0, 0, 65, 30, 44, 49,
	0, 0, 67, 64, 0, 0, 63, 0, 0, 0,
	25, 18, 19, 20, 21, 22, 23, 24, 0, 0,
	0, 16, 26, 0, 34, 0, 0, 0, 0, 17,
	27, 28,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	37, 38, 33, 31, 41, 32, 40, 34, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 42,
	25, 30, 23, 3, 39, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 35, 3, 36,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 24, 26, 27, 28, 29,
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
		//line parser/parser.go.y:56
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:61
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:65
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:71
		{
			yyVAL.Expr = &ast.Class{}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:80
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:84
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:89
		{
			yyVAL.Expr = &ast.Method{}
		}
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:93
		{
			yyVAL.Expr = &ast.Method{}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:101
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:105
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:110
		{
			yyVAL.Expr = &ast.Attribute{}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:114
		{
			yyVAL.Expr = &ast.ClassAttribute{}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:140
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, nil, yyDollar[5].Expr}
		}
	case 44:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:144
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[6].Expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:149
		{
			yyVAL.Expr = &ast.FuncCallExpr{yyDollar[1].String, nil}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:153
		{
			yyVAL.Expr = &ast.FuncCallExpr{yyDollar[1].String, yyDollar[3].Exprs}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:160
		{
			yyVAL.Strings = append(make([]string, 0), yyDollar[1].String)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:164
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:171
		{

			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:176
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:183
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:187
		{
			yyVAL.Expr = &ast.BlockExpr{}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:192
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:196
		{
			yyVAL.Expr = &ast.ReturnExpr{yyDollar[2].Expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:204
		{
			yyVAL.Expr = &ast.AssignExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:209
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Expr}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:213
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Expr}
		}
	case 63:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:219
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 64:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:223
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:227
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:231
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 67:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:235
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:242
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:246
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:255
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:259
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:263
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:267
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:271
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:275
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:279
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:283
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:287
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:291
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:296
		{
			yyVAL.Expr = &ast.Variable{yyDollar[1].String}
		}
	}
	goto yystack /* stack new state and value */
}
