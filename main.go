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

	countries := []struct {
		name          string
		positiveCases int
	}{
		{"USA", 100},
		{"IND", 10},
		{"SL", 10},
		{"ROC", 100},
		{"CAN", 10},
	}

	sort.SliceStable(countries, func(i, j int) bool {
		return countries[i].positiveCases < countries[j].positiveCases
	})

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].name < countries[j].name
	})
}
