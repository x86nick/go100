package main

import "fmt"

var a int      // it is zero
var b int = 10 // declare
var c = 10     // c become int var
var d, e, f bool

var (
	g    int
	h    string
	i    int = 1234
	j, k bool
)

func main() {
	var a int      // it is zero
	var b int = 10 // declare
	var c = 10     // c become int var
	var d, e, f bool

	var (
		g    int
		h    string
		i    int = 1234
		j, k bool
	)

	m := 1 // statement, only inside a functon
	n, o := 2, 3
	e, f = f, e
	a, p := 100, 200
	a = "wrong" // can not change value at run time
	fmt.Println("Hello world")
	fmt.Println(a)
}
