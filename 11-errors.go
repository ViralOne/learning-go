package main

import (
	"errors"
	"fmt"
)

// function that would throw an error when the errorCodeument is 13
func test_error(errorCode int) (int, error) {
	if errorCode == 13 {
		return -1, errors.New("Error, number recived: 13")
	}
	return errorCode, nil
}

// struct with an error code & message
type errorTemplate struct {
	errorCode    int
	errorMessage string
}

func (e *errorTemplate) Error() string {
	return fmt.Sprintf("%d - %s", e.errorCode, e.errorMessage)
}

func main() {

	for _, i := range []int{7, 13} {
		if r, e := test_error(i); e != nil {
			fmt.Println("Faild:", e)
		} else {
			fmt.Println("Passed:", r)
		}
	}
}
