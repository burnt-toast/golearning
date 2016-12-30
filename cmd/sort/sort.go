package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)

	fmt.Println(s)

	fruits := []string{"peach", "bananna", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
