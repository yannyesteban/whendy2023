package logic

import (
	"fmt"
	"unicode/utf8"
)

const eof = -1

type Lex struct {
	text string
	pos  int
	typ  int
}

func (Lex) Test() {

	println("lex test")
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

func (l *lexer) 


next() rune {
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
