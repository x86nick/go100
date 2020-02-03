package main

import "fmt"

func main() {
	ranks := map[string]int{"gold": 1, "silver": 2, "bronze": 3} // map literal
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])
	elements := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithium",
	}
	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])
	// - Exercise
	jewelry := make(map[string]float64)
	jewelry["neckless"] = 9.99
	jewelry["Earrings"] = 5.55
	cloathing := map[string]float64{"Pants": 59.99, "Shirt": 39.99}
	fmt.Println("Earrings", jewelry["Earrings"])
	fmt.Println("Pants", cloathing["Pants"])
	fmt.Println("------Zero value example------")
	counters := make(map[string]int)
	counters["a"] = 1
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["a"], counters["a"], counters["c"])
	fmt.Println("---- zero value if map is nil -------")
	var nilMap map[int]string
	fmt.Printf("%#v\n", nilMap)
	// nilMap[3] = "three" // panic: assignment to entry in nil map
	fmt.Println("---Before attempting to add keys and values, create a map using make or a map literal, and assign it to your map variable.---")
	var myMap map[int]string = make(map[int]string) // need to create a map first
	myMap[3] = "three"                              // then you can add values to it
	fmt.Printf("%#v\n", myMap)
	fmt.Println("-----How to tell zero values apart from assigned values------")
	status("Alma")
	status("Carl") // This code erroneously reports that the student "Carl" is failing, when in reality he just hasn’t had any grades logged:
	status("Rohit")
	fmt.Println("----To address situations like this, accessing a map key optionally returns a second, Boolean value--- ")
	fmt.Println("----Most Go developers assign this Boolean value to a variable named ok (because the name is nice and short)--- ")
	letsCount := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = letsCount["a"] // access an assigned value
	fmt.Println(value, ok)
	value, ok = letsCount["c"] // access an unassigned value
	fmt.Println(value, ok)
	fmt.Println("-----exercises-----")
	data := []string{"a", "c", "e", "a", "e"} // we'll count the number of times each letter occurs within this slice
	fmt.Println("Printing data------->:", data)
	counts := make(map[string]int) // a map to hold counts , letter as string and count as int
	fmt.Println("initilizing map to hold counts:", counts)
	for _, item := range data { // process each letter
		fmt.Println("processing item: ", item)
		counts[item]++ // increment the count for the current letter
		fmt.Println("printing counts and item: ", counts, item)
	}
	letters := []string{"a", "b", "c", "d", "e"} // we'll see if each of these letters exists as a key in the map
	fmt.Println("Printing letters-------->:", letters)
	for _, letter := range letters { // process each letter
		fmt.Println("// printing for _, letter--: ", letter)
		count, ok := counts[letter] // get the count for the current letter, as well as an indicator of wheter it was found at all
		fmt.Println("// block counts[letter] :", counts)
		fmt.Println("// checking count,ok block:", count, ok)
		if !ok { // if letter was not found
			fmt.Printf(" inside if okay block -- %s: not found\n", letter) // ...say so
		} else { // otherwise, letter was found...
			fmt.Printf(" eureka fount it -- %s: %d\n", letter, count) // so print the letter and the count that was recorded for it
		}
	}
	fmt.Println("-----extercise end-----")
}

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	// grade := grades[name]
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing\n", name)
	}

}
