// Area of square
package main

import "fmt"

var area int

func main() {
	var l int
	fmt.Println("enter lenght of square: ")
	fmt.Scan(&l)
	area = l * l
	fmt.Println("Area of square: ", area)
}
