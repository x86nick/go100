package main

import (
	"fmt"
	"strings"
)

// broken := "Hell# W#rld" // syntax error: non-declaration statement outside function body
func main() {
	broken := "Hell# W#rld"
	// var broken string = "Hell# W#rld"
	r := strings.NewReplacer("#", "o") // Returns strings.Replacer thats setup to replace every # with o.
	fmt.Println(r)                     // &{{{0 0} 0} <nil> [# o]}
	// replacer := strings.NewReplacer("#", "o")
	fmt.Println(broken) // Hell# W#rld
	// var fixed string = r.Replace(broken)
	fixed := r.Replace(broken)
	fmt.Println(fixed) // Hello World
}
