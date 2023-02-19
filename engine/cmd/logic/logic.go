package main

import (
	"fmt"
	"os"
	"text/scanner"
	"unicode/utf8"

	"sevian.com/whendy/logic"
)

func main() {

	// src is the input that we want to tokenize.
	//src := []byte(LoadFile("texto.txt")) //"日本語"//("cos(x) + 1i*sin(x) // Euler")

	ll := logic.Lex{}
	ll.Start(LoadFile("texto.txt"))
	/*
		// Initialize the scanner.
		var s sc.Scanner
		fset := token.NewFileSet()                      // positions are relative to fset
		file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
		s.Init(file, src, nil /* no error handler * /, scanner.ScanComments)

		// Repeated calls to Scan yield the token sequence found in the input.
		for {
			pos, tok, lit := s.Scan()
			if tok == token.EOF {
				break
			}
			fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
		}
	*/

}
func main2() {
	source := LoadFile("texto.txt") //"日本語"

	p := Parse{}
	p.start(source)

	var s scanner.Scanner
	fmt.Println(s)

	//r, w := utf8.DecodeRuneInString(nihongo[1:])

	//fmt.Println(r, w, nihongo[0:2])
}

func LoadFile(name string) string {

	dat, err := os.ReadFile(name)

	if err != nil {
		panic(err)
	}

	str := string(dat)

	return str

}

type Parse struct {
	text  string
	state int
	pos   int
	typ   int
}

func (p *Parse) start(text string) {
	p.text = text
	p.pos = 0

	for {

		fmt.Println("len ", len(text))

		rune, w := utf8.DecodeRuneInString(p.text[p.pos:])

		if p.isDelimeter(rune) {
			fmt.Printf("%#U starts at byte position %d\n", rune, w)

		}
		fmt.Printf("%#U starts at byte position %d\n", rune, w)
		p.pos++
		if p.pos >= len(p.text) {
			break
		}

	}

}

func (p *Parse) isDelimeter(s rune) bool {

	switch {
	case s == '@':
		p.typ = 1
		return true
	}

	return false
}

func (p Parse) isS(s rune) bool {

	switch {
	case s == '@':
		return true
	}

	return false
}


