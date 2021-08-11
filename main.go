package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{10, 0, 3, 6}
	sort.Ints(i)
	fmt.Println(i)

	s := []string{"SG", "USA", "BEL", "ROC", "IND"}
	sort.Strings(s)
	fmt.Println(s)

}
