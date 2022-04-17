package main

import "fmt"

func main() {
	var food string = "You are the food"
	println(food)

	chef := "You are the chef"
	fmt.Printf("%s\n", chef)

	var (
		first_dish              = "Pizza"
		first_dish_cooking_time = "30m"
		last_dish               = "Cake"
		last_dish_cooking_time  = "15m"
	)
	println(first_dish, first_dish_cooking_time, last_dish, last_dish_cooking_time)
}
