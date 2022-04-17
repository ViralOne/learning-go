package main

import (
	"fmt"
)

func main() {

	// create an array of 5 numbers and print all the numbers from array and print the first and the last number
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(numbers[0], numbers[4])

	// update the last number of the array with 13 and print the number from array + 12
	numbers[4] = 13
	fmt.Println(numbers[4] + 12)

}
