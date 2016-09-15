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

//line parser/parser.go.y:306

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 28,
	37, 47,
	-2, 87,
	-1, 66,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 75,
	-1, 67,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 76,
	-1, 68,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 77,
	-1, 69,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 78,
	-1, 74,
	37, 33,
	-2, 34,
}

const yyNprod = 88
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 431

var yyAct = [...]int{

	4, 42, 17, 140, 17, 91, 92, 85, 3, 72,
	95, 30, 80, 70, 84, 41, 141, 58, 131, 88,
	76, 17, 11, 9, 59, 7, 151, 54, 139, 57,
	73, 119, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 138, 6, 22, 21, 23, 17, 128, 28,
	56, 113, 17, 10, 75, 118, 19, 81, 119, 30,
	86, 22, 21, 23, 58, 93, 28, 8, 46, 47,
	78, 55, 22, 21, 23, 5, 159, 28, 93, 96,
	17, 106, 107, 45, 111, 112, 43, 155, 113, 86,
	19, 17, 158, 44, 135, 96, 17, 71, 116, 82,
	89, 103, 101, 115, 99, 121, 114, 90, 79, 111,
	111, 22, 21, 23, 86, 122, 28, 103, 101, 45,
	99, 127, 98, 35, 36, 111, 130, 33, 34, 35,
	36, 2, 102, 46, 47, 58, 110, 136, 98, 142,
	17, 142, 17, 93, 117, 83, 100, 87, 102, 74,
	109, 50, 142, 17, 97, 157, 53, 153, 142, 17,
	17, 17, 100, 126, 145, 49, 145, 161, 162, 29,
	97, 94, 129, 1, 153, 124, 125, 145, 133, 134,
	154, 156, 144, 145, 144, 52, 137, 148, 48, 104,
	25, 132, 147, 24, 147, 144, 16, 22, 21, 23,
	15, 144, 28, 12, 0, 147, 146, 0, 146, 0,
	0, 147, 0, 0, 143, 0, 143, 31, 32, 146,
	33, 34, 35, 36, 0, 146, 0, 143, 0, 37,
	38, 39, 40, 143, 31, 32, 0, 33, 34, 35,
	36, 37, 38, 39, 40, 0, 31, 32, 123, 33,
	34, 35, 36, 22, 21, 23, 13, 14, 28, 0,
	108, 26, 27, 0, 0, 0, 150, 22, 21, 23,
	13, 14, 28, 0, 0, 26, 27, 0, 0, 0,
	150, 0, 0, 0, 19, 160, 0, 0, 149, 0,
	0, 0, 22, 21, 23, 13, 14, 28, 19, 152,
	26, 27, 149, 0, 0, 150, 22, 21, 23, 13,
	14, 28, 0, 0, 26, 27, 0, 0, 0, 0,
	18, 20, 105, 19, 0, 0, 0, 149, 0, 0,
	0, 22, 21, 23, 13, 14, 28, 19, 120, 26,
	27, 0, 0, 0, 0, 18, 20, 0, 0, 0,
	22, 21, 23, 13, 14, 28, 0, 0, 26, 27,
	0, 0, 19, 77, 18, 20, 0, 0, 0, 22,
	21, 23, 13, 14, 28, 0, 0, 26, 27, 0,
	0, 19, 51, 18, 20, 105, 0, 22, 21, 23,
	13, 14, 28, 0, 0, 26, 27, 0, 0, 0,
	19, 18, 20, 37, 38, 39, 40, 0, 31, 32,
	0, 33, 34, 35, 36, 19, 0, 0, 19, 37,
	38, 39, 40, 0, 31, 32, 0, 33, 34, 35,
	36,
}
var yyPact = [...]int{

	383, -1000, 383, -1000, 396, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 193, 59, -1000, 53, 156, 346,
	147, -1000, -1000, -1000, -1000, -1000, 193, 8, -1000, -13,
	-1000, 193, 193, 193, 193, 193, 193, 193, 193, 193,
	193, 396, -27, 21, 140, 383, -1000, -1000, -17, -1000,
	327, -1000, 35, -1000, 380, -30, 57, 89, -1000, 107,
	96, 96, 90, 90, -1000, -1000, 189, 189, 189, 189,
	138, -1000, -18, 70, -1000, -1000, 69, -1000, 365, -1000,
	40, 218, 55, -1000, 47, -1000, 396, -1000, 68, 383,
	21, 17, -1000, -1000, 302, -1000, 396, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 138, 206, 126, 126, 21,
	-1000, 118, -1000, 193, -1000, 10, -1000, -1000, 21, 134,
	-1000, -1000, -19, 126, 21, 21, -1000, -1000, -1000, -1000,
	-1000, 56, 21, -1000, -1000, 7, -10, -1000, 288, -9,
	263, -1000, 396, -1000, -1000, -1000, -1000, -1000, -1000, 78,
	78, 288, -1000, -1000, 62, -1000, 46, 249, 383, 383,
	-1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 8, 43, 203, 0, 1, 200, 196, 75, 53,
	67, 23, 193, 7, 22, 10, 16, 190, 25, 189,
	187, 169, 185, 9, 30, 6, 5, 131, 173, 14,
	171, 3,
}
var yyR1 = [...]int{

	0, 28, 27, 27, 14, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 30, 30, 19, 19, 16, 16,
	16, 16, 16, 16, 16, 31, 31, 20, 20, 17,
	17, 18, 22, 23, 24, 1, 1, 1, 1, 1,
	1, 1, 1, 11, 11, 12, 12, 21, 26, 26,
	25, 29, 29, 13, 8, 8, 8, 8, 2, 10,
	10, 7, 7, 7, 7, 7, 6, 6, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 3,
	3, 3, 3, 3, 3, 9, 9, 5,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 5, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 7, 8, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 4, 4, 5,
	6, 5, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 5, 6, 3, 4, 1, 1, 3,
	1, 1, 3, 1, 1, 2, 1, 1, 3, 2,
	2, 7, 6, 5, 4, 6, 3, 3, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 1,
	1, 1, 1, 1, 1, 3, 2, 1,
}
var yyChk = [...]int{

	-1000, -28, -27, -1, -4, -8, -2, -18, -10, -11,
	-9, -14, -3, 7, 8, -6, -7, -5, 18, 35,
	19, 5, 4, 6, -12, -17, 12, 13, 9, -21,
	-1, 28, 29, 31, 32, 33, 34, 23, 24, 25,
	26, -4, -5, 27, 40, 30, 15, 16, -21, 9,
	-27, 36, -22, 9, -4, -2, 42, -5, 9, 37,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	40, -9, -23, -24, 9, -1, 37, 36, 35, -9,
	42, -4, 42, 38, -29, -13, -4, 9, 37, 30,
	38, -26, -25, 9, -30, -15, -4, -8, -2, -18,
	-10, -11, -9, -14, -19, 20, -4, 42, 42, -10,
	-9, -5, 38, 41, 38, -29, -1, -9, 38, 41,
	36, -15, -23, 42, -10, -10, -9, -13, 38, -9,
	-25, 37, -10, -9, -9, 38, -26, -9, 35, 38,
	-31, -16, -4, -8, -2, -18, -10, -9, -20, 39,
	17, 35, 36, -16, -24, 9, -24, -31, 30, 30,
	36, -1, -1,
}
var yyDef = [...]int{

	0, -2, 1, 2, 35, 36, 37, 38, 39, 40,
	41, 42, 68, 54, 0, 56, 57, 82, 0, 0,
	0, 79, 80, 81, 83, 84, 0, 0, -2, 0,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 55, 82, 0, 0, 0, 59, 60, 0, 47,
	0, 86, 0, 32, 0, 0, 0, 0, 87, 0,
	69, 70, 71, 72, 73, 74, -2, -2, -2, -2,
	0, 67, 0, 0, -2, 58, 0, 85, 0, 66,
	0, 0, 0, 45, 0, 51, 53, 33, 0, 0,
	0, 0, 48, 50, 0, 14, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 0, 0, 0, 0, 0,
	64, 0, 46, 0, 29, 0, 31, 43, 0, 0,
	4, 15, 0, 0, 0, 0, 63, 52, 30, 44,
	49, 0, 0, 65, 62, 0, 0, 61, 0, 0,
	0, 25, 18, 19, 20, 21, 22, 23, 24, 0,
	0, 0, 16, 26, 0, 34, 0, 0, 0, 0,
	17, 27, 28,
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:184
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:188
		{
			yyVAL.Expr = &ast.ReturnExpr{yyDollar[2].Expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:196
		{
			yyVAL.Expr = &ast.AssignExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:201
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Expr}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:205
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Expr}
		}
	case 61:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:211
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 62:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:215
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:219
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:223
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 65:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:227
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:234
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:238
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:249
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:253
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:257
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:261
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:265
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:269
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:273
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:277
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:281
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:285
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:292
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:296
		{
			yyVAL.Expr = &ast.BlockExpr{}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:301
		{
			yyVAL.Expr = &ast.Variable{yyDollar[1].String}
		}
	}
	goto yystack /* stack new state and value */
}
