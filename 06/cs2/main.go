package main

import "fmt"

func main() {
	/* ......
	err := f(x) // some func defined elsewhere
	if err != nil {
		fmt.Println(err.Error())
	*/
// other way to write
	if err := f(x); erro != nil {
		fmt.Println(err.Error())
	}
	}
}
