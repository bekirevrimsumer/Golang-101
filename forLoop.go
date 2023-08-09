package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Infinite loop
	// for {
	// 	fmt.Println("Hello World")
	// }

	// While loop
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//Break
	// for i := 0; i < 10; i++ {
	// 	if i > 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// Continue
	// for i := 0; i < 10; i++ {
	// 	if i < 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// Nested loop
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		fmt.Println(i, j)
	// 	}
	// }

	// Nested loop with break
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j > 5 {
	// 			break
	// 		}
	// 		fmt.Println(i, j)
	// 	}
	// }

	// Nested loop with continue
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j < 5 {
	// 			continue
	// 		}
	// 		fmt.Println(i, j)
	// 	}
	// }
}