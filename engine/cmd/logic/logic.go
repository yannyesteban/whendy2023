package main

import (
	"os"

	"sevian.com/whendy/logic"
)

func main() {

	ll := logic.Lex{}
	//ll.Start(LoadFile("texto.txt"))
	ll.Main(LoadFile("texto.txt"))

}

func LoadFile(name string) string {

	dat, err := os.ReadFile(name)

	if err != nil {
		panic(err)
	}

	str := string(dat)

	return str

}
