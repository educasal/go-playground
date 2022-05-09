package main

import (
	"fmt"
)

// func wrongTypes() {
// 	const price float32 = 300
// 	const vat float32 = 25.50
// 	const quantity int = 2

// 	fmt.Println("Total value:", (price+vat)*(quantity))
// }

func untypedConstant() {
	const price float32 = 300
	const vat float32 = 25.50
	const quantity = 2

	fmt.Println("Total value:", (price+vat)*quantity)
}

func multipleConstants() {
	const price, vat float32 = 300, 25.50
	const quantity, inStock = 2, true

	fmt.Println("Total value:", (price+vat)*quantity)
	fmt.Println("In stock:", inStock)
}
