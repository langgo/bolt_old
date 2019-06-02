package main

import (
	"fmt"
	"sort"
)

func main() {
	lead := []int{0, 1, 2, 3, 4, 5, 6}
	n := sort.Search(len(lead), func(i int) bool { return lead[i] > 4 })

	fmt.Println(n)
}
