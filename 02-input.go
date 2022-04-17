package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Don't do this
	// var user_name string
	// fmt.Scanln(&user_name)
	// fmt.Println("Hello " + user_name)

	// Do this instead
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	user_name, _ := reader.ReadString('\n')
	fmt.Println(user_name)

	// Change type
	fmt.Print("Enter your age: ")
	user_age, _ := reader.ReadString('\n')
	user_age_number, _ := strconv.ParseFloat(strings.TrimSpace(user_age), 64)
	fmt.Println(user_age_number * 10)

	if user_age_number <= 200 {
		fmt.Println("You are old")
	} else {
		fmt.Println("You are young")
	}

}
