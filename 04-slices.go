package main

import "fmt"

func main() {
	// create a slice of strings, maybe some food
	var food = []string{"pizza", "burger", "chips"}
	fmt.Println(food)

	// append a value to the slice
	food = append(food, "sandwich")
	fmt.Println(food)

	// using append show all the values from the second one to the last one minus one
	food = append(food[1 : len(food)-1])
	fmt.Println(food)

	//create a slice of crypto currencies and set a limit of 5
	crypto := make([]string, 5, 5)
	crypto[0] = "BTC"
	crypto[1] = "ETH"
	crypto[2] = "XRP"
	crypto[3] = "ADA"
	crypto[4] = "BCH"
	fmt.Println(crypto)

	// List allowed values for the slice
	fmt.Println(cap(crypto))
	//Append a value to the slice
	crypto = append(crypto, "LTC")
	fmt.Println(crypto)
	//Here we can append a new coin because it doesn't apply the above rule

	// List allowed values for the slice
	// The output is 10 because the slice is now bigger and GO adds for us more space to use it in the future
	fmt.Println(cap(crypto))
}
