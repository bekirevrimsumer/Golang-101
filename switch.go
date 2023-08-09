package main

import (
	"fmt"
	"time"
)

func main() {
	// for i :=0; i<10; i++ {
	// 	switch i {
	// 	case 0: fmt.Println("Zero")
	// 	case 1: fmt.Println("One")
	// 	case 2: fmt.Println("Two")
	// 	case 3: fmt.Println("Three")
	// 	case 4: fmt.Println("Four")
	// 	case 5: fmt.Println("Five")
	// 	case 6: fmt.Println("Six")
	// 	case 7: fmt.Println("Seven")
	// 	case 8: fmt.Println("Eight")
	// 	case 9: fmt.Println("Nine")
	// 	default: fmt.Println("Unknown Number")
	// 	}
	// }

	//if and switch examples
	//Program that prints all numbers 1-50 to the screen 4 by 4
	// for i := 1; i <= 50; i++ {
	// 	if i%4 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//Program that prints the day information according to the computer date to the screen

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Today is Monday")
	case time.Tuesday:
		fmt.Println("Today is Tuesday")
	case time.Wednesday:
		fmt.Println("Today is Wednesday")
	case time.Thursday:
		fmt.Println("Today is Thursday")
	case time.Friday:
		fmt.Println("Today is Friday")
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Unknown Day")
	}
}