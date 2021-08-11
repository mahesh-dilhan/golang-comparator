package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{10, 0, 3, 6}
	sort.Ints(s)
	fmt.Println(s)
}
