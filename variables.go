package main

import "fmt"

func omittingTypes() {
	var price = 132.05
	var price2 = price
	fmt.Println("Price:", price)
	fmt.Println("Price:", price2)
}

func defaultValues() {
	var defaultInt int
	var defaultUint uint
	var defaultByte byte
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	var defaultRune rune

	fmt.Println("Default value of an int:", defaultInt)
	fmt.Println("Default value of a uint:", defaultUint)
	fmt.Println("Default value of a byte:", defaultByte)
	fmt.Println("Default value of an float64:", defaultFloat64)
	fmt.Println("Default value of an bool:", defaultBool)
	fmt.Println("Default value of an string:", defaultString)
	fmt.Println("Default value of an rune:", defaultRune)
}

func multipleVariables() {
	var price, vat float32 = 300, 25.50
	var quantity, inStock = 2, true

	fmt.Println("Total value:", price+vat)
	// this wouldn't work because of mismatched types
	// fmt.Println("Total value:", (price+vat)*quantity)
	fmt.Println("Quantity:", quantity)
	fmt.Println("In stock:", inStock)
}

func shorthandAssignment() {
	price, vat, quantity, inStock := 300.0, 25.50, 2, true

	fmt.Println("Total value:", price+vat)
	// this wouldn't work because of mismatched types
	// fmt.Println("Total value:", (price+vat)*quantity)
	fmt.Println("Quantity:", quantity)
	fmt.Println("In stock:", inStock)
}

func defaultTypes() {
	price, quantity := 300.0, 2

	// This will cause problems: the default types are float64 and int
	// var vat float32 = 24
	// var extraUnits int16 = 4

	var vat float64 = 24
	var extraUnits int = 4

	fmt.Println("Total value:", price+vat)
	fmt.Println("Total quantity:", quantity+extraUnits)
}
