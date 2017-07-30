package main

import "fmt"

/*
// Ftoc prints two Fahrenheit-to-Celsius conversions.
*/


func main() {
	const freezingF, boilingF = 32.0, 212.0

	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) //32°C = 0°C
	fmt.Printf("%g°F= %g°C\n", boilingF, fToC(boilingF)) //32°C = 0°C
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
