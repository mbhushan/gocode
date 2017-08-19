package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

/*
Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads numbers from its
command-line arguments or from the standard input if there are no arguments, and converts each number into units like
temperature in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.

*/

func main() {
	input, err := readInut()
	if err != nil {
		fmt.Println("Not valid input")
	}

	fmt.Printf("%f feet = %f meters, %f meters = %f feet\n",
		input, feetToMeter(input), input, meterToFeet(input))

}

func feetToMeter(feet float64) float64 {
	return feet / 3.2808
}

func meterToFeet(meter float64) float64 {
	return meter * 3.2808
}

func readInut() (float64, error) {
	input := bufio.NewScanner(os.Stdin)
	str  := input.Text()
	return strconv.ParseFloat(str, 64)
}
