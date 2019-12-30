package main

import "fmt"

//How do you test this? It is good to separate your "domain" code from the outside world (side-effects).
// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
// let's separate these concerns so it's easier to test

// func main() {
// 	fmt.Println("Hello, world")
// }
// // --
// func Hello() string {
// 	return "Hello, world"
// }
// func main() {
//	fmt.Println(Hello())
//}

// // --

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
