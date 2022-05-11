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
	val1 := "999"
	val2 := "110"
	val3 := "edu"
	int1, int1err := strconv.ParseInt(val1, 10, 0)
	int2, int2err := strconv.ParseInt(val2, 2, 0)
	int3, int3err := strconv.ParseInt(val3, 2, 0)

	fmt.Println("Parsing integers from strings")
	fmt.Println(val1, int1, int1err)
	fmt.Println(val2, int2, int2err)
	fmt.Println(val3, int3, int3err)

	int4, int4err := strconv.Atoi(val1)
	int5, int5err := strconv.Atoi(val2)
	int6, int6err := strconv.Atoi(val3)
	fmt.Println(val1, int4, int4err)
	fmt.Println(val2, int5, int5err)
	fmt.Println(val3, int6, int6err)
}

func parseFloatsFromStrings() {
	val1 := "48.95"
	val2 := "4.895e+01"
	val3 := "edu"
	int1, int1err := strconv.ParseFloat(val1, 64)
	int2, int2err := strconv.ParseFloat(val2, 32)
	int3, int3err := strconv.ParseFloat(val3, 32)

	fmt.Println("Parsing floats from strings")
	fmt.Println(val1, int1, int1err)
	fmt.Println(val2, int2, int2err)
	fmt.Println(val3, int3, int3err)
}

func formatValuesAsStrings() {
	val1 := true
	val2 := 110
	val3 := 4.895e+01

	str1 := strconv.FormatBool(val1)
	str2 := strconv.FormatInt(int64(val2), 2)
	str3 := strconv.FormatFloat(val3, 'f', 2, 64)
	str4 := strconv.Itoa(val2)

	fmt.Println("Fomatting different types as strings")
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)

}
