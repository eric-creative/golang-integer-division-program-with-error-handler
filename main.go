package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator int = 12
	var denominator int = 9
	var result, remainder, err = intDivision(numerator, denominator)

	switch {
		case err!=nil:
			fmt.Println(err.Error())
		case remainder==0:
			fmt.Printf(" the division is perfect with %v as the answer", result)
		default :
			fmt.Printf("The result of the division is %v, and the remainder is %v", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator==0{
		err = errors.New("can not divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
