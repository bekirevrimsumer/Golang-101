package main

import "fmt"

func main() {
	var x string = "Hello World"
	var y = "Hello World"
	z := "Hello World"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Multiple variables
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Multiple variables of the same type
	var d, e, f int
	d = 5
	e = 10
	f = 15
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// Multiple variables of different types
	var (
		g = 5
		h = "Hello"
		i = 10
	)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	// Multiple variables of different types
	var (
		j int
		k string
		l int
	)
	j = 5
	k = "Hello"
	l = 10
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	// Multiple variables of different types
	var (
		m int = 5
		n string = "Hello"
		o int = 10
	)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)

	var x1 int = 30
	var y1 int = 10
	
	var sum int = x1 + y1
	fmt.Println("x1 + y1 = ", sum)

	var sub int = x1 - y1
	fmt.Println("x1 - y1 = ", sub)

	var mul int = x1 * y1
	fmt.Println("x1 * y1 = ", mul)

	var div int = x1 / y1
	fmt.Println("x1 / y1 = ", div)

	var mod int = x1 % y1
	fmt.Println("x1 % y1 = ", mod)
}