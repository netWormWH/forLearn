package main

import (
	"strings"
	"fmt"
	ts "text/scanner"
	gs "go/scanner"
	"go/token"
	
)

const src = `
func main() {
	var s scanner.Scanner
	s.Filename = "example"
	s.Init(strings.NewReader(src))
	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		fmt.Println("At position", s.Pos(), ":", s.TokenText())
	}
}
`
var srcBytes=[]byte(src)
func main() {
	testTextScanner()
	testGoScanner()
}
func testTextScanner(){
	var s ts.Scanner
	s.Filename = "example"
	s.Init(strings.NewReader(src))
	var tok rune
	for tok != ts.EOF {
		tok = s.Scan()
		fmt.Println("At position", s.Pos(), ":", s.TokenText())
	}
}
func testGoScanner(){
	var s gs.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(srcBytes)) // register input "file"
	s.Init(file, srcBytes, nil /* no error handler */, gs.ScanComments)
	//s.Init(file, srcBytes,func(pos token.Position, msg string){fmt.Println("==============================>",msg)}, gs.ScanComments)
	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
	fmt.Println("ErrorCount:",s.ErrorCount)
}

