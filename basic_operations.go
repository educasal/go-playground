package main

import (
	"fmt"
	"math"
	"strconv"
)

func conversionsFromFloatToInteger() {
	number := 25.5

	fmt.Println("Number:", number)
	fmt.Println("Ceil:", math.Ceil(number))
	fmt.Println("Floor:", math.Floor(number))
	fmt.Println("Round:", math.Round(number))
	fmt.Println("RoundToEven:", math.RoundToEven(number))
}

func parseBooleansFromStrings() {
	val1 := "true"
	val2 := "False"
	val3 := "0"
	val4 := "1"
	val5 := "faLSe"

	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)
	bool4, b4err := strconv.ParseBool(val4)
	bool5, b5err := strconv.ParseBool(val5)

	fmt.Println("Parsing booleans from strings")
	fmt.Println(val1, bool1, b1err)
	fmt.Println(val2, bool2, b2err)
	fmt.Println(val3, bool3, b3err)
	fmt.Println(val4, bool4, b4err)
	fmt.Println(val5, bool5, b5err)
}

func parseIntegersFromStrings() {
	val1 := "100"
	int1, int1err := strconv.ParseInt(val1, 10, 0)
	int2, int2err := strconv.ParseInt(val1, 2, 0)

	fmt.Println("Parsing integers from strings")
	fmt.Println(val1, int1, int1err)
	fmt.Println(val1, int2, int2err)
}
