package main

import "fmt"

func main() {
	// new  - only allocates memory for the map - no init of memory
	// make - allocate memory for the map and init of memory

	// create a map named points that maps string to int
	points := make(map[string]int)
	points["milu"] = 10
	fmt.Println(points)

	// add new values to the map
	points["adi"] = 7
	points["tavi"] = 5
	points["ivan"] = 9
	fmt.Println(points)

	// change value of a key
	points["tavi"] = 8
	fmt.Println(points)

	// delete a value from the map
	delete(points, "adi")
	fmt.Println(points)
}
