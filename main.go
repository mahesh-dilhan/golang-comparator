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

	fmt.Println(countries)

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].name < countries[j].name
	})
	fmt.Println(countries)

	covidCounties := []Country{
		{"USA", 100, 1},
		{"IND", 10, 2},
		{"SL", 10, 3},
		{"ROC", 100, 4},
		{"CAN", 10, 5},
	}

	sort.Sort(ByPositiveCases(covidCounties))
	fmt.Println(covidCounties)
}

type Country struct {
	Name          string
	PositiveCases int
	index         int
}
type ByPositiveCases []Country

func (c ByPositiveCases) Len() int {
	return len(c)
}
func (c ByPositiveCases) Less(a, b int) bool {
	return c[a].PositiveCases < c[b].PositiveCases
}

func (c ByPositiveCases) Swap(a, b int) {
	c[a], c[b] = c[b], c[a]
	c[a].index = a
	c[b].index = b
}
