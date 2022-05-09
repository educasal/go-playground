package main

import "fmt"

func simplePointer() {
	first := 100
	var second = first
	var third *int = &first

	fmt.Println("first++")
	first++

	fmt.Println("First:", first)
	fmt.Println("Second:", second)
	fmt.Println("Third:", third)
	fmt.Println("*Third:", *third)

	fmt.Println("*third++")
	*third++

	fmt.Println("First:", first)
	fmt.Println("Second:", second)
	fmt.Println("Third:", third)
	fmt.Println("*Third:", *third)
}

func nilPointerDereference() {
	first := 100
	var second *int

	fmt.Println(second)
	// This will break
	// fmt.Println(*second)
	second = &first
	fmt.Println(second)
	fmt.Println(*second)
}

func doublePointer() {
	first := 100
	second := &first
	third := &second

	fmt.Println("First:", first)
	fmt.Println("&First:", &first)
	fmt.Println("Second:", second)
	fmt.Println("*Second:", *second)
	fmt.Println("&Second:", &second)
	fmt.Println("Third:", third)
	fmt.Println("*Third:", *third)
	fmt.Println("**Third:", **third)
}
