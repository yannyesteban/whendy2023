package logic

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func lower(ch rune) rune { return ('a' - 'A') | ch } // returns lower-case ch iff ch is ASCII letter
// func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }
// func isHex(ch rune) bool     { return '0' <= ch && ch <= '9' || 'a' <= lower(ch) && lower(ch) <= 'f' }
func isLetter(ch rune) bool {
	return 'a' <= lower(ch) && lower(ch) <= 'z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

const eof = -1

type Data struct {
	token string
	data  map[string]interface{}
}

type Vars struct {
	input string // the string being scanned
	pos   int    // current position in the input
	atEOF bool   // we have hit the end of input and returned eof
	data  []Data
	line  int  // 1+number of newlines seen
	ch    rune // current character
	start int
}

func (vars *Vars) SetMap(pref string, data map[string]interface{}) {
	vars.data = append(vars.data, Data{token: pref, data: data})
	vars.start = 0

}

func (vars *Vars) Init(input string) {
	vars.input = input
}

func (vars *Vars) scanIdentifier() string {
	off := vars.start
	for rdOffset, b := range vars.input[vars.pos:] {
		vars.pos = rdOffset
		fmt.Println("..", off, b)
	}
	return ""
}

func (vars *Vars) isDataToken() bool {

	for index, v := range vars.data {
		sub := vars.input[vars.start : vars.start+len(v.token)]
		if sub == v.token {

			vars.next()
			rune := vars.ch
			if isLetter(rune) {
				lit := vars.scanIdentifier()

				fmt.Println("si", vars.ch, lit)
			}

		}
		fmt.Println(vars.ch, index, v.token, len(v.token), vars.input[vars.start:vars.start+len(v.token)])

	}
	return false
}

func (vars *Vars) Eval() {
	for {
		vars.next()

		if vars.atEOF {
			fmt.Println("bad")
			break
		}

		//fmt.Printf("%#U starts at byte position %d\n", f, 0)

		if vars.ch == '{' {
			vars.next()

			if vars.isDataToken() {

			}
			/*
				fmt.Printf("peekBack %c\n", g)
				if g >= 'a' && g <= 'z' || g >= 'A' && g <= 'Z' || g == '_' || g >= '0' && g <= '9' {

					fmt.Println("correo")
				}
			*/
		}
	}

}

func (vars *Vars) next() rune {
	if int(vars.pos) >= len(vars.input) {
		vars.atEOF = true
		return eof
	}
	r, w := utf8.DecodeRuneInString(vars.input[vars.pos:])
	vars.start = vars.pos
	vars.pos += int(w)
	if r == '\n' {
		vars.line++
	}
	vars.ch = r
	return r
}

func (vars *Vars) peek() rune {
	r := vars.next()
	vars.backup()
	return r
}

func (vars *Vars) backup() {
	if !vars.atEOF && vars.pos > 0 {
		r, w := utf8.DecodeLastRuneInString(vars.input[:vars.pos])
		vars.ch = r

		vars.pos -= int(w)
		// Correct newline count.
		if r == '\n' {
			vars.line--
		}
	}
}

type Lex struct {
	text string
	pos  int
	typ  int
}

func (Lex) Test() {

	println("lex test")
}
func (p *Lex) Main(text string) {
	mlex := lexer{input: text}

	for {
		f := mlex.next()

		if f == eof {
			fmt.Println("bad")
			break
		}

		//fmt.Printf("%#U starts at byte position %d\n", f, 0)

		if f == '@' {
			g := mlex.peekBack()
			fmt.Printf("peekBack %c\n", g)
			if g >= 'a' && g <= 'z' || g >= 'A' && g <= 'Z' || g == '_' || g >= '0' && g <= '9' {

				fmt.Println("correo")
			}

		}
	}

}

func (p *Lex) scanIdentifier() {

}

func (p *Lex) Start(text string) {

	p.text = text
	p.pos = 0

	mlex := lexer{input: text}

	fmt.Printf("%#U --- \n", mlex.next())

	for {

		fmt.Println("len ", len(text))

		rune, w := utf8.DecodeRuneInString(p.text[p.pos:])

		if p.isDelimeter(rune) {
			fmt.Printf("%#U starts at byte position %d\n", rune, w)

		}
		fmt.Printf("%#U starts at byte position... %d\n", rune, w)
		p.pos++
		break
		if p.pos >= len(p.text) {
			break
		}

	}
}

func (p *Lex) isDelimeter(s rune) bool {

	switch {
	case s == '@':
		p.typ = 1
		return true
	}

	return false
}

func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.atEOF = true
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += int(w)
	if r == '\n' {
		l.line++
	}
	return r
}

// lexer holds the state of the scanner.
type lexer struct {
	name       string // the name of the input; used only for error reports
	input      string // the string being scanned
	leftDelim  string // start of action marker
	rightDelim string // end of action marker
	pos        int    // current position in the input
	start      int    // start position of this item
	atEOF      bool   // we have hit the end of input and returned eof
	parenDepth int    // nesting depth of ( ) exprs
	line       int    // 1+number of newlines seen
	startLine  int    // start line of this item
	//item         item   // item to return to parser
	insideAction bool // are we inside an action?
	//options      lexOptions
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) peekBack() rune {
	if int(l.pos) == 0 {

		return 0
	}
	r, _ := utf8.DecodeRuneInString(l.input[l.pos-2:])

	return r
}

// backup steps back one rune.
func (l *lexer) backup() {
	if !l.atEOF && l.pos > 0 {
		r, w := utf8.DecodeLastRuneInString(l.input[:l.pos])
		l.pos -= int(w)
		// Correct newline count.
		if r == '\n' {
			l.line--
		}
	}
}
