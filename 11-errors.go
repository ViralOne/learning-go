package main

import (
	"fmt"
)

func main() {

	// testing error
	err := notFoundError()
	ErrorTemplate, ok := err.(*ErrorTemplate)
	if ok {
		fmt.Println("Error code: ", notFoundError.Code)
		fmt.Println("Error message: ", notFoundError.Message)
	}
}

// create a function for errors
func (e *ErrorTemplate) Error() string {
	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}

type ErrorTemplate struct {
	Code    string
	Message string
}

func notFoundError() {
	return &ErrorTemplate{
		Code:    "404",
		Message: "Not Found",
	}
}
