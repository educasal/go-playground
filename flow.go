package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func conditionals() {
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Int()

	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	numericString := "1234"

	if number%2 == 0 {
		numericString += "b"
	}

	if _, err := strconv.Atoi(numericString); err == nil {
		fmt.Println(numericString, "is a number")
	} else {
		fmt.Println(numericString, "is not a number")
	}

}

func loops() {
	count := 0

	fmt.Println("Simple for")
	for {
		fmt.Println(count)
		count++
		if count > 2 {
			break
		}
	}

	fmt.Println("For with only a condition")
	count = 0
	for count < 3 {
		fmt.Println(count)
		count++
	}

	fmt.Println("For with initialization and completion")
	for count2 := 0; count2 < 3; count2++ {
		fmt.Println(count2)
	}

	fmt.Println("For with continue")
	for count2 := 0; count2 < 3; count2++ {
		if count2%2 == 0 {
			continue
		}
		fmt.Println(count2)
	}

}

func enumerating() {
	name := "Gaspar Melchor de Jovellanos"

	for index, character := range name {
		fmt.Println(index, string(character))
	}
}

func switches() {

}
