package main

import "fmt"

func main() {
	// create a loop that prints out the numbers from 1 to 10
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// create a loop that prints out the elements from an array
	crypto := []string{"bitcoin", "ethereum", "ripple", "litecoin", "monero"}
	for i := 0; i < len(crypto); i++ {
		fmt.Println(crypto[i])
	}

	// make a loop using range
	for i := range crypto {
		// when i == 2 go to label
		if i == 2 {
			goto testlabel
		}
		fmt.Println(crypto[i])
	}

	// create a test label
testlabel:
	fmt.Println("Test label")
}
