package main

import (
	"fmt"
	"os"
	"text/template"

	//"go/scanner"
	"go/token"
	"regexp"

	scanner "sevian.com/whendy/element"
)

type Todo struct {
	Name        string
	Description string
}

func main() {

	td := Todo{"8", "G"}

	t, err := template.New("todos").Parse(`{{ "agua" }} ......`)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}

	// src is the input that we want to tokenize.
	src := []byte(`a:=2+3*/ *8+++5`) //cos(x) + 1i*sin(x) // Euler")

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

	re := regexp.MustCompile("(?P<first_char>.)(?P<middle_part>.*)(?P<last_char>.)")
	n1 := re.SubexpNames()
	r2 := re.FindAllStringSubmatch("Super", -1)[0]

	md := map[string]string{}
	for i, n := range r2 {
		fmt.Printf("%d. match='%s'\tname='%s'\n", i, n, n1[i])
		md[n1[i]] = n
	}
	fmt.Printf("The names are  : %v\n", n1)
	fmt.Printf("The matches are: %v\n", r2)
	fmt.Printf("The first character is %s\n", md["first_char"])
	fmt.Printf("The last  character is %s\n", md["last_char"])

	var str = `aaaaa`

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}
	type itemType int

	const (
		itemError        itemType = iota // error occurred; value is text of error
		itemBool                         // boolean constant
		itemChar                         // printable ASCII character; grab bag for comma etc.
		itemCharConstant                 // character constant
		itemComment                      // comment text
		itemComplex                      // complex constant (1+2i); imaginary is just a number
		itemAssign                       // equals ('=') introducing an assignment
		itemDeclare                      // colon-equals (':=') introducing a declaration
		itemEOF
		itemField      // alphanumeric identifier starting with '.'
		itemIdentifier // alphanumeric identifier not starting with '.'
		itemLeftDelim  // left action delimiter
		itemLeftParen  // '(' inside action
		itemNumber     // simple number, including imaginary
		itemPipe       // pipe symbol
		itemRawString  // raw quoted string (includes quotes)
		itemRightDelim // right action delimiter
		itemRightParen // ')' inside action
		itemSpace      // run of spaces separating arguments
		itemString     // quoted string (includes quotes)
		itemText       // plain text
		itemVariable   // variable starting with '$', such as '$' or  '$1' or '$hello'
		// Keywords appear after all the rest.
		itemKeyword  // used only to delimit the keywords
		itemBlock    // block keyword
		itemBreak    // break keyword
		itemContinue // continue keyword
		itemDot      // the cursor, spelled '.'
		itemDefine   // define keyword
		itemElse     // else keyword
		itemEnd      // end keyword
		itemIf       // if keyword
		itemNil      // the untyped nil constant, easiest to treat as a keyword
		itemRange    // range keyword
		itemTemplate // template keyword
		itemWith     // with keyword
	)

	println(itemError, itemWith)

}
