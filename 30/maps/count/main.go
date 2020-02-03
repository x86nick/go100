// count tallies the number of times each line occurs within a file

package main

import (
	"fmt"
	"github.com/x86nick/go100/30/maps/datafile"
	"log"
)

// // with slice
// func main() {
// 	lines, err := datafile.GetStrings("votes.txt") // read votes.txt and retrun a slice of strings with every line from the file
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var names []string           // this variable will hold a slice of candidates names
// 	var counts []int             // hold a slice with the number of times each name occur
// 	for _, line := range lines { // process each line from the file
// 		matched := false
// 		for i, name := range names { // loop over each value in the name slice
// 			if name == line { // if the line matches the current name...
// 				counts[i]++    // ... increment the corresponding count
// 				matched = true // mark that we found the match
// 			}
// 		}
// 		if matched == false { // if no match was found...
// 			names = append(names, line) // ... add it as a new name
// 			counts = append(counts, 1)  // and add a new count - this line will be the first occurance
// 		}
// 	}
// 	for i, name := range names { // print the results
// 		fmt.Printf("%s: %d\n", name, counts[i]) // print each element from the names slice, and the corresponding element from the counts slice
// 	}

// }

// // -- with MAP

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Print(counts)
}
