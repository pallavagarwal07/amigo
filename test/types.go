package main

import "fmt"

func main() {
	a := 0
	b := "abc"
	c := 0.0
	d := func(a int, b int) (float64, int) {
		return float64(a), a
	}
	e := func() func(int, int) (float64, int) {
		return d
	}
	f := struct {
		a int
		b string
	}{3, "hi"}
	fmt.Printf("'%T'\n'%T'\n'%T'\n'%T'\n'%T'\n'%T'\n", a, b, c, d, e, f)
}