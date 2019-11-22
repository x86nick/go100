package main

import (
	"fmt"
	"strings"
)

func main() {
	// broken := "Hell# W#rld"
	var broken string = "Hell# W#rld"
	r := strings.NewReplacer("#", "o")
	// replacer := strings.NewReplacer("#", "o")
	fixed := r.Replace(broken)
	fmt.Println(fixed)

}
