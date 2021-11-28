package main

import "fmt"

type S struct{}

func main() {
	var p1, p2 *S

	fmt.Printf("%p%f\n", p1, p1) // nil
	fmt.Printf("%p%f\n", p2, p2) // nil

	s1 := S{}
	s2 := S{}

	p1 = &s1
	p2 = &s2

	fmt.Printf("%s\n", s1)
	fmt.Printf("%s\n", s2)
	fmt.Printf("%p\n", p1) // 0x1e52bc
	fmt.Printf("%p\n", p2) // 0x1e52bc

	fmt.Println(p1)       // &{}
	fmt.Println(p2)       // &{}
	fmt.Println(&p1)      // 0x40c138
	fmt.Println(&p2)      // 0x40c140
	fmt.Println(*p1)      // {}
	fmt.Println(*p2)      // {}
	fmt.Println(p1 == p2) // true
	fmt.Println(s1 == s2)

	us := struct {
		name          string
		positiveCases int
	}{
		"USA", 100,
	}

	cn := struct {
		name          string
		positiveCases int
	}{
		"CNH", 100,
	}

	countries := []struct {
		name          string
		posativecases int
	}{
		{"US", 100},
		{"RUS", 200},
	}

	fmt.Println(cn == us)
	fmt.Println(&cn, &us)

	fmt.Println(countries)
	var m map[int]string = map[int]string{}
	fmt.Println(m)
	var n, q map[int]string

	fmt.Println(n, q)
	var i, o chan int
	fmt.Println(i, o)

	var r, t []string
	fmt.Println(r, t)
}
