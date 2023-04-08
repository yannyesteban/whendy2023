package main

import (
	"fmt"
	"os"

	"sevian.com/whendy/logic"
)

func main() {

	lg := logic.Vars{}
	lg.Init(LoadFile("texto2.txt"))
	//ll.Start(LoadFile("texto.txt"))
	data := map[string]string{

		"mode":     "update",
		"esteban":  "JIMENEZ",
		"yanny":    "NUñEZ",
		"apellido": "NEVER",
		"age":      "47",
	}
	lg.SetMap("@", data)

	lg.SetMap("&EX_", map[string]string{

		"apellido": "lopez",
	})
	fmt.Println(lg.Eval())

}
func Mmain2() {

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
