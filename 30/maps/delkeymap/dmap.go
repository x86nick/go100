package main

import "fmt"

func main() {
	var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3        // assign a value to bronze key
	rank, ok = ranks["bronze"] // "ok" wil be true becuase value is present
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	delete(ranks, "bronze") // delete bronze key and its corresponding value
	// fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	rank, ok = ranks["braonze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

}
