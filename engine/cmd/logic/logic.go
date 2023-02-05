package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	source := LoadFile("texto.txt") //"日本語"

	p := Parse{}
	p.start(source)

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
