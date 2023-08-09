package main

import "fmt"

func main() {
	for i:=0; i<10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	for i:=1; i<=10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " divisible by 2")
		} else if i%3 == 0 {
			fmt.Println(i, " divisible by 3")
		} else {
			fmt.Println(i, " not divisible by 2 or 3")
		}
	}
}