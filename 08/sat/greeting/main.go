// specify recipient of greeting

package main

import "fmt"

// Hello func
func Hello(name string) string {
	return "Hello, " + name
}
func main() {
	fmt.Println(Hello("john"))
}
