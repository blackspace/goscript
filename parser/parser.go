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
const LAMBDA = 57366
const GREATEREQUAL = 57367
const LESSEQUAL = 57368
const ELSE = 57369
const AND = 57370
const OR = 57371
const DOUBLEADD = 57372
const DOUBLESUB = 57373

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
	"LAMBDA",
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

//line parser/parser.go.y:325

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 65,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	-2, 72,
	-1, 66,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	-2, 73,
	-1, 67,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	-2, 74,
	-1, 68,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	-2, 75,
}

const yyNprod = 81
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 493

var yyAct = [...]int{

	4, 78, 140, 142, 10, 11, 29, 95, 83, 7,
	76, 139, 91, 127, 25, 24, 26, 41, 6, 29,
	25, 24, 26, 5, 75, 29, 111, 91, 55, 90,
	91, 51, 59, 60, 61, 62, 63, 64, 65, 66,
	67, 68, 57, 86, 87, 72, 148, 56, 69, 79,
	29, 37, 38, 39, 40, 104, 31, 32, 84, 70,
	82, 33, 34, 35, 36, 25, 24, 26, 15, 16,
	29, 119, 79, 27, 28, 25, 24, 26, 23, 102,
	29, 22, 96, 134, 103, 2, 89, 100, 109, 3,
	107, 99, 30, 113, 96, 106, 138, 22, 79, 100,
	98, 116, 115, 99, 46, 97, 88, 44, 52, 124,
	112, 122, 98, 110, 120, 121, 85, 97, 92, 45,
	47, 48, 45, 81, 46, 129, 130, 46, 29, 128,
	44, 43, 77, 136, 118, 73, 35, 36, 135, 143,
	79, 143, 30, 147, 150, 147, 74, 54, 146, 143,
	146, 151, 143, 147, 19, 150, 147, 145, 146, 145,
	50, 146, 144, 93, 144, 71, 1, 145, 117, 49,
	145, 42, 144, 47, 48, 144, 101, 13, 114, 9,
	46, 8, 42, 58, 14, 18, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 123, 17, 125, 42,
	31, 32, 12, 0, 30, 33, 34, 35, 36, 0,
	132, 0, 42, 30, 0, 30, 33, 34, 35, 36,
	0, 0, 30, 37, 38, 39, 40, 0, 31, 32,
	0, 0, 0, 33, 34, 35, 36, 0, 42, 0,
	108, 0, 42, 105, 25, 24, 26, 15, 16, 29,
	0, 0, 27, 28, 0, 0, 20, 23, 0, 108,
	108, 0, 0, 0, 21, 0, 37, 38, 39, 40,
	0, 31, 32, 0, 108, 0, 33, 34, 35, 36,
	22, 137, 25, 24, 26, 15, 16, 29, 0, 0,
	27, 28, 0, 0, 20, 23, 25, 24, 26, 15,
	16, 29, 21, 0, 27, 28, 0, 0, 20, 23,
	25, 24, 26, 0, 0, 29, 21, 0, 22, 133,
	0, 0, 25, 24, 26, 15, 16, 29, 0, 0,
	27, 28, 22, 131, 20, 23, 25, 24, 26, 15,
	16, 29, 21, 0, 27, 28, 0, 0, 20, 23,
	0, 0, 0, 0, 0, 0, 21, 0, 22, 126,
	0, 0, 25, 24, 26, 15, 16, 29, 0, 0,
	27, 28, 22, 80, 20, 23, 25, 24, 26, 15,
	16, 29, 21, 0, 27, 28, 0, 0, 20, 23,
	0, 0, 0, 0, 0, 0, 21, 0, 22, 53,
	25, 24, 26, 15, 16, 29, 0, 0, 27, 28,
	0, 0, 22, 23, 102, 25, 24, 26, 15, 16,
	29, 0, 0, 27, 28, 25, 24, 26, 15, 16,
	29, 0, 0, 27, 28, 0, 0, 94, 25, 24,
	26, 15, 16, 29, 0, 0, 27, 28, 0, 0,
	0, 22, 153, 0, 0, 25, 24, 26, 15, 16,
	29, 22, 152, 27, 28, 0, 37, 38, 39, 40,
	0, 31, 32, 0, 22, 149, 33, 34, 35, 36,
	0, 22, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 22, 141,
}
var yyPact = [...]int{

	372, -1000, 372, -1000, 241, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 306, 102, -1000, 88,
	151, -11, 358, 138, -1000, -1000, -1000, 306, -3, -1000,
	-1000, 306, 306, 306, 306, 306, 306, 306, 306, 306,
	306, 241, 65, 57, 16, 372, 137, -1000, -1000, -18,
	-1000, 89, 332, -1000, 83, 441, -37, 71, 85, 181,
	181, 99, 99, -1000, -1000, 170, 170, 170, 170, -1000,
	-1000, 0, 241, -1000, -1000, 63, -14, 78, -1000, -1000,
	-1000, 396, -1000, 10, 198, 41, -1000, 306, 73, -17,
	70, 131, 372, 61, -1000, -1000, 241, -1000, -1000, -1000,
	-1000, -1000, 125, 26, 119, 119, 57, -1000, 141, 241,
	372, 69, 372, -1000, 318, -1000, -1000, -29, -1000, 119,
	57, 57, -1000, 292, 372, 278, -1000, 40, 57, -1000,
	-1000, -1000, 240, -1000, 56, -32, -1000, -1000, 451, 6,
	434, -1000, -1000, 241, -1000, -1000, -1000, -1000, 421, -1000,
	-1000, 411, -1000, -1000,
}
var yyPgo = [...]int{

	0, 89, 202, 0, 197, 185, 23, 4, 9, 184,
	18, 181, 179, 5, 7, 3, 177, 176, 169, 168,
	1, 10, 154, 85, 166, 165, 163, 2,
}
var yyR1 = [...]int{

	0, 24, 23, 23, 13, 13, 14, 14, 14, 14,
	14, 14, 26, 26, 17, 17, 17, 17, 19, 15,
	15, 15, 15, 15, 27, 27, 16, 16, 9, 10,
	22, 22, 11, 11, 12, 12, 18, 21, 21, 20,
	25, 25, 1, 1, 1, 1, 1, 1, 1, 1,
	6, 6, 6, 6, 5, 5, 5, 5, 5, 4,
	4, 8, 8, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 2, 7,
	7,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 5, 4, 1, 1, 1, 1,
	1, 1, 1, 2, 7, 6, 8, 7, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 4, 1, 3,
	1, 3, 7, 8, 7, 6, 1, 1, 3, 1,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 1, 1, 7, 6, 5, 4, 6, 3,
	3, 2, 2, 1, 1, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 1, 1, 3,
	2,
}
var yyChk = [...]int{

	-1000, -24, -23, -1, -3, -6, -10, -8, -11, -12,
	-7, -13, -2, -16, -9, 7, 8, -4, -5, -22,
	16, 24, 40, 17, 5, 4, 6, 12, 13, 9,
	-1, 30, 31, 35, 36, 37, 38, 25, 26, 27,
	28, -3, -22, 29, 42, 34, 39, 32, 33, -18,
	9, 42, -23, 41, 9, -3, -10, 45, -22, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -7,
	43, -25, -3, -1, 9, 42, -21, 43, -20, 9,
	41, 40, -7, 45, -3, 45, 43, 44, 43, -21,
	43, 44, 40, -26, 41, -14, -3, -6, -10, -8,
	-13, -17, 18, -3, 45, 45, -8, -7, -22, -3,
	40, 43, 40, -20, -23, 41, -14, -19, 9, 45,
	-8, -8, -7, -23, 40, -23, 41, 42, -8, -7,
	-7, 41, -23, 41, 43, -21, -7, 41, 40, 43,
	-27, 41, -15, -3, -6, -10, -8, -7, 40, 41,
	-15, -27, 41, 41,
}
var yyDef = [...]int{

	0, -2, 1, 2, 42, 43, 44, 45, 46, 47,
	48, 49, 63, 64, 65, 50, 0, 52, 53, 28,
	0, 0, 0, 0, 76, 77, 78, 0, 0, 30,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 51, 28, 0, 0, 0, 0, 61, 62, 0,
	36, 0, 0, 80, 0, 0, 0, 0, 0, 66,
	67, 68, 69, 70, 71, -2, -2, -2, -2, 60,
	26, 0, 40, 29, 31, 0, 0, 0, 37, 39,
	79, 0, 59, 0, 0, 0, 27, 0, 0, 0,
	0, 0, 0, 0, 5, 12, 6, 7, 8, 9,
	10, 11, 0, 0, 0, 0, 0, 57, 0, 41,
	0, 0, 0, 38, 0, 4, 13, 0, 18, 0,
	0, 0, 56, 0, 0, 0, 35, 0, 0, 58,
	55, 32, 0, 34, 0, 0, 54, 33, 0, 0,
	0, 15, 24, 19, 20, 21, 22, 23, 0, 14,
	25, 0, 17, 16,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	42, 43, 37, 35, 44, 36, 39, 38, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 45,
	27, 34, 25, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 40, 3, 41,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 26, 28, 29, 30, 31, 32, 33,
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
	case 34:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:174
		{
			yyVAL.Expr = &ast.LambdaDefineExpr{yyDollar[3].Strings, yyDollar[6].Exprs}
		}
	case 35:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:178
		{
			yyVAL.Expr = &ast.LambdaDefineExpr{nil, yyDollar[5].Exprs}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:185
		{
			yyVAL.Strings = append(make([]string, 0), yyDollar[1].String)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:189
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:196
		{

			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:201
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:218
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:222
		{
			yyVAL.Expr = &ast.ReturnExpr{yyDollar[2].Expr}
		}
	case 54:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:229
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 55:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:233
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 56:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:237
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:241
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 58:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:245
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:250
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:254
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:262
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Strings}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:266
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Strings}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:275
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:279
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:283
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:287
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:291
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:295
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:299
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:303
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:307
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:311
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:318
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:322
		{
			yyVAL.Expr = &ast.BlockExpr{}
		}
	}
	goto yystack /* stack new state and value */
}
