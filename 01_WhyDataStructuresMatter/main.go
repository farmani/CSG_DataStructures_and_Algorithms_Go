package main

import "fmt"

func print_numbers_version_one() {
	number := 2
	for number <= 100 {
		// If number is even, print it:
		if (number % 2) == 0 {
			fmt.Println(number)
		}

		number++
	}
}

func print_numbers_version_two() {
	number := 2
	for number <= 100 {
		fmt.Println(number)

		// Increase number by 2, which, by definition
		// is the next even number.
		number += 2
	}
}
