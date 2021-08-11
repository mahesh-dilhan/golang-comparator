package main

import "fmt"

func changeVlaue(a []int) {
	a[0] = 10
}
func main() {
	s := []int{1, 2}
	changeVlaue(s)
	fmt.Println(s)
}
