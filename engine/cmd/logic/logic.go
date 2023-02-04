package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	nihongo := LoadFile("texto.txt") //"日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)

	}
	//nihongo[1]
	r, w := utf8.DecodeRuneInString(nihongo[1:])

	fmt.Println(r, w, nihongo[0:2])
}

func LoadFile(name string) string {

	dat, err := os.ReadFile(name)

	if err != nil {
		panic(err)
	}

	str := string(dat)

	return str

}
