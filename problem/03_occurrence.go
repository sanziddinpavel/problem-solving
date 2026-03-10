package problem

import "fmt"

func occurrence(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, v := range s {
		counts[v]++
		fmt.Println(counts)
	}
	return counts
}
