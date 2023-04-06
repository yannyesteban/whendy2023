package main

import (
	"os"

	"sevian.com/whendy/logic"
)

func main() {

	lg := logic.Vars{}
	lg.Init(LoadFile("texto.txt"))
	//ll.Start(LoadFile("texto.txt"))
	data := map[string]interface{}{

		"mode": "update",
	}
	lg.SetMap("@", data)
	lg.Eval()

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
