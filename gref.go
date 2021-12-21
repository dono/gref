package main

import (
	"fmt"

	"github.com/sahilm/fuzzy"
)

// ToDo: -v, -r, file reading, highlight
func main() {
	cases := []string{"hoge", "fuga", "foo"}
	matches := fuzzy.Find("zab", cases)

	const bold = "\033[1m%s\033[0m"

	for _, match := range matches {
		for i := 0; i < len(match.Str); i++ {
			if contains(i, match.MatchedIndexes) {
				fmt.Printf(bold, string(match.Str[i]))
			} else {
				fmt.Print(string(match.Str[i]))
			}
		}
		fmt.Println()
	}
}

func contains(needle int, haystack []int) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}