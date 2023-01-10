package main

import (
	"sevian.com/whendy"
	"sevian.com/whendy/element"
	"sevian.com/whendy/wc"
)

func Start() {
	element.Register("unit2", whendy.Unit{})
	element.Register("app", wc.App{})

}
