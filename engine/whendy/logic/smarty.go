package logic

import (
	"fmt"
	"strconv"
)

type Item struct {
	Typ  int    // The type of this item.
	Pos  int    // The starting position, in bytes, of this item in the input string.
	Val  string // The value of this item.
	Line int    // The line number at the start of this item.
}

type Smarty struct {
	input string
	items []Item
}

func (s *Smarty) Init() {

	s.items = []Item{
		{Typ: 1, Pos: 1, Val: "8", Line: 1},
		{Typ: 2, Pos: 1, Val: "+", Line: 1},
		{Typ: 1, Pos: 1, Val: "4", Line: 1},
	}
}

func (s *Smarty) Eval() string {

	for index, i := range s.items {
		fmt.Println(index, i.Val)
	}

	return ""
}

func (s *Smarty) resolve() string {

	return ""
}

func (s *Smarty) prod(a string, b string, typ int) string {

	aa, _ := strconv.Atoi(a)
	bb, _ := strconv.Atoi(b)
	if typ == 1 {
		return string(aa * bb)
	} else if typ == 2 {
		return string(aa / bb)
	}
	return ""
}
