package main

import (
	"fmt"

	"github.com/ktr0731/go-fuzzyfinder/matching"
)

// ToDo: -v, -r, file reading, highlight
func main() {
	cases := []string{"hoge", "fuga", "test"}
	matched := matching.FindAll("test", cases, matching.WithMode(matching.ModeSmart))
	for _, v := range matched {
		fmt.Println(cases[v.Idx])
	}
}