package main

import "fmt"

type S struct{}

func main() {
	var p1, p2 *S
	s1 := S{}
	s2 := S{}
	p1 = &s1
	p2 = &s2
	fmt.Printf("%p\n", p1) // 0x1e52bc
	fmt.Printf("%p\n", p2) // 0x1e52bc
	fmt.Println(p1)        // &{}
	fmt.Println(p2)        // &{}
	fmt.Println(&p1)       // 0x40c138
	fmt.Println(&p2)       // 0x40c140
	fmt.Println(*p1)       // {}
	fmt.Println(*p2)       // {}
	fmt.Println(p1 == p2)  // true
	fmt.Println(s1 == s2)

	cn := struct {
		name          string
		positiveCases int
	}{
		"USA", 100,
	}

	us := struct {
		name          string
		positiveCases int
	}{
		"CNH", 100,
	}

	fmt.Println(cn == us)
	fmt.Println(&cn, &us)
}
