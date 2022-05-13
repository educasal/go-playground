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

}
