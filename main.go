package main

import (
	"errors"
	"fmt"
)

func main() {
	divResult, _ := div(5, 2)
	fmt.Println(divResult, "Hello, world!")
}

func div(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero!")
	}
	return a / b, nil
}
