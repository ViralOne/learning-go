package main

import "fmt"

func main() {
	var number = add(9, 16)
	square(number)
	oddEven(number)

	var numbers = []int{13, 8, 6, 2, 25, 87}
	addArray(numbers)

	fmt.Println(returnMultipleValus())
}

// create a function that takes two numbers and adds them together
func add(x int, y int) int {
	fmt.Println(x + y)
	return x + y
}

// create a function that takes a number and returns the square of that number
func square(x int) int {
	fmt.Println(x * x)
	return x * x
}

// create a function that takes a number and returns if the number is odd or even
func oddEven(x int) string {
	if x%2 == 0 {
		fmt.Println("even")
		return "even"
	} else {
		fmt.Println("odd")
		return "odd"
	}
}

// create a function that adds all the numbers from an array
func addArray(x []int) int {
	var sum int
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	fmt.Println(sum)
	return sum
}

// create a function that returns multiple values
func returnMultipleValus() (int, int, string) {
	lenght := 10
	width := 5
	name := "milu"
	return lenght, width, name
}
