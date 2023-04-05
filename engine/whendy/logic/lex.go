package logic

import (
	"fmt"
	"unicode/utf8"
)

const eof = -1

type Vars struct {
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

func (vars *Vars) Init(input string) {
	vars.input = input
}

func (vars *Vars) Eval() {
	for {
		f := vars.next()

		if f == eof {
			fmt.Println("bad1")
			break
		}

		//fmt.Printf("%#U starts at byte position %d\n", f, 0)

		if f == '{' {
			g := vars.next()
			fmt.Printf("peekBack %c\n", g)
			if g >= 'a' && g <= 'z' || g >= 'A' && g <= 'Z' || g == '_' || g >= '0' && g <= '9' {

				fmt.Println("correo")
			}

		}
	}

}

func (vars *Vars) next() rune {
	if int(vars.pos) >= len(vars.input) {
		vars.atEOF = true
		return eof
	}
	r, w := utf8.DecodeRuneInString(vars.input[vars.pos:])
	vars.pos += int(w)
	if r == '\n' {
		vars.line++
	}
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
