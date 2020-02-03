package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"alise": 23.2, "hena": 76.2, "binku": 12.7}
	var names []string
	for name := range grades { //build a slice with all map keys
		fmt.Println("printing --for name := -->:", name)
		names = append(names, name)
		fmt.Println("printing appending to names--->:", names)
	}
	sort.Strings(names) // sort the slice alphabetically
	fmt.Println("printing sorted names --->:", names)
	for _, name := range names { // process the names alphabetically
		fmt.Println(name)
		fmt.Printf("%s has grade of %0.1f%%\n", name, grades[name]) // use the current student name to get the grade from the map
	}
}
